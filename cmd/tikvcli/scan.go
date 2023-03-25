package cmd

import (
	"fmt"

	"github.com/ieooo/tikv-cli/pkg/client"
	"github.com/spf13/cobra"
	"github.com/tikv/client-go/v2/rawkv"
)

var (
	startKey string
	endKey   string
	limit    uint32
	reverse  bool
)

var ScanCommand = &cobra.Command{
	Use:              "scan",
	Short:            "scan keys and values from tikv",
	SilenceUsage:     true,
	Run:              scan,
	PersistentPreRun: initConfig,
}

func init() {
	ScanCommand.Flags().BoolVarP(&reverse, "reverse", "r", false, "reverse scan")
}

func init() {
	ScanCommand.Flags().StringVarP(&startKey, "start", "s", "", "set startkey for scan")
	ScanCommand.Flags().StringVarP(&endKey, "end", "e", "", "set startkey for scan")
	ScanCommand.Flags().Uint32VarP(&limit, "limit", "l", uint32(rawkv.MaxRawKVScanLimit), "limit for scan result")
}

func scan(cmd *cobra.Command, args []string) {
	cli, err := client.NewTikvClient(cmd.Context(), conf)
	if err != nil {
		errorExit("client error:%v\n", err)
	}
	defer cli.Close()

	if limit > uint32(rawkv.MaxRawKVScanLimit) {
		errorExit("limit count is too large\n")
	}

	var keys, values [][]byte
	if reverse {
		keys, values, err = cli.ReverseScan(cmd.Context(), []byte(startKey), []byte(endKey), int(limit))
		if err != nil {
			errorExit("scan error:%v\n", err)
		}
	} else {
		keys, values, err = cli.Scan(cmd.Context(), []byte(startKey), []byte(endKey), int(limit))
		if err != nil {
			errorExit("scan error:%v\n", err)
		}
	}
	for i := range keys {
		fmt.Println(string(keys[i]), string(values[i]))
	}
}
