package main

import (
	"embed"
	"htmx-templ/database"
	"htmx-templ/routes"
	"io/fs"
	"log"
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/filesystem"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

//go:embed static
var staticFS embed.FS

func main() {
	if err := database.ConnectDB(); err != nil {
		panic(err)
	}
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
	app.Use(logger.New())

	routes.SetupControllers(app)

	log.Fatal(app.Listen(":3000"))
}
