package main

import (
	"fmt"
	"log"
	"time"

	coda "github.com/spdd/coda-go-client/client"
)

func main() {
	defer elapsed("Status")()
	codaClient := coda.NewClient("http://192.168.100.100:3085/graphql", nil, nil)
	statusCunc(codaClient)
	//statusSync(codaClient)
}

func statusCunc(codaClient *coda.Client) {
	da := codaClient.GetDaemonStatusCh()
	time.Sleep(5 * time.Second)
	d := <-da
	log.Println(d.DaemonStatus.BlockchainLength)
}

func statusSync(codaClient *coda.Client) {
	ds, _ := codaClient.GetDaemonStatus()
	time.Sleep(5 * time.Second)
	log.Println(ds.DaemonStatus.BlockchainLength)
}

// 1. da := <-codaClient.GetDaemonStatusCh()
//	Status took 7.181462596s
// Status took 7.10634263s
// Status took 7.424387586s

// 2. da := codaClient.GetDaemonStatusCh()
// d := <-da
// 1 run Status took 7.04974913s
// 2 run Status took 7.050652611s
// 3 run Status took 7.172364076s
// 4 run

func elapsed(what string) func() {
	start := time.Now()
	return func() {
		fmt.Printf("%s took %v\n", what, time.Since(start))
	}
}
