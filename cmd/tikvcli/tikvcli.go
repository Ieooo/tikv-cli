package main

import (
	"fmt"
	"os"

	"github.com/ieooo/tikv-cli/pkg/config"
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

}

func main() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func initConfig() error {
	c := new(config.Config)
	if err := c.Load(); err != nil {
		return err
	}

	if c.CurrentTikv == "" {
		errorExit("please select one config first\n")
	}
	for _, v := range c.Tikvs {
		if v.Name == c.CurrentTikv {
			conf = v
		}
	}
	return nil
}

func errorExit(format string, a ...any) {
	fmt.Printf(format, a...)
	os.Exit(1)
}
