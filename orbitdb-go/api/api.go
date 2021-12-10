package api

import (
	"debridge-finance/orbitdb-go/app/emitent/api/emission"
	"debridge-finance/orbitdb-go/app/emitent/meta"
	"debridge-finance/orbitdb-go/errors"
	"debridge-finance/orbitdb-go/handler/info"
	"debridge-finance/orbitdb-go/http"
	"debridge-finance/orbitdb-go/http/spec"
	"debridge-finance/orbitdb-go/log"
	"debridge-finance/orbitdb-go/services"
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
		spec.NewEndpoint("post", "/emission/request", "Request tokens emission", // FIXME: more concrete examples
			spec.EndpointHandler(handlers.Get("emissionRequest")),
			spec.EndpointDescription("This handler creates a request for token emission which awaits approval from operator role"),
			spec.EndpointResponse(http.StatusCreated, emission.RequestResult{}, "Successfully created an emission request"),
			spec.EndpointResponse(http.StatusBadRequest, http.Error{}, "Body parsing was failed"),
			spec.EndpointResponse(http.StatusInternalServerError, http.Error{}, "Internal error occured while creating an emission request in blockchain"),
			spec.EndpointTags("emission"),
			spec.EndpointBody(emission.RequestParameters{}, "", true),
			spec.EndpointConsumes(encodingMime),
			spec.EndpointProduces(encodingMime),
		),
		info.CreateGetInfoEndpoint(handlers.Get("getInfo")),
	}

}

func Create(c Config, sc http.Config, l log.Logger, s *services.Services) (*API, error) {
	handlers := spec.HandlerRegistry{}

	getInfo, err := info.CreateGetInfo(s.Blockchain.Masterchain, s.Setting)
	if err != nil {
		return nil, wrapErr(err, "info")
	}

	emissionRequest, err := emission.CreateRequest(
		s.Blockchain.Masterchain,
		s.Setting,
		s.Notification,
	)
	if err != nil {
		return nil, wrapErr(err, "emission request")
	}

	//

	handlers.
		Add("getInfo", getInfo).
		Add("emissionRequest", emissionRequest)

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
			spec.ContactEmail(meta.ContactEmail),
		),
		Config: c,
	}, nil
}
