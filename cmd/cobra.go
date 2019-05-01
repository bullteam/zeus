package cmd

import (
	api_server "github.com/bullteam/zeus/cmd/api-server"
	"github.com/spf13/cobra"
	"os"
)

var rootCmd = &cobra.Command{
	Use:               "zeus",
	Short:             "zeus API server",
	SilenceUsage:      true,
	DisableAutoGenTag: true,
	Long:              `Start zeus API server`,
	PersistentPreRunE: func(*cobra.Command, []string) error { return nil },
}

func init() {
	rootCmd.AddCommand(api_server.StartCmd)
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		os.Exit(-1)
	}
}
