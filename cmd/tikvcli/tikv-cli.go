package main

import (
	"fmt"
	"os"

	"github.com/ieooo/tikv-cli/pkg/config"
	plog "github.com/pingcap/log"
	"github.com/spf13/cobra"
)

var (
	version string = "latest"
	conf    config.TikvConfig
)

var rootCmd = &cobra.Command{
	Use:          "tikv-cli",
	Short:        "tikv-cli is a client for tikv",
	Version:      version,
	SilenceUsage: true,
}

func init() {
	rootCmd.AddCommand(ConfigCommand)
	rootCmd.AddCommand(GetCommand)
	rootCmd.AddCommand(PutCommand)
	rootCmd.AddCommand(ScanCommand)
	rootCmd.AddCommand(DeleteCommand)

	ignoreTikvLog()
}

func main() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func initConfig(cmd *cobra.Command, args []string) {
	c := new(config.Config)
	if err := c.Load(); err != nil {
		errorExit("load config err:%v\n", err)
	}

	if c.CurrentTikv == "" {
		errorExit("please select one config first\n")
	}
	for _, v := range c.Tikvs {
		if v.Name == c.CurrentTikv {
			conf = v
		}
	}
}

func errorExit(format string, a ...any) {
	fmt.Printf(format, a...)
	os.Exit(1)
}

// ignore tikv log
func ignoreTikvLog() {
	conf := &plog.Config{Level: "error", File: plog.FileLogConfig{Filename: "/dev/null"}}
	log, p, _ := plog.InitLogger(conf)
	plog.ReplaceGlobals(log, p)
}
