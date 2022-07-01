package config

type Config struct {
	Port     string
	UserHost string
	UserPort string
}

func NewConfig() *Config {
	return &Config{
		Port:     "9090",
		UserHost: "localhost",
		UserPort: "9091",
	}
}
