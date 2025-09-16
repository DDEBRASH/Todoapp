-- name: CreateTeamProject :one
INSERT INTO team_projects (name, code, created_by)
VALUES ($1, $2, $3)
RETURNING *;

-- name: GetTeamProjectByCode :one
SELECT * FROM team_projects WHERE code = $1;

-- name: GetTeamProjectByID :one
SELECT * FROM team_projects WHERE id = $1;

-- name: GetUserTeamProjects :many
SELECT tp.*, tpm.joined_at
FROM team_projects tp
JOIN team_project_members tpm ON tp.id = tpm.project_id
WHERE tpm.user_id = $1
ORDER BY tpm.joined_at DESC;

-- name: AddTeamProjectMember :one
INSERT INTO team_project_members (project_id, user_id)
VALUES ($1, $2)
RETURNING *;

-- name: GetTeamProjectMembers :many
SELECT u.id, u.username, u.email, tpm.joined_at
FROM team_project_members tpm
JOIN users u ON tpm.user_id = u.id
WHERE tpm.project_id = $1
ORDER BY tpm.joined_at ASC;

-- name: CheckTeamProjectMember :one
SELECT EXISTS(
    SELECT 1 FROM team_project_members 
    WHERE project_id = $1 AND user_id = $2
) as is_member;

-- name: CreateTeamTask :one
INSERT INTO team_tasks (project_id, title, description, done, created_by, deadline)
VALUES ($1, $2, $3, false, $4, $5)
RETURNING *;

-- name: GetTeamTasks :many
SELECT tt.*, u.username as created_by_username
FROM team_tasks tt
JOIN users u ON tt.created_by = u.id
WHERE tt.project_id = $1
ORDER BY tt.created_at DESC;

-- name: UpdateTeamTask :one
UPDATE team_tasks 
SET title = $2, description = $3, deadline = $5, updated_at = CURRENT_TIMESTAMP
WHERE id = $1 AND project_id = $4
RETURNING *;

-- name: SetTeamTaskDone :one
UPDATE team_tasks 
SET done = $2, updated_at = CURRENT_TIMESTAMP
WHERE id = $1 AND project_id = $3
RETURNING *;

-- name: DeleteTeamTask :exec
DELETE FROM team_tasks WHERE id = $1 AND project_id = $2;

-- name: GetTeamTaskByID :one
SELECT tt.*, u.username as created_by_username
FROM team_tasks tt
JOIN users u ON tt.created_by = u.id
WHERE tt.id = $1 AND tt.project_id = $2;
