package main

import (
	"log"

	coda "github.com/spdd/coda-go-client/client"
)

func main2() {
	codaClient := coda.NewClient("http://192.168.100.100:3085/graphql")
	ds, err := codaClient.GetDaemonStatus()
	if err != nil {
		log.Fatalln(err)
	}
	log.Println(ds.DaemonStatus.NumAccounts)
	log.Println(ds.DaemonStatus.BlockchainLength)
	log.Println(ds.DaemonStatus.Peers)
	log.Println(ds.DaemonStatus.SnarkWorker)

	v, err := codaClient.GetDaemonVersion()
	if err != nil {
		log.Fatalln(err)
	}
	log.Println("Version:", v.Version)

	ur, err := codaClient.GetDaemonStatus()
	if err != nil {
		log.Fatalln(err)
	}
	log.Println("Universal Status:", ur.DaemonStatus.BlockchainLength)

	ur2, err := codaClient.GetDaemonVersion()
	if err != nil {
		log.Fatalln(err)
	}
	log.Println("Universal Version:", ur2.Version)

	go codaClient.SubscribeForNewBlock()
	blockCount := 0
	for {
		select {
		case r := <-codaClient.SubscriptionEvents.NewBlock.Response:
			if blockCount == 2 {
				codaClient.SubscriptionEvents.NewBlock.Unsubscribe <- true
				log.Println("NewBlock Creator:", r.Payload.Data.Block.Creator)
				log.Println("NewBlock Count:", blockCount)
				return
			}
			blockCount += 1
		}
	}
}
