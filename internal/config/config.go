package config

import "os"

type Config struct {
	App      AppConfig
	HTTP     HTTPConfig
	Database DatabaseConfig
}

type AppConfig struct {
	Name        string
	Environment string
}

type HTTPConfig struct {
	Port string
}

type DatabaseConfig struct {
	URL string
}

func Load() Config {
	return Config{
		App: AppConfig{
			Name:        getEnv("APP_NAME", "go-starter"),
			Environment: getEnv("APP_ENV", "development"),
		},
		HTTP: HTTPConfig{
			Port: getEnv("HTTP_PORT", "8080"),
		},
		Database: DatabaseConfig{
			URL: getEnv("DATABASE_URL", ""),
		},
	}
}

func getEnv(key, fallback string) string {
	value := os.Getenv(key)

	if value == "" {
		return fallback
	}

	return value
}
