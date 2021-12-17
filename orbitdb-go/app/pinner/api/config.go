package api

import (
	"github.com/debridge-finance/orbitdb-go/api/auth"
	"github.com/debridge-finance/orbitdb-go/api/eventlog"
	"github.com/debridge-finance/orbitdb-go/pkg/errors"
)

var DefaultConfig = Config{
	Auth:     &auth.DefaultConfig,
	EventLog: &eventlog.DefaultConfig,
}

type Config struct {
	Auth     *auth.Config
	EventLog *eventlog.Config
}

func (c *Config) SetDefaults() {
loop:
	for {
		switch {
		case c.Auth == nil:
			c.Auth = DefaultConfig.Auth
		case c.EventLog == nil:
			c.EventLog = DefaultConfig.EventLog
		default:
			break loop
		}
	}
	c.EventLog.SetDefaults()
}

func (c Config) Validate() error {
	wrapErr := func(err error, name string) error {
		return errors.Wrapf(err, "failed to validate %q", name)
	}

	var err error
	err = c.EventLog.Validate()
	if err != nil {
		return wrapErr(err, "eventlog")
	}

	return nil
}
