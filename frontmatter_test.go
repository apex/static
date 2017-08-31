package static

import (
	"io/ioutil"
	"strings"
	"testing"

	"github.com/tj/assert"
)

func TestFrontmatter(t *testing.T) {
	t.Run("valid", func(t *testing.T) {
		in := strings.NewReader(`---
title: Something
---

Hello __World__.
    `)

		var meta struct {
			Title string `yml:"title"`
		}

		out := Markdown(Frontmatter(in, &meta))
		b, err := ioutil.ReadAll(out)
		assert.NoError(t, err, "reading")
		assert.NoError(t, out.Close())
		assert.Equal(t, "<p>Hello <strong>World</strong>.</p>\n", string(b))
	})

	t.Run("read error", func(t *testing.T) {
		in := &errorReader{}

		var meta struct {
			Title string `yml:"title"`
		}

		out := Markdown(Frontmatter(in, &meta))

		_, err := ioutil.ReadAll(out)
		assert.EqualError(t, err, `boom`)
	})
}
