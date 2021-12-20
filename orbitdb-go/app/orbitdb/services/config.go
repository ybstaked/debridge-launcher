package services

import (
	"github.com/debridge-finance/orbitdb-go/pkg/errors"
	"github.com/debridge-finance/orbitdb-go/pkg/services/eventlog"
	"github.com/debridge-finance/orbitdb-go/pkg/services/ipfs"
	"github.com/debridge-finance/orbitdb-go/pkg/services/orbitdb"
)

var DefaultConfig = Config{
	IPFS:       &ipfs.DefaultConfig,
	OrbitDB:    &orbitdb.DefaultConfig,
	Submission: &eventlog.DefaultConfig,
	Asset:      &eventlog.DefaultConfig,
}

//

type Config struct {
	IPFS       *ipfs.Config
	OrbitDB    *orbitdb.Config
	Submission *eventlog.Config
	Asset      *eventlog.Config
}

func (c *Config) SetDefaults() {
loop:
	for {
		switch {
		case c.IPFS == nil:
			c.IPFS = DefaultConfig.IPFS
		case c.OrbitDB == nil:
			c.OrbitDB = DefaultConfig.OrbitDB
		case c.Submission == nil:
			c.Submission = DefaultConfig.Submission
		case c.Asset == nil:
			c.Asset = DefaultConfig.Asset
		default:
			break loop
		}
	}

	c.IPFS.SetDefaults()
	c.OrbitDB.SetDefaults()
	c.Submission.SetDefaults()
	c.Asset.SetDefaults()
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
	err = c.Submission.Validate()
	if err != nil {
		return errors.Wrap(err, "failed to validate submission eventlog configuration")
	}
	err = c.Asset.Validate()
	if err != nil {
		return errors.Wrap(err, "failed to validate newAsset eventlog configuration")
	}
	return nil
}
