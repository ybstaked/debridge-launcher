package http

import (
	"net/http"

	"github.com/google/uuid"

	"github.com/debridge-finance/orbitdb-go/pkg/context"
	"github.com/debridge-finance/orbitdb-go/pkg/log"
)

const (
	RequestIdLogKey    = "requestId"
	RequestIdHeaderKey = "X-Request-ID"
)

var DefaultRequestIdMiddlewareConfig = RequestIdMiddlewareConfig{}

//

type RequestIdMiddlewareConfig struct{}

func (c *RequestIdMiddlewareConfig) SetDefaults() {
loop:
	for {
		switch {
		default:
			break loop
		}
	}
}

func (c RequestIdMiddlewareConfig) Validate() error { return nil }

//

// CreateRequestIdMiddleware creates a middleware to set unique request Id for request.
// If header `X-Request-ID` is already present in the request, it is considered the
// request id. Otherwise, generates a new unique Id
func CreateRequestIdMiddleware(c MiddlewareConfig, l log.Logger) (Middleware, error) {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			ctx := r.Context()
			rid := r.Header.Get(RequestIdHeaderKey)

			if rid != "" {
				_, err := uuid.Parse(rid)
				if err != nil {
					// XXX: malformed request id, discard it
					rid = ""
				}
			}

			if rid == "" {
				rid = uuid.New().String()
				r.Header.Set(RequestIdHeaderKey, rid)
			}

			l, _ := context.Get(ctx, RequestLoggerContextKey).(log.Logger)
			l = l.With().Str(RequestIdLogKey, rid).Logger()
			ctx = context.Put(ctx, RequestLoggerContextKey, l)

			next.ServeHTTP(
				w,
				r.WithContext(context.Put(
					ctx,
					RequestIdContextKey,
					rid,
				)),
			)
		})
	}, nil
}
