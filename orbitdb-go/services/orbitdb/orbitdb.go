package orbitdb

import (
	"github.com/debridge-finance/orbitdb-go/pkg/context"
	i "github.com/debridge-finance/orbitdb-go/pkg/ipfs"
	"github.com/debridge-finance/orbitdb-go/pkg/log"
	o "github.com/debridge-finance/orbitdb-go/pkg/orbitdb"
)

type OrbitDB struct {
	Config Config

	log     log.Logger
	OrbitDB o.OrbitDB
}

func Create(ctx context.Context, c Config, l log.Logger, ipfs i.CoreAPI) (*OrbitDB, error) {
	orbit, err := o.Create(ctx, ipfs, c.Repo)
	if err != nil {
		return nil, err
	}

	l = l.With().Str("component", "orbitdbService").Logger()

	return &OrbitDB{
		Config:  c,
		log:     l,
		OrbitDB: orbit,
	}, nil
}
