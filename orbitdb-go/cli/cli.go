package cli

import (
	"context"
	"encoding/json"
	"fmt"
	"os"
	"os/signal"
	"strings"
	"sync"
	"syscall"
	"time"

	watchdog "github.com/cloudflare/tableflip"
	spew "github.com/davecgh/go-spew/spew"
	"github.com/debridge-finance/orbitdb-go/pkg/bus"
	"github.com/debridge-finance/orbitdb-go/pkg/config"
	"github.com/debridge-finance/orbitdb-go/pkg/crypto"
	"github.com/debridge-finance/orbitdb-go/pkg/errors"
	"github.com/debridge-finance/orbitdb-go/pkg/log"
	"github.com/debridge-finance/orbitdb-go/pkg/meta"

	cli "github.com/urfave/cli/v2"
	di "go.uber.org/dig"
)

var (
	Flags = []cli.Flag{
		&cli.StringFlag{
			Name:    "pid-file",
			Aliases: []string{"p"},
			EnvVars: []string{config.EnvironPrefix + "_PID_FILE"},
			Usage:   "path to pid file to report into",
			Value:   meta.Name + ".pid",
		},
		&cli.StringFlag{
			Name:    "log-level",
			Aliases: []string{"l"},
			Usage:   "logging level (debug, info, warn, error)",
		},
		&cli.StringSliceFlag{
			Name:    "config",
			Aliases: []string{"c"},
			EnvVars: []string{config.EnvironPrefix + "_CONFIG"},
			Usage:   "path to application configuration file/files",
			Value:   cli.NewStringSlice("config.yml"),
		},

		//

		&cli.DurationFlag{
			Name:  "duration",
			Usage: "exit after duration",
		},
		&cli.BoolFlag{
			Name:  "profile",
			Usage: "write profile information for debugging (cpu.prof, heap.prof)",
		},
		&cli.BoolFlag{
			Name:  "trace",
			Usage: "write trace information for debugging (trace.prof)",
		},
	}
	Commands = []*cli.Command{
		{
			Name:    "config",
			Aliases: []string{"c"},
			Usage:   "Configuration tools",
			Subcommands: []*cli.Command{
				{
					Name:    "validate",
					Aliases: []string{"v"},
					Usage:   "Validate configuration and exit",
					Action:  ConfigValidateAction,
				},
			},
		},
	}

	c *di.Container
)

func Before(ctx *cli.Context) error {
	var err error

	c = di.New()

	//

	err = c.Provide(func() *cli.Context { return ctx })
	if err != nil {
		return err
	}

	err = c.Provide(func() *spew.ConfigState {
		return &spew.ConfigState{
			DisableMethods:          false,
			DisableCapacities:       true,
			DisablePointerAddresses: true,
			Indent:                  "  ",
			SortKeys:                true,
			SpewKeys:                false,
		}
	})
	if err != nil {
		return err
	}

	err = c.Provide(func() *json.Encoder {
		enc := json.NewEncoder(os.Stdout)
		enc.SetIndent("", "  ")
		return enc
	})
	if err != nil {
		return err
	}

	err = c.Provide(func(ctx *cli.Context) (*config.Config, error) {
		c, err := config.Load(ctx.StringSlice("config"))
		if err != nil {
			return nil, err
		}

		return c, nil
	})
	if err != nil {
		return err
	}

	err = c.Provide(func(ctx *cli.Context, c *config.Config) (log.Logger, error) {
		lc := *c.Log
		level := ctx.String("log-level")
		if level != "" {
			lc.Level = level
		}

		return log.Create(lc)
	})
	if err != nil {
		return err
	}

	err = c.Provide(func() crypto.Rand { return crypto.DefaultRand })
	if err != nil {
		return err
	}

	err = c.Provide(func(ctx *cli.Context, c *config.Config) (*watchdog.Upgrader, error) {
		return watchdog.New(watchdog.Options{
			UpgradeTimeout: c.ShutdownGraceTime,
			PIDFile:        ctx.String("pid-file"),
		})
	})
	if err != nil {
		return err
	}

	//

	err = c.Provide(func() *sync.WaitGroup { return &sync.WaitGroup{} })
	if err != nil {
		return err
	}

	err = c.Provide(func() chan error { return make(chan error, 1) })
	if err != nil {
		return err
	}

	err = c.Provide(func() chan os.Signal {
		sig := make(chan os.Signal, 1)
		signal.Notify(
			sig,
			syscall.SIGQUIT,
			syscall.SIGTERM,
			syscall.SIGINT,
			syscall.SIGUSR1,
			syscall.SIGUSR2,
			syscall.SIGHUP,
		)
		return sig
	})
	if err != nil {
		return err
	}

	//

	duration := ctx.Duration("duration")
	if duration == 0 {
		err = c.Provide(func(ctx *cli.Context) context.Context {
			return context.Background()
		})
	} else {
		err = c.Provide(func(ctx *cli.Context) context.Context {
			c, cancel := context.WithTimeout(context.Background(), duration)
			go func() {
				<-c.Done()
				cancel()
			}()
			return c
		})
	}
	if err != nil {
		return err
	}

	//

	if ctx.Bool("profile") {
		err = c.Invoke(writeProfile)
		if err != nil {
			return err
		}
	}

	if ctx.Bool("trace") {
		err = c.Invoke(writeTrace)
		if err != nil {
			return err
		}
	}

	return nil
}

func ConfigValidateAction(ctx *cli.Context) error {
	return c.Invoke(func(l log.Logger) error {
		configs := ctx.StringSlice("config")
		c, err := config.Load(
			configs,
			config.InitPostprocessors...,
		)
		if err != nil {
			return err
		}

		err = config.Validate(c)
		if err != nil {
			return err
		}

		l.Info().
			Strs("configs", configs).
			Msg("configuration validation is ok")

		return nil
	})
}

func RootAction(ctx *cli.Context) error {
	components := c.String()
	_ = c.Invoke(func(l log.Logger) {
		l.Trace().Msgf(
			"component graph: %s",
			strings.TrimSpace(components),
		)
	})

	return c.Invoke(func(
		ctx context.Context,
		cfg *config.Config,
		w *watchdog.Upgrader,
		l log.Logger,
		running *sync.WaitGroup,
		errc chan error,
		sig chan os.Signal,
	) error {
		l.Info().Msg("running")
		spew.Dump([]interface{}{"probably we need spew"})

		err := w.Ready()
		if err != nil {
			return err
		}
		// ipfs
		// // sh := shell.NewShell("localhost:5001")
		// // spew.Dump([]interface{}{"go-ipfs-shell:", sh})

		// ipfs, err := pinner.NewIPFS(ctx, "./ipfs")
		// if err != nil {
		// 	fmt.Println(fmt.Errorf("failed to spawn IPFS node: %s", err))
		// }

		// fmt.Println("creating orbitdb instance...")
		// orbitdb, err := pinner.CreateOrbitdb(ctx, ipfs, "orbitdb")
		// if err != nil {
		// 	errc <- (fmt.Errorf("failed to spawn orbitdb: %s", err))
		// }
		// fmt.Println("orbitdb instance created")

		// // fmt.Println("openning orbitdb db...")
		// db, err := orbitdb.Log(ctx, "eventlog", nil)
		// if err != nil {
		// 	errc <- fmt.Errorf("failed to spawn orbitdb2: %s", err)
		// }
		// fmt.Println("orbitdb db opened...")
		// fmt.Printf("db.Address()\t%v\n", db)

	loop:
		for {
			select {
			case <-w.Exit():
				break loop
			case <-ctx.Done():
				w.Stop()
				break loop

			case err := <-errc:
				if err != nil {
					return err
				}
			case si := <-sig:
				l.Info().Str("signal", si.String()).Msg("received signal")
				switch si {
				case syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT:
					w.Stop()
				case syscall.SIGUSR1, syscall.SIGUSR2, syscall.SIGHUP:
					err = w.Upgrade()
					if err != nil {
						return err
					}
				}
			case <-bus.Config:
				err = w.Upgrade()
				if err != nil {
					return err
				}
			}
		}

		//

		defer os.Exit(0)
		l.Info().Msg("shutdown watchdog")

		time.AfterFunc(cfg.ShutdownGraceTime, func() {
			l.Warn().
				Dur("graceTime", cfg.ShutdownGraceTime).
				Msg("graceful shutdown timed out")
			os.Exit(1)
		})

		running.Wait() // wait for other running components to finish

		return nil
	})
}

//

func NewApp() *cli.App {
	app := &cli.App{}

	app.Before = Before
	app.Flags = Flags
	app.Action = RootAction
	app.Commands = Commands
	app.Version = meta.Version

	return app
}

func Run() {
	err := NewApp().Run(os.Args)
	if err != nil {
		errors.Fatal(errors.Wrap(
			err, fmt.Sprintf(
				"pid: %d, ppid: %d",
				os.Getpid(), os.Getppid(),
			),
		))
	}
}
