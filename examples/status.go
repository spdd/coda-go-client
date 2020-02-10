package main

import (
	"fmt"
	"log"
	"time"

	coda "github.com/spdd/coda-go-client/client"
)

func main() {
	defer elapsed("Status")()
	client := coda.NewClient("http://192.168.100.100:3085/graphql", nil, nil)
	d, err := client.GetDaemonStatus()
	if err != nil {
		log.Fatalln(err)
	}
	log.Println("Block Number: ", d.DaemonStatus.BlockchainLength)
	log.Println("CommitId Number: ", d.DaemonStatus.CommitId)
}

// Async version
func asyncVersion() {
	defer elapsed("Async Status")()
	client := coda.NewClient("http://192.168.100.100:3085/graphql", nil, nil)
	da := client.GetDaemonStatusCh()
	if da != nil {
		d := <-da
		log.Println("Block Number: ", d.DaemonStatus.BlockchainLength)
		log.Println("CommitId Number: ", d.DaemonStatus.CommitId)
	}
}

func elapsed(what string) func() {
	start := time.Now()
	return func() {
		fmt.Printf("%s took %v\n", what, time.Since(start))
	}
}
