package config

import (
	sdkAPI "github.com/SebastianJ/elrond-sdk/api"
	sdkTxs "github.com/SebastianJ/elrond-sdk/transactions"
	sdkWallet "github.com/SebastianJ/elrond-sdk/wallet"
)

var Configuration Config

// Config - general config
type Config struct {
	BasePath     string
	Endpoints    []string
	Concurrency  int
	Verbose      bool
	Transactions TxConfig
}

// TxConfig - tx configuration
type TxConfig struct {
	Wallets              []sdkWallet.Wallet
	Receivers            []string
	Client               sdkAPI.Client
	Amount               float64
	Nonce                int64
	GasParams            sdkTxs.GasParams
	Data                 string
	ForceAPINonceLookups bool
}
