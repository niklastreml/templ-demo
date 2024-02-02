-- name: NewProject :one
INSERT INTO project (name)
VALUES ($1)
RETURNING *;
-- name: GetProjects :many
SELECT *
FROM project;
-- name: GetProject :one
SELECT *
FROM project
WHERE id = $1;
-- name: UpdateProject :one
UPDATE project
SET name = $1
WHERE id = $2
RETURNING *;
-- name: DeleteProject :exec
DELETE FROM project
WHERE id = $1;