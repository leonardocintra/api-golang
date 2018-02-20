package main

import (
	"net/http"

	"github.com/labstack/echo"
)

func main() {
	e := echo.New()

	e.GET("/", home)

	e.Start(":3000")
}

func home(c echo.Context) error {
	return c.String(http.StatusOK, "Hello World!")
}
