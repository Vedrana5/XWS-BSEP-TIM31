package config

type Config struct {
	Port       string
	PostDBPort string
	PostDBHost string
}

func NewConfig() *Config {
	return &Config{

		Port:       "9092",
		PostDBPort: "27017",
		PostDBHost: "localhost",
	}
}
