package cmd

import (
	"fmt"
	"github.com/rasouliali1379/terrax/cmd/server"
	"github.com/spf13/cobra"
	"os"
)

var rootCmd = &cobra.Command{
	Use:   "terrax",
	Short: "Terrax is a server for tap to earn telegram mini app",
	Run:   server.Start,
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
