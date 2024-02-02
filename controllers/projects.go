package controllers

import (
	"htmx-templ/services/project"
	"htmx-templ/views"

	"github.com/gofiber/fiber/v2"
)

type ProjectController struct {
	ps project.ProjectService
}

func NewProjects(ps project.ProjectService) *ProjectController {
	return &ProjectController{ps}
}

func (p *ProjectController) Register(app *fiber.App) {
	app.Get("/projects", p.Index)
}

func (p *ProjectController) Index(c *fiber.Ctx) error {
	projects, err := p.ps.ListProjects(c.Context())
	if err != nil {
		return err
	}

	return views.Render(c, views.Projects(projects))
}
