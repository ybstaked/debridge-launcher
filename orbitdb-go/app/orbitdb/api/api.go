package api

import (
	"github.com/debridge-finance/orbitdb-go/app/orbitdb/api/asset"
	"github.com/debridge-finance/orbitdb-go/app/orbitdb/api/auth"
	submission "github.com/debridge-finance/orbitdb-go/app/orbitdb/api/submission"
	"github.com/go-chi/jwtauth/v5"

	"github.com/debridge-finance/orbitdb-go/app/orbitdb/services"
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
		spec.NewEndpoint("post", "/submission", "Add new entry to submission", //
			spec.EndpointHandler(handlers.Get("submissionAddReq")),
			spec.EndpointDescription("This handler creates a request for token emission which awaits approval from operator role"),
			spec.EndpointResponse(http.StatusCreated, submission.AddRequestResult{}, "Successfully created an submission ADD request"),
			spec.EndpointResponse(http.StatusBadRequest, http.Error{}, "Body parsing was failed"),
			spec.EndpointResponse(http.StatusInternalServerError, http.Error{}, "Internal error occured while creating an orbitdb request"),
			spec.EndpointTags("orbitdb"),
			spec.EndpointBody(submission.AddRequestResult{}, "", true),
			spec.EndpointConsumes(encodingMime),
			spec.EndpointProduces(encodingMime),
		),
		spec.NewEndpoint("get", "/submission/{hash}", "Get entry from submission by hash",
			spec.EndpointHandler(handlers.Get("submissionGetEntryReq")),
			spec.EndpointDescription("Get submission entry by hash"),
			spec.EndpointPath("hash", "string", "IPFS hash of entry, entry id in submission", true),
			spec.EndpointResponse(http.StatusOk, submission.GetEntryRequestResult{}, "Successful operation"),
			spec.EndpointResponse(http.StatusInternalServerError, http.Error{}, "Internal error occured while creating a get submission by hash request"),
			spec.EndpointBody(submission.GetEntryRequestResult{}, "", true),
			spec.EndpointConsumes(encodingMime),
			spec.EndpointProduces(encodingMime),
		),
		spec.NewEndpoint("get", "/submission/stats", "Get submission stats",
			spec.EndpointHandler(handlers.Get("submissionStatsReq")),
			spec.EndpointDescription("Get submission stats"),
			spec.EndpointResponse(http.StatusOk, submission.StatsRequestResult{}, "Successful operation"),
			spec.EndpointResponse(http.StatusInternalServerError, http.Error{}, "Internal error occured while creating get stats request"),
			spec.EndpointBody(submission.StatsRequestResult{}, "", true),
			spec.EndpointConsumes(encodingMime),
			spec.EndpointProduces(encodingMime),
		),
		spec.NewEndpoint("get", "/submission/address", "Get submission address",
			spec.EndpointHandler(handlers.Get("submissionAddressReq")),
			spec.EndpointDescription("Get submission address"),
			spec.EndpointResponse(http.StatusOk, submission.StatsRequestResult{}, "Successful operation"),
			spec.EndpointResponse(http.StatusInternalServerError, http.Error{}, "Internal error occured while creating get stats request"),
			spec.EndpointBody(submission.StatsRequestResult{}, "", true),
			spec.EndpointConsumes(encodingMime),
			spec.EndpointProduces(encodingMime),
		),
		spec.NewEndpoint("post", "/asset", "Add new entry to asset", //
			spec.EndpointHandler(handlers.Get("assetAddReq")),
			spec.EndpointDescription("This handler creates a request for token emission which awaits approval from operator role"),
			spec.EndpointResponse(http.StatusCreated, asset.AddRequestResult{}, "Successfully created an asset ADD request"),
			spec.EndpointResponse(http.StatusBadRequest, http.Error{}, "Body parsing was failed"),
			spec.EndpointResponse(http.StatusInternalServerError, http.Error{}, "Internal error occured while creating an orbitdb request"),
			spec.EndpointTags("orbitdb"),
			spec.EndpointBody(asset.AddRequestResult{}, "", true),
			spec.EndpointConsumes(encodingMime),
			spec.EndpointProduces(encodingMime),
		),
		spec.NewEndpoint("get", "/asset/{hash}", "Get entry from asset by hash",
			spec.EndpointHandler(handlers.Get("assetGetReq")),
			spec.EndpointDescription("Get assetn by hash"),
			spec.EndpointPath("hash", "string", "IPFS hash of entry, entry id in asset", true),
			spec.EndpointResponse(http.StatusOk, asset.GetRequestResult{}, "Successful operation"),
			spec.EndpointResponse(http.StatusInternalServerError, http.Error{}, "Internal error occured while creating a get asset by hash request"),
			spec.EndpointBody(asset.GetRequestResult{}, "", true),
			spec.EndpointConsumes(encodingMime),
			spec.EndpointProduces(encodingMime),
		),
		spec.NewEndpoint("get", "/asset/stats", "Get asset stats",
			spec.EndpointHandler(handlers.Get("assetStatsReq")),
			spec.EndpointDescription("Get asset stats"),
			spec.EndpointResponse(http.StatusOk, asset.StatsRequestResult{}, "Successful operation"),
			spec.EndpointResponse(http.StatusInternalServerError, http.Error{}, "Internal error occured while creating get stats request"),
			spec.EndpointBody(asset.StatsRequestResult{}, "", true),
			spec.EndpointConsumes(encodingMime),
			spec.EndpointProduces(encodingMime),
		),
		spec.NewEndpoint("get", "/asset/address", "Get asset address",
			spec.EndpointHandler(handlers.Get("assetAddressReq")),
			spec.EndpointDescription("Get asset address"),
			spec.EndpointResponse(http.StatusOk, asset.StatsRequestResult{}, "Successful operation"),
			spec.EndpointResponse(http.StatusInternalServerError, http.Error{}, "Internal error occured while creating get stats request"),
			spec.EndpointBody(asset.StatsRequestResult{}, "", true),
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
	submissionAddReq, err := submission.CreateAddRequest(
		*c.Submission, l,
		s.Submission,
	)
	if err != nil {
		return nil, wrapErr(err, "failed to create addReq")
	}

	submissionGetEntryReq, err := submission.CreateGetEntryRequest(
		s.Submission,
	)
	if err != nil {
		return nil, wrapErr(err, "failed to create getReq")
	}
	submissionStatsReq, err := submission.CreateStatsRequest(
		s.Submission,
	)
	if err != nil {
		return nil, wrapErr(err, "failed to create statsReq")
	}
	submissionAddressReq, err := submission.CreateAddressRequest(
		s.Submission,
	)
	if err != nil {
		return nil, wrapErr(err, "failed to create addressReq")
	}
	assetAddReq, err := asset.CreateAddRequest(
		*c.Asset, l,
		s.Asset,
	)
	if err != nil {
		return nil, wrapErr(err, "failed to create addReq")
	}

	assetGetReq, err := asset.CreateGetRequest(
		s.Asset,
	)
	if err != nil {
		return nil, wrapErr(err, "failed to create getReq")
	}
	assetStatsReq, err := asset.CreateStatsRequest(
		s.Asset,
	)
	if err != nil {
		return nil, wrapErr(err, "failed to create statsReq")
	}
	assetAddressReq, err := asset.CreateAddressRequest(
		s.Asset,
	)

	if err != nil {
		return nil, wrapErr(err, "failed to create addressReq")
	}

	//

	handlers.
		Add("authReq", authReq).
		Add("submissionGetEntryReq", submissionGetEntryReq).
		Add("submissionAddReq", submissionAddReq).
		Add("submissionStatsReq", submissionStatsReq).
		Add("submissionAddressReq", submissionAddressReq).
		Add("assetGetReq", assetGetReq).
		Add("assetAddReq", assetAddReq).
		Add("assetStatsReq", assetStatsReq).
		Add("assetAddressReq", assetAddressReq)

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
