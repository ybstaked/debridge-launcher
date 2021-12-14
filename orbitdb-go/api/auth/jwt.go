package auth

import (
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
	Username string `json:"username"     swag_example:"f9872d1840D7322E4476C4C08c625Ab9E04d3960"`
	Password string `json:"password"`
}

type JWTRequestResult struct {
	JWT string `json:"jwt"             swag_example:"zdpuA"  swag_description:"OrbitDB hash"`
}

func (h *JWTRequest) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	u, p, ok := r.BasicAuth()
	if !ok {
		http.WriteError(
			w, r, http.StatusUnauthorized,
			errors.New("Error parsing basic auth"),
		)
		return
	}
	if u != h.Config.Username {
		http.WriteError(
			w, r, http.StatusUnauthorized,
			errors.New("Username is not correct"),
		)
		return
	}
	if p != h.Config.Password {
		http.WriteError(
			w, r, http.StatusUnauthorized,
			errors.New("Password is not correct"),
		)
		return
	}

	_, tokenString, err := h.tokenAuth.Encode(map[string]interface{}{"username": u})
	if err != nil {
		http.WriteError(
			w, r, http.StatusUnauthorized,
			errors.Wrapf(err, "Failed to encode jwt for username %s\n", u),
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
