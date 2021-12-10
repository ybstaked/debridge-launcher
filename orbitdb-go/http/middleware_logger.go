package http

import (
	"net/http"

	"github.com/debridge-finance/orbitdb-go/pkg/context"
	"github.com/debridge-finance/orbitdb-go/pkg/log"
)

var DefaultLoggerMiddlewareConfig = LoggerMiddlewareConfig{}

//

type LoggerMiddlewareConfig struct{}

func (c *LoggerMiddlewareConfig) SetDefaults() {
loop:
	for {
		switch {
		default:
			break loop
		}
	}
}

func (c LoggerMiddlewareConfig) Validate() error {
	return nil
}

//

func CreateLoggerMiddleware(c MiddlewareConfig, l log.Logger) (Middleware, error) {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			next.ServeHTTP(w, r.WithContext(context.Put(
				r.Context(),
				RequestLoggerContextKey,
				l,
			)))
		})
	}, nil
}
