package cmd

import (
	"fmt"

	"github.com/ieooo/tikv-cli/pkg/client"
	"github.com/spf13/cobra"
)

var GetTTLCommand = &cobra.Command{
	Use:              "gett",
	Short:            "get key ttl",
	SilenceUsage:     true,
	Run:              getttl,
	PersistentPreRun: initConfig,
}

func getttl(cmd *cobra.Command, args []string) {
	if len(args) < 1 {
		errorExit("lack argument\n")
	}

	key := args[0]

	cli, err := client.NewTikvClient(cmd.Context(), conf)
	if err != nil {
		errorExit("client error:%v\n", err)
	}
	defer cli.Close()

	ttl, err := cli.GetKeyTTL(cmd.Context(), []byte(key))
	if err != nil {
		errorExit("get error:%v\n", err)
	}

	fmt.Printf("%s:%d\n", key, ttl)
}
