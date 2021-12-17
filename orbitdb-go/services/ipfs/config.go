package ipfs

var DefaultConfig = Config{
	Repo:       "./data",
	IPFSConfig: "./ipfs_config",
}

type Config struct {
	Repo       string
	IPFSConfig string
}

func (c *Config) SetDefaults() {
loop:
	for {
		switch {
		case c.Repo == "":
			c.Repo = DefaultConfig.Repo
		case c.IPFSConfig == "":
			c.IPFSConfig = DefaultConfig.IPFSConfig
		default:
			break loop
		}
	}
}
func (c Config) Validate() error {
	return nil
}
