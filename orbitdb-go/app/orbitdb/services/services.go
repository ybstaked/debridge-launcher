package services

import (
	"github.com/debridge-finance/orbitdb-go/pkg/context"
	"github.com/debridge-finance/orbitdb-go/pkg/errors"
	"github.com/debridge-finance/orbitdb-go/pkg/log"
	se "github.com/debridge-finance/orbitdb-go/pkg/services/eventlog"
	si "github.com/debridge-finance/orbitdb-go/pkg/services/ipfs"
	so "github.com/debridge-finance/orbitdb-go/pkg/services/orbitdb"
)

type Services struct {
	Config Config
	Ctx    context.Context

	IPFS       *si.IPFS
	OrbitDB    *so.OrbitDB
	Submission *se.Eventlog
	Asset      *se.Eventlog
}

func Create(c Config, l log.Logger, ctx context.Context) (*Services, error) {
	ipfs, err := si.Create(ctx, *c.IPFS, l)
	if err != nil {
		return nil, errors.Wrap(err, "failed to create IPFS service")
	}
	l.Info().Msgf("ipfs service was created:%v", ipfs.PeerAddrs())

	orbitdb, err := so.Create(ctx, *c.OrbitDB, l, ipfs.CoreAPI)
	if err != nil {
		return nil, errors.Wrap(err, "failed to create orbitdb services")
	}
	submissions, err := se.Create(ctx, *c.Submission, l, orbitdb.OrbitDB)
	if err != nil {
		return nil, errors.Wrap(err, "failed to create orbitdb services")
	}
	assets, err := se.Create(ctx, *c.Asset, l, orbitdb.OrbitDB)
	if err != nil {
		return nil, errors.Wrap(err, "failed to create orbitdb services")
	}

	return &Services{
		Config: c,
		Ctx:    ctx,

		IPFS:       ipfs,
		OrbitDB:    orbitdb,
		Submission: submissions,
		Asset:      assets,
	}, nil
}
