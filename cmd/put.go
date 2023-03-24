package cmd

import "github.com/spf13/cobra"

var PutCommand = &cobra.Command{
	Use:   "put",
	Short: "put key value pair to tikv",
	Run:   put,
}

func put(cmd *cobra.Command, args []string) {

}
