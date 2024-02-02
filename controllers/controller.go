package controllers

import "github.com/gofiber/fiber/v2"

type Interface interface {
	Register(app *fiber.App)
}
