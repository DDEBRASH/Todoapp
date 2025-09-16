-- name: CreateUser :one
INSERT INTO users (username, email, password_hash)
VALUES ($1, $2, $3)
RETURNING id, username, email, password_hash;

-- name: GetUserByUsername :one
SELECT id, username, password_hash, failed_attempts, is_blocked, email
FROM users
WHERE username = $1;

-- name: GetUserByID :one
SELECT id, username, password_hash
FROM users
WHERE id = $1;

-- name: ResetFailedAttempts :exec
UPDATE users SET failed_attempts = 0 WHERE id = $1;

-- name: IncrementFailedAttempts :exec
UPDATE users 
SET failed_attempts = failed_attempts + 1, last_failed = NOW() 
WHERE id = $1;

-- name: BlockUser :exec
UPDATE users
SET is_blocked = TRUE
WHERE id = $1;

-- name: SetPasswordHash :exec
UPDATE users SET password_hash = $2 WHERE id = $1;

-- name: GetUserByEmail :one
SELECT id, username, password_hash, email, is_blocked, failed_attempts
FROM users
WHERE email = $1;


-- name: SetUserPassword :exec
UPDATE users SET password_hash = $2 WHERE id = $1;

-- name: GetPasswordReset :one
SELECT * FROM password_reset_tokens WHERE token = $1;

