package cmd

import (
	"github.com/elliot14A/tcorp/assessment/server"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(serveCmd)
}

var serveCmd = &cobra.Command{
	Use:   "server",
	Short: "Starts the ruspie server",
	Run: func(cmd *cobra.Command, args []string) {
		serve()
	},
}

func serve() {
	server.StartServer()
}
