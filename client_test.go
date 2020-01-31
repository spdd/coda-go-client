package main

import (
	"log"
	"testing"

	coda "github.com/spdd/coda-go-client/client"
)

func TestGetDeamonStatus(t *testing.T) {
	codaClient := coda.NewClient()
	ds, err := codaClient.GetDeamonStatus()
	if err != nil {
		log.Fatalln(err)
	}
	if ds.DaemonStatus.ConsensusMechanism != "proof_of_stake" {
		t.Errorf("DaemonStatus Fail")
	}
}
