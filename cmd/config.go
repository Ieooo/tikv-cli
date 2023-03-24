package cmd

import "github.com/spf13/cobra"

var ConfigCommand = &cobra.Command{
	Use:   "config",
	Short: "config for tikv",
	Run:   configTikv,
}

func configTikv(cmd *cobra.Command, args []string) {
}
