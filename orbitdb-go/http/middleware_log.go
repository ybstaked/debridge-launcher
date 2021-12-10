package http

import (
	"net/http"

	"github.com/rs/zerolog/hlog"

	"github.com/debridge-finance/orbitdb-go/pkg/context"
	"github.com/debridge-finance/orbitdb-go/pkg/log"
	"github.com/debridge-finance/orbitdb-go/pkg/time"
)

var DefaultLogMiddlewareConfig = LogMiddlewareConfig{}

//

type LogMiddlewareConfig struct{}

func (c *LogMiddlewareConfig) SetDefaults() {
loop:
	for {
		switch {
		default:
			break loop
		}
	}
}

func (c LogMiddlewareConfig) Validate() error { return nil }

//

func CreateLogMiddleware(c MiddlewareConfig, l log.Logger) (Middleware, error) {
	return hlog.AccessHandler(
		func(r *http.Request, status, size int, duration time.Duration) {
			l, _ := context.Get(r.Context(), RequestLoggerContextKey).(log.Logger)
			hs := make(map[string]interface{}, len(r.Header))
			for k, v := range r.Header {
				hs[k] = v
			}

			l.Debug().Fields(hs).Msg(r.URL.String())
			l.Info().
				Str("method", r.Method).
				Int("status", status).
				Int("size", size).
				Dur("duration", duration).
				Str("address", r.RemoteAddr). // FIXME: toggle x-real-ip, x-forwarded-for trust with configuration file?
				Msg(r.URL.String())
		},
	), nil
}
