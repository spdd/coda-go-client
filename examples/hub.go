package main

import (
	"log"

	coda "github.com/spdd/coda-go-client/client"
)

func main() {
	events := []string{
		"NewBlock",
		"SyncUpdate",
	}

	hub := coda.NewHub()
	go hub.Run()
	codaClient := coda.NewClient("http://192.168.100.100:3085/graphql", hub, events)
	codaClient2 := coda.NewClient("http://graphql.o1test.net/graphql", hub, events)
	hub.Subscribe <- codaClient
	hub.Subscribe <- codaClient2

	for {
		select {
		case status := <-hub.Result:
			log.Println("Status:", status)
		}
	}

}
