package cmd

import (
	"fmt"
	"strconv"

	"github.com/ieooo/tikv-cli/pkg/client"
	"github.com/spf13/cobra"
)

var PutTTLCommand = &cobra.Command{
	Use:              "putt",
	Short:            "put key with ttl",
	SilenceUsage:     true,
	Run:              putttl,
	PersistentPreRun: initConfig,
}

func putttl(cmd *cobra.Command, args []string) {
	if len(args) < 3 {
		errorExit("lack arguments\n")
	}

	key, value := args[0], args[1]

	ttl, err := strconv.Atoi(args[2])
	if err != nil {
		errorExit("invalid argument ttl\n")
	}

	cli, err := client.NewTikvClient(cmd.Context(), conf)
	if err != nil {
		errorExit("client error:%v\n", err)
	}
	defer cli.Close()

	if err := cli.PutWithTTL(cmd.Context(), []byte(key), []byte(value), uint64(ttl)); err != nil {
		fmt.Println(err)
	}

	fmt.Printf("%s:%s %d\n", key, value, ttl)
}
