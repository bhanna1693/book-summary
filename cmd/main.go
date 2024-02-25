package main

import (
	"fmt"

	"github.com/bhanna1693/book-summary/internal/handlers"
	"github.com/labstack/echo/v4"
)

func main() {
	fmt.Println("Book Summary App INIT...")
	e := echo.New()
	e.Static("/static", "static")
	e.GET("/", handlers.NewHomeHandler().ServeHTTP)
	e.GET("/book-autocomplete", handlers.NewGetBookAutocompleteHandler().ServeHTTP)
	e.POST("/book-summary", handlers.NewPostBookSummaryHandler().ServeHTTP)

	e.Logger.Fatal(e.Start(":8080"))
}
