package config

type Config struct {
	DatabaseURL string
}

func LoadConfig() *Config {
	return &Config{
		DatabaseURL: "host=localhost user=root password=123456 dbname=simple_bank port=5432 sslmode=disable",
	}
}
