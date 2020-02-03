package main

import (
	"log"
	"testing"

	coda "github.com/spdd/coda-go-client/client"
)

func TestGetDaemonStatus(t *testing.T) {
	codaClient := coda.NewClient()
	ds, err := codaClient.GetDaemonStatus()
	if err != nil {
		log.Fatalln(err)
	}
	if ds.DaemonStatus.ConsensusMechanism != "proof_of_stake" {
		t.Errorf("DaemonStatus Fail")
	}
}

func TestGetDaemonVersion(t *testing.T) {
	codaClient := coda.NewClient()
	ds, err := codaClient.GetDaemonVersion()
	if err != nil {
		log.Fatalln(err)
	}
	if ds.Version != "99d1e1f7a03be70f22d1a56acf38d2dd262b0d88" {
		t.Errorf("%v not %v", ds.Version, "99d1e1f7a03be70f22d1a56acf38d2dd262b0d88")
	}
}
