package db

// Config for db config
type Config struct {
	Read       *Database
	Write      *Database
	Secrets    string `yaml:"secrets"`
	WithColor  bool   `yaml:"with_color"`
	WithCaller bool   `yaml:"with_caller"`
}

// NewInjection ...
func (c *Config) NewInjection() *Config {
	return c
}
