package renderer_test

import (
	"bytes"
	"io"
	"testing"

	approvals "github.com/approvals/go-approval-tests"
	renderer "github.com/thrashwerk/learn-go-with-tests-code/20-html-templates"
)

func TestRender(t *testing.T) {
	var (
		aPost = renderer.Post{
			Title:       "hello world",
			Body:        "This is a post",
			Description: "This is a description",
			Tags:        []string{"go", "tdd"},
		}
	)

	postRenderer, err := renderer.NewPostRenderer()
	if err != nil {
		t.Fatal(err)
	}

	t.Run("convert single post to HTML", func(t *testing.T) {
		buf := bytes.Buffer{}
		err := postRenderer.Render(&buf, aPost)

		if err != nil {
			t.Fatal(err)
		}

		approvals.VerifyString(t, buf.String())
	})

	t.Run("renders index of posts", func(t *testing.T) {
		buf := bytes.Buffer{}
		posts := []renderer.Post{{Title: "Hello world"}, {Title: "Hello world 2"}}

		if err := postRenderer.RenderIndex(&buf, posts); err != nil {
			t.Fatal(err)
		}

		approvals.VerifyString(t, buf.String())
	})
}

func BenchmarkRender(b *testing.B) {
	var (
		aPost = renderer.Post{
			Title:       "hello world",
			Body:        "This is a post",
			Description: "This is a description",
			Tags:        []string{"go", "tdd"},
		}
	)

	postRenderer, err := renderer.NewPostRenderer()
	if err != nil {
		b.Fatal(err)
	}

	b.ResetTimer()
	// for i := 0; i < b.N; i++ {
	// 	renderer.Render(io.Discard, aPost)
	// }
	for b.Loop() { // introduced in Go 1.24 for use instead of b.N
		postRenderer.Render(io.Discard, aPost)
	}
}
