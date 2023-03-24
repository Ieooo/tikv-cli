package cmd

import "github.com/spf13/cobra"

var ScanCommand = &cobra.Command{
	Use:   "scan",
	Short: "scan keys and values from tikv",
	Run:   scan,
}

func scan(cmd *cobra.Command, args []string) {

}
