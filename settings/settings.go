package settings

import (
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

var (
	APP_PORT string

	// PostgreSQL
	PG_HOST    string
	PG_PORT    string
	PG_USER    string
	PG_PASS    string
	PG_DB      string
	PG_SSLMODE string

	// JWT
	JWT_SECRET string

	// SMTP
	SMTP_HOST string
	SMTP_PORT string
	SMTP_USER string
	SMTP_PASS string
	SMTP_FROM string

	// Приложение
	APP_BASE_URL string

	// Админ-токен
	ADMIN_TOKEN string

	// Порог попыток входа
	LOGIN_MAX_FAIL string
)

func init() {
	// Загружаем .env файл если он существует
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found, using environment variables")
	}

	// Инициализируем переменные из окружения
	APP_PORT = getEnv("APP_PORT", "7540")
	PG_HOST = getEnv("PG_HOST", "localhost")
	PG_PORT = getEnv("PG_PORT", "5432")
	PG_USER = getEnv("PG_USER", "postgres")
	PG_PASS = getEnv("PG_PASS", "")
	PG_DB = getEnv("PG_DB", "todoapp")
	PG_SSLMODE = getEnv("PG_SSLMODE", "disable")
	JWT_SECRET = getEnv("JWT_SECRET", "")
	SMTP_HOST = getEnv("SMTP_HOST", "smtp.gmail.com")
	SMTP_PORT = getEnv("SMTP_PORT", "587")
	SMTP_USER = getEnv("SMTP_USER", "")
	SMTP_PASS = getEnv("SMTP_PASS", "")
	SMTP_FROM = getEnv("SMTP_FROM", "")
	APP_BASE_URL = getEnv("APP_BASE_URL", "http://localhost:7540")
	ADMIN_TOKEN = getEnv("ADMIN_TOKEN", "")
	LOGIN_MAX_FAIL = getEnv("LOGIN_MAX_FAIL", "4")

	// Проверяем обязательные переменные
	validateRequiredVars()
}

// getEnv получает переменную окружения или возвращает значение по умолчанию
func getEnv(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}

// getEnvAsInt получает переменную окружения как int или возвращает значение по умолчанию
func getEnvAsInt(key string, defaultValue int) int {
	if value := os.Getenv(key); value != "" {
		if intValue, err := strconv.Atoi(value); err == nil {
			return intValue
		}
	}
	return defaultValue
}

// validateRequiredVars проверяет обязательные переменные окружения
func validateRequiredVars() {
	required := map[string]string{
		"PG_PASS":    PG_PASS,
		"JWT_SECRET": JWT_SECRET,
		"ADMIN_TOKEN": ADMIN_TOKEN,
	}

	for name, value := range required {
		if value == "" {
			log.Fatalf("Required environment variable %s is not set", name)
		}
	}
}
