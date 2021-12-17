package auth

import "errors"

var DefaultConfig = Config{
	Username: "test_user",
	Password: "test_password",
	JWT:      "sldkfnsnlkdfnsldfnksldkfnlsdfns",
}

type Config struct {
	Username string `json:"username"`
	Password string `json:"password"`
	JWT      string `json:"jwt"`
}

func (c *Config) SetDefaults() {
loop:
	for {
		switch {
		case c.JWT == "":
			c.JWT = DefaultConfig.JWT
		default:
			break loop
		}
	}
}
func (c Config) Validate() error {
	if c.JWT == "" {
		return errors.New("jwt secret should not be empty")
	}
	return nil
}
