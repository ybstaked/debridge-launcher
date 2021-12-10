package http

import (
	"path"
	"strings"

	"github.com/go-chi/chi"

	"github.com/debridge-finance/orbitdb-go/pkg/http/spec"
)

type Router struct {
	basePath  string
	router    chi.Router
	endpoints spec.Endpoints
}

func (r *Router) ServeHTTP(rw ResponseWriter, req *Request) { r.router.ServeHTTP(rw, req) }

////

func (r *Router) AddEndpoint(e *spec.Endpoint) {
	var (
		fullPath = path.Join(r.basePath, e.Path)
		handler  = e.Handler.(Handler)
	)

	r.endpoints = append(r.endpoints, e)

	switch {
	// FIXME: any better convention for url which should be mounted?
	case len(fullPath) > 1 && strings.HasSuffix(e.Path, "/"):
		r.router.Mount(fullPath+"/", handler)
	default:
		r.router.Method(
			e.Method,
			fullPath,
			handler,
		)
	}
}

func (r *Router) AddEndpoints(es ...*spec.Endpoint) {
	for _, e := range es {
		r.AddEndpoint(e)
	}
}

func (r *Router) Endpoints() spec.Endpoints {
	return r.endpoints
}

////

func NewRouter(basePath string, ms Middlewares, es spec.Endpoints) *Router {
	r := chi.NewRouter()

	for _, m := range ms {
		r.Use(m)
	}

	r.MethodNotAllowed(func(w ResponseWriter, r *Request) {
		WriteError(w, r, StatusMethodNotAllowed, nil)
	})
	r.NotFound(func(w ResponseWriter, r *Request) {
		WriteNotFound(w, r, nil)
	})

	rr := &Router{
		basePath: basePath,
		router:   r,
	}
	rr.AddEndpoints(es...)

	return rr
}

//

func PathParam(r *Request, key string) string { return chi.URLParam(r, key) }

func Redirect(w ResponseWriter, url string) {
	w.Header().Add("Location", url)
	w.WriteHeader(StatusFound)
}
