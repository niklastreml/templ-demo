package controllers

import (
	"htmx-templ/views"

	"github.com/gofiber/fiber/v2"
)


type Login struct {
}

func NewLogin() *Login {
	return &Login{}
}

func (i *Login) Register(app *fiber.App) {
	app.Get("/login", func(c *fiber.Ctx) error {
		return views.Render(c, views.Login())
	})
}
