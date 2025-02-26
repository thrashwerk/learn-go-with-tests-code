package renderer

import (
	"embed"
	"html/template"
	"io"
)

var (
	//go:embed "templates/*"
	postTemplates embed.FS
)

type Post struct {
	Title       string
	Body        string
	Description string
	Tags        []string
}

func Render(w io.Writer, p Post) error {
	templ, err := template.ParseFS(postTemplates, "templates/*.tmpl.html")
	if err != nil {
		return err
	}

	if err := templ.ExecuteTemplate(w, "blog.tmpl.html", p); err != nil {
		return err
	}

	return nil
}
