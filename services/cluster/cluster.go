package cluster

import (
	"context"
	"htmx-templ/database/sqlc"
	"htmx-templ/repo"
)

type Service struct {
	repo repo.Cluster
}

func NewClusterService(repo repo.Cluster) *Service {
	return &Service{
		repo: repo,
	}
}

func (s *Service) OrderCluster(ctx context.Context, p sqlc.NewClusterParams) (sqlc.Cluster, error) {
	return s.repo.Create(ctx, p)
}

func (s *Service) ListClusters(ctx context.Context) ([]sqlc.Cluster, error) {
	return s.repo.ListAll(ctx)
}

func (s *Service) CountClusters(ctx context.Context) (int64, error) {
	return s.repo.Count(ctx)
}

func (s *Service) FindClusterById(ctx context.Context, id int32) (sqlc.Cluster, error) {
	return s.repo.Find(ctx, id)
}
