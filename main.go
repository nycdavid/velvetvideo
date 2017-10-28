package main

import (
	"net/http"

	"gopkg.in/labstack/echo.v3"
)

func main() {
	e := echo.New()
	e.GET("/", Home)
	e.Start(":1323")
}

func Home(ctx echo.Context) error {
	return ctx.String(http.StatusOK, "Homepage")
}
