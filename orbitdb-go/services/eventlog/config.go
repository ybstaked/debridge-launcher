package eventlog

var DefaultConfig = Config{
	Repo: "./data/orbitdb",
}

type Config struct {
	Repo string
}

func (c *Config) SetDefaults() {
loop:
	for {
		switch {
		case c.Repo == "":
			c.Repo = DefaultConfig.Repo
		default:
			break loop
		}
	}
}
func (c Config) Validate() error {
	return nil
}
