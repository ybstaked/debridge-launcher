package services

import (
	"github.com/debridge-finance/orbitdb-go/pkg/context"
	"github.com/debridge-finance/orbitdb-go/pkg/errors"
	"github.com/debridge-finance/orbitdb-go/pkg/log"
	si "github.com/debridge-finance/orbitdb-go/services/ipfs"
	//  "github.com/debridge-finance/orbitdb-go/services/ipfs"
)

type Services struct {
	Config Config
	Ctx    context.Context

	IPFS *si.IPFS
	// OrbitDB *so.OrbitDB
}

func Create(c Config, l log.Logger, ctx context.Context) (*Services, error) {
	ipfs, err := si.Create(ctx, *c.IPFS, l)
	if err != nil {
		return nil, errors.Wrap(err, "failed to create IPFS service")
	}
	l.Log().Msg("ipfs was created")

	// orbitdb, err := so.Create(ctx, *c.OrbitDB, l, ipfs.CoreAPI)
	// if err != nil {
	// 	return nil, errors.Wrap(err, "failed to create orbitdb services")
	// }

	return &Services{
		Config: c,
		Ctx:    ctx,

		IPFS: ipfs,
		// OrbitDB: orbitdb,
	}, nil
}
