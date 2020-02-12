package main

import (
	"log"

	coda "github.com/spdd/coda-go-client/client"
)

const (
	pk  = "4vsRCVzBeSxp3iBQ1C3ahHyKjKVbPd93JLSAsqtRtmjB9Xhn29NBdnzT4o6Hb3iNwaFECrh18YsxhAkqMY8nZQrN8jRX5LfbB9h4p5csrRe8xza4VWToXnFaHtGx6gB9FKAr1eKebSiPyH5c"
	pk2 = "4vsRCVVDuQkLqY1n4thRw5omsyXxLGSydqXsXzgjFsLdb14q3AHFGXNFSj8cQrWMAgd2QKLHpMTLwnLRfPrJaSvbnjLXPjRVtki1YWVNNr6XoRPs98MyBtEmbVda5siKRFzDtMxrhsvKHd8D"
)

func main() {
	client := coda.NewClient("http://192.168.100.100:3085/graphql", nil, nil)
	setSnarkWorker(client, pk, "1")
	//setSnarkWorker(client, pk2, "1")
	GetCurrentSnarkWorker(client)
}

func setSnarkWorker(client *coda.Client, pk, fee string) {
	r, err := client.SetSnarkWorker(pk, fee)
	if err != nil {
		log.Fatalln(err)
	}
	log.Println("LastSnarkWorker Pk: ", r.SetSnarkWorker.LastSnarkWorker)
}

func GetCurrentSnarkWorker(client *coda.Client) {
	r, err := client.GetCurrentSnarkWorker()
	if err != nil {
		log.Fatalln(err)
	}
	log.Println("CurrentSnarkWorker key: ", r.CurrentSnarkWorker.Key)
	log.Println("CurrentSnarkWorker fee: ", r.CurrentSnarkWorker.Fee)
}
