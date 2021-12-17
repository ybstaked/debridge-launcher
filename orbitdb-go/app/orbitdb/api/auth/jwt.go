package auth

import (
	"encoding/json"

	"github.com/davecgh/go-spew/spew"
	"github.com/debridge-finance/orbitdb-go/http"
	"github.com/debridge-finance/orbitdb-go/pkg/errors"
	"github.com/debridge-finance/orbitdb-go/pkg/log"
	"github.com/go-chi/jwtauth/v5"
)

type JWTRequest struct {
	Config Config

	log       log.Logger
	tokenAuth *jwtauth.JWTAuth
}

type JWTRequestParams struct {
	Username string `json:"login"     swag_example:"login"`
	Password string `json:"password"  swag_example:"password"`
}

type JWTRequestResult struct {
	JWT string `json:"accessToken"             swag_example:"zdpuA"  swag_description:"OrbitDB hash"`
}

func (h *JWTRequest) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	params := &JWTRequestParams{}

	err := json.NewDecoder(r.Body).Decode(params)
	if err != nil {
		http.WriteError(w, r, http.StatusBadRequest, err)
		return
	}
	r.Body.Close()

	err = params.Validate()
	if err != nil {
		http.WriteError(
			w, r, http.StatusBadRequest,
			errors.Wrap(err, "failed to validate request parameters"),
		)
		return
	}
	spew.Dump([]interface{}{"params", params})

	if params.Username != h.Config.Username {
		http.WriteError(
			w, r, http.StatusUnauthorized,
			errors.New("Username is not correct"),
		)
		return
	}
	if params.Password != h.Config.Password {
		http.WriteError(
			w, r, http.StatusUnauthorized,
			errors.New("Password is not correct"),
		)
		return
	}

	_, tokenString, err := h.tokenAuth.Encode(map[string]interface{}{"username": params.Username})
	if err != nil {
		http.WriteError(
			w, r, http.StatusUnauthorized,
			errors.Wrapf(err, "Failed to encode jwt for username %s\n", params.Username),
		)
		return
	}

	http.Write(
		w, r, http.StatusOk,
		&JWTRequestResult{
			JWT: tokenString,
		},
	)
}

func (p *JWTRequestParams) SetDefaults() {
loop:
	for {
		switch {
		default:
			break loop
		}
	}
}

func (p *JWTRequestParams) Validate() error {
	// var err error
	if p.Username == "" {
		return errors.New("usernname should not be empty")
	}
	if p.Username == "" {
		return errors.New("usernname should not be empty")
	}

	return nil
}

func CreateJWTRequest(c Config, l log.Logger, t *jwtauth.JWTAuth) (*JWTRequest, error) {
	return &JWTRequest{
		Config:    c,
		log:       l,
		tokenAuth: t,
	}, nil
}
