package config

import (
	"path/filepath"

	sdkTxs "github.com/SebastianJ/elrond-sdk/transactions"
	sdkUtils "github.com/SebastianJ/elrond-sdk/utils"
	sdkWallet "github.com/SebastianJ/elrond-sdk/wallet"
	cmd "github.com/SebastianJ/elrond-stress/config/cmd"
	"github.com/SebastianJ/elrond-stress/utils"
)

func Configure(basePath string) (err error) {
	Configuration.BasePath = basePath
	Configuration.Endpoints = cmd.Persistent.Endpoints
	Configuration.Concurrency = cmd.Persistent.Concurrency

	if err := configureReceivers(); err != nil {
		return err
	}

	if err := configureData(); err != nil {
		return err
	}

	if err := configureGasParams(); err != nil {
		return err
	}

	if err := configureWallets(); err != nil {
		return err
	}

	Configuration.Transactions.Amount = cmd.Tx.Amount
	Configuration.Transactions.Nonce = cmd.Tx.Nonce
	Configuration.Transactions.ForceAPINonceLookups = cmd.Tx.ForceAPINonceLookups

	return nil
}

func configureReceivers() error {
	receiversPath := filepath.Join(Configuration.BasePath, cmd.Tx.ReceiversPath)

	receivers, err := utils.ReadFileToSlice(receiversPath)
	if err != nil {
		return err
	}

	if len(receivers) == 0 {
		receivers = []string{"erd1mp543xj384uzehwzp360wy2y86q22svdm022lwxaryg8cqmxqwvszjnrf7"}
	}

	Configuration.Transactions.Receivers = receivers

	return nil
}

func configureData() error {
	dataPath := filepath.Join(Configuration.BasePath, cmd.Tx.DataPath)

	data, err := utils.ReadFileToString(dataPath)
	if err != nil {
		return err
	}

	Configuration.Transactions.Data = data

	return nil
}

func configureGasParams() error {
	defaultGasParams, err := sdkTxs.ParseGasSettings(cmd.Tx.ConfigPath)
	if err != nil {
		return err
	}

	Configuration.Transactions.GasParams = defaultGasParams

	if cmd.Tx.GasLimit != -1 {
		Configuration.Transactions.GasParams.GasLimit = uint64(cmd.Tx.GasLimit)
	}

	return nil
}

func configureWallets() error {
	walletsPath := filepath.Join(Configuration.BasePath, cmd.Tx.WalletsPath)

	pemFiles, err := sdkUtils.IdentifyPemFiles(walletsPath)
	if err != nil {
		return err
	}

	var wallets []sdkWallet.Wallet

	for _, pemFile := range pemFiles {
		wallet, err := sdkWallet.Decrypt(pemFile)
		if err != nil {
			return err
		}

		wallets = append(wallets, wallet)
	}

	Configuration.Transactions.Wallets = wallets

	return nil
}
