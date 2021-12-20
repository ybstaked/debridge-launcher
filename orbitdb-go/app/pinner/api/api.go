package api

import (
	"github.com/debridge-finance/orbitdb-go/app/pinner/api/auth"
	"github.com/go-chi/jwtauth/v5"

	"github.com/debridge-finance/orbitdb-go/app/pinner/services"
	"github.com/debridge-finance/orbitdb-go/http"
	"github.com/debridge-finance/orbitdb-go/http/spec"
	"github.com/debridge-finance/orbitdb-go/pkg/errors"
	"github.com/debridge-finance/orbitdb-go/pkg/log"
	"github.com/debridge-finance/orbitdb-go/pkg/meta"
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
		// spec.NewEndpoint("post", "/pinnner", "Add new entry to pinnner", //
		// 	spec.EndpointHandler(handlers.Get("pinnnerAddReq")),
		// 	spec.EndpointDescription("This handler creates a request for token emission which awaits approval from operator role"),
		// 	spec.EndpointResponse(http.StatusCreated, pinnner.AddRequestResult{}, "Successfully created an pinnner ADD request"),
		// 	spec.EndpointResponse(http.StatusBadRequest, http.Error{}, "Body parsing was failed"),
		// 	spec.EndpointResponse(http.StatusInternalServerError, http.Error{}, "Internal error occured while creating an orbitdb request"),
		// 	spec.EndpointTags("orbitdb"),
		// 	spec.EndpointBody(pinnner.AddRequestResult{}, "", true),
		// 	spec.EndpointConsumes(encodingMime),
		// 	spec.EndpointProduces(encodingMime),
		// ),
		// spec.NewEndpoint("get", "/pinnner/{hash}", "Get entry from pinnner by hash",
		// 	spec.EndpointHandler(handlers.Get("pinnnerGetReq")),
		// 	spec.EndpointDescription("Get submissionn by hash"),
		// 	spec.EndpointPath("hash", "string", "IPFS hash of entry, entry id in pinnner", true),
		// 	spec.EndpointResponse(http.StatusOk, pinnner.GetRequestResult{}, "Successful operation"),
		// 	spec.EndpointResponse(http.StatusInternalServerError, http.Error{}, "Internal error occured while creating a get submission by hash request"),
		// 	spec.EndpointBody(pinnner.GetRequestResult{}, "", true),
		// 	spec.EndpointConsumes(encodingMime),
		// 	spec.EndpointProduces(encodingMime),
		// ),
		// spec.NewEndpoint("get", "/pinnner/stats", "Get pinnner stats",
		// 	spec.EndpointHandler(handlers.Get("pinnnerStatsReq")),
		// 	spec.EndpointDescription("Get pinnner stats"),
		// 	spec.EndpointResponse(http.StatusOk, pinnner.StatsRequestResult{}, "Successful operation"),
		// 	spec.EndpointResponse(http.StatusInternalServerError, http.Error{}, "Internal error occured while creating get stats request"),
		// 	spec.EndpointBody(pinnner.StatsRequestResult{}, "", true),
		// 	spec.EndpointConsumes(encodingMime),
		// 	spec.EndpointProduces(encodingMime),
		// ),
		// spec.NewEndpoint("get", "/pinnner/address", "Get pinnner address",
		// 	spec.EndpointHandler(handlers.Get("pinnnerAddressReq")),
		// 	spec.EndpointDescription("Get pinnner address"),
		// 	spec.EndpointResponse(http.StatusOk, pinnner.StatsRequestResult{}, "Successful operation"),
		// 	spec.EndpointResponse(http.StatusInternalServerError, http.Error{}, "Internal error occured while creating get stats request"),
		// 	spec.EndpointBody(pinnner.StatsRequestResult{}, "", true),
		// 	spec.EndpointConsumes(encodingMime),
		// 	spec.EndpointProduces(encodingMime),
		// ),
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
	// pinnnerAddReq, err := pinnner.CreateAddRequest(
	// 	*c.EventLog, l,
	// 	s.Pinner,
	// )
	// if err != nil {
	// 	return nil, wrapErr(err, "failed to create addReq")
	// }

	// pinnnerGetReq, err := pinnner.CreateGetRequest(
	// 	s.Pinner,
	// )
	// if err != nil {
	// 	return nil, wrapErr(err, "failed to create getReq")
	// }
	// pinnnerStatsReq, err := pinnner.CreateStatsRequest(
	// 	s.Pinner,
	// )
	// if err != nil {
	// 	return nil, wrapErr(err, "failed to create statsReq")
	// }
	// pinnnerAddressReq, err := pinnner.CreateAddressRequest(
	// 	s.Pinner,
	// )
	// if err != nil {
	// 	return nil, wrapErr(err, "failed to create addressReq")
	// }

	//

	handlers.
		Add("authReq", authReq)
		// Add("pinnnerGetReq", pinnnerGetReq).
		// Add("pinnnerAddReq", pinnnerAddReq).
		// Add("pinnnerStatsReq", pinnnerStatsReq).
		// Add("pinnnerAddressReq", pinnnerAddressReq)

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
