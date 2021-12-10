package http

import (
	"fmt"
	"net/http"
	"os"
	"strings"

	console "github.com/mattn/go-isatty"

	"github.com/debridge-finance/orbitdb-go/pkg/context"
	"github.com/debridge-finance/orbitdb-go/pkg/http/spec"
	"github.com/debridge-finance/orbitdb-go/pkg/log"
)

type (
	Client         = http.Client
	Handler        = http.Handler
	HandlerFunc    = http.HandlerFunc
	ResponseWriter = http.ResponseWriter
	Request        = http.Request

	Context = context.Context
)

//

type Server struct {
	*Router
	Config Config

	log    log.Logger
	server *http.Server
	spec   *spec.Spec
}

func (s *Server) Spec() spec.Spec {
	return *s.spec
}

func (s *Server) ListenAndServe() error {
	s.log.Info().
		Str("address", s.Config.Address).
		Msg("starting")

	if console.IsTerminal(os.Stderr.Fd()) {
		s.log.Debug().Msg("printing route list (terminal is attached)")

		basePath := s.Config.BasePath
		if basePath == "/" {
			basePath = ""
		}

		for _, e := range s.Router.Endpoints() {
			fmt.Fprintf(
				os.Stderr,
				"           %-5s %5s%-40s %q\n",
				strings.ToUpper(e.Method),
				basePath,
				e.Path,
				e.Summary,
			)
		}
	} else {
		s.log.Debug().Msg("not printing route list because terminal is not attached")
	}

	return s.server.ListenAndServe()
}

//

func New(c Config, l log.Logger, ms Middlewares, es spec.Endpoints, specOpts ...spec.Option) *Server {
	l = l.With().Str("component", "server").Logger()
	r := NewRouter(c.BasePath, ms, es)
	s := spec.New(append(specOpts, spec.BasePath(c.BasePath))...)

	if *c.Swagger {
		l.Debug().Msg("enabling swagger endpoints")
		for _, p := range spec.NewSwaggerEndpoints(
			s.Handler(*c.SwaggerCORS),
			spec.NewSwaggerUIHandler(c.BasePath),
		) {
			r.AddEndpoint(p)
		}

	}

	for _, e := range r.Endpoints() {
		s.AddEndpoint(e)
	}

	//

	return &Server{
		Router: r,
		Config: c,
		log:    l,
		server: &http.Server{
			Addr:         c.Address,
			ReadTimeout:  c.ReadTimeout,
			WriteTimeout: c.WriteTimeout,
			Handler:      r,
		},
		spec: s,
	}
}
