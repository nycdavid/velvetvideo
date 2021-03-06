package main

import (
	"bytes"
	"fmt"
	"html/template"
	"io"
	"net/http"
	"net/url"
	"os"
	"strconv"

	"github.com/fatih/color"
	"github.com/spf13/afero"
	"gopkg.in/labstack/echo.v3"
)

type HomeView struct {
	templates *template.Template
}

func (h *HomeView) Render(w io.Writer, name string, data interface{}, ctx echo.Context) error {
	htmlTmpl, err := template.New("home.html").ParseFiles("src/templates/home.html")
	if err != nil {
		fmt.Println("Error parsing template string")
	}
	return htmlTmpl.ExecuteTemplate(w, name, data)
}

func main() {
	e := echo.New()
	e.Renderer = &HomeView{}
	e.Use(VelvetLogger)
	e.GET("/", Home)
	e.Static("/dist", "dist")
	e.Static("/files", "files")
	port := fmt.Sprintf(":%s", os.Getenv("PORT"))
	e.Start(port)
}

func Home(ctx echo.Context) error {
	AppFs := afero.NewOsFs()
	files, _ := afero.Glob(AppFs, "files/*")
	return ctx.Render(http.StatusOK, "home.html", files)
}

func VelvetLogger(next echo.HandlerFunc) echo.HandlerFunc {
	return func(ctx echo.Context) error {
		req := ctx.Request()
		res := ctx.Response()
		var buf bytes.Buffer
		buf.WriteString(req.Method)
		buf.WriteString(" ")
		fullUrl := parseReqUrl(req.URL)
		buf.WriteString(fullUrl)
		statuscode := strconv.Itoa(res.Status)
		buf.WriteString(" ")
		fmt.Print(buf.String())
		if res.Status == 200 {
			color.Green("%s %s", statuscode, http.StatusText(res.Status))
		} else {
		}

		next(ctx)
		return nil
	}
}

func parseReqUrl(url *url.URL) string {
	var buf bytes.Buffer
	buf.WriteString(url.Path)
	if url.RawQuery != "" {
		buf.WriteString("?")
		buf.WriteString(url.RawQuery)
	}
	return buf.String()
}
