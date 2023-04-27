package appData

import(
	"github.com/vidarsolutions/edKeys"
	"github.com/vidarsolutions/Transfer"
	"github.com/vidarsolutions/Transaction"

)

var (
    	
	//Enable / Disable Logging
	Logging 	  	bool
	
	//Setup Tor and I2p network dialers to communicate with other nodes
	Dial_Tor		Transfer
	Dial_I2p		Transfer
	TorProxy		string
	I2pProxy		string
	
	// Get the Bootstrap Rings map
	appRings = 	Ring.Rings.loadRings()
	
	//Set up the working directories for the App
	Wd			string
	TorDir =	"tor"
	I2pDir =	"i2p"
	TxDir  =	"tx"
	ItemsDir =	"items"
	
	//setup map of transactions
	Transactions		map[uint64]Transaction			 //Known Transactions
	notValid := 		make(map[uint64][]transaction)   //Map of potential transactions ID to node IDs indicating a bad transaction was approved. We do not track valid transaction in a map that is track in each bundle, but this is a way to track and vote a bad transaction be reversed and removed.
	authorized := 		make(map[uint64][]uint64)  		 //Map of bundle IDs to a list of Authorized Node IDs

	
	//setup Mutex to lock global variable to ensure app stability
	Mu sync.Mutex
)

const(
	TorDefault = "127.0.0.1:9050"
	I2pDefault = "127.0.0.1:4444"
	Usage = "sets the proxy ports for accessing dark net networks"
	WdUsage = "sets the working directory for the application"

)