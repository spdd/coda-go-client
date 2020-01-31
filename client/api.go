package coda

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/spdd/coda-go-client/client/types"
)

var graphql_endpoint = "https://graphql.o1test.net/graphql"

//var graphql_endpoint = "http://localhost:3085/graphql"

func removeFromJsonString(jsonString string) string {
	jsonString = strings.Replace(jsonString, "'", `"`, -1)
	jsonString = strings.Replace(jsonString, "None", "0", -1)
	jsonString = strings.Replace(jsonString, "null", "0", -1)
	jsonString = strings.Replace(jsonString, `{"data":`, "", -1)
	jsonString = jsonString[:len(jsonString)-1]
	return jsonString
}

// Client struct
type Client struct {
	httpClient *http.Client
}

// NewClient create new client object
func NewClient() *Client {
	timeout := time.Duration(5 * time.Second)
	httpClient := &http.Client{
		Timeout: timeout,
	}
	return &Client{httpClient: httpClient}
}

// Request HTTP request helper
func (c *Client) makeRequest(url string, query string) (string, error) {
	requestBody, err := json.Marshal(map[string]string{
		"query": query,
	})
	request, err := http.NewRequest("POST", url, bytes.NewBuffer(requestBody))
	request.Header.Set("Content-Type", "application/json")
	fmt.Println(request)
	if err != nil {
		log.Fatalln(err)
	}

	resp, err := c.httpClient.Do(request)
	if err != nil {
		log.Fatalln(err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}
	return string(body), nil
}

// Coda API
// GetDeamonStatus
func (c *Client) GetDeamonStatus() (*types.DaemonStatusResult, error) {
	response, err := c.makeRequest(graphql_endpoint, types.DaemonStatusQuery)
	if err != nil {
		return nil, err
	}
	var ds types.DaemonStatusResult
	response = removeFromJsonString(response)
	//log.Println(string(response))
	r := bytes.NewReader([]byte(response))
	err2 := json.NewDecoder(r).Decode(&ds)
	if err2 != nil {
		log.Fatalln(err2)
	}
	if err2 != nil {
		return nil, err2
	}
	return &ds, nil
}
