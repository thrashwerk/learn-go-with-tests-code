package renderer_test

import (
	"bytes"
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

	t.Run("convert single post to HTML", func(t *testing.T) {
		buf := bytes.Buffer{}
		err := renderer.Render(&buf, aPost)

		if err != nil {
			t.Fatal(err)
		}

		approvals.VerifyString(t, buf.String())
	})
}
