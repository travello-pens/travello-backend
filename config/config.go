package config

import "os"

type Config struct {
	SERVER_ADDRESS string
	DB_USERNAME    string
	DB_PASSWORD    string
	DB_PORT        string
	DB_HOST        string
	DB_NAME        string
	JWT_KEY        string
}

func InitConfig() Config {
	return Config{
		SERVER_ADDRESS: GetOrDefault("SERVER_ADDRESS", "0.0.0.0:8888"),
		DB_USERNAME:    GetOrDefault("DB_USERNAME", "travello-backend"),
		DB_PASSWORD:    GetOrDefault("DB_PASSWORD", "travelmhHyUEf8h"),
		DB_NAME:        GetOrDefault("DB_NAME", "travello"),
		DB_HOST:        GetOrDefault("DB_HOST", "10.0.0.4"),
		DB_PORT:        GetOrDefault("DB_PORT", "3306"),
		JWT_KEY:        GetOrDefault("JWT_KEY", "AlphaWolf"),
		// DB_USERNAME:    GetOrDefault("DB_USERNAME", "root"),
		// DB_PASSWORD:    GetOrDefault("DB_PASSWORD", ""),
		// DB_HOST:        GetOrDefault("DB_HOST", "127.0.0.1"),
	}
}

func GetOrDefault(envName string, defaultValue string) string {
	if value, ok := os.LookupEnv(envName); ok {
		return value
	}

	return defaultValue
}
