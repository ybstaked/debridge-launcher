package eventlog

var DefaultConfig = Config{
	Repo:  "./data/orbitdb",
	Name:  "debridgeEventlog",
	Limit: -1,
}

type Config struct {
	Repo  string
	Name  string
	Limit int
}

func (c *Config) SetDefaults() {
loop:
	for {
		switch {
		case c.Repo == "":
			c.Repo = DefaultConfig.Repo
		case c.Name == "":
			c.Name = DefaultConfig.Name
		case c.Limit == 0:
			c.Limit = DefaultConfig.Limit
		default:
			break loop
		}
	}
}
func (c Config) Validate() error {
	return nil
}
