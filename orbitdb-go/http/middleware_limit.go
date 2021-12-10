package http

import (
	"net/http"

	bytesize "github.com/inhies/go-bytesize"

	"github.com/debridge-finance/orbitdb-go/pkg/errors"
	"github.com/debridge-finance/orbitdb-go/pkg/io"
	"github.com/debridge-finance/orbitdb-go/pkg/log"
)

var DefaultLimitMiddlewareConfig = LimitMiddlewareConfig{Body: "10MB"}

type LimitMiddlewareConfig struct{ Body string }

func (c *LimitMiddlewareConfig) SetDefaults() {
loop:
	for {
		switch {
		case c.Body == "":
			c.Body = DefaultLimitMiddlewareConfig.Body
		default:
			break loop
		}
	}
}

func (c LimitMiddlewareConfig) Validate() error {
	if c.Body == "" {
		return errors.New("body should not be empty")
	}
	return nil
}

//

func CreateLimitMiddleware(c MiddlewareConfig, l log.Logger) (Middleware, error) {
	bodyLimit, err := bytesize.Parse(c.Limit.Body)
	if err != nil {
		return nil, errors.Wrap(err, "failed to parse body limit")
	}
	bodyLimitInt64 := int64(bodyLimit)

	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			r.Body = io.NewReadCloser(
				io.NewLimitReader(r.Body, bodyLimitInt64),
				r.Body,
			)
			next.ServeHTTP(w, r)
		})
	}, nil
}
