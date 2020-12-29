package apiserver

type Config struct {
	DatabaseURL string `json:"database_url"`
}

func NewConfig() *Config {
	return &Config{}
}