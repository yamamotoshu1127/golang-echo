package main

import (
	"github.com/labstack/echo/middleware"
	"net/http"

	"github.com/labstack/echo"
)

func main() {
	e := echo.New()
	// can be used to serve static files from the provided root directory.
	// e.g. request /js/main.js -> static/js/main.js
	e.Use(middleware.Static("/static"))
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello echo World!")
	})
	e.Logger.Fatal(e.Start(":8080"))
}
