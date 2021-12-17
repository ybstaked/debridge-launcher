package orbitdb

import (
	"github.com/debridge-finance/orbitdb-go/pkg/berty.tech/go-orbit-db/iface"
	"github.com/debridge-finance/orbitdb-go/pkg/context"
	"github.com/debridge-finance/orbitdb-go/pkg/errors"
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

func (o *OrbitDB) Log(ctx context.Context, address string, options *iface.CreateDBOptions) (iface.EventLogStore, error) {
	if options == nil {
		options = &iface.CreateDBOptions{}
	}

	options.Create = boolPtr(true)
	options.StoreType = stringPtr("eventlog")
	store, err := o.OrbitDB.Open(ctx, address, options)
	if err != nil {
		return nil, errors.Wrap(err, "unable to open database")
	}

	logStore, ok := store.(iface.EventLogStore)
	if !ok {
		return nil, errors.New("unable to cast store to log")
	}

	return logStore, nil
}

func stringPtr(s string) *string {
	return &s
}

func boolPtr(b bool) *bool {
	return &b
}
