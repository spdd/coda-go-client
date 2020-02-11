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
		subEvents[item] = event
	}
	return &Client{
		SubscriptionEvents: subEvents,
		Endpoint:           endpoint,
		httpClient:         httpClient,
		hub:                hub}
}

// Request HTTP request helper
func (c *Client) makeHttpRequest(query string, variables interface{}) (string, error) {
	payload, err := json.Marshal(map[string]string{
		"query": query,
	})

	if variables != "" {
		type Payload struct {
			Query     string      `json:"query"`
			Variables interface{} `json:"variables"`
		}
		p := Payload{
			Query:     query,
			Variables: variables,
		}
		payload, err = json.Marshal(p)
		if err != nil {
			log.Println(err)
		}
	}
	//log.Println(bytes.NewBuffer(payload))
	request, err := http.NewRequest("POST", c.Endpoint, bytes.NewBuffer(payload))
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

func getResponse(c *Client, query string, variables interface{}, ch chan *types.UniversalHttpResult) (*types.UniversalHttpResult, error) {
	response, err := c.makeHttpRequest(query, variables)
	if err != nil {
		if ch != nil {
			ch <- nil
		}
		return nil, err
	}
	var ds types.UniversalHttpResult
	response = removeFromJsonString(response)
	//log.Println("Result Universal2:", response)
	r := bytes.NewReader([]byte(response))
	err2 := json.NewDecoder(r).Decode(&ds)
	if err2 != nil {
		if ch != nil {
			ch <- nil
		}
		log.Println(err2)
		return nil, err2
	}
	if ch != nil {
		ch <- &ds
		close(ch)
	}
	return &ds, nil
}

func (c *Client) getUniversalCh(query string, variables interface{}) <-chan *types.UniversalHttpResult {
	ch := make(chan *types.UniversalHttpResult, 1)
	go func() {
		getResponse(c, query, variables, ch)
	}()
	return ch
}

// GraphQL http/s query
func (c *Client) getUniversal(query string, variables interface{}) (*types.UniversalHttpResult, error) {
	return getResponse(c, query, variables, nil)
}

func (c *Client) subscribe(event *types.Event) {
	if event == nil {
		log.Println("Event is nil")
		return
	}
	defer func() {
		log.Println("Exit Subscribtion: ", event.Type)
	}()
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
			}

			defer conn.Close()
			log.Printf("Subscription Type: %s", event.Type)
			d2 := types.SubscribeDataQuery{
				Type:    "start",
				Id:      "1",
				Payload: types.SubscribeQuery{Query: event.Query},
			}
			// send message
			err2 := websocket.JSON.Send(conn, d2)
			if err2 != nil {
				log.Println("websocket.JSON.", err2)
			}

			var m types.SubscriptionResponse
			// receive message
			// messageType initializes some type of message
			err3 := websocket.JSON.Receive(conn, &m)
			if err3 != nil {
				log.Println("Error Receive", err3)
			}
			conn.Close()
			log.Println("Receive type:", m.Type)

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
			event.Unsubscribe <- true
			return
		}
	}
}

// Coda API
// GetDaemonStatus

// get status concurrenly
func (c *Client) GetDaemonStatusCh() <-chan *types.UniversalHttpResult {
	return c.getUniversalCh(types.DaemonStatusQuery, "")
}

func (c *Client) GetDaemonStatus() (*types.UniversalHttpResult, error) {
	return c.getUniversal(types.DaemonStatusQuery, "")
}

// GetDaemonVersion
func (c *Client) GetDaemonVersion() (*types.UniversalHttpResult, error) {
	return c.getUniversal(types.DaemonVersionQuery, "")
}

// Get Owned Wallets
func (c *Client) GetWallets() (*types.UniversalHttpResult, error) {
	return c.getUniversal(types.GetWalletsQuery, "")
}

// Get Wallet
func (c *Client) GetWallet(pk string) (*types.UniversalHttpResult, error) {
	type PublicKey struct {
		Pk string `json:"publicKey"`
	}
	return c.getUniversal(types.GetWalletQuery, PublicKey{Pk: pk})
}

// Unlock wallet with password
func (c *Client) UnlockWallet(pk string, password string) (*types.UniversalHttpResult, error) {
	type UnlockWallet struct {
		Pk       string `json:"publicKey"`
		Password string `json:"password"`
	}
	return c.getUniversal(types.UnlockWalletQuery, UnlockWallet{Pk: pk, Password: password})
}

func (c *Client) CreateWallet(password string) (*types.UniversalHttpResult, error) {
	type CreateWallet struct {
		Password string `json:"password"`
	}
	return c.getUniversal(types.CreateWalletQuery, CreateWallet{Password: password})
}

func (c *Client) SendPayment(from, to string, amount, fee int, memo string) (*types.UniversalHttpResult, error) {
	type SendPayment struct {
		From   string `json:"from"`
		To     string `json:"to"`
		Amount int    `json:"amount"`
		Fee    int    `json:"fee"`
		Memo   string `json:"memo"`
	}
	return c.getUniversal(types.SendPaymentQuery,
		SendPayment{
			From:   from,
			To:     to,
			Amount: amount,
			Fee:    fee,
			Memo:   memo,
		})
}

// Subscription API

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
