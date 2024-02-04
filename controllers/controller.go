package controllers

import "github.com/gofiber/fiber/v2"

type Controller interface {
	Register(app *fiber.App)
}
