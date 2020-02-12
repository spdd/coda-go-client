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
	Unsubscribe chan bool
	Subscribed  bool
	Count       int32
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
	DaemonStatus struct {
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
	} `json:"daemonStatus"`
}

// Daemon Version

type DaemonVersionResult struct {
	Version string `json:"version"`
}

// Get Wallets

type GetWalletsResult struct {
	OwnedWallets []Wallet `json:"ownedWallets"`
}

type GetWalletResult struct {
	Wallet Wallet `json:"wallet"`
}

type Wallet struct {
	PublicKey string `json:"publicKey"`
	Balance   struct {
		Total   string `json:"total"`
		Unknown string `json:"unknown"`
	} `json:"balance"`
	Nonce            string `json:"nonce"`
	ReceiptChainHash string `json:"receiptChainHash"`
	Delegate         string `json:"delegate"`
	VotingFor        string `json:"votingFor"`
	StakingActive    bool   `json:"stakingActive"`
	PrivateKeyPath   string `json:"privateKeyPath"`
}

// Unlock Wallet

type UnlockWalletResult struct {
	UnlockWallet struct {
		Account struct {
			Balance struct {
				Total string `json:"total"`
			} `json:"balance"`
		} `json:"account"`
	} `json:"unlockWallet"`
}

// Create Account

type CreateWalletResult struct {
	CreateAccount struct {
		PublicKey string `json:"publicKey"`
	} `json:"createAccount"`
}

// Send Payment

type SendPaymentResult struct {
	SendPayment struct {
		Payment Payment `json:"payment"`
	} `json:"sendPayment"`
}

type Payment struct {
	Id           string `json:"id"`
	IsDelegation bool   `json:"isDelegation"`
	Nonce        int    `json:"nonce"`
	From         string `json:"from"`
	To           string `json:"to"`
	Amount       string `json:"amount"`
	Fee          string `json:"fee"`
	Memo         string `json:"memo"`
}

// Pooled Payment
type GetPooledPaymentResult struct {
	PooledPayments []Payment `json:"pooledUserCommands"`
}

type GetTransactionStatusResult struct {
	TransactionStatus string `json:"transactionStatus"`
}

// Snark Worker
type SetSnarkWorkerResult struct {
	SetSnarkWorker struct {
		LastSnarkWorker interface{} `json:"lastSnarkWorker"`
	} `json:"setSnarkWorker"`
	SetSnarkWorkFee struct{} `json:"setSnarkWorkFee"`
}

type GetCurrentSnarkWorkerResult struct {
	CurrentSnarkWorker struct {
		Key string `json:"key"`
		Fee string `json:"fee"`
	} `json:"currentSnarkWorker"`
}

type UniversalHttpResult struct {
	DaemonStatusResult
	DaemonVersionResult
	GetWalletsResult
	GetWalletResult
	UnlockWalletResult
	CreateWalletResult
	SendPaymentResult
	GetPooledPaymentResult
	GetTransactionStatusResult
	SetSnarkWorkerResult
	GetCurrentSnarkWorkerResult
	SyncStatus string `json:"syncStatus"`
}

// Subscription Query

type SubscribeDataQuery struct {
	Type    string         `json:"type"`
	Id      string         `json:"id"`
	Payload SubscribeQuery `json:"payload"`
}

type SubscribeQuery struct {
	Query string `json:"query"`
}

// Subscription Common

type SubscriptionResponse struct {
	BaseResponse
}

type BaseResponse struct {
	Type    string `json:"type"`
	Id      string `json:"id"`
	Payload struct {
		Data SubData `json:"data"`
	} `json:"payload"`
}

type SubData struct {
	NewBlock
	SyncUpdate
}

// New Block Subscription

type NewBlock struct {
	Block struct {
		Creator       string `json:"creator"`
		StateHash     string `json:"stateHash"`
		ProtocolState struct {
			PreviousStateHash string `json:"previousStateHash"`
			BlockchainState   struct {
				Date              string `json:"date"`
				SnarkedLedgerHash string `json:"snarkedLedgerHash"`
				StagedLedgerHash  string `json:"stagedLedgerHash"`
			} `json:"blockchainState"`
		} `json:"protocolState"`
		Transactions struct {
			UserCommands []struct {
				Id           string `json:"id"`
				IsDelegation bool   `json:"isDelegation"`
				Nonce        int    `json:"nonce"`
				From         string `json:"from"`
				To           string `json:"to"`
				Amount       string `json:"amount"`
				Fee          string `json:"fee"`
				Memo         string `json:"memo"`
			} `json:"userCommands"`
			FeeTransfer []struct {
				Recipient string `json:"recipient"`
				Fee       string `json:"fee"`
			} `json:"feeTransfer"`
			Coinbase string `json:"coinbase"`
		} `json:"transactions"`
	} `json:"newBlock"`
}

// Sync Update Subscription

type SyncUpdate struct {
	NewSyncUpdate string `json:"newSyncUpdate"`
}

// Block Confirmation Subscription
