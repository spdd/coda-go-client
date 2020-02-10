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
|      Subscription for new blocks      |  :ok:   |   :ok:   |      SubscribeForNewBlocks     |
| Subscription for Network Sync Updates |  :ok:   |   :x:    |     SubscribeForSyncUpdates    |
| Subscription for Block Confirmations  |  :ok:   |   :x:    | SubscribeForBlockConfirmations |

# Installation 

`go get -u github.com/spdd/coda-go-client`