package cmd

import (
	"fmt"
	"os"

	"github.com/ieooo/tikv-cli/pkg/config"
	"github.com/spf13/cobra"
)

var (
	conf config.TikvConfig
)

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
