// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.25.0
// source: project.sql

package sqlc

import (
	"context"
)

const countProjects = `-- name: CountProjects :one
SELECT COUNT(*)
FROM project
`

func (q *Queries) CountProjects(ctx context.Context) (int64, error) {
	row := q.db.QueryRowContext(ctx, countProjects)
	var count int64
	err := row.Scan(&count)
	return count, err
}

const deleteProject = `-- name: DeleteProject :exec
DELETE FROM project
WHERE id = $1
`

func (q *Queries) DeleteProject(ctx context.Context, id int32) error {
	_, err := q.db.ExecContext(ctx, deleteProject, id)
	return err
}

const fuzzyFindProjects = `-- name: FuzzyFindProjects :many
SELECT id, name, cpu, memory, storage, cluster FROM project WHERE similarity(name, $1) > 0.15
`

func (q *Queries) FuzzyFindProjects(ctx context.Context, similarity string) ([]Project, error) {
	rows, err := q.db.QueryContext(ctx, fuzzyFindProjects, similarity)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Project
	for rows.Next() {
		var i Project
		if err := rows.Scan(
			&i.ID,
			&i.Name,
			&i.Cpu,
			&i.Memory,
			&i.Storage,
			&i.Cluster,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getProject = `-- name: GetProject :one
SELECT id, name, cpu, memory, storage, cluster
FROM project
WHERE id = $1
`

func (q *Queries) GetProject(ctx context.Context, id int32) (Project, error) {
	row := q.db.QueryRowContext(ctx, getProject, id)
	var i Project
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.Cpu,
		&i.Memory,
		&i.Storage,
		&i.Cluster,
	)
	return i, err
}

const getProjects = `-- name: GetProjects :many
SELECT id, name, cpu, memory, storage, cluster
FROM project
`

func (q *Queries) GetProjects(ctx context.Context) ([]Project, error) {
	rows, err := q.db.QueryContext(ctx, getProjects)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Project
	for rows.Next() {
		var i Project
		if err := rows.Scan(
			&i.ID,
			&i.Name,
			&i.Cpu,
			&i.Memory,
			&i.Storage,
			&i.Cluster,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const newProject = `-- name: NewProject :one
INSERT INTO project (name, cpu, memory, storage, cluster)
VALUES ($1, $2, $3, $4, $5)
RETURNING id, name, cpu, memory, storage, cluster
`

type NewProjectParams struct {
	Name    string
	Cpu     int64
	Memory  int64
	Storage int64
	Cluster string
}

func (q *Queries) NewProject(ctx context.Context, arg NewProjectParams) (Project, error) {
	row := q.db.QueryRowContext(ctx, newProject,
		arg.Name,
		arg.Cpu,
		arg.Memory,
		arg.Storage,
		arg.Cluster,
	)
	var i Project
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.Cpu,
		&i.Memory,
		&i.Storage,
		&i.Cluster,
	)
	return i, err
}

const updateProject = `-- name: UpdateProject :one
UPDATE project
SET name = $1
WHERE id = $2
RETURNING id, name, cpu, memory, storage, cluster
`

type UpdateProjectParams struct {
	Name string
	ID   int32
}

func (q *Queries) UpdateProject(ctx context.Context, arg UpdateProjectParams) (Project, error) {
	row := q.db.QueryRowContext(ctx, updateProject, arg.Name, arg.ID)
	var i Project
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.Cpu,
		&i.Memory,
		&i.Storage,
		&i.Cluster,
	)
	return i, err
}
