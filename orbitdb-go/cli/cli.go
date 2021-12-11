package cli

import (
	"os"
	"time"

	"github.com/debridge-finance/orbitdb-go/api"
	appConfig "github.com/debridge-finance/orbitdb-go/config"
	config "github.com/debridge-finance/orbitdb-go/pkg/config"
	"github.com/debridge-finance/orbitdb-go/pkg/errors"
	"github.com/debridge-finance/orbitdb-go/pkg/log"
	"github.com/debridge-finance/orbitdb-go/pkg/meta"
	"github.com/debridge-finance/orbitdb-go/supervisor"

	cli "github.com/urfave/cli/v2"
	"gopkg.in/yaml.v2"

	"github.com/debridge-finance/orbitdb-go/services"
)

var (
	Flags = []cli.Flag{
		&cli.StringFlag{
			Name:    "config",
			Aliases: []string{"c"},
			// EnvVars: []string{config.Config.EnvPrefix + "_CONFIG"},
			Usage: "path to service configuration file",
			Value: "config.yml",
		},
		&cli.StringFlag{
			Name:    "address",
			Aliases: []string{"a"},
			Usage:   "address to listen on, example: ':8080'",
			Value:   "",
		},
		&cli.StringFlag{
			Name:    "log-level",
			Aliases: []string{"l"},
			// EnvVars: []string{config.EnvPrefix + "_LOG_LEVEL"},
			Usage: "logging level which must be one of: debug, info, warn, error, panic, fatal",
			Value: "debug",
		},
	}
	Commands = []*cli.Command{
		{
			Name:    "config",
			Aliases: []string{"c"},
			Usage:   "Configuration tools",
			Subcommands: []*cli.Command{
				{
					Name:    "show-default",
					Aliases: []string{"sd"},
					Usage:   "Show default configuration",
					Action:  ConfigShowDefaultAction,
				},
				{
					Name:    "show",
					Aliases: []string{"s"},
					Usage:   "Show defaults merged with configuration from --config file",
					Action:  ConfigShowAction,
				},
			},
		},
	}
)

func ApplyContextToConfig(ctx *cli.Context, c *appConfig.Config) {
	// XXX: to control some configuration keys from command-line arguments
	address := ctx.String("address")
	if address != "" {
		c.Server.Address = address
	}
	logLevel := ctx.String("log-level")
	if logLevel != "" {
		c.Log.Level = logLevel
	}
}

//

func ConfigShowDefaultAction(ctx *cli.Context) error {
	buf, err := yaml.Marshal(appConfig.DefaultConfig)
	if err != nil {
		return err
	}

	_, err = os.Stdout.Write(buf)
	return err
}

func ConfigShowAction(ctx *cli.Context) error {
	c := &appConfig.Config{}
	err := config.
		NewLoader(appConfig.EnvPrefix).
		Load(ctx.String("config"), c)
	if err != nil {
		return err
	}
	// FIXME: may fail with NPE on invalid config
	ApplyContextToConfig(ctx, c)

	enc := yaml.NewEncoder(os.Stdout)
	defer enc.Close()
	return enc.Encode(c)
}

//

func RootAction(ctx *cli.Context) error {
	c := &appConfig.Config{}
	err := config.
		NewLoader(appConfig.EnvPrefix).
		Load(ctx.String("config"), c)
	if err != nil {
		return errors.Wrapf(err, "failed to load config from %q", ctx.String("config"))
	}
	err = c.Validate()
	if err != nil {
		return errors.Wrap(err, "failed to validate config")
	}
	ApplyContextToConfig(ctx, c)

	//

	l, err := log.Create(*c.Log)
	if err != nil {
		return errors.Wrap(err, "failed to create logger")
	}

	// FIXME: configuration for supervisor?
	su := supervisor.New("root", supervisor.NewDelayRestartStrategy(l, 5*time.Second))

	//
	l.Log().Msg("tratata")

	return su.Supervise(func() error {
		ss, err := services.Create(*c.Services, l, ctx.Context)
		if err != nil {
			return errors.Wrap(err, "failed to create services")
		}
		s, err := api.Create(*c.Api, *c.Server, l, ss)
		if err != nil {
			return errors.Wrap(err, "failed to create API handler")
		}

		return s.ListenAndServe()
	})

}

//

func NewApp() *cli.App {
	app := &cli.App{}
	app.Name = meta.CommandName
	app.Usage = meta.Description
	app.Version = meta.Version
	app.Flags = Flags
	app.Commands = Commands
	app.Action = RootAction
	return app
}

func Run() {
	err := NewApp().Run(os.Args)
	if err != nil {
		errors.Fatal(err)
	}
}
