package main

import (
	"fmt"
	"os"
	"tikv-cli/cmd"

	"github.com/spf13/cobra"
)

func main() {
	rootCmd := &cobra.Command{
		Use:   "tikv-cli",
		Short: "tikv-cli is a client for tikv",
		Run: func(cmd *cobra.Command, args []string) {
			// Do Stuff Here
		},
	}

	rootCmd.AddCommand(cmd.ConfigCommand, cmd.GetCommand, cmd.GetCommand, cmd.ScanCommand)
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
