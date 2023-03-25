package cmd

import (
	"fmt"

	"github.com/ieooo/tikv-cli/pkg/client"
	"github.com/spf13/cobra"
)

var PutCommand = &cobra.Command{
	Use:              "put",
	Short:            "put key value pair to tikv",
	SilenceUsage:     true,
	Run:              put,
	PersistentPreRun: initConfig,
}

func put(cmd *cobra.Command, args []string) {
	if len(args) < 2 || len(args)%2 != 0 {
		errorExit("lack arguments\n")
	}

	keys := make([][]byte, (len(args) / 2))
	values := make([][]byte, (len(args) / 2))
	for i := range args {
		if i%2 == 0 {
			keys[i/2] = []byte(args[i])
		} else {
			values[i/2] = []byte(args[i])
		}
	}

	cli, err := client.NewTikvClient(cmd.Context(), conf)
	if err != nil {
		errorExit("client error:%v\n", err)
	}
	defer cli.Close()

	if err := cli.BatchPut(cmd.Context(), keys, values); err != nil {
		fmt.Println(err)
	}

	for i := range keys {
		fmt.Printf("%s:%s\n", keys[i], values[i])
	}
}
