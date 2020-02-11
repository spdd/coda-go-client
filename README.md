# Coda API Client for Go
Simple API Client for [Coda](https://codaprotocol.com/) GraphQL API written in Go

<p align="center">
    <img src="go.png" alt="gocoda" height="200" />
</p>

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
|        get_pooled_payments            |  :ok:   |   :ok:   |         -                      |
|        get_transaction_status         |  :ok:   |   :ok:   |         -                      |
|        set_current_snark_worker       |  :ok:   |   :ok:   |         -                      |
|        get_current_snark_worker       |  :ok:   |   :ok:   |         -                      |
|        get_sync_status                |  :ok:   |   :ok:   |         -                      |
|        get_blocks                     |  :ok:   |   :ok:   |         -                      |
|      Subscription for new blocks      |  :ok:   |   :ok:   |      SubscribeForNewBlocks     |
| Subscription for Network Sync Updates |  :ok:   |   :ok:   |     SubscribeForSyncUpdates    |
| Subscription for Block Confirmations  |  :x:    |   :x:    | SubscribeForBlockConfirmations |

# Installation 

`go get -u github.com/spdd/coda-go-client`