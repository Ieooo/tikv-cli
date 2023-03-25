package main

import (
	"fmt"
	"os"

	cmd "github.com/ieooo/tikv-cli/cmd/tikvcli"
	plog "github.com/pingcap/log"
	"github.com/spf13/cobra"
)

var (
	version string = "latest"
)

var rootCmd = &cobra.Command{
	Use:          "tikv-cli",
	Short:        "tikv-cli is a client for tikv",
	Version:      version,
	SilenceUsage: true,
}

func init() {
	rootCmd.AddCommand(cmd.ConfigCommand)
	rootCmd.AddCommand(cmd.GetCommand)
	rootCmd.AddCommand(cmd.PutCommand)
	rootCmd.AddCommand(cmd.ScanCommand)
	rootCmd.AddCommand(cmd.DeleteCommand)

	ignoreTikvLog()
}

func main() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

// ignore tikv log
func ignoreTikvLog() {
	conf := &plog.Config{Level: "error", File: plog.FileLogConfig{Filename: "/dev/null"}}
	log, p, _ := plog.InitLogger(conf)
	plog.ReplaceGlobals(log, p)
}
