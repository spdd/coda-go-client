package main

import (
	"log"
	"testing"

	coda "github.com/spdd/coda-go-client/client"
)

const (
	pk   = "4vsRCVzBeSxp3iBQ1C3ahHyKjKVbPd93JLSAsqtRtmjB9Xhn29NBdnzT4o6Hb3iNwaFECrh18YsxhAkqMY8nZQrN8jRX5LfbB9h4p5csrRe8xza4VWToXnFaHtGx6gB9FKAr1eKebSiPyH5c"
	pk2  = "4vsRCVVDuQkLqY1n4thRw5omsyXxLGSydqXsXzgjFsLdb14q3AHFGXNFSj8cQrWMAgd2QKLHpMTLwnLRfPrJaSvbnjLXPjRVtki1YWVNNr6XoRPs98MyBtEmbVda5siKRFzDtMxrhsvKHd8D"
	pID1 = "H5r6fvJjWciiAEod17QEdbsVzMaX8GPuan1cSpBmdRkrLd3aTd7rMZF4CSxSDZ5djKvyr5rXRAi2642kMFYVguQP2gFDfa5UCL4j3y2VyPS6z4cfGq47ECkHSCGoTeKTZXqjZtegMKfUqrbnZt8sYANaZEh9CjAVB5zkJitEozeznF3Z19VfTzfxpVsz9YfZd5dQjKDPak7BTBtvDgjgLramh8coS9jTxmvYJgfLPSSgHfoPXoemLGW4yG8bMTHXeGvWDokzWb8Xtr84pvGVqC2VRYN16KAaRVMs57uzBUi5avifCWdUE31MF7qwdYinizPhx6Qy86hg4aVoCNAVcFaY4BwMt181wL3zrQo6hzVrW8VGv7LysvvyHDV7MGQFk7LaUtCGnTzM4hNNpzbMQ9ZvRdpUjEAbr5q3h8qpD4btCpQun8HX4owSzT4YT8Eue8A9rYAVxZGtHTtayyeeZnuZr4aR2KryACxvFbuQZJWBjaNcMUedi2QW96vBb8jMXrXTonYA9eFsvUU24nK8KdQtGNeKAntvuGSPZooQGsTEHG2GAcCE6GZKboiZuC1UoD4xuWcsrD6e9XdjvusfjFENxvYhVYKMxoEX4juRb"
	pID2 = "H5r6fvJjWciiAEkoUJWbP8HUCfrXGUG7GHfWnfcUd3ZVQ7LvCRVTE6r3xjucedciGuHvpWkUGREzREPkZD9uQq9Bhg8NW4Lpsj9zDhfXBXCUB7Cnxv5QAPUCB3d33zanyXjx492vW3YryreCLQz3uwqkQP6ZheA7mwk3BNtbydXfRhhkVgCxBhf4MjhxGTPMsuXuM3MK7a1SBXpGYapA1BDSXPxaiTc73NRBbS4CxNfTEmiojeQ9yQDRHDV54sN8j2esFwoPwrLqDhfzKo8ZZoau9cCY71RtZFUEs3mbb2C1XFe4Rhjc74YxnWbtnNjMGptNugMc59mkEQ9KK76MRSeoMD8xa6N838zQLcLY2Nshmmb88qDPRZYPQv376NMksJcgG4XWPc8YjrFQYHTryUAizw4sEsPLvk6ZXPXTkMY3P5AiFciQWST9AhoYDxGXo1tpFDNt3shzjvyV5Vs418VNw7ididwmj7HrZ59KFrmXgDrY5UhJbi5kpvbdEVHsUnbjzoZoMStBVub4Dy2zQT9hw1xTG2W35mUteuVV5naLf3WvsQsUdHs8F4Cxa6K1L22z8N7nFVt4TBsmDtvNR3D6cPJ9A7Z1SmBJrDAy3"
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

func TestGetWallets(t *testing.T) {
	client := coda.NewClient(endpoint, nil, nil)
	ds, err := client.GetWallets()
	if err != nil {
		log.Fatalln(err)
	}
	if len(ds.OwnedWallets) == 0 {
		t.Errorf("GetWallets Fail")
	}
}

func TestGetWallet(t *testing.T) {
	client := coda.NewClient(endpoint, nil, nil)
	ds, err := client.GetWallet(pk)
	if err != nil {
		log.Fatalln(err)
	}
	if ds.Wallet.PublicKey != pk {
		t.Errorf("GetWallet Fail")
	}
}

func TestTransactionStatus(t *testing.T) {
	client := coda.NewClient(endpoint, nil, nil)
	ds, err := client.GetTransactionStatus(pID1)
	if err != nil {
		log.Fatalln(err)
	}
	if ds.TransactionStatus != "INCLUDED" {
		t.Errorf("TransactionStatus Fail")
	}
}
