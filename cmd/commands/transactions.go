package commands

import (
	"path/filepath"

	"github.com/SebastianJ/elrond-stress/config"
	cmdConfig "github.com/SebastianJ/elrond-stress/config/cmd"
	"github.com/SebastianJ/elrond-stress/transactions"
	"github.com/spf13/cobra"
)

func init() {
	cmdTx := &cobra.Command{
		Use:   "txs",
		Short: "Stress transactions",
		Long:  "Stress transactions",
		RunE: func(cmd *cobra.Command, args []string) error {
			return sendTransactions(cmd)
		},
	}

	cmdConfig.Tx = cmdConfig.TxFlags{}
	cmdTx.Flags().StringVar(&cmdConfig.Tx.WalletsPath, "wallets", "./keys", "Path to wallet PEM files")
	cmdTx.Flags().StringVar(&cmdConfig.Tx.ReceiversPath, "receivers", "./data/receivers.txt", "Path to receivers file")
	cmdTx.Flags().Float64Var(&cmdConfig.Tx.Amount, "amount", 0.0, "How many tokens to send")
	cmdTx.Flags().Int64Var(&cmdConfig.Tx.Nonce, "nonce", -1, "What nonce to use for sending the transaction")
	cmdTx.Flags().StringVar(&cmdConfig.Tx.DataPath, "data", "./data/data.txt", "File containing data to use for sending transactions")
	cmdTx.Flags().StringVar(&cmdConfig.Tx.ConfigPath, "config", "./configs/economics.toml", "The economics configuration file to load")
	cmdTx.Flags().BoolVar(&cmdConfig.Tx.ForceAPINonceLookups, "force-api-nonce-lookups", true, "Force the usage of https://wallet-api.elrond.com for checking nonces when using local node endpoints")

	RootCmd.AddCommand(cmdTx)
}

func sendTransactions(cmd *cobra.Command) error {
	basePath, err := filepath.Abs(cmdConfig.Persistent.Path)
	if err != nil {
		return err
	}

	if err := config.Configure(basePath); err != nil {
		return err
	}

	if err := transactions.SendTransactions(); err != nil {
		return err
	}

	return nil
}
