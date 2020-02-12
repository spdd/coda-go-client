package types

var (
	// Daemon Status
	DaemonStatusQuery = `
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

	// Daemon Version
	DaemonVersionQuery = `
	{
		version
	}
`

	// Get Wallets
	GetWalletsQuery = `
	{
		ownedWallets {
		publicKey
		balance {
			total
		}
		}
	}
`

	// GetWallet
	GetWalletQuery = `
	query($publicKey:PublicKey!){
		wallet(publicKey:$publicKey) {
		publicKey
			balance {
				total
				unknown
			}
		nonce
		receiptChainHash
		delegate
		votingFor
		stakingActive
		privateKeyPath
		}
	}
`

	UnlockWalletQuery = `
	mutation ($publicKey: PublicKey!, $password: String!) {
		unlockWallet(input: {publicKey: $publicKey, password: $password}) {
		account {
			balance {
			total
			}
		}
		}
	}
`

	CreateWalletQuery = `
	mutation ($password: String!) {
		createAccount(input: {password: $password}) {
		publicKey
		}
	}
`

	SendPaymentQuery = `
	mutation($from:PublicKey!, $to:PublicKey!, $amount:UInt64!, $fee:UInt64!, $memo:String){
		sendPayment(input: {
		from:$from,
		to:$to,
		amount:$amount,
		fee:$fee,
		memo:$memo
		}) {
		payment {
			id,
			isDelegation,
			nonce,
			from,
			to,
			amount,
			fee,
			memo
		}
		}
	}
`

	GetPooledPaymentsQuery = `
	query ($publicKey:String!){
		pooledUserCommands(publicKey:$publicKey) {
		id,
		isDelegation,
		nonce,
		from,
		to,
		amount,
		fee,
		memo
		}
	}
`

	GetTransactionStatusQuery = `
	query($paymentId:ID!){
		transactionStatus(payment:$paymentId)
	}
`

	SetSnarkWorkerQuery = `
	mutation($worker_pk:PublicKey!, $fee:UInt64!){
	setSnarkWorker(input: {publicKey:$worker_pk}) {
		lastSnarkWorker
	}
	setSnarkWorkFee(input: {fee:$fee})
	}
`

	GetCurrentSnarkWorkerQuery = `
	{
		currentSnarkWorker{
		key
		fee
		}
	}
	`
	GetSyncStatusQuery = `
	{
		syncStatus
	}
	`

	// New Block Subscription

	NewBlockSubscriptionQuery = `
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

	SyncUpdateSubscriptionQuery = `
	subscription{
		newSyncUpdate 
	}
`

	BlockConfirmationSubscriptionQuery = `
	subscription{
		blockConfirmation {
		stateHash
		numConfirmations
		}
	}
`
)
