package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"

	coda "github.com/spdd/coda-go-client/client"
)

func main() {
	hub := coda.NewHub()
	client := coda.NewClient("http://192.168.100.100:3085/graphql", hub, nil)

	newBlock(client, hub)
}

func newBlock(client *coda.Client, hub *coda.Hub) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	sigc := make(chan os.Signal, 1)
	signal.Notify(sigc, syscall.SIGINT, syscall.SIGTERM)

	go client.SubscribeForNewBlocks(ctx)

	for {
		select {
		case r := <-hub.SubscriptionData:
			log.Println("We have a new block!")
			log.Println("NewBlock Creator:", r.Data.Payload.Data.Block.Creator)
			// Unsubscribe from event
			client.SubscriptionEvents["NewBlock"].Unsubscribe <- true
			log.Printf("Unsubscribed successfully %v", <-client.SubscriptionEvents["NewBlock"].Unsubscribe)
			return
		case <-sigc:
			cancel()
			log.Println("System kill")
			os.Exit(0)
		}
	}
}

func syncUpdate(client *coda.Client, hub *coda.Hub) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	go client.SubscribeForSyncUpdates(ctx)

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
