package handlers

import (
	"github.com/bhanna1693/book-summary/internal/templates"
	"github.com/bhanna1693/book-summary/internal/utils"
	"github.com/labstack/echo/v4"
)

type HomeHandler struct{}

func NewHomeHandler() *HomeHandler {
	return &HomeHandler{}
}

func (h *HomeHandler) ServeHTTP(e echo.Context) error {
	return utils.Render(e, templates.Home("Home", ""))
}
