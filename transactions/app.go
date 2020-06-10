package transactions

import (
	"fmt"
	"sync"

	"github.com/SebastianJ/elrond-sdk/api"
	sdkAPI "github.com/SebastianJ/elrond-sdk/api"
	sdkTxs "github.com/SebastianJ/elrond-sdk/transactions"
	sdkWallet "github.com/SebastianJ/elrond-sdk/wallet"
	"github.com/SebastianJ/elrond-stress/config"
	"github.com/SebastianJ/elrond-stress/utils"
)

func SendTransactions() error {
	var walletWaitGroup sync.WaitGroup

	for _, wallet := range config.Configuration.Transactions.Wallets {
		walletWaitGroup.Add(1)
		go SendTransactionFor(wallet, &walletWaitGroup)
	}

	walletWaitGroup.Wait()

	return nil
}

func SendTransactionFor(wallet sdkWallet.Wallet, walletWaitGroup *sync.WaitGroup) error {
	defer walletWaitGroup.Done()
	var receiverWaitGroup sync.WaitGroup

	endpoint := utils.RandomElementFromSlice(config.Configuration.Endpoints)

	client := api.Client{
		Host:                 endpoint,
		ForceAPINonceLookups: true,
	}
	client.Initialize()

	account, err := client.GetAccount(wallet.Address)
	if err != nil {
		return err
	}
	nonce := int64(account.Nonce)

	// Make a copy of the default gas params that can be modified when processing the tx
	gasParams := config.Configuration.Transactions.GasParams

	for i := 0; i < config.Configuration.Concurrency; i++ {
		receiverWaitGroup.Add(1)
		receiver := utils.RandomElementFromSlice(config.Configuration.Transactions.Receivers)
		go SendTransactionToReceiver(wallet, receiver, nonce, gasParams, client, &receiverWaitGroup)
		nonce++
	}

	receiverWaitGroup.Wait()

	return nil
}

func SendTransactionToReceiver(wallet sdkWallet.Wallet, receiver string, nonce int64, gasParams sdkTxs.GasParams, client sdkAPI.Client, receiverWaitGroup *sync.WaitGroup) (string, error) {
	defer receiverWaitGroup.Done()

	fmt.Printf("Sending tx from: %s, to: %s, nonce: %d\n", wallet.Address, receiver, nonce)

	txHash, err := sdkTxs.SendTransaction(
		wallet,
		receiver,
		config.Configuration.Transactions.Amount,
		false,
		nonce,
		config.Configuration.Transactions.Data,
		gasParams,
		client,
	)
	if err != nil {
		return "", err
	}

	fmt.Println(fmt.Sprintf("Success! Your pending transaction hash is: %s", txHash))

	return txHash, nil
}
