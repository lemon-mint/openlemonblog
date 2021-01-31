package main

import (
	"html/template"
	"io"

	"github.com/labstack/echo/v4"
	"github.com/microcosm-cc/bluemonday"
)

//Template wrapper
type Template struct {
	templates *template.Template
}

//Render template
func (t *Template) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	return t.templates.ExecuteTemplate(w, name, data)
}

func main() {
	t := &Template{
		templates: template.Must(template.ParseGlob("views/*.html")),
	}
	p := bluemonday.UGCPolicy()
	e := echo.New()
	e.Renderer = t
	e.Static("/public", "public")
	e.GET("/", func(c echo.Context) error {
		return c.Render(
			200,
			"index.html",
			map[string]interface{}{
				"navtitle": "open lemon blog",
				"gbody":    "",
				"ghead":    "",
				"gfooter":  "",

				"body":        template.HTML(p.Sanitize("body")),
				"tag":         []string{},
				"ogimg":       false,
				"description": "",
				"title":       "",
				"section":     "",
			},
		)
	})
	e.Logger.Fatal(e.Start(":19601"))
}
