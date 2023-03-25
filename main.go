package main

import (
	"fmt"
	"os"
	"time"

	cmd "github.com/ieooo/tikv-cli/cmd/tikvcli"
	plog "github.com/pingcap/log"
	"github.com/spf13/cobra"
	"golang.org/x/net/context"
)

var (
	version string = "latest"
	commit  string = "HEAD"
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
	rootCmd.AddCommand(cmd.PutTTLCommand)
	rootCmd.AddCommand(cmd.GetTTLCommand)

	ignoreTikvLog()
}

func main() {
	if version == "" {
		rootCmd.Version = commit
	}
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()
	rootCmd.SetContext(ctx)
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		cancel()
		os.Exit(1)
	}
}

// ignore tikv log
func ignoreTikvLog() {
	conf := &plog.Config{Level: "error", File: plog.FileLogConfig{Filename: "/dev/null"}}
	log, p, _ := plog.InitLogger(conf)
	plog.ReplaceGlobals(log, p)
}
