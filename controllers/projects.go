package controllers

import (
	"htmx-templ/database/sqlc"
	"htmx-templ/services/project"
	"htmx-templ/views"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

type ProjectController struct {
	ps *project.ProjectService
}

func NewProjects(ps *project.ProjectService) *ProjectController {
	return &ProjectController{ps}
}

func (p *ProjectController) Register(app *fiber.App) {
	app.Get("/projects", p.Index)
	app.Get("/projects/:id<int>", p.Details)
	app.Post("/projects", p.NewProject)
}

func (p *ProjectController) Index(c *fiber.Ctx) error {
	projects, err := p.ps.ListProjects(c.Context())
	if err != nil {
		return err
	}

	return views.Render(c, views.Projects(projects))
}

func (p *ProjectController) NewProject(c *fiber.Ctx) error {
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

	return views.Render(c, views.ProjectOrdered(proj))
}

func (p *ProjectController) Details(c *fiber.Ctx) error {
	projectId := c.Params("id")
	id, err := strconv.ParseInt(projectId, 10, 32)
	if err != nil {
		return err
	}

	project, err := p.ps.FindProjectById(c.Context(), int32(id))
	if err != nil {
		return err
	}

	return views.Render(c, views.ProjectDetails(project))
}
