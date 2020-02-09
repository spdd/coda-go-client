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

var subscriptionEventsQueries = map[string]string{
	"NewBlock":          types.NewBlockSubscriptionQuery,
	"SyncUpdate":        types.SyncUpdateSubscriptionQuery,
	"BlockConfirmation": types.BlockConfirmationSubscriptionQuery,
}

// Client struct
type Client struct {
	SubscriptionEvents map[string]*types.Event
	httpClient         *http.Client
	Endpoint           string
	hub                *Hub
}

func createEvent(t string) *types.Event {
	return &types.Event{
		Response:    make(chan *types.ResponseData),
		Type:        t,
		Query:       subscriptionEventsQueries[t],
		Unsubscribe: make(chan bool),
		Subscribed:  false,
		Count:       0,
	}
}

func (c *Client) getEvent(t string) *types.Event {
	if event, ok := c.SubscriptionEvents[t]; ok {
		return event
	} else {
		event := createEvent(t)
		c.SubscriptionEvents[t] = event
		return event
	}
}

// NewClient create new client object
func NewClient(endpoint string, hub *Hub, eventsIt []string) *Client {
	timeout := time.Duration(5 * time.Second)
	httpClient := &http.Client{
		Timeout: timeout,
	}
	subEvents := make(map[string]*types.Event)
	for _, item := range eventsIt {
		event := createEvent(item)
		//log.Printf("Event: %v", event)
		subEvents[item] = event
	}
	return &Client{
		SubscriptionEvents: subEvents,
		Endpoint:           endpoint,
		httpClient:         httpClient,
		hub:                hub}
}

// Request HTTP request helper
func (c *Client) makeHttpRequest(query string) (string, error) {
	requestBody, err := json.Marshal(map[string]string{
		"query": query,
	})
	request, err := http.NewRequest("POST", c.Endpoint, bytes.NewBuffer(requestBody))
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
func (c *Client) GetDaemonStatusRepeat() {
	c.getUniversalRepeat(types.DaemonStatusQuery)
}

func (c *Client) GetDaemonStatusCh() <-chan *types.UniversalHttpResult {
	return c.getUniversalCh(types.DaemonStatusQuery)
}

func (c *Client) GetDaemonStatus() (*types.UniversalHttpResult, error) {
	return c.getUniversal(types.DaemonStatusQuery)
}

// GetDaemonVersion
func (c *Client) GetDaemonVersion() (*types.UniversalHttpResult, error) {
	return c.getUniversal(types.DaemonVersionQuery)
}

func (c *Client) getUniversalCh(query string) <-chan *types.UniversalHttpResult {
	ch := make(chan *types.UniversalHttpResult, 1)
	go func() {
		response, err := c.makeHttpRequest(query)
		if err != nil {
			ch <- nil
		}
		var ds types.UniversalHttpResult
		response = removeFromJsonString(response)
		log.Println("Result Universal2:", response)
		r := bytes.NewReader([]byte(response))
		err2 := json.NewDecoder(r).Decode(&ds)
		if err2 != nil {
			log.Fatalln(err2)
			ch <- nil
		}
		ch <- &ds
	}()
	return ch
}

// GraphQL http/s query
func (c *Client) getUniversal(query string) (*types.UniversalHttpResult, error) {
	response, err := c.makeHttpRequest(query)
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
		return nil, err2
	}
	return &ds, nil
}

func (c *Client) getUniversalRepeat(query string) {
	for {
		select {
		default:
			response, err := c.makeHttpRequest(query)
			if err != nil {
				log.Println(err)
				c.hub.Status <- nil
				return
			}
			var ds types.UniversalHttpResult
			response = removeFromJsonString(response)
			//log.Println("Result UniversalCh:", response)
			r := bytes.NewReader([]byte(response))
			err2 := json.NewDecoder(r).Decode(&ds)
			if err2 != nil {
				log.Fatalln(err2)
				c.hub.Status <- nil
			}
			log.Printf("%s Sync Status: %s", c.Endpoint, ds.DaemonStatus.SyncStatus)

			c.hub.Status <- &Status{Client: c, Status: &ds}
			time.Sleep(60 * time.Second)
		}
	}
}

// Subscription API
func (c *Client) subscribe(event *types.Event) {
	if event == nil {
		log.Println("Event is nil")
		return
	}
	for {
		select {
		default:
			event.Subscribed = true
			url := strings.Replace(c.Endpoint, "http", `ws`, -1)
			interrupt := make(chan os.Signal, 1)
			signal.Notify(interrupt, os.Interrupt)

			log.Printf("connecting to %s", url)
			origin := "http://localhost/"
			conn, err := websocket.Dial(url, "", origin)
			if err != nil {
				log.Println("dial:", err)
				log.Printf("Trying to connect to %s after %v seconds", url, 60)
				time.Sleep(60 * time.Second)
				continue
			}

			defer conn.Close()
			log.Printf("Subscription Type: %s", event.Type)
			d2 := types.SubscribeData{
				Type:    "start",
				Id:      "1",
				Payload: types.SubscribeQuery{Query: event.Query},
			}
			//s2, _ := json.Marshal(d2)
			//log.Println(string(s2))

			// send message
			err2 := websocket.JSON.Send(conn, d2)
			if err2 != nil {
				log.Println("websocket.JSON.", err2)
			}

			var m types.SubscriptionResponse
			// receive message
			// messageType initializes some type of message
			err3 := websocket.JSON.Receive(conn, &m)
			conn.Close()
			log.Println("Receive type:", m.Type)
			s3, _ := json.Marshal(m)
			if err3 != nil {
				log.Println("Error Receive", err3)
			}
			log.Printf("recv: %s", string(s3))
			responseData := &types.ResponseData{
				Host: c.Endpoint,
				Type: event.Type,
				Data: &m,
			}
			event.Count += 1
			if c.hub == nil {
				event.Response <- responseData
			} else {
				c.hub.SubscriptionData <- responseData
			}
			time.Sleep(1 * time.Second)
		case <-event.Unsubscribe:
			log.Printf("%s Unsubscribed from %s", c.Endpoint, event.Type)
			return
		}
	}
}

func (c *Client) SubscribeForEvent(event *types.Event) {
	c.subscribe(event)
}

func (c *Client) SubscribeForNewBlocks() {
	c.subscribe(c.getEvent("NewBlock"))
}

func (c *Client) SubscribeForSyncUpdates() {
	c.subscribe(c.getEvent("SyncUpdate"))
}

func (c *Client) SubscribeForBlockConfirmations() {
	c.subscribe(c.getEvent("BlockConfirmation"))
}
