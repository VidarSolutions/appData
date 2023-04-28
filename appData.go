package appData

import(
	"sync"
	"github.com/vidarsolutions/Ring"
	"github.com/vidarsolutions/Transfer"


)


type Bytes32 [32]byte
type Bytes64 [64]byte

var (
    	
	//Enable / Disable Logging
	Logging 	  	bool
	
	//Setup Tor and I2p network dialers to communicate with other nodes
	Dial_Tor		Transfer.TransferClient
	Dial_I2p		Transfer.TransferClient
	TorProxy		string
	I2pProxy		string
	
	// Get the Bootstrap Rings map
	appRings = 	Ring.Rings.LoadRings()
	
	//Set up the working directories for the App
	Wd			string
	TorDir =	"tor"
	I2pDir =	"i2p"
	TxDir  =	"tx"
	ItemsDir =	"items"
	
		
	//setup Mutex to lock global variable to ensure app stability
	Mu sync.Mutex
)

const(
	TorDefault = "127.0.0.1:9050"
	I2pDefault = "127.0.0.1:4444"
	Usage = "sets the proxy ports for accessing dark net networks"
	WdUsage = "sets the working directory for the application"

)