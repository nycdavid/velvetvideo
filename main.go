package main

import (
	"bytes"
	"fmt"
	"net/http"
	"strconv"

	"gopkg.in/labstack/echo.v2"
	"gopkg.in/labstack/echo.v2/engine"
	"gopkg.in/labstack/echo.v2/engine/standard"
)

func main() {
	e := echo.New()
	e.Use(VelvetLogger)
	e.GET("/", Home)
	e.Static("/assets", "assets")
	e.Run(standard.New(":1323"))
}

func Home(ctx echo.Context) error {
	return ctx.String(http.StatusOK, "Homepage")
}

func VelvetLogger(next echo.HandlerFunc) echo.HandlerFunc {
	return func(ctx echo.Context) error {
		req := ctx.Request()
		res := ctx.Response()
		var buf bytes.Buffer
		buf.WriteString(req.Method())
		buf.WriteString(" ")
		fullUrl := parseReqUrl(req.URL())
		buf.WriteString(fullUrl)
		statuscode := strconv.Itoa(res.Status())
		buf.WriteString(" ")
		buf.WriteString(statuscode)
		buf.WriteString(" ")
		buf.WriteString(http.StatusText(res.Status()))
		fmt.Println(buf.String())
		next(ctx)
		return nil
	}
}

func parseReqUrl(url engine.URL) string {
	var buf bytes.Buffer
	buf.WriteString(url.Path())
	buf.WriteString("?")
	buf.WriteString(url.QueryString())
	return buf.String()
}
