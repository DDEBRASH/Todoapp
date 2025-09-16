package i18n

import (
	"net/http"
)

// ErrorResponse отправляет переведенную ошибку клиенту
func ErrorResponse(w http.ResponseWriter, r *http.Request, message string, statusCode int) {
	translatedMessage := TranslateFromRequest(message, r)
	http.Error(w, translatedMessage, statusCode)
}

