package repo

import (
	"context"
	"htmx-templ/database"
	"htmx-templ/database/sqlc"
)

var _ Cluster = (*cluster)(nil)

type Cluster interface {
	Find(context.Context, int32) (sqlc.Cluster, error)
	ListAll(context.Context) ([]sqlc.Cluster, error)
	Create(context.Context, sqlc.NewClusterParams) (sqlc.Cluster, error)
	Update(context.Context, sqlc.UpdateClusterParams) (sqlc.Cluster, error)
	Delete(context.Context, int32) error
	Count(context.Context) (int64, error)
}

type cluster struct{}

func NewClusterRepo() Cluster {
	return &cluster{}
}

func (c *cluster) Find(ctx context.Context, id int32) (sqlc.Cluster, error) {
	return sqlc.New(database.DB).GetCluster(ctx, id)
}

func (c *cluster) ListAll(ctx context.Context) ([]sqlc.Cluster, error) {
	return sqlc.New(database.DB).GetClusters(ctx)
}

func (c *cluster) Create(ctx context.Context, arg sqlc.NewClusterParams) (sqlc.Cluster, error) {
	return sqlc.New(database.DB).NewCluster(ctx, arg)
}

func (c *cluster) Update(ctx context.Context, arg sqlc.UpdateClusterParams) (sqlc.Cluster, error) {
	return sqlc.New(database.DB).UpdateCluster(ctx, arg)
}

func (c *cluster) Delete(ctx context.Context, id int32) error {
	return sqlc.New(database.DB).DeleteCluster(ctx, id)
}

func (c *cluster) Count(ctx context.Context) (int64, error) {
	return sqlc.New(database.DB).CountClusters(ctx)
}
