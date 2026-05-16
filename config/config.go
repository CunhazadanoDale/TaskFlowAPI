package config

import "os"

type Config struct {
	PostgresURL string
	JwtSecret   string
}

func LoadConfig() (*Config, error) {
	databseurl := os.Getenv("DATABASEURL")
	jwtsecret := os.Getenv("JWTSECRET")

	return &Config{
		PostgresURL: databseurl,
		JwtSecret:   jwtsecret,
	}, nil
}