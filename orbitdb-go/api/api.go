package api

import (
	eventlog "github.com/debridge-finance/orbitdb-go/api/eventlog"
	// "github.com/debridge-finance/orbitdb-go/handler/info"
	"github.com/debridge-finance/orbitdb-go/http"
	"github.com/debridge-finance/orbitdb-go/http/spec"
	"github.com/debridge-finance/orbitdb-go/pkg/errors"
	"github.com/debridge-finance/orbitdb-go/pkg/log"
	"github.com/debridge-finance/orbitdb-go/pkg/meta"
	"github.com/debridge-finance/orbitdb-go/services"
)

type API struct {
	*http.Server

	Config Config
}

func wrapErr(err error, ctrName string) error {
	return errors.Wrapf(err, "failed to create %q handler", ctrName)
}

func Endpoints(handlers spec.HandlerRegistry) spec.Endpoints {
	encodingMime := "application/json"
	return spec.Endpoints{
		spec.NewEndpoint("post", "/submission", "Request tokens emission", // FIXME: more concrete examples
			spec.EndpointHandler(handlers.Get("eventlogAddReq")),
			spec.EndpointDescription("This handler creates a request for token emission which awaits approval from operator role"),
			spec.EndpointResponse(http.StatusCreated, eventlog.RequestResult{}, "Successfully created an eventlog ADD request"),
			spec.EndpointResponse(http.StatusBadRequest, http.Error{}, "Body parsing was failed"),
			spec.EndpointResponse(http.StatusInternalServerError, http.Error{}, "Internal error occured while creating an orbitdb request in blockchain"),
			spec.EndpointTags("orbitdb"),
			spec.EndpointBody(eventlog.RequestResult{}, "", true),
			spec.EndpointConsumes(encodingMime),
			spec.EndpointProduces(encodingMime),
		),
		// info.CreateGetInfoEndpoint(handlers.Get("getInfo")),
	}

}

func Create(c Config, sc http.Config, l log.Logger, s *services.Services) (*API, error) {
	handlers := spec.HandlerRegistry{}

	eventlogAddReq, err := eventlog.CreateAddRequest(
		*c.EventLog, l,
		s.OrbitDB,
	)
	if err != nil {
		return nil, wrapErr(err, "emission request")
	}

	//

	handlers.
		Add("eventlogAddReq", eventlogAddReq)

	//

	ms, err := http.CreateMiddlewares(
		http.DefaultMiddlewareRegistry,
		http.DefaultMiddlewareConfig,
		l,
	)
	if err != nil {
		return nil, errors.Wrap(err, "failed to create middlewares")
	}

	return &API{
		Server: http.New(
			sc, l, ms, Endpoints(handlers),
			spec.Title(meta.Name),
			spec.Version(meta.Version),
			spec.Description(meta.Description),
		),
		Config: c,
	}, nil
}
