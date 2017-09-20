package static

import (
	"io/ioutil"
	"strings"
	"testing"

	"github.com/tj/assert"
)

func TestHeadingAnchors(t *testing.T) {
	in := ioutil.NopCloser(strings.NewReader(`<html><head></head><body>
	<h1>One</h1>
	<h2>Two</h2>
  <h3>Three</h3>
</body></html>`))

	in = HeadingAnchors(in)

	b, err := ioutil.ReadAll(in)
	assert.NoError(t, err, "reading")

	expected := `<html><head></head><body>
	<h1 id="one">One</h1>
	<h2 id="two">Two</h2>
  <h3 id="three">Three</h3>
</body></html>`

	assert.Equal(t, expected, string(b))
}
