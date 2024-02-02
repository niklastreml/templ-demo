-- name: NewProject :one
INSERT INTO project (name, cpu, memory, storage, cluster)
VALUES ($1, $2, $3, $4, $5)
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
-- name: CountProjects :one
SELECT COUNT(*)
FROM project;