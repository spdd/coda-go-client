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
	//getWallets(client)
	//getWallet(client, pk)
	//unlockWallet(client, pk)
	//createWallet(client, "")
	sendPayment(client, 10, 5)
}

func getWallets(client *coda.Client) {
	r, err := client.GetWallets()
	if err != nil {
		log.Fatalln(err)
	}
	log.Println("Wallet address: ", r.OwnedWallets[0].PublicKey)
	log.Println("Balance total: ", r.OwnedWallets[0].Balance.Total)
}

func getWallet(client *coda.Client, p string) {
	r, err := client.GetWallet(p)
	if err != nil {
		log.Fatalln(err)
	}
	log.Println("Wallet address: ", r.Wallet.PublicKey)
	log.Println("Balance total: ", r.Wallet.Balance.Total)
	log.Println("Delegate: ", r.Wallet.Delegate)
}

func unlockWallet(client *coda.Client, p string) {
	r, err := client.UnlockWallet(p, "")
	if err != nil {
		log.Fatalln(err)
	}
	log.Println("Balance total: ", r.UnlockWallet.Account.Balance.Total)
}

func createWallet(client *coda.Client, ps string) {
	r, err := client.CreateWallet(ps)
	if err != nil {
		log.Fatalln(err)
	}
	log.Println("Created New PublicKey: ", r.CreateAccount.PublicKey)
}

func sendPayment(client *coda.Client, amount, fee int) {
	r, err := client.SendPayment(pk, pk2, amount, fee, "123")
	if err != nil {
		log.Fatalln(err)
	}
	log.Println("Payment ID: ", r.SendPayment.Payment.Id)
	log.Println("Memo: ", r.SendPayment.Payment.Memo)
}
