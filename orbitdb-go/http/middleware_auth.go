package http

import (
	"net/http"

	"github.com/debridge-finance/orbitdb-go/pkg/errors"
	"github.com/debridge-finance/orbitdb-go/pkg/log"
)

var DefaultAuthMiddlewareConfig = AuthMiddlewareConfig{
	Username: "test_user",
	Password: "test_password",
}

//

type AuthMiddlewareConfig struct {
	Username string
	Password string
}

func (c *AuthMiddlewareConfig) SetDefaults() {
loop:
	for {
		switch {
		case c.Username == "":
			c.Username = DefaultAuthMiddlewareConfig.Username
		case c.Password == "":
			c.Password = DefaultAuthMiddlewareConfig.Password
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
	return nil
}

//

func CreateAuthMiddleware(c MiddlewareConfig, l log.Logger) (Middleware, error) {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			u, p, ok := r.BasicAuth()
			if !ok {
				l.Error().Msg("Error parsing basic auth")
				w.WriteHeader(401)
				return
			}
			if u != c.Auth.Username {
				l.Error().Msgf("Username provided is not correct: %s\n", u)
				w.WriteHeader(401)
				return
			}
			if p != c.Auth.Password {
				l.Error().Msgf("Password provided is not correct: %s\n", p)
				w.WriteHeader(401)
				return
			}

			next.ServeHTTP(w, r)
		})
	}, nil
}
