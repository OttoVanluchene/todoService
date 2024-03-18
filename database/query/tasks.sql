-- name: GetTasks :many
SELECT id, title, description, completed FROM tasks;

-- name: CreateTask :exec
INSERT INTO tasks (title, description, completed) VALUES ($1, $2, $3);

-- name: GetTask :one
SELECT id, title, description, completed FROM tasks WHERE id = $1;

-- name: UpdateTask :exec
UPDATE tasks SET title = $2, description = $3, completed = $4 WHERE id = $1;

-- name: DeleteTask :exec
DELETE FROM tasks WHERE id = $1;
