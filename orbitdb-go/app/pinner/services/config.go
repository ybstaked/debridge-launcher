package services

import (
	"github.com/debridge-finance/orbitdb-go/pkg/errors"
	"github.com/debridge-finance/orbitdb-go/pkg/services/ipfs"
	"github.com/debridge-finance/orbitdb-go/pkg/services/orbitdb"
	"github.com/debridge-finance/orbitdb-go/pkg/services/pinner"
)

var DefaultConfig = Config{
	IPFS:    &ipfs.DefaultConfig,
	OrbitDB: &orbitdb.DefaultConfig,
	Pinner:  &pinner.DefaultConfig,
}

//

type Config struct {
	IPFS    *ipfs.Config
	OrbitDB *orbitdb.Config
	Pinner  *pinner.Config
}

func (c *Config) SetDefaults() {
loop:
	for {
		switch {
		case c.IPFS == nil:
			c.IPFS = DefaultConfig.IPFS
		case c.OrbitDB == nil:
			c.OrbitDB = DefaultConfig.OrbitDB
		case c.Pinner == nil:
			c.Pinner = DefaultConfig.Pinner
		default:
			break loop
		}
	}

	c.IPFS.SetDefaults()
	c.OrbitDB.SetDefaults()
	c.Pinner.SetDefaults()
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
	err = c.Pinner.Validate()
	if err != nil {
		return errors.Wrap(err, "failed to validate pinner configuration")
	}
	return nil
}
