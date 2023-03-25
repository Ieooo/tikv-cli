package cmd

import (
	"github.com/ieooo/tikv-cli/pkg/client"
	"github.com/spf13/cobra"
)

var (
	rang bool
)

var DeleteCommand = &cobra.Command{
	Use:              "delete",
	Short:            "delete key value",
	SilenceUsage:     true,
	Run:              delete,
	PersistentPreRun: initConfig,
}

func init() {
	DeleteCommand.Flags().BoolVarP(&rang, "range", "r", false, "delete range")
}

func delete(cmd *cobra.Command, args []string) {
	if len(args) < 1 || rang && len(args) < 2 {
		errorExit("lack argument\n")
	}

	if rang {
		cli, err := client.NewTikvClient(cmd.Context(), conf)
		if err != nil {
			errorExit("client error:%v\n", err)
		}
		if err := cli.DeleteRange(cmd.Context(), []byte(args[0]), []byte(args[1])); err != nil {
			errorExit("delete error:%v\n", err)
		}
		return
	}

	keys := make([][]byte, len(args))
	for i := range args {
		keys[i] = []byte(args[i])
	}

	cli, err := client.NewTikvClient(cmd.Context(), conf)
	if err != nil {
		errorExit("client error:%v\n", err)
	}
	cli.Close()

	if err := cli.BatchDelete(cmd.Context(), keys); err != nil {
		errorExit("delete error:%v\n", err)
	}
}
