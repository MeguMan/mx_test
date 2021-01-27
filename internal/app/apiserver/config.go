package apiserver

type Config struct {
	DatabaseURL string `json:"database_url"`
	Authorization string `json:"authorization"`
}

func NewConfig() *Config {
	return &Config{}
}