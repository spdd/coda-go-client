package main

import (
	"log"
	"time"

	coda "github.com/spdd/coda-go-client/client"
)

func main() {
	getStatus()
	getStatusAsync()
}

func getStatus() {
	defer elapsed("Status")()
	client := coda.NewClient("http://192.168.100.100:3085/graphql", nil, nil)
	d, err := client.GetDaemonStatus()

	time.Sleep(time.Second)

	if err != nil {
		log.Fatalln(err)
	}
	log.Println("Block Number: ", d.DaemonStatus.BlockchainLength)
	log.Println("CommitId Number: ", d.DaemonStatus.CommitId)
	log.Println("")
}

// Async version
func getStatusAsync() {
	defer elapsed("Async Status")()
	client := coda.NewClient("http://192.168.100.100:3085/graphql", nil, nil)
	da := client.GetDaemonStatusCh()

	time.Sleep(time.Second)

	if da != nil {
		d := <-da
		log.Println("Block Number: ", d.DaemonStatus.BlockchainLength)
		log.Println("CommitId Number: ", d.DaemonStatus.CommitId)
		log.Println("")
	}
}

func elapsed(what string) func() {
	start := time.Now()
	return func() {
		log.Printf("%s took %v\n", what, time.Since(start))
		log.Println("")
	}
}
