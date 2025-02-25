package utils

import (
	"github.com/a-h/templ"
	"github.com/labstack/echo/v4"
)

func Render(e echo.Context, component templ.Component) error {
	return component.Render(e.Request().Context(), e.Response())
}
