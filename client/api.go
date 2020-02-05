package coda

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"os/signal"
	"strings"
	"time"

	"github.com/spdd/coda-go-client/client/types"
	"golang.org/x/net/websocket"
)

var graphqlEndpoint = "http://192.168.100.100:3085/graphql"

//var graphqlEndpoint = "https://graphql.o1test.net/graphql"
var websocketEndpoint = "ws://graphql.o1test.net/graphql"

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
	wsIn       chan types.NewBlockSubscribeResponse
	httpClient *http.Client
}

// NewClient create new client object
func NewClient(ws chan types.NewBlockSubscribeResponse) *Client {
	timeout := time.Duration(5 * time.Second)
	httpClient := &http.Client{
		Timeout: timeout,
	}
	return &Client{wsIn: ws, httpClient: httpClient}
}

func (client *Client) SubscribeNewBlock(url string, query string) {
	for {
		interrupt := make(chan os.Signal, 1)
		signal.Notify(interrupt, os.Interrupt)

		log.Printf("connecting to %s", url)
		origin := "http://localhost/"
		conn, err := websocket.Dial(url, "", origin)
		if err != nil {
			log.Fatal("dial:", err)
		}

		defer conn.Close()

		d2 := types.NewBlockSubscribeQuery{
			Type:    "start",
			Id:      "1",
			Payload: types.NewBlockQuery{Query: types.NewBlockSubscription},
		}
		s2, _ := json.Marshal(d2)
		log.Println(string(s2))
		// send message
		err3 := websocket.JSON.Send(conn, d2)
		if err3 != nil {
			log.Println("websocket.JSON.", err3)
		}
		//var m2 types.HelloSend
		//var m types.HelloReceive
		var m types.NewBlockSubscribeResponse
		// receive message
		// messageType initializes some type of message
		err4 := websocket.JSON.Receive(conn, &m)
		log.Println("Receive type:", m.Type)
		s3, _ := json.Marshal(m)
		if err4 != nil {
			log.Println("Error Receive", err4)
		}
		log.Printf("recv: %s", string(s3))
		client.wsIn <- m
	}
}

// Request HTTP request helper
func (c *Client) makeHttpRequest(url string, query string) (string, error) {
	requestBody, err := json.Marshal(map[string]string{
		"query": query,
	})
	request, err := http.NewRequest("POST", url, bytes.NewBuffer(requestBody))
	request.Header.Set("Content-Type", "application/json")
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
// GetDaemonStatus
func (c *Client) GetDaemonStatus() (*types.UniversalHttpResult, error) {
	return c.GetUniversal(types.DaemonStatusQuery)
}

// GetDaemonVersion
func (c *Client) GetDaemonVersion() (*types.UniversalHttpResult, error) {
	return c.GetUniversal(types.DaemonVersionQuery)
}

// GraphQL http/s query
func (c *Client) GetUniversal(query string) (*types.UniversalHttpResult, error) {
	response, err := c.makeHttpRequest(graphqlEndpoint, query)
	if err != nil {
		return nil, err
	}
	var ds types.UniversalHttpResult
	response = removeFromJsonString(response)
	log.Println("Result Universal:", response)
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
