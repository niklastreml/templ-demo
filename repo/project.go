package repo

import (
	"context"
	"htmx-templ/database"
	"htmx-templ/database/sqlc"
)

var _ ProjectRepo = (*Project)(nil)

type ProjectRepo interface {
	Find(context.Context, int32) (sqlc.Project, error)
	ListAll(context.Context) ([]sqlc.Project, error)
	Create(context.Context, sqlc.NewProjectParams) (sqlc.Project, error)
	Update(context.Context, sqlc.UpdateProjectParams) (sqlc.Project, error)
	Delete(context.Context, int32) error
}

type Project struct {
}

func NewProjectRepo() *Project {
	return &Project{}
}

func (p *Project) Find(ctx context.Context, id int32) (sqlc.Project, error) {
	return sqlc.New(database.DB).GetProject(ctx, id)
}

func (p *Project) ListAll(ctx context.Context) ([]sqlc.Project, error) {
	return sqlc.New(database.DB).GetProjects(ctx)
}

func (p *Project) Create(ctx context.Context, arg sqlc.NewProjectParams) (sqlc.Project, error) {
	return sqlc.New(database.DB).NewProject(ctx, arg)
}

func (p *Project) Update(ctx context.Context, arg sqlc.UpdateProjectParams) (sqlc.Project, error) {
	return sqlc.New(database.DB).UpdateProject(ctx, arg)
}

func (p *Project) Delete(ctx context.Context, id int32) error {
	return sqlc.New(database.DB).DeleteProject(ctx, id)
}
