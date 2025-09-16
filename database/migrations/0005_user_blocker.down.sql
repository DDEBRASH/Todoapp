-- name: UnblockUser :exec
UPDATE users SET is_blocked = false, failed_attempts = 0 WHERE id = $1;