package cmd

import (
	"fmt"

	"github.com/ieooo/tikv-cli/pkg/client"
	"github.com/spf13/cobra"
	"github.com/tikv/client-go/v2/rawkv"
)

var (
	prefix bool
)

var GetCommand = &cobra.Command{
	Use:              "get",
	Short:            "get value by key from tikv",
	SilenceUsage:     true,
	Run:              get,
	PersistentPreRun: initConfig,
}

func init() {
	GetCommand.Flags().BoolVarP(&prefix, "prefix", "p", false, "get by prefix")
}

func get(cmd *cobra.Command, args []string) {
	if len(args) < 1 {
		errorExit("lack argument\n")
	}

	cli, err := client.NewTikvClient(cmd.Context(), conf)
	if err != nil {
		errorExit("client error:%v\n", err)
	}
	defer cli.Close()

	if prefix {
		startKey := args[0]
		endKey := startKey[:len(startKey)-1] + string(startKey[len(startKey)-1]+1)
		keys, values, err := cli.Scan(cmd.Context(), []byte(startKey), []byte(endKey), rawkv.MaxRawKVScanLimit)
		if err != nil {
			errorExit("get error:%v\n", err)
		}

		for i := range keys {
			fmt.Printf("%s:%s\n", keys[i], values[i])
		}
		return
	}

	keys := make([][]byte, len(args))
	for i, v := range args {
		keys[i] = []byte(v)
	}

	values, err := cli.BatchGet(cmd.Context(), keys)
	if err != nil {
		errorExit("get error:%v\n", err)
	}

	for i := range keys {
		fmt.Printf("%s:%s\n", keys[i], values[i])
	}
}
