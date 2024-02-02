package routes

import (
	"htmx-templ/controllers"
	"htmx-templ/repo"
	"htmx-templ/services/project"
	"htmx-templ/views"
	"net/http"

	"github.com/a-h/templ"
	"github.com/gofiber/fiber/v2"
)

func SetupControllers(app *fiber.App) {
	ps := project.NewProjectService(repo.NewProjectRepo())
	controllers.NewIndex(ps).Register(app)
	controllers.NewProjects(ps).Register(app)
	app.Use(NotFoundMiddleware)
}

func NotFoundMiddleware(c *fiber.Ctx) error {
	return views.Render(c, views.NotFound(), templ.WithStatus(http.StatusNotFound))
}
