package api

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
	"time"
	"todoapp/database/generated"
	"todoapp/i18n"
	"todoapp/settings"

	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

type ctxKey string

const userIDKey ctxKey = "userID"

// Middleware
func WithAuth(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authHeader := r.Header.Get("Authorization")
		fmt.Printf("AUTH DEBUG: Authorization header: '%s'\n", authHeader)
		if authHeader == "" {
			fmt.Printf("AUTH DEBUG: No Authorization header found\n")
			i18n.ErrorResponse(w, r, "Unauthorized", http.StatusUnauthorized)
			return
		}

		// "Bearer <token>"
		tokenString := strings.TrimPrefix(authHeader, "Bearer ")
		fmt.Printf("AUTH DEBUG: Token string after trim: '%s'\n", tokenString)
		if tokenString == authHeader {
			// No "Bearer " prefix found
			fmt.Printf("AUTH DEBUG: No 'Bearer ' prefix found\n")
			i18n.ErrorResponse(w, r, "Invalid authorization header format", http.StatusUnauthorized)
			return
		}

		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
			}
			return []byte(settings.JWT_SECRET), nil
		})
		if err != nil || !token.Valid {
			fmt.Printf("AUTH DEBUG: JWT parse error: %v, valid: %v\n", err, token.Valid)
			i18n.ErrorResponse(w, r, "Unauthorized", http.StatusUnauthorized)
			return
		}

		claims, ok := token.Claims.(jwt.MapClaims)
		if !ok {
			i18n.ErrorResponse(w, r, "Invalid token claims", http.StatusUnauthorized)
			return
		}

		userIDFloat, exists := claims["user_id"]
		if !exists {
			i18n.ErrorResponse(w, r, "Missing user_id in token", http.StatusUnauthorized)
			return
		}

		uid, ok := userIDFloat.(float64)
		if !ok {
			i18n.ErrorResponse(w, r, "Invalid user_id format in token", http.StatusUnauthorized)
			return
		}

		if uid <= 0 {
			i18n.ErrorResponse(w, r, "Invalid user_id value in token", http.StatusUnauthorized)
			return
		}

		fmt.Printf("AUTH DEBUG: Successfully authenticated user ID: %d\n", int32(uid))
		ctx := context.WithValue(r.Context(), userIDKey, int32(uid))
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

// Достаём user_id из контекста
func MustUserID(r *http.Request) (int32, bool) {
	if v := r.Context().Value(userIDKey); v != nil {
		if id, ok := v.(int32); ok && id > 0 {
			return id, true
		}
	}
	return 0, false
}

// ===== REQUEST PASSWORD RESET =====
func (a *UserAPI) RequestPasswordReset(w http.ResponseWriter, r *http.Request) {
	var in struct {
		Email string `json:"email"`
	}
	if err := json.NewDecoder(r.Body).Decode(&in); err != nil || in.Email == "" {
		i18n.ErrorResponse(w, r, "bad json", http.StatusBadRequest)
		return
	}

	user, err := a.Q.GetUserByEmail(r.Context(), sql.NullString{String: in.Email, Valid: true})
	if err != nil {
		i18n.ErrorResponse(w, r, "user not found", http.StatusNotFound)
		return
	}

	token := generateToken(32) // генерим случайный токен
	expires := time.Now().Add(1 * time.Hour)

	_, err = a.Q.CreatePasswordReset(r.Context(), generated.CreatePasswordResetParams{
		UserID:    user.ID,
		Token:     token,
		ExpiresAt: expires,
	})
	if err != nil {
		i18n.ErrorResponse(w, r, "db err", http.StatusInternalServerError)
		return
	}

	// For testing: log reset URL to console
	fmt.Printf("\nPASSWORD RESET LINK (для тестирования):\n")
	fmt.Printf("   %s/reset-password?token=%s\n\n", settings.APP_BASE_URL, token)

	w.Header().Set("Content-Type", "application/json")
	_ = json.NewEncoder(w).Encode(map[string]string{
		"message": "Password reset email sent successfully",
	})
}

// ===== RESET PASSWORD =====
func (a *UserAPI) ResetPassword(w http.ResponseWriter, r *http.Request) {
	var in struct {
		Token    string `json:"token"`
		Password string `json:"password"`
	}
	if err := json.NewDecoder(r.Body).Decode(&in); err != nil || in.Token == "" || in.Password == "" {
		i18n.ErrorResponse(w, r, "bad json", http.StatusBadRequest)
		return
	}

	reset, err := a.Q.GetPasswordResetByToken(r.Context(), in.Token)
	if err != nil || reset.Used.Bool || reset.ExpiresAt.Before(time.Now()) {
		i18n.ErrorResponse(w, r, "invalid or expired token", http.StatusBadRequest)
		return
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(in.Password), bcrypt.DefaultCost)
	if err != nil {
		i18n.ErrorResponse(w, r, "hash err", http.StatusInternalServerError)
		return
	}

	// меняем пароль
	if err := a.Q.SetUserPassword(r.Context(), generated.SetUserPasswordParams{
		ID:           reset.UserID,
		PasswordHash: string(hash),
	}); err != nil {
		i18n.ErrorResponse(w, r, "db err", http.StatusInternalServerError)
		return
	}

	// помечаем токен как использованный
	_ = a.Q.MarkPasswordResetUsed(r.Context(), reset.ID)

	w.WriteHeader(http.StatusNoContent)
}

// ===== Helper =====
func generateToken(n int) string {
	const letters = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	b := make([]byte, n)
	for i := range b {
		b[i] = letters[time.Now().UnixNano()%int64(len(letters))]
	}
	return string(b)
}
