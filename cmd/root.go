package cmd

import (
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "gameserver",
	Short: "Simple engine server",
	Long:  "A simple WebSocket engine server",
}

func Execute() error {
	return rootCmd.Execute()
}
