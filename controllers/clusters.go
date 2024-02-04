package controllers

import (
	"fmt"
	"htmx-templ/database/sqlc"
	"htmx-templ/services/cluster"
	"htmx-templ/views"
	"htmx-templ/views/clusters"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

var _ Controller = (*Cluster)(nil)

type Cluster struct {
	cs *cluster.Service
}

func NewClusters(cs *cluster.Service) *Cluster {
	return &Cluster{cs: cs}
}

func (p *Cluster) Register(app *fiber.App) {
	app.Get("/clusters", p.Index)
	app.Get("/clusters/:id<int>", p.Details)
	app.Post("/clusters", p.NewCluster)
	app.Get("/clusters/count", p.CountOnly)
}

func (p *Cluster) Index(c *fiber.Ctx) error {
	clus, err := p.cs.ListClusters(c.Context())
	if err != nil {
		return err
	}

	return views.Render(c, clusters.Clusters(clus))
}

func (p *Cluster) NewCluster(c *fiber.Ctx) error {
	var nc sqlc.NewClusterParams
	if err := c.BodyParser(&nc); err != nil {
		return err
	}

	c.Context().Logger().Printf("New project: %v", nc)

	nc.LonghornStorage = nc.LonghornStorage * 1024
	clus, err := p.cs.OrderCluster(c.Context(), nc)
	if err != nil {
		return err
	}

	c.Response().Header.Add("HX-Trigger", "project-created")
	return views.Render(c, clusters.ClusterOrdered(clus))
}

func (p *Cluster) Details(c *fiber.Ctx) error {
	projectId := c.Params("id")
	id, err := strconv.ParseInt(projectId, 10, 32)
	if err != nil {
		return err
	}

	cluster, err := p.cs.FindClusterById(c.Context(), int32(id))
	if err != nil {
		return err
	}

	return views.Render(c, clusters.ClusterDetails(cluster))
}

func (p *Cluster) CountOnly(c *fiber.Ctx) error {
	count, err := p.cs.CountClusters(c.Context())
	if err != nil {
		return err
	}

	c.WriteString(fmt.Sprintf("%d", count))
	return nil
}
