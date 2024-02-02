package controllers

import (
	"htmx-templ/services/project"
	"htmx-templ/views"

	"github.com/gofiber/fiber/v2"
)

type Index struct {
	ps *project.ProjectService
}

func NewIndex(ps *project.ProjectService) *Index {
	return &Index{ps}
}

func (i *Index) Register(app *fiber.App) {
	app.Get("/", func(c *fiber.Ctx) error {
		nProjects, err := i.ps.CountProjects(c.Context())
		if err != nil {
			return err
		}
		return views.Render(c, views.Index(nProjects))
	})
}
