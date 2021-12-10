package orbitdb

import (
	"context"

	"github.com/debridge-finance/orbitdb-go/pkg/log"
	o "github.com/debridge-finance/orbitdb-go/pkg/orbitdb"
)

type (
	OrbitDB = o.OrbitDB
)
type OrbitDBService struct {
	Config Config

	log     log.Logger
	OrbitDB *OrbitDB
}

func Create(c Config, ctx context.Context, ipfs *ipfs.IPFS, l log.Logger) (*OrbitDBService, error) {
	o, err := o.Create(ctx, ipfs, c.Repo)
	if err != nil {
		return nil, err
	}

	l = l.With().Str("component", "orbitdbService").Logger()

	return &OrbitDBService{
		Config:  c,
		log:     l,
		OrbitDB: o,
	}, nil
}
