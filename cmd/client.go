package cmd

import (
	"github.com/elliot14A/tcorp/assessment/client"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(clientCmd)
}

var clientCmd = &cobra.Command{
	Use:   "client",
	Short: "Starts the ruspie server",
	Run: func(cmd *cobra.Command, args []string) {
		clientC()
	},
}

func clientC() {
	client.StartClient()
}
