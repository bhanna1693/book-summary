package handlers

import (
	"github.com/bhanna1693/book-summary/internal/templates"
	"github.com/bhanna1693/book-summary/internal/utils"
	"github.com/labstack/echo/v4"
)

type GetBookAutocompleteHandler struct{}

func NewGetBookAutocompleteHandler() *GetBookAutocompleteHandler {
	return &GetBookAutocompleteHandler{}
}

func (h *GetBookAutocompleteHandler) ServeHTTP(e echo.Context) error {
	bookName := e.QueryParam("q")
	options := []string{bookName}
	// https://openlibrary.org/dev/docs/api/search
	return utils.Render(e, templates.BookAutocompleteOptions(options))
}
