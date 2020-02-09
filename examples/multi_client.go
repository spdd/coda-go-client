package main

import (
	"log"
	"os"
	"os/signal"

	coda "github.com/spdd/coda-go-client/client"
)

func main() {

	hub := coda.NewHub()
	codaClient := coda.NewClient("http://192.168.100.100:3085/graphql", hub, nil)
	codaClient2 := coda.NewClient("http://graphql.o1test.net/graphql", hub, nil)

	go codaClient.SubscribeForNewBlocks()
	go codaClient.SubscribeForSyncUpdates()
	go codaClient2.SubscribeForNewBlocks()

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
						codaClient.SubscriptionEvents["NewBlock"].Unsubscribe <- true
						return
					}
				}
				if blockCount == 1 {
					if r.Host == "http://graphql.o1test.net/graphql" {
						codaClient2.SubscriptionEvents["NewBlock"].Unsubscribe <- true
					}
				}
				blockCount += 1
			}
			if r.Type == "SyncUpdate" {
				log.Println("syncUpdate Arrived")
				if r.Host == "http://192.168.100.100:3085/graphql" {
					//hub.Unsubscribe <- codaClient
					codaClient.SubscriptionEvents["SyncUpdate"].Unsubscribe <- true
				}
			}
		}
	}

	ch := make(chan os.Signal, 1)
	signal.Notify(ch, os.Interrupt)
	<-ch
}
