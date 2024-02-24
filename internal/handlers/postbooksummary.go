package handlers

import (
	"github.com/bhanna1693/book-summary/internal/templates"
	"github.com/bhanna1693/book-summary/internal/utils"
	"github.com/labstack/echo/v4"
)

type PostBookSummaryHandler struct{}

func NewPostBookSummaryHandler() *PostBookSummaryHandler {
	return &PostBookSummaryHandler{}
}

func (h *PostBookSummaryHandler) ServeHTTP(e echo.Context) error {
	return utils.Render(e, templates.Home("POST!"))
}
