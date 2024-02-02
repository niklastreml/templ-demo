package controllers

import (
	"htmx-templ/views"

	"github.com/gofiber/fiber/v2"
)

type Index struct {
}

func NewIndex() *Index {
	return &Index{}
}

func (i *Index) Register(app *fiber.App) {
	app.Get("/", func(c *fiber.Ctx) error {
		return views.Render(c, views.Index())
	})
}
