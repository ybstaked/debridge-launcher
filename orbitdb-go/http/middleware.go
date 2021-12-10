package http

import (
	"fmt"
	"net/http"

	"github.com/debridge-finance/orbitdb-go/pkg/errors"
	"github.com/debridge-finance/orbitdb-go/pkg/log"
)

const (
	LimitMiddlewareName      MiddlewareName = "limit"
	LogMiddlewareName        MiddlewareName = "log"
	LoggerMiddlewareName     MiddlewareName = "logger"
	RequestIdMiddlewareName  MiddlewareName = "requestid"
	PaginationMiddlewareName MiddlewareName = "pagination"
)

var (
	MiddlewareNames = []MiddlewareName{
		LimitMiddlewareName,
		LogMiddlewareName,
		LoggerMiddlewareName,
		RequestIdMiddlewareName,
		PaginationMiddlewareName,
	}

	DefaultMiddlewareRegistry = MiddlewareRegistry{}
)

type (
	Middleware            = func(http.Handler) http.Handler
	Middlewares           = []Middleware
	MiddlewareName        = string
	MiddlewareConstructor = func(MiddlewareConfig, log.Logger) (Middleware, error)
)

type MiddlewareRegistry map[MiddlewareName]MiddlewareConstructor

func (r MiddlewareRegistry) Register(k MiddlewareName, ctr MiddlewareConstructor) MiddlewareRegistry {
	if _, ok := r[k]; ok {
		panic(fmt.Sprintf("key collision: middleware %q already registered", k))
	}

	r[k] = ctr

	return r
}

func (r MiddlewareRegistry) Get(k MiddlewareName) MiddlewareConstructor {
	if ctr, ok := r[k]; !ok {
		panic(fmt.Sprintf("key error: middleware %q is not registered", k))
	} else {
		return ctr
	}
}

func init() {
	DefaultMiddlewareRegistry.
		Register(LimitMiddlewareName, CreateLimitMiddleware).
		Register(LogMiddlewareName, CreateLogMiddleware).
		Register(LoggerMiddlewareName, CreateLoggerMiddleware).
		Register(RequestIdMiddlewareName, CreateRequestIdMiddleware).
		Register(PaginationMiddlewareName, CreatePaginationMiddleware)
}

//

func CreateMiddlewares(r MiddlewareRegistry, c MiddlewareConfig, l log.Logger) (Middlewares, error) {
	var (
		m   Middleware
		err error

		ms = make(Middlewares, len(c.Enable))
	)

	for k, name := range c.Enable {
		m, err = r.Get(name)(c, l)
		if err != nil {
			return nil, errors.Wrapf(err, "failed to create middleware %q", name)
		}

		ms[k] = m
	}

	return ms, nil
}
