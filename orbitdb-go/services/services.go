package services

import (
	"github.com/debridge-finance/orbitdb-go/pkg/errors"
	"github.com/debridge-finance/orbitdb-go/pkg/log"
	"github.com/debridge-finance/orbitdb-go/services/ipfs"
	"github.com/debridge-finance/orbitdb-go/services/orbitdb"
)

type Services struct {
	Config Config

	IPFS *ipfs.IPFS
	ODB  *orbitdb.Orbitdb
}

func Create(c Config, l log.Logger) (*Services, error) {
	ipfs, err := ipfs.Create(*c.IPFS, l)
	if err != nil {
		return nil, errors.Wrap(err, "failed to create IPFS service")
	}

	orbitdb, err := orbitdb.Create(*c.Orbitdb, l)
	if err != nil {
		return nil, errors.Wrap(err, "failed to create orbitdb services")
	}

	return &Services{
		Config: c,

		IPFS: ipfs,
		ODB:  orbitdb,
	}, nil
}
