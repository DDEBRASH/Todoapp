-- name: CreateEmailVerification :one
INSERT INTO email_verification_tokens (user_id, token, expires_at)
VALUES ($1, $2, $3)
RETURNING id, user_id, token, expires_at, used, created_at;

-- name: GetEmailVerificationByToken :one
SELECT id, user_id, token, expires_at, used, created_at 
FROM email_verification_tokens 
WHERE token = $1;

-- name: MarkEmailVerificationUsed :exec
UPDATE email_verification_tokens 
SET used = TRUE 
WHERE id = $1;

-- name: VerifyUserEmail :exec
UPDATE users 
SET email_verified = TRUE 
WHERE id = $1;

