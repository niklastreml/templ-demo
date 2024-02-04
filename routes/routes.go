package routes

import (
	"htmx-templ/controllers"
	"htmx-templ/repo"
	"htmx-templ/services/cluster"
	"htmx-templ/services/project"
	"htmx-templ/views"
	"net/http"

	"github.com/a-h/templ"
	"github.com/gofiber/fiber/v2"
)

type Handlers struct {
	index    *controllers.Index
	projects *controllers.Project
	clusters *controllers.Cluster
}

func NewHandlers() *Handlers {
	ps := project.NewProjectService(repo.NewProjectRepo())
	cs := cluster.NewClusterService(repo.NewClusterRepo())
	return &Handlers{
		index:    controllers.NewIndex(ps),
		projects: controllers.NewProjects(ps),
		clusters: controllers.NewClusters(cs),
	}
}

func (h *Handlers) Setup(app *fiber.App) {
	h.index.Register(app)
	h.projects.Register(app)
	h.clusters.Register(app)
}

func NotFoundMiddleware(c *fiber.Ctx) error {
	return views.Render(c, views.NotFound(), templ.WithStatus(http.StatusNotFound))
}
