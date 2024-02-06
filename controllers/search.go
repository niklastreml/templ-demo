package controllers

import (
	"htmx-templ/services/search"
	"htmx-templ/views"

	"github.com/gofiber/fiber/v2"
)

type SearchController struct {
	ss *search.SearchService
}

func NewSearch(ps *search.SearchService) *SearchController {
	return &SearchController{ps}
}

func (s *SearchController) Register(app *fiber.App) {
	app.Get("/search", s.Search)
}

func (s *SearchController) Search(c *fiber.Ctx) error {
	query := c.Query("q")
	projects, err := s.ss.Projects(c.Context(), query)
	if err != nil {
		return err
	}

	return views.Render(c, views.Projects(projects))
}
