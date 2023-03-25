package main

import (
	"fmt"

	"github.com/ieooo/tikv-cli/pkg/client"
	"github.com/spf13/cobra"
)

var GetCommand = &cobra.Command{
	Use:          "get",
	Short:        "get value by key from tikv",
	SilenceUsage: true,
	Run:          get,
}

func get(cmd *cobra.Command, args []string) {
	initConfig()
	cli, err := client.NewTikvClient(cmd.Context(), conf)
	if err != nil {
		fmt.Printf("client error:%v\n", err)
	}
	b, err := cli.Get(cmd.Context(), []byte("key"))
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(b))
}
