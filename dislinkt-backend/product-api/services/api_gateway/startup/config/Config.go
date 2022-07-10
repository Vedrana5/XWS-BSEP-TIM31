package config

type Config struct {
	Port           string
	UserHost       string
	UserPort       string
	PostHost       string
	PostPort       string
	ConnectionHost string
	ConnectionPort string
}

func NewConfig() *Config {
	return &Config{
		Port:           "9090",
		UserHost:       "user_service",
		UserPort:       "9091",
		PostHost:       "post_service",
		PostPort:       "9092",
		ConnectionHost: "connection_service",
		ConnectionPort: "9099",
	}
}
