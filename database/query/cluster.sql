-- name: NewCluster :one
INSERT INTO clusters (
    name, psp_element, operational_model, platform, worker_type,
    is_gpu_worker, max_worker_count, worker_count_number, longhorn_storage,
    nfs_storage, reason, notes_text
)
VALUES (
    $1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12
)
RETURNING *;

-- name: GetClusters :many
SELECT *
FROM clusters;

-- name: GetCluster :one
SELECT *
FROM clusters
WHERE id = $1;

-- name: UpdateCluster :one
UPDATE clusters
SET
    name = $1, psp_element = $2, operational_model = $3, platform = $4,
    worker_type = $5, is_gpu_worker = $6, max_worker_count = $7, worker_count_number = $8,
    longhorn_storage = $9, nfs_storage = $10, reason = $11, notes_text = $12
WHERE id = $13
RETURNING *;

-- name: DeleteCluster :exec
DELETE FROM clusters
WHERE id = $1;

-- name: CountClusters :one
SELECT COUNT(*)
FROM clusters;
