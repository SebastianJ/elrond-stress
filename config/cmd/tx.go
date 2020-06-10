package cmd

// Tx is a collection of tx related options
var Tx TxFlags

// TxFlags represents the tx flags
type TxFlags struct {
	WalletsPath          string
	ReceiversPath        string
	Amount               float64
	Nonce                int64
	DataPath             string
	ConfigPath           string
	ForceAPINonceLookups bool
}
