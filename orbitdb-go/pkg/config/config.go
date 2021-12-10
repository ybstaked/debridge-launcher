package config

import (
	"os"
	"strings"

	"github.com/mitchellh/mapstructure"
	"github.com/spf13/viper"

	"github.com/debridge-finance/orbitdb-go/pkg/errors"
	"github.com/debridge-finance/orbitdb-go/pkg/path"
	"github.com/debridge-finance/orbitdb-go/pkg/time"
)

type Config interface {
	SetDefaults()
	Validate() error
}

type Loader struct {
	EnvPrefix string
}

func (c *Loader) Unmarshal(config *viper.Viper, v interface{}) error {
	err := config.Unmarshal(v, func(c *mapstructure.DecoderConfig) {
		c.DecodeHook = mapstructure.ComposeDecodeHookFunc(
			mapstructure.StringToTimeDurationHookFunc(),
			mapstructure.StringToSliceHookFunc(","),
			mapstructure.StringToTimeHookFunc(time.LayoutRFC3339),
		)
	})
	if err != nil {
		return errors.Wrapf(err, "failed to unmarshal viper config %s", config.ConfigFileUsed())
	}

	return nil
}

func (c *Loader) Load(p string, v Config) error {
	if p == "" {
		return errors.New("path to the configuration file should not be empty")
	}

	var (
		wd, dir, name string
		err           error
		config        = viper.New()
		keyReplacer   = strings.NewReplacer(".", "_")
	)

	if p == "" {
		return errors.New("failed to load configuration: path is empty")
	}

	wd, err = os.Getwd()
	if err != nil {
		return errors.Wrap(err, "failed to get cwd while loading config")
	}

	p = path.Resolve(wd, p)

	if _, err := os.Stat(p); os.IsNotExist(err) {
		// XXX: viper does not check this and fails with confusing error message
		return errors.Errorf("configuration file %s does not exists", p)
	}

	dir, name, _ = path.Explode(p)

	config.AddConfigPath(dir)
	config.SetConfigName(name)
	config.SetEnvPrefix(c.EnvPrefix)
	config.SetEnvKeyReplacer(keyReplacer)
	config.AutomaticEnv()

	err = config.ReadInConfig()
	if err != nil {
		return errors.Wrapf(err, "failed to read configuration from %s", p)
	}

	//

	err = c.Unmarshal(config, v)
	if err != nil {
		return errors.Wrap(err, "failed to merge loaded configuration with default")
	}
	v.SetDefaults()

	return nil
}

func NewLoader(envPrefix string) *Loader { return &Loader{envPrefix} }
