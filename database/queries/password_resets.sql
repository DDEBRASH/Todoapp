-- name: CreatePasswordReset :one
INSERT INTO password_reset_tokens (user_id, token, expires_at)
VALUES ($1, $2, $3)
RETURNING id, user_id, token, expires_at, used;

-- name: GetPasswordResetByToken :one
SELECT * FROM password_reset_tokens
WHERE token = $1
LIMIT 1;

-- name: MarkPasswordResetUsed :exec
UPDATE password_reset_tokens
SET used = TRUE
WHERE id = $1;
