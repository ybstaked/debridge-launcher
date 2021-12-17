package config

import (
	"github.com/debridge-finance/orbitdb-go/app/orbitdb/api"
	"github.com/debridge-finance/orbitdb-go/http"
	"github.com/debridge-finance/orbitdb-go/pkg/errors"
	"github.com/debridge-finance/orbitdb-go/pkg/log"
	"github.com/debridge-finance/orbitdb-go/services"
)

const (
	EnvPrefix = "OrbitDB_API"
)

// Default carries the default configuration values
var DefaultConfig = Config{
	Log:      &log.DefaultConfig,
	Server:   &http.DefaultConfig,
	Api:      &api.DefaultConfig,
	Services: &services.DefaultConfig,
}

//

type Config struct {
	Log      *log.Config
	Server   *http.Config
	Api      *api.Config
	Services *services.Config
}

func (c *Config) SetDefaults() {
loop:
	for {
		switch {
		case c.Log == nil:
			c.Log = DefaultConfig.Log
		case c.Server == nil:
			c.Server = DefaultConfig.Server
		case c.Api == nil:
			c.Api = DefaultConfig.Api
		case c.Services == nil:
			c.Services = DefaultConfig.Services
		default:
			break loop
		}
	}

	c.Log.SetDefaults()
	c.Server.SetDefaults()
	c.Api.SetDefaults()
	c.Services.SetDefaults()
}

func (c Config) Validate() error {
	err := c.Log.Validate()
	if err != nil {
		return errors.Wrap(err, "failed to validate log configuration")
	}
	err = c.Server.Validate()
	if err != nil {
		return errors.Wrap(err, "failed to validate server configuration")
	}
	err = c.Api.Validate()
	if err != nil {
		return errors.Wrap(err, "failed to validate API configuration")
	}
	err = c.Services.Validate()
	if err != nil {
		return errors.Wrap(err, "failed to validate services configuration")
	}
	return nil
}
