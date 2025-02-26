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

type PostRenderer struct {
	tmpl *template.Template
}

func NewPostRenderer() (*PostRenderer, error) {
	tmpl, err := template.ParseFS(postTemplates, "templates/*.tmpl.html")
	if err != nil {
		return nil, err
	}

	return &PostRenderer{tmpl: tmpl}, nil
}

func (r *PostRenderer) Render(w io.Writer, p Post) error {
	if err := r.tmpl.ExecuteTemplate(w, "blog.tmpl.html", p); err != nil {
		return err
	}

	return nil
}
