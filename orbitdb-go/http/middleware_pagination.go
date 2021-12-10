package http

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/debridge-finance/orbitdb-go/pkg/context"
	"github.com/debridge-finance/orbitdb-go/pkg/errors"
	"github.com/debridge-finance/orbitdb-go/pkg/log"
)

var DefaultPaginationMiddlewareConfig = PaginationMiddlewareConfig{}

//

type PaginationMiddlewareConfig struct{}

func (c *PaginationMiddlewareConfig) SetDefaults() {
loop:
	for {
		switch {
		default:
			break loop
		}
	}
}

func (c PaginationMiddlewareConfig) Validate() error { return nil }

//

func CreatePaginationMiddleware(c MiddlewareConfig, l log.Logger) (Middleware, error) {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			p := Pagination{}
			ctx := r.Context()
			query := r.URL.Query()

			limit := query.Get("limit")
			if limit != "" {
				if v, err := strconv.ParseUint(limit, 10, 64); err == nil && v > 0 {
					p.Limit = uint64(v)
				} else {
					p.Limit = uint64(0)
				}
			}

			offset := query.Get("offset")
			if offset != "" {
				if v, err := strconv.ParseUint(offset, 10, 64); err == nil {
					p.Offset = v
				} else {
					fmt.Println(errors.Wrap(err, "failed to parse offset param"))
				}
			}
			next.ServeHTTP(
				w,
				r.WithContext(context.Put(
					ctx,
					PaginationContextKey,
					p,
				)),
			)
		})
	}, nil
}
