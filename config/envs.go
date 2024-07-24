package config

import (
	"os"
	"strconv"
)

type Config struct {
	Addr 					string
	JWTSecret 		string
	JWTExpiration string
	DB 						DBConfig
}

type DBConfig struct {
	Addr 					string
	MaxOpenConns	int
	MaxIdleConns	int
	MaxIdleTime		string
}

var Envs = initConfig()

func initConfig() Config {
	return Config{
		Addr: getEnvString("ADDR", ":8080"),
		JWTSecret: getEnvString("JWT_SECRET", "secret123"),
		JWTExpiration: getEnvString("JWT_EXPIRATION", "168h"),
		DB: DBConfig{
			Addr: getEnvString("DB_ADDR", "postgres://root:password@localhost/taskmanager?sslmode=disable"),
			MaxOpenConns: getEnvInt("DB_MAX_OPEN_CONNS", 30),
			MaxIdleConns: getEnvInt("DB_MAX_IDLE_CONNS", 30),
			MaxIdleTime:  getEnvString("DB_MAX_IDLE_TIME", "15m"),
		},
	}
}

func getEnvString(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}

func getEnvInt(key string, fallback int) int {
	if value, ok := os.LookupEnv(key); ok {
		int, err := strconv.Atoi(value)
		if err != nil {
			return fallback
		}
		return int
	}
	return fallback
}