package types

type QueryResult interface {
	PrintResult()
}

type Query interface {
	PrintQuery()
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
