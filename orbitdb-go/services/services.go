package services

import (
	"github.com/debridge-finance/orbitdb-go/pkg/context"
	"github.com/debridge-finance/orbitdb-go/pkg/errors"
	"github.com/debridge-finance/orbitdb-go/pkg/log"
	si "github.com/debridge-finance/orbitdb-go/services/ipfs"
	so "github.com/debridge-finance/orbitdb-go/services/orbitdb"
)

type Services struct {
	Config Config

	IPFS    *si.IPFS
	OrbitDB *so.OrbitDB
	Ctx     context.Context
}

func Create(c Config, l log.Logger, ctx context.Context) (*Services, error) {
	ipfs, err := si.Create(ctx, *c.IPFS, l)
	if err != nil {
		return nil, errors.Wrap(err, "failed to create IPFS service")
	}

	orbitdb, err := so.Create(ctx, *c.OrbitDB, l, ipfs.CoreAPI)
	if err != nil {
		return nil, errors.Wrap(err, "failed to create orbitdb services")
	}

	return &Services{
		Config: c,

		IPFS:    ipfs,
		OrbitDB: orbitdb,
		Ctx:     ctx,
	}, nil
}
