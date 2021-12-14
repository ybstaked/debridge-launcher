package http

import (
	"net/http"

	"github.com/debridge-finance/orbitdb-go/pkg/errors"
	"github.com/debridge-finance/orbitdb-go/pkg/log"
	"github.com/go-chi/jwtauth/v5"
)

var DefaultAuthMiddlewareConfig = AuthMiddlewareConfig{
	Username:    "test_user",
	Password:    "test_password",
	JWT:         "wuT7ZSIMe4qkfalknf23rlks__nRyXfYweJhcwEq",
	Unprotected: map[string]struct{}{"/api/auth": {}},
}
var tokenAuth *jwtauth.JWTAuth

//

type AuthMiddlewareConfig struct {
	Username    string
	Password    string
	JWT         string
	Unprotected map[string]struct{}
}

func (c *AuthMiddlewareConfig) SetDefaults() {
loop:
	for {
		switch {
		case c.Username == "":
			c.Username = DefaultAuthMiddlewareConfig.Username
		case c.Password == "":
			c.Password = DefaultAuthMiddlewareConfig.Password
		case c.JWT == "":
			c.JWT = DefaultAuthMiddlewareConfig.JWT
		default:
			break loop
		}
	}
}

func (c AuthMiddlewareConfig) Validate() error {
	if c.Username == "" {
		return errors.New("usernname should not be empty")
	}
	if c.Password == "" {
		return errors.New("password should not be empty")
	}
	if c.JWT == "" {
		return errors.New("jwt should not be empty")
	}
	return nil
}

//

func CreateAuthMiddleware(c MiddlewareConfig, l log.Logger) (Middleware, error) {
	tokenAuth = jwtauth.New("HS256", []byte(c.Auth.JWT), nil)
	verify := jwtauth.Verifier(tokenAuth)
	authenticator := jwtauth.Authenticator
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

			if _, ok := c.Auth.Unprotected[r.URL.Path]; !ok {
				verify(authenticator(next)).ServeHTTP(w, r)
				return
			}
			// spew.Dump([]interface{}{"c.Auth.Unprotected", c.Auth.Unprotected})
			// if r.URL.Path != "/api/auth" {
			// 	verify(authenticator(next)).ServeHTTP(w, r)
			// 	return
			// }

			next.ServeHTTP(w, r)
		})
	}, nil
}
