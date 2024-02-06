package search

import (
	"context"
	"htmx-templ/database/sqlc"
	"htmx-templ/repo"
)

type SearchService struct {
	repo repo.SearchRepo
}

func NewSearchService(repo repo.SearchRepo) *SearchService {
	return &SearchService{
		repo: repo,
	}
}

func (s *SearchService) Projects(ctx context.Context, name string) ([]sqlc.Project, error) {
	return s.repo.Project(ctx, name)
}
