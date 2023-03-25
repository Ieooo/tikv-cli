package cmd

import (
	"fmt"

	"github.com/ieooo/tikv-cli/pkg/client"
	"github.com/spf13/cobra"
)

var GetCommand = &cobra.Command{
	Use:              "get",
	Short:            "get value by key from tikv",
	SilenceUsage:     true,
	Run:              get,
	PersistentPreRun: initConfig,
}

func get(cmd *cobra.Command, args []string) {
	if len(args) < 1 {
		errorExit("lack argument\n")
	}
	key := args[0]

	cli, err := client.NewTikvClient(cmd.Context(), conf)
	if err != nil {
		errorExit("client error:%v\n", err)
	}
	defer cli.Close()

	b, err := cli.Get(cmd.Context(), []byte(key))
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(b))
}
