package api

import (
	"github.com/debridge-finance/orbitdb-go/app/pinner/api/auth"
	"github.com/debridge-finance/orbitdb-go/app/pinner/api/eventlog"
	"github.com/go-chi/jwtauth/v5"

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
		spec.NewEndpoint("post", "/auth", "Auth and get jwt token for given user", //
			spec.EndpointHandler(handlers.Get("authReq")),
			spec.EndpointDescription("This handler creates a request for jwt token emission for given user"),
			spec.EndpointResponse(http.StatusCreated, auth.JWTRequestResult{}, "Successfully created an auth request"),
			spec.EndpointResponse(http.StatusBadRequest, http.Error{}, "Body parsing was failed"),
			spec.EndpointResponse(http.StatusInternalServerError, http.Error{}, "Internal error occured while creating an orbitdb request"),
			spec.EndpointTags("orbitdb"),
			spec.EndpointBody(auth.JWTRequestResult{}, "", true),
			spec.EndpointConsumes(encodingMime),
			spec.EndpointProduces(encodingMime),
		),
		spec.NewEndpoint("post", "/eventlog", "Add new entry to eventlog", //
			spec.EndpointHandler(handlers.Get("eventlogAddReq")),
			spec.EndpointDescription("This handler creates a request for token emission which awaits approval from operator role"),
			spec.EndpointResponse(http.StatusCreated, eventlog.AddRequestResult{}, "Successfully created an eventlog ADD request"),
			spec.EndpointResponse(http.StatusBadRequest, http.Error{}, "Body parsing was failed"),
			spec.EndpointResponse(http.StatusInternalServerError, http.Error{}, "Internal error occured while creating an orbitdb request"),
			spec.EndpointTags("orbitdb"),
			spec.EndpointBody(eventlog.AddRequestResult{}, "", true),
			spec.EndpointConsumes(encodingMime),
			spec.EndpointProduces(encodingMime),
		),
		spec.NewEndpoint("get", "/eventlog/{hash}", "Get entry from eventlog by hash",
			spec.EndpointHandler(handlers.Get("eventlogGetReq")),
			spec.EndpointDescription("Get submissionn by hash"),
			spec.EndpointPath("hash", "string", "IPFS hash of entry, entry id in eventlog", true),
			spec.EndpointResponse(http.StatusOk, eventlog.GetRequestResult{}, "Successful operation"),
			spec.EndpointResponse(http.StatusInternalServerError, http.Error{}, "Internal error occured while creating a get submission by hash request"),
			spec.EndpointBody(eventlog.GetRequestResult{}, "", true),
			spec.EndpointConsumes(encodingMime),
			spec.EndpointProduces(encodingMime),
		),
		spec.NewEndpoint("get", "/eventlog/stats", "Get eventlog stats",
			spec.EndpointHandler(handlers.Get("eventlogStatsReq")),
			spec.EndpointDescription("Get eventlog stats"),
			spec.EndpointResponse(http.StatusOk, eventlog.StatsRequestResult{}, "Successful operation"),
			spec.EndpointResponse(http.StatusInternalServerError, http.Error{}, "Internal error occured while creating get stats request"),
			spec.EndpointBody(eventlog.StatsRequestResult{}, "", true),
			spec.EndpointConsumes(encodingMime),
			spec.EndpointProduces(encodingMime),
		),
		spec.NewEndpoint("get", "/eventlog/address", "Get eventlog address",
			spec.EndpointHandler(handlers.Get("eventlogAddressReq")),
			spec.EndpointDescription("Get eventlog address"),
			spec.EndpointResponse(http.StatusOk, eventlog.StatsRequestResult{}, "Successful operation"),
			spec.EndpointResponse(http.StatusInternalServerError, http.Error{}, "Internal error occured while creating get stats request"),
			spec.EndpointBody(eventlog.StatsRequestResult{}, "", true),
			spec.EndpointConsumes(encodingMime),
			spec.EndpointProduces(encodingMime),
		),
		// info.CreateGetInfoEndpoint(handlers.Get("getInfo")),
	}

}

func Create(c Config, sc http.Config, l log.Logger, s *services.Services) (*API, error) {
	handlers := spec.HandlerRegistry{}
	c.Auth = &auth.Config{
		Password: sc.Middlewares.Auth.Password,
		Username: sc.Middlewares.Auth.Username,
		JWT:      sc.Middlewares.Auth.JWT,
	}
	tokenAuth := jwtauth.New("HS256", []byte(c.Auth.JWT), nil)

	authReq, err := auth.CreateJWTRequest(
		*c.Auth, l, tokenAuth,
	)
	if err != nil {
		return nil, wrapErr(err, "failed to create addReq")
	}
	eventlogAddReq, err := eventlog.CreateAddRequest(
		*c.EventLog, l,
		s.Eventlog,
	)
	if err != nil {
		return nil, wrapErr(err, "failed to create addReq")
	}

	eventlogGetReq, err := eventlog.CreateGetRequest(
		s.Eventlog,
	)
	if err != nil {
		return nil, wrapErr(err, "failed to create getReq")
	}
	eventlogStatsReq, err := eventlog.CreateStatsRequest(
		s.Eventlog,
	)
	if err != nil {
		return nil, wrapErr(err, "failed to create statsReq")
	}
	eventlogAddressReq, err := eventlog.CreateAddressRequest(
		s.Eventlog,
	)
	if err != nil {
		return nil, wrapErr(err, "failed to create addressReq")
	}

	//

	handlers.
		Add("authReq", authReq).
		Add("eventlogGetReq", eventlogGetReq).
		Add("eventlogAddReq", eventlogAddReq).
		Add("eventlogStatsReq", eventlogStatsReq).
		Add("eventlogAddressReq", eventlogAddressReq)

	//

	ms, err := http.CreateMiddlewares(
		http.DefaultMiddlewareRegistry,
		*sc.Middlewares,
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
