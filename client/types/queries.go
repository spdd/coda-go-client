package types

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

// Daemon Version
var DaemonVersionQuery = `
{
	version
}
`

// Get Wallets
var GetWalletsQuery = `
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
var GetWalletQuery = `
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

var UnlockWalletQuery = `
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

var CreateWalletQuery = `
	mutation ($password: String!) {
		createAccount(input: {password: $password}) {
		publicKey
		}
	}
`

var SendPaymentQuery = `
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

// New Block Subscription

var NewBlockSubscriptionQuery = `
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

var SyncUpdateSubscriptionQuery = `
	subscription{
		newSyncUpdate 
	}
`

var BlockConfirmationSubscriptionQuery = `
	subscription{
		blockConfirmation {
		stateHash
		numConfirmations
		}
	}
`
