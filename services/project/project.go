package project

import (
	"context"
	"htmx-templ/database/sqlc"
	"htmx-templ/repo"
)

type Service struct {
	repo repo.ProjectRepo
}

func NewProjectService(repo repo.ProjectRepo) *Service {
	return &Service{
		repo: repo,
	}
}

func (s *Service) OrderProject(ctx context.Context, p sqlc.NewProjectParams) (sqlc.Project, error) {
	return s.repo.Create(ctx, p)
}

func (s *Service) ListProjects(ctx context.Context) ([]sqlc.Project, error) {
	return s.repo.ListAll(ctx)
}

func (s *Service) CountProjects(ctx context.Context) (int64, error) {
	return s.repo.Count(ctx)
}

func (s *Service) FindProjectById(ctx context.Context, id int32) (sqlc.Project, error) {
	return s.repo.Find(ctx, id)
}
