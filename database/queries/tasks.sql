-- name: ListTasks :many
SELECT id, title, done, user_id, deadline
FROM tasks
WHERE user_id = $1
ORDER BY id DESC;

-- name: CreateTask :one
INSERT INTO tasks (title, done, user_id, deadline)
VALUES ($1, false, $2, $3)
RETURNING id, title, done, user_id, deadline;

-- name: UpdateTask :one
UPDATE tasks
SET title = $2, deadline = $4
WHERE id = $1 AND user_id = $3
RETURNING id, title, done, user_id, deadline;

-- name: SetDone :one
UPDATE tasks
SET done = $2
WHERE id = $1 AND user_id = $3
RETURNING id, title, done, user_id;

-- name: DeleteTask :exec
DELETE FROM tasks
WHERE id = $1 AND user_id = $2;
