-- name: BlockUser :exec
UPDATE users SET is_blocked = true WHERE id = $1;