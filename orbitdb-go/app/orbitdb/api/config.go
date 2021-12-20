package api

import (
	"github.com/debridge-finance/orbitdb-go/app/orbitdb/api/asset"
	"github.com/debridge-finance/orbitdb-go/app/orbitdb/api/auth"
	"github.com/debridge-finance/orbitdb-go/app/orbitdb/api/submission"
	"github.com/debridge-finance/orbitdb-go/pkg/errors"
)

var DefaultConfig = Config{
	Auth:       &auth.DefaultConfig,
	Submission: &submission.DefaultConfig,
	Asset:      &asset.DefaultConfig,
}

type Config struct {
	Auth       *auth.Config
	Submission *submission.Config
	Asset      *asset.Config
}

func (c *Config) SetDefaults() {
loop:
	for {
		switch {
		case c.Auth == nil:
			c.Auth = DefaultConfig.Auth
		case c.Submission == nil:
			c.Submission = DefaultConfig.Submission
		case c.Asset == nil:
			c.Asset = DefaultConfig.Asset
		default:
			break loop
		}
	}
	c.Submission.SetDefaults()
}

func (c Config) Validate() error {
	wrapErr := func(err error, name string) error {
		return errors.Wrapf(err, "failed to validate %q", name)
	}

	var err error
	err = c.Submission.Validate()
	if err != nil {
		return wrapErr(err, "submission")
	}

	return nil
}
