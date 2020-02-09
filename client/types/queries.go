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
