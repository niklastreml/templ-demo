package main

import (
	"context"
	"embed"
	"htmx-templ/database"
	"htmx-templ/routes"
	"io/fs"
	"net/http"
	"os"
	"time"

	"github.com/charmbracelet/log"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/filesystem"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

const port = ":3000"

//go:embed static
var staticFS embed.FS

func main() {
	l := log.New(os.Stdout)
	ctx := log.WithContext(context.Background(), l)

	if err := database.Connect(ctx); err != nil {
		l.Error("Error while connecting to the database", "error", err)
		panic(err)
	}
	app := fiber.New()
	staticfs, err := fs.Sub(staticFS, "static")
	if err != nil {
		l.Error("Error while opening static filesystem", "error", err)
		panic(err)
	}

	app.Use("/static", filesystem.New(filesystem.Config{
		Root:   http.FS(staticfs),
		Browse: true,
		MaxAge: -1,
	}))
	app.Use(logger.New())
	app.Use(func(c *fiber.Ctx) error {
		time.Sleep(500 * time.Millisecond)
		return c.Next()
	})

	hs := routes.NewHandlers()
	hs.Setup(app)
	app.Use(routes.NotFoundMiddleware)

	if err := app.Listen(port); err != nil {
		l.Error("Error while running web server", "port", port, "error", err)
		os.Exit(1)
	}
}
