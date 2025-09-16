package i18n

import (
	"net/http"
	"strings"
)

const (
	LangRussian   = "ru"
	LangEnglish   = "en"
	LangMongolian = "mn"
	LangDefault   = LangRussian
)

// Переводы ошибок
var translations = map[string]map[string]string{
	LangRussian: {
		// Общие ошибки
		"Unauthorized":                        "Не авторизован",
		"Invalid authorization header format": "Неверный формат заголовка авторизации",
		"Invalid token claims":                "Неверные данные токена",
		"Missing user_id in token":            "Отсутствует user_id в токене",
		"Invalid user_id format in token":     "Неверный формат user_id в токене",
		"Invalid user_id value in token":      "Неверное значение user_id в токене",
		"bad json":                            "Неверный JSON",
		"user not found":                      "Пользователь не найден",
		"db err":                              "Ошибка базы данных",
		"hash error":                          "Ошибка хеширования",
		"hash err":                            "Ошибка хеширования",
		"token error":                         "Ошибка токена",

		// Ошибки регистрации и входа
		"username, email and password are required": "Требуются имя пользователя, email и пароль",
		"invalid email format":                      "Неверный формат email",
		"username or email required":                "Требуется имя пользователя или email",
		"account blocked":                           "Аккаунт заблокирован",
		"invalid credentials":                       "Неверные учетные данные",

		// Ошибки задач
		"invalid id":   "Неверный ID",
		"invalid json": "Неверный JSON",

		// Ошибки сброса пароля
		"invalid or expired token": "Неверный или истекший токен",

		// Ошибки базы данных (общие)
		"database error": "Ошибка базы данных",
	},

	LangEnglish: {
		// General errors
		"Unauthorized":                        "Unauthorized",
		"Invalid authorization header format": "Invalid authorization header format",
		"Invalid token claims":                "Invalid token claims",
		"Missing user_id in token":            "Missing user_id in token",
		"Invalid user_id format in token":     "Invalid user_id format in token",
		"Invalid user_id value in token":      "Invalid user_id value in token",
		"bad json":                            "Bad JSON",
		"user not found":                      "User not found",
		"db err":                              "Database error",
		"hash error":                          "Hash error",
		"hash err":                            "Hash error",
		"token error":                         "Token error",

		// Registration and login errors
		"username, email and password are required": "Username, email and password are required",
		"invalid email format":                      "Invalid email format",
		"username or email required":                "Username or email required",
		"account blocked":                           "Account blocked",
		"invalid credentials":                       "Invalid credentials",

		// Task errors
		"invalid id":   "Invalid ID",
		"invalid json": "Invalid JSON",

		// Password reset errors
		"invalid or expired token": "Invalid or expired token",

		// Database errors (general)
		"database error": "Database error",
	},

	LangMongolian: {
		"Unauthorized":                        "Зөвшөөрөлгүй",
		"Invalid authorization header format": "Буруу зөвшөөрлийн толгойн формат",
		"Invalid token claims":                "Буруу токенын мэдээлэл",
		"Missing user_id in token":            "Токенд user_id байхгүй",
		"Invalid user_id format in token":     "Токенд user_id-ийн буруу формат",
		"Invalid user_id value in token":      "Токенд user_id-ийн буруу утга",
		"bad json":                            "Буруу JSON",
		"user not found":                      "Хэрэглэгч олдсонгүй",
		"db err":                              "Өгөгдлийн сангийн алдаа",
		"hash error":                          "Хеш алдаа",
		"hash err":                            "Хеш алдаа",
		"token error":                         "Токенын алдаа",

		"username, email and password are required": "Хэрэглэгчийн нэр, и-мэйл болон нууц үг шаардлагатай",
		"invalid email format":                      "Буруу и-мэйл формат",
		"username or email required":                "Хэрэглэгчийн нэр эсвэл и-мэйл шаардлагатай",
		"account blocked":                           "Хэрэглэгчийн эрх хаагдсан",
		"invalid credentials":                       "Буруу нэвтрэх мэдээлэл",

		"invalid id":   "Буруу ID",
		"invalid json": "Буруу JSON",

		"invalid or expired token": "Буруу эсвэл хугацаа дууссан токен",

		"database error": "Өгөгдлийн сангийн алдаа",
	},
}

// GetLanguageFromRequest извлекает язык из заголовков запроса
func GetLanguageFromRequest(r *http.Request) string {
	acceptLang := r.Header.Get("Accept-Language")
	if acceptLang != "" {
		langs := strings.Split(acceptLang, ",")
		if len(langs) > 0 {
			lang := strings.Split(langs[0], ";")[0]
			lang = strings.TrimSpace(lang)

			if _, exists := translations[lang]; exists {
				return lang
			}

			if len(lang) >= 2 {
				shortLang := lang[:2]
				if _, exists := translations[shortLang]; exists {
					return shortLang
				}
			}
		}
	}

	return LangDefault
}

// Translate переводит сообщение об ошибке на указанный язык
func Translate(message, language string) string {
	if _, exists := translations[language]; !exists {
		language = LangDefault
	}

	if translation, exists := translations[language][message]; exists {
		return translation
	}

	return message
}

// TranslateFromRequest переводит сообщение об ошибке на язык из запроса
func TranslateFromRequest(message string, r *http.Request) string {
	language := GetLanguageFromRequest(r)
	return Translate(message, language)
}
