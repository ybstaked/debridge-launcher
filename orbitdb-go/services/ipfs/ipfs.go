package ipfs

import (
	"context"

	i "github.com/debridge-finance/orbitdb-go/pkg/ipfs"
	"github.com/debridge-finance/orbitdb-go/pkg/log"
)

type IPFS struct {
	Config Config

	log     log.Logger
	CoreAPI i.CoreAPI
}

func Create(ctx context.Context, c Config, l log.Logger) (*IPFS, error) {
	coreAPI, err := i.Create(ctx, l, c.Repo)
	if err != nil {
		return nil, err
	}

	l = l.With().Str("component", "IPFS").Logger()

	return &IPFS{
		Config:  c,
		log:     l,
		CoreAPI: coreAPI,
	}, nil
}
