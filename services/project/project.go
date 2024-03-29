package project

import (
	"context"
	"htmx-templ/database/sqlc"
	"htmx-templ/repo"
)

type ProjectService struct {
	repo repo.ProjectRepo
}

func NewProjectService(repo repo.ProjectRepo) *ProjectService {
	return &ProjectService{
		repo: repo,
	}
}

func (s *ProjectService) OrderProject(ctx context.Context, p sqlc.NewProjectParams) (sqlc.Project, error) {
	return s.repo.Create(ctx, p)
}

func (s *ProjectService) ListProjects(ctx context.Context) ([]sqlc.Project, error) {
	return s.repo.ListAll(ctx)
}

func (s *ProjectService) CountProjects(ctx context.Context) (int64, error) {
	return s.repo.Count(ctx)
}

func (s *ProjectService) FindProjectById(ctx context.Context, id int32) (sqlc.Project, error) {
	return s.repo.Find(ctx, id)
}
