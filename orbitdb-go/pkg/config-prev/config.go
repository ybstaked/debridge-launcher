package config

import (
	"net/url"
	"strings"
	"time"

	"github.com/debridge-finance/orbitdb-go/pkg/revip"

	"github.com/debridge-finance/orbitdb-go/pkg/bus"
	"github.com/debridge-finance/orbitdb-go/pkg/log"
	"github.com/debridge-finance/orbitdb-go/pkg/meta"
)

const (
	Subsystem = "config"
)

var (
	EnvironPrefix = meta.EnvNamespace

	LocalPostprocessors = []revip.Option{
		revip.WithDefaults(),
		revip.WithValidation(),
	}
	InitPostprocessors = []revip.Option{
		revip.WithDefaults(),
	}
)

var (
	Unmarshaler = revip.YamlUnmarshaler
	Marshaler   = revip.YamlMarshaler
)

type Config struct {
	Log *log.Config

	ShutdownGraceTime time.Duration
	EnvPrefix         string
}

func (c *Config) Default() {
loop:
	for {
		switch {
		case c.Log == nil:
			c.Log = &log.Config{}
		// TODO: add config for orbitdb and ipfs
		case c.ShutdownGraceTime == 0:
			c.ShutdownGraceTime = 120 * time.Second
		case c.EnvPrefix == "":
			c.EnvPrefix = "production"
		default:
			break loop
		}
	}
}

func (c *Config) Update(cc interface{}) error {
	bus.Config <- bus.ConfigUpdate{
		Subsystem: Subsystem,
		Config:    cc,
	}
	return nil
}

//

func Default() (*Config, error) {
	c := &Config{}
	err := revip.Postprocess(
		c,
		revip.WithDefaults(),
	)
	if err != nil {
		return nil, err
	}

	return c, nil
}

func Load(paths []string, postprocessors ...revip.Option) (*Config, error) {
	var (
		c   = &Config{}
		err error
	)

	_, err = log.Create(log.Config{Level: "info"})
	if err != nil {
		return nil, err
	}

	//

	if len(postprocessors) == 0 {

		postprocessors = append(postprocessors, LocalPostprocessors...)
		for _, path := range paths {
			if strings.HasPrefix(path, revip.SchemeEtcd+":") {
				_, err := url.Parse(path)
				if err != nil {
					return nil, err
				}

			}
		}
	}

	loaders := make([]revip.Option, len(paths))
	for n, path := range paths {
		loaders[n], err = revip.FromURL(
			strings.TrimSpace(path),
			Unmarshaler,
		)
		if nil != err {
			return nil, err
		}
	}

	//

	// FIXME: this is because etcd loader is bad at handling pointers,
	// we need default values for this to work
	err = revip.Postprocess(
		c,
		InitPostprocessors...,
	)
	if err != nil {
		return nil, err
	}

	_, err = revip.Load(
		c,
		append(loaders, revip.FromEnviron(EnvironPrefix))...,
	)
	if err != nil {
		return nil, err
	}

	err = revip.Postprocess(
		c,
		postprocessors...,
	)
	if err != nil {
		return nil, err
	}

	return c, nil
}

func Validate(c *Config) error {
	return revip.WithValidation()(c)
}
