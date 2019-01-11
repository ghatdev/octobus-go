package octobus

// Config for logger
type Config struct {
	Host    string
	Service string
}

func DefaultConfig(host, service string) *Config {
	if len(host) < 2 {
		host = "localhost:17000"
	}

	return &Config{Host: host, Service: service}
}
