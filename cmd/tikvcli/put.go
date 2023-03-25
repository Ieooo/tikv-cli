package cmd

import (
	"context"
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
	if len(args) < 2 {
		errorExit("lack arguments\n")
	}
	key, value := args[0], args[1]

	cli, err := client.NewTikvClient(cmd.Context(), conf)
	if err != nil {
		errorExit("client error:%v\n", err)
	}
	defer cli.Close()

	if err := cli.Put(context.TODO(), []byte(key), []byte(value)); err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(key), string(value))
}
