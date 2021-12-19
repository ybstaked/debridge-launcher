package services

import (
	"github.com/debridge-finance/orbitdb-go/pkg/errors"
	"github.com/debridge-finance/orbitdb-go/pkg/services/eventlog"
	"github.com/debridge-finance/orbitdb-go/pkg/services/ipfs"
	"github.com/debridge-finance/orbitdb-go/pkg/services/orbitdb"
)

var DefaultConfig = Config{
	IPFS:     &ipfs.DefaultConfig,
	OrbitDB:  &orbitdb.DefaultConfig,
	Eventlog: &eventlog.DefaultConfig,
}

//

type Config struct {
	IPFS     *ipfs.Config
	OrbitDB  *orbitdb.Config
	Eventlog *eventlog.Config
}

func (c *Config) SetDefaults() {
loop:
	for {
		switch {
		case c.IPFS == nil:
			c.IPFS = DefaultConfig.IPFS
		case c.OrbitDB == nil:
			c.OrbitDB = DefaultConfig.OrbitDB
		case c.Eventlog == nil:
			c.Eventlog = DefaultConfig.Eventlog
		default:
			break loop
		}
	}

	c.IPFS.SetDefaults()
	c.OrbitDB.SetDefaults()
	c.Eventlog.SetDefaults()
}

func (c Config) Validate() error {
	err := c.IPFS.Validate()
	if err != nil {
		return errors.Wrap(err, "failed to validate ipfs configuration")
	}
	err = c.OrbitDB.Validate()
	if err != nil {
		return errors.Wrap(err, "failed to validate orbitdb configuration")
	}
	err = c.Eventlog.Validate()
	if err != nil {
		return errors.Wrap(err, "failed to validate eventlog configuration")
	}
	return nil
}
