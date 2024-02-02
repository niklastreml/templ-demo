package main

import (
	"embed"
	"htmx-templ/views"
	"io/fs"
	"net/http"

	"github.com/a-h/templ"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
	"github.com/gofiber/fiber/v2/middleware/adaptor"
	"github.com/gofiber/fiber/v2/middleware/filesystem"
)

//go:embed static
var staticFS embed.FS

func main() {
	app := fiber.New()
	staticfs, err := fs.Sub(staticFS, "static")
	if err != nil {
		panic(err)
	}

	app.Use("/static", filesystem.New(filesystem.Config{
		Root:   http.FS(staticfs),
		Browse: true,
		MaxAge: -1,
	}))

	app.Get("/", func(c *fiber.Ctx) error {
		return Render(c, views.Index())
	})
	app.Use(NotFoundMiddleware)

	log.Fatal(app.Listen(":3000"))
}

func NotFoundMiddleware(c *fiber.Ctx) error {
	return Render(c, views.NotFound(), templ.WithStatus(http.StatusNotFound))
}

func Render(c *fiber.Ctx, component templ.Component, options ...func(*templ.ComponentHandler)) error {
	componentHandler := templ.Handler(component)
	for _, o := range options {
		o(componentHandler)
	}
	return adaptor.HTTPHandler(componentHandler)(c)
}
