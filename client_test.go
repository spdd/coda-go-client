package main

import (
	"context"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/http/httptest"
	"os"
	"os/signal"
	"syscall"
	"testing"

	"github.com/gorilla/websocket"
	coda "github.com/spdd/coda-go-client/client"
)

const (
	pk   = "4vsRCVzBeSxp3iBQ1C3ahHyKjKVbPd93JLSAsqtRtmjB9Xhn29NBdnzT4o6Hb3iNwaFECrh18YsxhAkqMY8nZQrN8jRX5LfbB9h4p5csrRe8xza4VWToXnFaHtGx6gB9FKAr1eKebSiPyH5c"
	pk2  = "4vsRCVVDuQkLqY1n4thRw5omsyXxLGSydqXsXzgjFsLdb14q3AHFGXNFSj8cQrWMAgd2QKLHpMTLwnLRfPrJaSvbnjLXPjRVtki1YWVNNr6XoRPs98MyBtEmbVda5siKRFzDtMxrhsvKHd8D"
	pID1 = "H5r6fvJjWciiAEod17QEdbsVzMaX8GPuan1cSpBmdRkrLd3aTd7rMZF4CSxSDZ5djKvyr5rXRAi2642kMFYVguQP2gFDfa5UCL4j3y2VyPS6z4cfGq47ECkHSCGoTeKTZXqjZtegMKfUqrbnZt8sYANaZEh9CjAVB5zkJitEozeznF3Z19VfTzfxpVsz9YfZd5dQjKDPak7BTBtvDgjgLramh8coS9jTxmvYJgfLPSSgHfoPXoemLGW4yG8bMTHXeGvWDokzWb8Xtr84pvGVqC2VRYN16KAaRVMs57uzBUi5avifCWdUE31MF7qwdYinizPhx6Qy86hg4aVoCNAVcFaY4BwMt181wL3zrQo6hzVrW8VGv7LysvvyHDV7MGQFk7LaUtCGnTzM4hNNpzbMQ9ZvRdpUjEAbr5q3h8qpD4btCpQun8HX4owSzT4YT8Eue8A9rYAVxZGtHTtayyeeZnuZr4aR2KryACxvFbuQZJWBjaNcMUedi2QW96vBb8jMXrXTonYA9eFsvUU24nK8KdQtGNeKAntvuGSPZooQGsTEHG2GAcCE6GZKboiZuC1UoD4xuWcsrD6e9XdjvusfjFENxvYhVYKMxoEX4juRb"
	pID2 = "H5r6fvJjWciiAEkoUJWbP8HUCfrXGUG7GHfWnfcUd3ZVQ7LvCRVTE6r3xjucedciGuHvpWkUGREzREPkZD9uQq9Bhg8NW4Lpsj9zDhfXBXCUB7Cnxv5QAPUCB3d33zanyXjx492vW3YryreCLQz3uwqkQP6ZheA7mwk3BNtbydXfRhhkVgCxBhf4MjhxGTPMsuXuM3MK7a1SBXpGYapA1BDSXPxaiTc73NRBbS4CxNfTEmiojeQ9yQDRHDV54sN8j2esFwoPwrLqDhfzKo8ZZoau9cCY71RtZFUEs3mbb2C1XFe4Rhjc74YxnWbtnNjMGptNugMc59mkEQ9KK76MRSeoMD8xa6N838zQLcLY2Nshmmb88qDPRZYPQv376NMksJcgG4XWPc8YjrFQYHTryUAizw4sEsPLvk6ZXPXTkMY3P5AiFciQWST9AhoYDxGXo1tpFDNt3shzjvyV5Vs418VNw7ididwmj7HrZ59KFrmXgDrY5UhJbi5kpvbdEVHsUnbjzoZoMStBVub4Dy2zQT9hw1xTG2W35mUteuVV5naLf3WvsQsUdHs8F4Cxa6K1L22z8N7nFVt4TBsmDtvNR3D6cPJ9A7Z1SmBJrDAy3"
	pID3 = "H5r6fvJjWciiAErSXvHuFrQk2Zdm68e1pXhgBVyVm2acojCgE6YQwc1BiMBGrQgxTyGeydWo1dMDGsdzzazCKyfMdJThomHnKdY64WErGEsGLHUhMb88sS7Rm6UsK6kW9NZUrxzhqZFvPXezgY6LfBvgbokEFAv3cNWu7AS3Hwo9LaRdavaW8Z4JaJ3HxrSWviicc85dXZCzaSuxWePi1VLLZeLiqPaCjxXxTKaNbxbfxJ1Y1WDGTnRWmpUWHLSzcaDhiaWWaSArWTF3na1mNQVfrK9iJopsEqjspYhfmh2xZLXF9S2eyN4vf9hYHhG9aGkCaaP25tnb7gZNvRoNkXQerHkVhgEqtGq49bvnWEnSEes14JFzaicqpm48sBnR9fcUYVmh81CkughXDonC9XCvVodGTcd3BbrBDZ3PBXEmNtF9h5X7wRk3vNsExtBJ98fxG6w3zox6XiJ1bCSfHR69Fbe3GXP556TM8La88qfqLyU3bDA1YRQEvhU23eiik7YvhjMEqK1ZQGD1LwYEhKsguPaEQGmZRPDNtgKoiU2t7WYiUoENAka2BHWX8ZNXds5WNx4oqhpfVcL5i9Rvtc3S3VghMiMyzHq7M6oPW"
)

func getHandler(filename string) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		fmt.Fprintln(w, getTestData(filename))
		return
	})
}

func getTestData(filename string) string {
	b, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Print(err)
	}
	return string(b)
}

func getCodaClient(filename string, hub *coda.Hub) (*coda.Client, string) {
	server := httptest.NewServer(getHandler(filename))
	client := coda.NewClientWith(server.Client(), server.URL, hub, nil)
	return client, server.URL
}

func TestGetDaemonStatus(t *testing.T) {
	client, _ := getCodaClient("test_data/status.json", nil)
	ds, err := client.GetDaemonStatus()
	if err != nil {
		log.Fatalln(err)
	}
	if ds.DaemonStatus.BlockchainLength != 1000000 {
		t.Errorf("DaemonStatus BlockchainLength != %d", 1000000)
	}
}

func TestGetDaemonVersion(t *testing.T) {
	client, _ := getCodaClient("test_data/daemon_version.json", nil)
	ds, err := client.GetDaemonVersion()
	if err != nil {
		log.Fatalln(err)
	}
	if ds.Version != "911631273ee694596407d1a181064ca6174cd8a3" {
		t.Errorf("%v not %v", ds.Version, "911631273ee694596407d1a181064ca6174cd8a3")
	}
}

func TestGetWallets(t *testing.T) {
	client, _ := getCodaClient("test_data/get_wallets.json", nil)
	ds, err := client.GetWallets()
	if err != nil {
		log.Fatalln(err)
	}
	if len(ds.OwnedWallets) == 0 {
		t.Errorf("GetWallets Fail")
	}
}

func TestGetWallet(t *testing.T) {
	client, _ := getCodaClient("test_data/get_wallet.json", nil)
	ds, err := client.GetWallet(pk)
	if err != nil {
		log.Fatalln(err)
	}
	if ds.Wallet.PublicKey != pk {
		t.Errorf("GetWallet Fail")
	}
}

func TestUnlockWallet(t *testing.T) {
	client, _ := getCodaClient("test_data/unlock_wallet.json", nil)
	ds, err := client.UnlockWallet(pk, "password")
	if err != nil {
		log.Fatalln(err)
	}
	if ds.UnlockWallet.Account.Balance.Total != "1501" {
		t.Errorf("UnlockWallet Fail")
	}
}

func TestSendPayment(t *testing.T) {
	client, _ := getCodaClient("test_data/send_payment.json", nil)
	ds, err := client.SendPayment(pk, pk2, 10, 5, "")
	if err != nil {
		log.Fatalln(err)
	}
	if ds.SendPayment.Payment.Id != pID3 {
		t.Errorf("SendPayment Fail")
	}
}

func TestTransactionStatus(t *testing.T) {
	client, _ := getCodaClient("test_data/transaction_status.json", nil)
	ds, err := client.GetTransactionStatus(pID1)
	if err != nil {
		log.Fatalln(err)
	}
	if ds.TransactionStatus != "UNKNOWN" {
		t.Errorf("TransactionStatus Fail")
	}
}

// Subscriptions test
var upgrader = websocket.Upgrader{}

func getWsHandler(filename string) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		upgrader.CheckOrigin = func(r *http.Request) bool { return true }
		c, err := upgrader.Upgrade(w, r, nil)
		if err != nil {
			log.Println("getWsHandler error", err)
			return
		}
		defer c.Close()
		for {
			mt, _, err := c.ReadMessage()
			if err != nil {
				break
			}
			err = c.WriteMessage(mt, []byte(getTestData(filename)))
			if err != nil {
				break
			}
		}
	})
}

func TestNewBlock(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	sigc := make(chan os.Signal, 1)
	signal.Notify(sigc, syscall.SIGINT, syscall.SIGTERM)

	hub := coda.NewHub()
	server := httptest.NewServer(getWsHandler("test_data/new_block.json"))
	defer server.Close()
	client := coda.NewClientWith(server.Client(), server.URL, hub, nil)
	go client.SubscribeForNewBlocks(ctx)

	for {
		select {
		case r := <-hub.SubscriptionData:
			// Unsubscribe from event
			client.SubscriptionEvents["NewBlock"].Unsubscribe <- true
			log.Printf("Unsubscribed successfully %v", <-client.SubscriptionEvents["NewBlock"].Unsubscribe)
			if r.Data.Payload.Data.Block.Creator != pk {
				t.Errorf("NewBlock Fail")
			}
			return
		case <-sigc:
			cancel()
			log.Println("System kill")
			os.Exit(0)
		}
	}
}
