package main

import (
	"log"

	coda "github.com/spdd/coda-go-client/client"
)

var response = `
{
	'daemonStatus':{
	   'numAccounts':10000,
	   'blockchainLength':35,
	   'highestBlockLengthReceived':35,
	   'uptimeSecs':113612,
	   'ledgerMerkleRoot':'SBKtucGLzzz2r9jBDSuXwCAS2VUqtqmEYAzsctWod3xY5HTrTT4d6xVCeeV29qxeiXvx1tJQYLjFpqZCpuDRxDzricXfttCiPKX4oBCBsLJEJPNBL4a8Z6GEjpQtUKfqQv7cYH2EPhq73D',
	   'stateHash':'994hGvSU4kGrX1xCdyb2YVe7qSwJ9s6tftsvDsNKexe6RSvxYMvpv5WzGvfqA1tYBTQbgBMgnMeESkkMp2Z5CEfVocTPwz6SbMsC8KtyuWk9EhZRt1mE98xbjg5zXS6RVfoVFgyYQTQnR8',
	   'commitId':'1dffa7cc1438d3e39eb7182b812e882e4a94a7d7',
	   'peers':[
		  '144.208.127.129:8303',
		  '78.46.86.168:8303',
		  '77.56.116.29:8303',
		  '13.90.208.62:8303',
		  '185.224.249.189:8303',
		  '159.69.72.206:8303',
		  '82.206.30.6:8303',
		  '64.227.88.59:8303',
		  '163.172.81.135:8303',
		  '149.202.87.123:8303',
		  '3.14.73.25:8303',
		  '121.166.50.110:8303',
		  '178.128.124.69:8303',
		  '142.44.139.163:8303',
		  '165.22.241.104:8303',
		  '195.201.159.253:8303',
		  '34.68.45.202:8303',
		  '45.77.142.132:8303',
		  '62.171.131.155:8303',
		  '47.56.137.241:8303',
		  '139.99.237.65:8303',
		  '134.209.49.152:8303',
		  '108.61.78.47:8303',
		  '34.89.223.219:8303',
		  '54.37.17.216:8303',
		  '138.68.13.220:8303',
		  '198.50.168.190:8303',
		  '195.201.159.252:8303',
		  '64.225.114.171:8303',
		  '85.159.212.72:8303',
		  '35.234.91.192:8303',
		  '3.15.204.131:8303',
		  '35.233.131.48:8303',
		  '108.61.203.157:8303',
		  '62.171.131.202:8303',
		  '74.105.145.200:8303',
		  '104.194.8.175:8303',
		  '35.197.9.68:8303',
		  '54.183.66.160:8303'
	   ],
	   'userCommandsSent':0,
	   'snarkWorker':None,
	   'snarkWorkFee':1,
	   'syncStatus':'SYNCED',
	   'proposePubkeys':[
		  '4vsRCVzBeSxp3iBQ1C3ahHyKjKVbPd93JLSAsqtRtmjB9Xhn29NBdnzT4o6Hb3iNwaFECrh18YsxhAkqMY8nZQrN8jRX5LfbB9h4p5csrRe8xza4VWToXnFaHtGx6gB9FKAr1eKebSiPyH5c'
	   ],
	   'consensusMechanism':'proof_of_stake',
	   'confDir':'/home/rut/.coda-config',
	   'consensusConfiguration':{
		  'delta':3,
		  'k':10,
		  'c':8,
		  'cTimesK':80,
		  'slotsPerEpoch':240,
		  'slotDuration':180000,
		  'epochDuration':43200000,
		  'acceptableNetworkDelay':540000
	   }
	}
 }
`

func main() {
	codaClient := coda.NewClient()
	ds, err := codaClient.GetDeamonStatus()
	if err != nil {
		log.Fatalln(err)
	}
	log.Println(ds.DaemonStatus.NumAccounts)
	log.Println(ds.DaemonStatus.BlockchainLength)
	log.Println(ds.DaemonStatus.Peers)
	log.Println(ds.DaemonStatus.SnarkWorker)
}
