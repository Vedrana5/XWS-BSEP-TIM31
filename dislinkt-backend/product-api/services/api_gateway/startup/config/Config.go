package config

type Config struct {
	Port     string
	UserHost string
	UserPort string
	PostHost string
	PostPort string
}

func NewConfig() *Config {
	return &Config{
		Port:     "9090",
		UserHost: "localhost",
		UserPort: "9091",
		PostHost: "localhost",
		PostPort: "9092",
	}
}
