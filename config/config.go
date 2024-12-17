package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	DBHost string
	DBPort string
	DBUser string
	DBPass string
	DBName string
}

func LoadConfig(path string) *Config {
	err := godotenv.Load(path)
	if err != nil {
		log.Println("Aviso: Não foi possível carregar o arquivo .env. Usando variáveis de ambiente do sistema.")
	}

	return &Config{
		DBHost: getEnv("DB_HOST", "localhost"),
		DBPort: getEnv("DB_PORT", "5432"),
		DBUser: getEnv("DB_USER", "postgres"),
		DBPass: getEnv("DB_PASS", "123"),
		DBName: getEnv("DB_NAME", "postgresqllocal"),
	}
}

func getEnv(key, fallback string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return fallback
}
