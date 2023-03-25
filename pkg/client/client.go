package client

import (
	"context"

	cliconfig "github.com/ieooo/tikv-cli/pkg/config"
	"github.com/tikv/client-go/v2/config"
	"github.com/tikv/client-go/v2/rawkv"
)

func NewTikvClient(ctx context.Context, conf cliconfig.TikvConfig) (*rawkv.Client, error) {
	return rawkv.NewClient(ctx, conf.Address, config.Security(conf.Security))
}
