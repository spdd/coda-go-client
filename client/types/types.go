package types

type HelloSend struct {
	Type    string         `json:"type"`
	Payload map[string]int `json:"payload"`
}

type HelloReceive struct {
	Type    string         `json:"type"`
	Id      string         `json:"id"`
	Payload map[string]int `json:"payload"`
}

// Daemon Status
var DaemonStatusQuery = `
	query {
		daemonStatus {
		numAccounts
		blockchainLength
		highestBlockLengthReceived
		uptimeSecs
		ledgerMerkleRoot
		stateHash
		commitId
		peers
		userCommandsSent
		snarkWorker
		snarkWorkFee
		syncStatus
		proposePubkeys
		consensusMechanism
		confDir
		commitId
		consensusConfiguration {
			delta
			k
			c
			cTimesK
			slotsPerEpoch
			slotDuration
			epochDuration
			acceptableNetworkDelay
		}
		}
	}
  `

type DaemonStatusResult struct {
	DaemonStatus DaemonStatusObj `json:"daemonStatus"`
}

func (ds DaemonStatusResult) PrintResult() {
}

type DaemonStatusObj struct {
	NumAccounts                int32            `json:"numAccounts"`
	BlockchainLength           int32            `json:"blockchainLength"`
	HighestBlockLengthReceived int32            `json:"highestBlockLengthReceived"`
	UptimeSecs                 int32            `json:"uptimeSecs"`
	LedgerMerkleRoot           string           `json:"ledgerMerkleRoot"`
	StateHash                  string           `json:"stateHash"`
	CommitId                   string           `json:"commitId"`
	Peers                      []string         `json:"peers"`
	UserCommandsSent           int32            `json:"userCommandsSent"`
	SnarkWorker                int32            `json:"snarkWorker"`
	SnarkWorkFee               int32            `json:"snarkWorkFee"`
	SyncStatus                 string           `json:"syncStatus "`
	ProposePubkeys             []string         `json:"proposePubkeys"`
	ConsensusMechanism         string           `json:"consensusMechanism"`
	ConfDir                    string           `json:"confDir"`
	ConsensusConfiguration     map[string]int32 `json:"consensusConfiguration"`
}

// Daemon Version
var DaemonVersionQuery = `
	{
		version
	}
`

type DaemonVersionResult struct {
	Version string `json:"version"`
}

type UniversalHttpResult struct {
	DaemonStatusResult
	DaemonVersionResult
}

// New Block Subscription

var NewBlockSubscription = `
	subscription(){
		newBlock(){
		creator
		stateHash
		protocolState {
			previousStateHash
			blockchainState {
			date
			snarkedLedgerHash
			stagedLedgerHash
			}
		},
		transactions {
			userCommands {
			id
			isDelegation
			nonce
			from
			to
			amount
			fee
			memo
			}
			feeTransfer {
			recipient
			fee
			}
			coinbase
		}
		}
	}
`

type NewBlockSubscribeQuery struct {
	Type    string        `json:"type"`
	Id      string        `json:"id"`
	Payload NewBlockQuery `json:"payload"`
}

type NewBlockQuery struct {
	Query string `json:"query"`
}

type NewBlockSubscribeResponse struct {
	Type    string       `json:"type"`
	Id      string       `json:"id"`
	Payload NewBlockData `json:"payload"`
}

type NewBlockData struct {
	Data NewBlock `json:"data"`
}

type NewBlock struct {
	Block InsideBlockObj `json:"newBlock"`
}

type InsideBlockObj struct {
	Creator       string           `json:"creator"`
	StateHash     string           `json:"stateHash"`
	ProtocolState ProtocolStateObj `json:"protocolState"`
	Transactions  TransactionsObj  `json:"transactions"`
}

type ProtocolStateObj struct {
	PreviousStateHash string             `json:"previousStateHash"`
	BlockchainState   BlockchainStateObj `json:"blockchainState"`
}

type TransactionsObj struct {
	UserCommands []UserCommandsObj `json:"userCommands"`
	FeeTransfer  []FeeTransferObj  `json:"feeTransfer"`
	Coinbase     string            `json:"coinbase"`
}

type UserCommandsObj struct {
	Id           string `json:"id"`
	IsDelegation bool   `json:"isDelegation"`
	Nonce        int    `json:"nonce"`
	From         string `json:"from"`
	To           string `json:"to"`
	Amount       string `json:"amount"`
	Fee          string `json:"fee"`
	Memo         string `json:"memo"`
}

type FeeTransferObj struct {
	Recipient string `json:"recipient"`
	Fee       string `json:"fee"`
}

type BlockchainStateObj struct {
	Date              string `json:"date"`
	SnarkedLedgerHash string `json:"snarkedLedgerHash"`
	StagedLedgerHash  string `json:"stagedLedgerHash"`
}
