# Coda API Client for Go
Simple API Client for [Coda](https://codaprotocol.com/) GraphQL API written in Go

<p align="center">
    <img src="go.png" alt="gocoda" height="200" />
</p>

## Requirements

[Go](http://golang.org) 1.11 or newer.

# Installation 

`go get -u github.com/spdd/coda-go-client`

## Available GraphQL API

[Refer official API](https://codaprotocol.com/docs/developers/graphql-api)

|            Api description            | Status  | Testing  |            Function            |
| :---------------------------------:   | :----:  | :------: | :----------------------------: |
|        Get daemon status              |  :ok:   |   :ok:   |         GetDaemonStatus        |
|        Get daemon version             |  :ok:   |   :ok:   |         GetDaemonVersion       |
|        Get wallets                    |  :ok:   |   :ok:   |         GetWallets             |
|        Get wallet                     |  :ok:   |   :ok:   |         GetWallet(pk)          |
|        Unlock wallet                  |  :ok:   |   :ok:   |         UnlockWallet(pk, ps)   |
|        Create wallet                  |  :ok:   |   :ok:   |         CreateWallet(ps)       |
|        Send payment                   |  :ok:   |   :ok:   |         SendPayment(r,s,a,f,m) |
|        Get pooled payments            |  :ok:   |   :ok:   |         GetPooledPayments(pk)  |
|        Get transaction status         |  :ok:   |   :ok:   |    GetTransactionStatus(pID)   |
|        Set snark worker               |  :ok:   |   :ok:   |    SetSnarkWorker(workerPk,fee)|
|        Get current snark worker       |  :ok:   |   :ok:   |    GetCurrentSnarkWorker()     |
|        Get sync status                |  :ok:   |   :ok:   |    GetSyncStatus()             |
|        Get blocks                     |  :x:    |   :x:    |         -                      |
|      Subscription for new blocks      |  :ok:   |   :ok:   |      SubscribeForNewBlocks     |
| Subscription for Network Sync Updates |  :ok:   |   :ok:   |     SubscribeForSyncUpdates    |
| Subscription for Block Confirmations  |  :x:    |   :x:    | SubscribeForBlockConfirmations |
