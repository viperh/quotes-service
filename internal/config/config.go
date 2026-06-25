package config

import "os"

type Config struct {
	DbHost string `env:"DB_HOST"`
	DbPort string `env:"DB_PORT"`
	DbUser string `env:"DB_USER"`
	DbPass string `env:"DB_PASS"`
	DbName string `env:"DB_NAME"`
	DbSSL  string `env:"DB_SSL"`

	Port string `env:"PORT"`
}

func New() *Config {
	return &Config{
		DbHost: get("DB_HOST", "localhost"),
		DbPort: get("DB_PORT", "5432"),
		DbUser: get("DB_USER", "postgres"),
		DbPass: get("DB_PASS", "postgres"),
		DbName: get("DB_NAME", "postgres"),
		DbSSL:  get("DB_SSL", "disable"),
		Port:   get("PORT", "3000"),
	}
}

func get(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok && value != "" {
		return value
	}
	return fallback
}
