package renderer

import (
	"embed"
	"html/template"
	"io"
	"strings"
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

func (p Post) SanitizedTitle() string {
	return strings.ToLower(strings.Replace(p.Title, " ", "-", -1))
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
	return r.tmpl.ExecuteTemplate(w, "blog.tmpl.html", p)
}

func (r *PostRenderer) RenderIndex(w io.Writer, posts []Post) error {
	return r.tmpl.ExecuteTemplate(w, "index.tmpl.html", posts)
}

// type PostViewModel struct {
// 	Title          string
// 	SanitizedTitle string
// 	Description    string
// 	Body           string
// 	Tags           []string
// }
