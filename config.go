package logger

// Config for logger
type Config struct {
	Host string
}

func DefaultConfig(host string) *Config {
	if len(host) < 2 {
		host = "localhost:17000"
	}

	return &Config{Host: host}
}
