package main

import (
	"context"
	"fmt"

	"github.com/ieooo/tikv-cli/pkg/client"
	"github.com/spf13/cobra"
)

var PutCommand = &cobra.Command{
	Use:          "put",
	Short:        "put key value pair to tikv",
	SilenceUsage: true,
	Run:          put,
}

func put(cmd *cobra.Command, args []string) {
	key := "key"
	value := "value"
	cli, err := client.NewTikvClient(cmd.Context(), conf)
	if err != nil {
		fmt.Printf("client error:%v\n", err)
	}
	if err := cli.Put(context.TODO(), []byte(key), []byte(value)); err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(key), string(value))
}
