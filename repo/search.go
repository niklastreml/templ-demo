package repo

import (
	"context"
	"htmx-templ/database"
	"htmx-templ/database/sqlc"
)

type SearchRepo interface {
	Project(context.Context, string) ([]sqlc.Project, error)
}

func NewSearchRepo() SearchRepo {
	return &Search{}
}

type Search struct {
}

func (s *Search) Project(ctx context.Context, name string) ([]sqlc.Project, error) {
	return sqlc.New(database.DB).FuzzyFindProjects(ctx, name)
}
