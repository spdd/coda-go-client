package coda

import (
	"log"

	"github.com/spdd/coda-go-client/client/types"
)

type Status struct {
	Client *Client
	Status *types.UniversalHttpResult
}

type Hub struct {
	// Registered clients.
	Clients map[*Client]bool

	// Inbound statuses from the clients.
	Status chan *Status

	// Inbound subscription data from the clients.
	SubscriptionData chan *types.ResponseData
	// Unsubscribe client from events
	Unsubscribe chan *Client

	Subscribe chan *Client

	Broadcast chan string

	Result chan string
}

func NewHub() *Hub {
	return &Hub{
		Clients:          make(map[*Client]bool),
		Status:           make(chan *Status),
		SubscriptionData: make(chan *types.ResponseData),
		Unsubscribe:      make(chan *Client),
		Subscribe:        make(chan *Client),
		Broadcast:        make(chan string),
		Result:           make(chan string),
	}
}

func (h *Hub) Run() {
	for {
		select {
		case client := <-h.Subscribe:
			log.Printf("Subscribed: %v", client)
			h.Clients[client] = true
			if client.Endpoint == "http://graphql.o1test.net/graphql" {
				continue
			}
			//go client.GetDaemonStatusRepeat()
			//time.Sleep(5 * time.Second)
		case s := <-h.Status:
			client := s.Client
			if _, ok := h.Clients[client]; ok {
				for _, event := range client.SubscriptionEvents {
					if event.Subscribed {
						log.Printf("Already subscribed")
						if event.Count == 2 {
							log.Printf("Event Type: %v", event.Type)
							log.Printf("Event count: %v", event.Count)
							log.Printf("Endpoint: %v", client.Endpoint)
							log.Printf("Trying to unsubscribe")
							event.Unsubscribe <- true
						}
						continue
					}
					//log.Printf("Status s: %v", s.Status)
					log.Println("Blockchain Height:", s.Status.DaemonStatus.BlockchainLength)
					if s.Status.DaemonStatus.SyncStatus == "SYNCED" {
						log.Printf("Subscribe client for : %v", s.Client)
						if event.Subscribed {
							log.Printf("You %s already subscribed for %s", client.Endpoint, event.Type)
						} else {
							go client.SubscribeForEvent(event)
						}

						//time.Sleep(10 * time.Second)
					} else {
						if event.Subscribed {
							event.Unsubscribe <- true
						}
					}
				}
			}
		case r := <-h.SubscriptionData:
			log.Println("Response Host:", r.Host)
			log.Println("Response Type:", r.Type)
			if r.Type == "NewBlock" {
				log.Println("NewBlock Creator:", r.Data.Payload.Data.Block.Creator)
			}
			if r.Type == "syncUpdate" {
				log.Println("syncUpdate Arrived")
			}
		}
	}
}
