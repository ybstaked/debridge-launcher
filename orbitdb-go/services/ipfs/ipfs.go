package ipfs

import (
	"context"

	i "github.com/debridge-finance/orbitdb-go/pkg/ipfs"
	"github.com/debridge-finance/orbitdb-go/pkg/log"
)

type IPFSService struct {
	Config Config

	log     log.Logger
	CoreAPI i.CoreAPI
}

func Create(c Config, ctx context.Context, l log.Logger) (*IPFSService, error) {
	coreAPI, err := i.Create(ctx, l, c.Repo)
	if err != nil {
		return nil, err
	}

	l = l.With().Str("component", "IPFSService").Logger()

	return &IPFSService{
		Config:  c,
		log:     l,
		CoreAPI: coreAPI,
	}, nil
}
