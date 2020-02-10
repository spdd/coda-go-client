package main

import (
	"log"
	"testing"

	coda "github.com/spdd/coda-go-client/client"
)

var (
	endpoint = "http://192.168.100.100:3085/graphql"
)

func TestGetDaemonStatus(t *testing.T) {
	client := coda.NewClient(endpoint, nil, nil)
	ds, err := client.GetDaemonStatus()
	if err != nil {
		log.Fatalln(err)
	}
	if ds.DaemonStatus.ConsensusMechanism != "proof_of_stake" {
		t.Errorf("DaemonStatus Fail")
	}
}

func TestGetDaemonVersion(t *testing.T) {
	codaClient := coda.NewClient(endpoint, nil, nil)
	ds, err := codaClient.GetDaemonVersion()
	if err != nil {
		log.Fatalln(err)
	}
	if ds.Version != "911631273ee694596407d1a181064ca6174cd8a3" {
		t.Errorf("%v not %v", ds.Version, "911631273ee694596407d1a181064ca6174cd8a3")
	}
}
