package types

// Coda types

type ResponseData struct {
	Host string
	Type string
	Data *SubscriptionResponse
}

type Event struct {
	Response    chan *ResponseData
	Type        string
	Query       string
	Unsubscribe <-chan bool
	Subscribed  bool
	Count       int32
}

type Events struct {
	NewBlock          *Event
	SyncUpdate        *Event
	BlockConfirmation *Event
}

// Coda Api objects
// Daemon Status

type HelloSend struct {
	Type    string         `json:"type"`
	Payload map[string]int `json:"payload"`
}

type HelloReceive struct {
	Type    string         `json:"type"`
	Id      string         `json:"id"`
	Payload map[string]int `json:"payload"`
}

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
	SyncStatus                 string           `json:"syncStatus"`
	ProposePubkeys             []string         `json:"proposePubkeys"`
	ConsensusMechanism         string           `json:"consensusMechanism"`
	ConfDir                    string           `json:"confDir"`
	ConsensusConfiguration     map[string]int32 `json:"consensusConfiguration"`
}

// Daemon Version

type DaemonVersionResult struct {
	Version string `json:"version"`
}

type UniversalHttpResult struct {
	DaemonStatusResult
	DaemonVersionResult
}

// New Block Subscription

type SubscribeData struct {
	Type    string         `json:"type"`
	Id      string         `json:"id"`
	Payload SubscribeQuery `json:"payload"`
}

type SubscribeQuery struct {
	Query string `json:"query"`
}

type SubscriptionResponse struct {
	NewBlockSubscriptionResponse
	SyncUpdateSubscriptionResponse
}

type SyncUpdateSubscriptionResponse struct {
}

type NewBlockSubscriptionResponse struct {
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

// Sync Update Subscription

// Block Confirmation Subscription
