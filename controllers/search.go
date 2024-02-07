package controllers

import (
	"htmx-templ/services/search"
	"htmx-templ/views"
	"htmx-templ/views/components"

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
	if query == "" {
		return views.Render(c, components.SearchResults(nil))
	}
	projects, err := s.ss.Projects(c.Context(), query)
	if err != nil {
		return err
	}

	return views.Render(c, components.SearchResults(&projects))
}
