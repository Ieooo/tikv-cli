package main

import (
	"context"
	"fmt"

	"github.com/ieooo/tikv-cli/pkg/client"
	"github.com/spf13/cobra"
)

var ScanCommand = &cobra.Command{
	Use:          "scan",
	Short:        "scan keys and values from tikv",
	SilenceUsage: true,
	Run:          scan,
}

func scan(cmd *cobra.Command, args []string) {
	startKey := ""
	endKey := ""
	limit := 10
	cli, err := client.NewTikvClient(cmd.Context(), conf)
	if err != nil {
		fmt.Printf("client error:%v\n", err)
	}
	keys, values, err := cli.Scan(context.TODO(), []byte(startKey), []byte(endKey), limit)
	if err != nil {
		fmt.Println(err)
	}
	for i := range keys {
		fmt.Println(string(keys[i]), string(values[i]))
	}
}
