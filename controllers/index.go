package controllers

import (
	"htmx-templ/services/cluster"
	"htmx-templ/services/project"
	"htmx-templ/views"

	"github.com/gofiber/fiber/v2"
)

var _ Controller = (*Index)(nil)

type Index struct {
	ps *project.Service
	cs *cluster.Service
}

func NewIndex(ps *project.Service) *Index {
	return &Index{ps: ps}
}

func (i *Index) Register(app *fiber.App) {
	app.Get("/", i.View)
}

func (i *Index) View(c *fiber.Ctx) error {
	nProjects, err := i.ps.CountProjects(c.Context())
	if err != nil {
		return err
	}
	nClusters, err := i.cs.CountClusters(c.Context())
	if err != nil {
		return err
	}
	return views.Render(c, views.Index(nProjects, nClusters))
}
