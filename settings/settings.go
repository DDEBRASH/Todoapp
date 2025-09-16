package settings

var (
	APP_PORT = "7540"

	// PostgreSQL
	PG_HOST    = "localhost"
	PG_PORT    = "5432"
	PG_USER    = "postgres"
	PG_PASS    = "1234"
	PG_DB      = "todoapp"
	PG_SSLMODE = "disable"

	// JWT
	JWT_SECRET = "super_secret_key_change_me"

	// SMTP
	SMTP_HOST = "smtp.gmail.com"
	SMTP_PORT = "587"
	SMTP_USER = "test.t0d0.dimash@gmail.com"
	SMTP_PASS = "nres wobi oouf ecvb"
	SMTP_FROM = "test.t0d0.dimash@gmail.com"

	// Приложение
	APP_BASE_URL = "http://localhost:7540"

	// Админ-токен
	ADMIN_TOKEN = "change_me_admin_token"

	// Порог попыток входа
	LOGIN_MAX_FAIL = "4"
)
