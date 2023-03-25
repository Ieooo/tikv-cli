package cmd

import (
	"github.com/ieooo/tikv-cli/pkg/client"
	"github.com/spf13/cobra"
)

var DeleteCommand = &cobra.Command{
	Use:              "delete",
	Short:            "delete key value",
	SilenceUsage:     true,
	Run:              delete,
	PersistentPreRun: initConfig,
}

func delete(cmd *cobra.Command, args []string) {
	if len(args) < 1 {
		errorExit("lack argument\n")
	}
	key := args[0]

	cli, err := client.NewTikvClient(cmd.Context(), conf)
	if err != nil {
		errorExit("client error:%v\n", err)
	}
	cli.Close()

	if err := cli.Delete(cmd.Context(), []byte(key)); err != nil {
		errorExit("delete error:%v\n", err)
	}
}
