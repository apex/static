package static

import (
	"errors"
	"io/ioutil"
	"strings"
	"testing"

	"github.com/tj/assert"
)

type errorReader struct{}

func (r *errorReader) Read(b []byte) (int, error) {
	return 0, errors.New("boom")
}

func TestMarkdown(t *testing.T) {
	t.Run("valid", func(t *testing.T) {
		in := strings.NewReader(`Hello __World__.`)
		out := Markdown(in)
		b, err := ioutil.ReadAll(out)
		assert.NoError(t, err, "reading")
		assert.NoError(t, out.Close())
		assert.Equal(t, "<p>Hello <strong>World</strong>.</p>\n", string(b))
	})

	t.Run("read error", func(t *testing.T) {
		in := &errorReader{}
		out := Markdown(in)
		_, err := ioutil.ReadAll(out)
		assert.EqualError(t, err, `boom`)
	})
}
