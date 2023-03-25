package cmd

import (
	"errors"
	"fmt"
	"os"
	"strings"

	"github.com/ieooo/tikv-cli/pkg/config"
	"github.com/spf13/cobra"
)

var ConfigCommand = &cobra.Command{
	Use:          "config",
	Short:        "config for tikv",
	SilenceUsage: true,
}

var ConfigAddCommand = &cobra.Command{
	Use:          "add",
	Short:        "add config item",
	SilenceUsage: true,
	Run:          addConfig,
}

var ConfigRemoveCommand = &cobra.Command{
	Use:          "rm",
	Short:        "remove config item",
	SilenceUsage: true,
	Run:          removeConfig,
}

var ConfigUseCommand = &cobra.Command{
	Use:          "use",
	Short:        "select a tikv config",
	SilenceUsage: true,
	Run:          useConfig,
}

var ConfigListCommand = &cobra.Command{
	Use:          "list",
	Short:        "list all tikv config",
	SilenceUsage: true,
	Run:          listConfig,
}

func init() {
	ConfigCommand.AddCommand(ConfigAddCommand)
	ConfigCommand.AddCommand(ConfigRemoveCommand)
	ConfigCommand.AddCommand(ConfigUseCommand)
	ConfigCommand.AddCommand(ConfigListCommand)
}

func addConfig(cmd *cobra.Command, args []string) {
	if len(args) < 2 {
		errorExit("lack argument")
	}
	name := args[0]

	address := args[1]

	c := new(config.Config)
	if err := c.Load(); err != nil && !errors.Is(err, os.ErrNotExist) {
		errorExit("add config:%v\n", err)
	}

	for _, v := range c.Tikvs {
		if v.Name == name {
			errorExit("config %s is already exist\n", name)
		}
	}

	addresses := strings.Split(address, ",")
	c.Tikvs = append(c.Tikvs, config.TikvConfig{
		Name:     name,
		Address:  addresses,
		Security: config.Security{},
	})

	if err := c.Save(); err != nil {
		errorExit("add config:%v\n", err)
	}
}

func removeConfig(cmd *cobra.Command, args []string) {
	if len(args) < 1 {
		errorExit("lack argument")
	}
	name := args[0]

	c := new(config.Config)
	if err := c.Load(); err != nil {
		errorExit("load config:%v\n", err)
	}

	index := 0
	for ; index < len(c.Tikvs); index++ {
		if name == c.Tikvs[index].Name {
			break
		}
	}

	if index >= len(c.Tikvs) {
		errorExit("config named %s is not found\n", name)
	}

	if name == conf.Name {
		c.CurrentTikv = ""
	}

	c.Tikvs = append(c.Tikvs[:index], c.Tikvs[index+1:]...)

	if err := c.Save(); err != nil {
		errorExit("save config:%v\n", err)
	}
}

func useConfig(cmd *cobra.Command, args []string) {
	if len(args) < 1 {
		errorExit("lack argument")
	}
	name := args[0]

	c := new(config.Config)
	if err := c.Load(); err != nil {
		errorExit("load config:%v\n", err)
	}

	for _, v := range c.Tikvs {
		if name == v.Name {
			c.CurrentTikv = name
			conf = v
			break
		}
	}

	if err := c.Save(); err != nil {
		errorExit("save config:%v\n", err)
	}

}

func listConfig(cmd *cobra.Command, args []string) {
	c := new(config.Config)
	if err := c.Load(); err != nil {
		errorExit("load config:%v\n", err)
	}

	for _, v := range c.Tikvs {
		if c.CurrentTikv == v.Name {
			fmt.Print("> ")
		} else {
			fmt.Print("  ")
		}
		fmt.Printf("%s %v\n", v.Name, v.Address)
	}
}
