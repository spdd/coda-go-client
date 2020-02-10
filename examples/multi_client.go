package main

import (
	"log"

	coda "github.com/spdd/coda-go-client/client"
)

func main() {

	hub := coda.NewHub()
	client1 := coda.NewClient("http://192.168.100.100:3085/graphql", hub, nil)
	client2 := coda.NewClient("http://graphql.o1test.net/graphql", hub, nil)

	go client1.SubscribeForNewBlocks()
	go client1.SubscribeForSyncUpdates()
	go client2.SubscribeForNewBlocks()

	blockCount := 0
	for {
		select {
		case r := <-hub.SubscriptionData:
			log.Println("Response Host:", r.Host)
			log.Println("Response Type:", r.Type)
			if r.Type == "NewBlock" {
				log.Println("NewBlock Creator:", r.Data.Payload.Data.Block.Creator)
				log.Println("NewBlock Count:", blockCount)
			}
			if r.Type == "NewBlock" {
				if blockCount == 2 {
					if r.Host == "http://192.168.100.100:3085/graphql" {
						//hub.Unsubscribe <- codaClient
						client1.SubscriptionEvents["NewBlock"].Unsubscribe <- true
						return
					}
				}
				if blockCount == 1 {
					if r.Host == "http://graphql.o1test.net/graphql" {
						client2.SubscriptionEvents["NewBlock"].Unsubscribe <- true
					}
				}
				blockCount += 1
			}
			if r.Type == "SyncUpdate" {
				log.Println("syncUpdate Arrived")
				if r.Host == "http://192.168.100.100:3085/graphql" {
					//hub.Unsubscribe <- codaClient
					client1.SubscriptionEvents["SyncUpdate"].Unsubscribe <- true
				}
			}
		}
	}
}
