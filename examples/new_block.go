package main

import (
	"log"

	coda "github.com/spdd/coda-go-client/client"
)

func main1() {

	hub := coda.NewHub()
	client := coda.NewClient("http://192.168.100.100:3085/graphql", hub, nil)

	go client.SubscribeForNewBlocks()

	for {
		select {
		case r := <-hub.SubscriptionData:
			log.Println("We have a new block!")
			log.Println("NewBlock Creator:", r.Data.Payload.Data.Block.Creator)
			// Unsubscribe from event
			client.SubscriptionEvents["NewBlock"].Unsubscribe <- true
			log.Printf("Unsubscribed successfully %v", <-client.SubscriptionEvents["NewBlock"].Unsubscribe)
			return
		}
	}
}

func main() {

	hub := coda.NewHub()
	client := coda.NewClient("http://192.168.100.100:3085/graphql", hub, nil)

	go client.SubscribeForSyncUpdates()

	for {
		select {
		case r := <-hub.SubscriptionData:
			log.Println("We have a new Sync Status!")
			log.Println("NewSyncUpdate:", r.Data.Payload.Data.SyncUpdate.NewSyncUpdate)
			// Unsubscribe from event
			client.SubscriptionEvents["SyncUpdate"].Unsubscribe <- true
			log.Printf("Unsubscribed successfully %v", <-client.SubscriptionEvents["SyncUpdate"].Unsubscribe)
			return
		}
	}
}
