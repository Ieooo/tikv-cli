package cmd

import (
	"context"
	"fmt"

	"github.com/spf13/cobra"
	"github.com/tikv/client-go/v2/config"
	"github.com/tikv/client-go/v2/rawkv"
)

var GetCommand = &cobra.Command{
	Use:   "get",
	Short: "get value by key from tikv",
	Run:   get,
}

func get(cmd *cobra.Command, args []string) {
	tikvCli, err := rawkv.NewClient(context.Background(), []string{"127.0.0.1:2379"}, config.Security{})
	if err != nil {
		fmt.Printf("client error:%v\n", err)
	}
	b, err := tikvCli.Get(context.TODO(), []byte("key"))
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(b))
}
