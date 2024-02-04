package controllers

import (
	"fmt"
	"htmx-templ/database/sqlc"
	"htmx-templ/services/project"
	"htmx-templ/views"
	"htmx-templ/views/projects"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

var _ Controller = (*Project)(nil)

type Project struct {
	ps *project.Service
}

func NewProjects(ps *project.Service) *Project {
	return &Project{ps: ps}
}

func (p *Project) Register(app *fiber.App) {
	app.Get("/projects", p.Index)
	app.Get("/projects/:id<int>", p.Details)
	app.Post("/projects", p.NewProject)
	app.Get("/projects/count", p.CountOnly)
}

func (p *Project) Index(c *fiber.Ctx) error {
	pros, err := p.ps.ListProjects(c.Context())
	if err != nil {
		return err
	}

	return views.Render(c, projects.Projects(pros))
}

func (p *Project) NewProject(c *fiber.Ctx) error {
	var np sqlc.NewProjectParams
	if err := c.BodyParser(&np); err != nil {
		return err
	}

	c.Context().Logger().Printf("New project: %v", np)

	np.Storage = np.Storage * 1024
	proj, err := p.ps.OrderProject(c.Context(), np)
	if err != nil {
		return err
	}

	c.Response().Header.Add("HX-Trigger", "project-created")
	return views.Render(c, projects.ProjectOrdered(proj))
}

func (p *Project) Details(c *fiber.Ctx) error {
	projectId := c.Params("id")
	id, err := strconv.ParseInt(projectId, 10, 32)
	if err != nil {
		return err
	}

	project, err := p.ps.FindProjectById(c.Context(), int32(id))
	if err != nil {
		return err
	}

	return views.Render(c, projects.ProjectDetails(project))
}

func (p *Project) CountOnly(c *fiber.Ctx) error {
	count, err := p.ps.CountProjects(c.Context())
	if err != nil {
		return err
	}

	c.WriteString(fmt.Sprintf("%d", count))
	return nil
}
