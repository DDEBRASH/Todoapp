package api

import (
	"database/sql"
	"encoding/json"
	"net/http"
	"regexp"
	"time"

	"todoapp/database/generated"
	"todoapp/i18n"
	"todoapp/settings"

	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

type UserAPI struct {
	Q *generated.Queries
}

// JWT secret is now imported from settings package

// ===== REGISTER =====
func (a *UserAPI) Register(w http.ResponseWriter, r *http.Request) {
	var in struct {
		Username string `json:"username"`
		Email    string `json:"email"`
		Password string `json:"password"`
	}
	if err := json.NewDecoder(r.Body).Decode(&in); err != nil || in.Username == "" || in.Email == "" || in.Password == "" {
		i18n.ErrorResponse(w, r, "username, email and password are required", http.StatusBadRequest)
		return
	}

	// Email validation with regex
	emailRegex := regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`)
	if !emailRegex.MatchString(in.Email) {
		i18n.ErrorResponse(w, r, "invalid email format", http.StatusBadRequest)
		return
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(in.Password), bcrypt.DefaultCost)
	if err != nil {
		i18n.ErrorResponse(w, r, "hash error", http.StatusInternalServerError)
		return
	}

	u, err := a.Q.CreateUser(r.Context(), generated.CreateUserParams{
		Username:     in.Username,
		Email:        sql.NullString{String: in.Email, Valid: true},
		PasswordHash: string(hash),
	})
	if err != nil {
		i18n.ErrorResponse(w, r, "db err", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	_ = json.NewEncoder(w).Encode(struct {
		ID       int32  `json:"id"`
		Username string `json:"username"`
		Email    string `json:"email"`
		Message  string `json:"message"`
	}{
		ID:       u.ID,
		Username: u.Username,
		Email:    u.Email.String,
		Message:  "Registration successful!",
	})
}

// ===== LOGIN =====
func (a *UserAPI) Login(w http.ResponseWriter, r *http.Request) {
	var in struct {
		Username string `json:"username"`
		Email    string `json:"email"`
		Password string `json:"password"`
	}
	if err := json.NewDecoder(r.Body).Decode(&in); err != nil {
		i18n.ErrorResponse(w, r, "bad json", http.StatusBadRequest)
		return
	}

	// Support both username and email login
	var userID int32
	var passwordHash string
	var isBlocked bool
	var failedAttempts int32

	if in.Email != "" {
		emailUser, err := a.Q.GetUserByEmail(r.Context(), sql.NullString{String: in.Email, Valid: true})
		if err != nil {
			i18n.ErrorResponse(w, r, "user not found", http.StatusUnauthorized)
			return
		}
		userID = emailUser.ID
		passwordHash = emailUser.PasswordHash
		isBlocked = emailUser.IsBlocked
		failedAttempts = emailUser.FailedAttempts
	} else if in.Username != "" {
		usernameUser, err := a.Q.GetUserByUsername(r.Context(), in.Username)
		if err != nil {
			i18n.ErrorResponse(w, r, "user not found", http.StatusUnauthorized)
			return
		}
		userID = usernameUser.ID
		passwordHash = usernameUser.PasswordHash
		isBlocked = usernameUser.IsBlocked
		failedAttempts = usernameUser.FailedAttempts
	} else {
		i18n.ErrorResponse(w, r, "username or email required", http.StatusBadRequest)
		return
	}

	// Проверяем блокировку
	if isBlocked {
		i18n.ErrorResponse(w, r, "account blocked", http.StatusForbidden)
		return
	}

	// Проверяем пароль
	if bcrypt.CompareHashAndPassword([]byte(passwordHash), []byte(in.Password)) != nil {
		// увеличить счетчик
		_ = a.Q.IncrementFailedAttempts(r.Context(), userID)

		// проверить, не превысил ли лимит
		if failedAttempts+1 >= 4 {
			_ = a.Q.BlockUser(r.Context(), userID)
		}

		i18n.ErrorResponse(w, r, "invalid credentials", http.StatusUnauthorized)
		return
	}

	// Если пароль верный → сбрасываем счётчик
	_ = a.Q.ResetFailedAttempts(r.Context(), userID)

	// Генерируем JWT
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id": userID,
		"exp":     time.Now().Add(24 * time.Hour).Unix(),
	})

	tokenString, err := token.SignedString([]byte(settings.JWT_SECRET))
	if err != nil {
		i18n.ErrorResponse(w, r, "token error", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	_ = json.NewEncoder(w).Encode(map[string]string{
		"token": tokenString,
	})
}

// ===== LOGOUT =====
func (a *UserAPI) Logout(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNoContent)
}
