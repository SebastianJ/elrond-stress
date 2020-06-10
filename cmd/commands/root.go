package commands

import (
	"fmt"
	"os"

	cmdConfig "github.com/SebastianJ/elrond-stress/config/cmd"
	"github.com/pkg/errors"
	"github.com/spf13/cobra"
)

var (
	// VersionWrap - version displayed in case of errors
	VersionWrap = ""

	// RootCmd - main entry point for Cobra commands
	RootCmd = &cobra.Command{
		Use:          "stress",
		Short:        "Elrond Stress",
		Long:         "Elrond Stress",
		SilenceUsage: true,
		RunE: func(cmd *cobra.Command, args []string) error {
			cmd.Help()
			return nil
		},
	}
)

func init() {
	cmdConfig.Persistent = cmdConfig.PersistentFlags{}
	RootCmd.PersistentFlags().StringVar(&cmdConfig.Persistent.Path, "path", ".", "Base path for resolving files etc")
	RootCmd.PersistentFlags().StringSliceVar(&cmdConfig.Persistent.Endpoints, "endpoint", []string{"https://wallet-api.elrond.com"}, "Which API endpoint to use for API commands")
}

// Execute starts the actual app
func Execute() {
	RootCmd.SilenceErrors = true
	if err := RootCmd.Execute(); err != nil {
		fmt.Println(errors.Wrapf(err, "commit: %s, error", VersionWrap).Error())
		os.Exit(1)
	}
}
