package chainparams

// For now we will just use this to store the
// peers, to test peer data
// Once complete, it will store the genesis params for testnet and mainnet

var mainnetSeedList = []string{
	"seed1.neo.org:10333",  // NOP
	"seed2.neo.org:10333",  // YEP
	"seed3.neo.org:10333",  // YEP
	"seed4.neo.org:10333",  // NOP
	"seed5.neo.org:10333",  // YEP
	"13.59.52.94:10333",    // NOP
	"18.220.214.143:10333", // NOP
	"13.58.198.112:10333",  // NOP
	"13.59.14.206:10333",   // NOP
	"18.216.9.7:10333",     // NOP
}

//MainnetSeedList is a string slice containing the initial seeds from protocol.mainnet
// That are replying
var MainnetSeedList = []string{
	"seed2.neo.org:10333",
	"seed3.neo.org:10333",
	"seed5.neo.org:10333",
	"http://seed1.ngd.network:10333",
	"http://seed2.ngd.network:10333",
	"http://seed3.ngd.network:10333",
	"http://seed4.ngd.network:10333",
	"http://seed5.ngd.network:10333",
	"http://seed6.ngd.network:10333",
	"http://seed7.ngd.network:10333",
	"http://seed8.ngd.network:10333",
	"http://seed9.ngd.network:10333",
	"http://seed10.ngd.network:10333",
}

var testNetSeedList = []string{

	"18.218.97.227:20333",
	"18.219.30.120:20333",
	"18.219.13.91:20333",
	"13.59.116.121:20333",
	"18.218.255.178:20333",
}

// GenesisHash for mainnet
var GenesisHash = "d42561e3d30e15be6400b6df2f328e02d2bf6354c41dce433bc57687c82144bf"

// GenesisBlock for mainnet
var GenesisBlock = "000000000000000000000000000000000000000000000000000000000000000000000000f41bc036e39b0d6b0579c851c6fde83af802fa4e57bec0bc3365eae3abf43f8065fc8857000000001dac2b7c0000000059e75d652b5d3827bf04c165bbe9ef95cca4bf55010001510400001dac2b7c00000000400000455b7b226c616e67223a227a682d434e222c226e616d65223a22e5b08fe89a81e882a1227d2c7b226c616e67223a22656e222c226e616d65223a22416e745368617265227d5d0000c16ff28623000000da1745e9b549bd0bfa1a569971c77eba30cd5a4b00000000400001445b7b226c616e67223a227a682d434e222c226e616d65223a22e5b08fe89a81e5b881227d2c7b226c616e67223a22656e222c226e616d65223a22416e74436f696e227d5d0000c16ff286230008009f7fd096d37ed2c0e3f7f0cfc924beef4ffceb680000000001000000019b7cffdaa674beae0f930ebe6085af9093e5fe56b34a5c220ccdcf6efc336fc50000c16ff28623005fa99d93303775fe50ca119c327759313eccfa1c01000151"
