package api

import (
	"debridge-finance/orbitdb-go/app/emitent/api/emission"
	"debridge-finance/orbitdb-go/errors"
)

var DefaultConfig = Config{
	Emission: &emission.DefaultConfig,
}

type Config struct {
	Emission *emission.Config
}

func (c *Config) SetDefaults() {
loop:
	for {
		switch {
		case c.Emission == nil:
			c.Emission = DefaultConfig.Emission
		default:
			break loop
		}
	}
	c.Emission.SetDefaults()
}

func (c Config) Validate() error {
	wrapErr := func(err error, name string) error {
		return errors.Wrapf(err, "failed to validate %q", name)
	}

	var err error
	err = c.Emission.Validate()
	if err != nil {
		return wrapErr(err, "emission")
	}

	return nil
}
