package eventlog

var DefaultConfig = Config{}

type Config struct{}

func (c *Config) SetDefaults() {
loop:
	for {
		switch {
		default:
			break loop
		}
	}
}
func (c Config) Validate() error {
	return nil
}
