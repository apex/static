package static

import (
	"io/ioutil"
	"strings"
	"testing"

	"github.com/tj/assert"
)

func TestHeadingAnchors(t *testing.T) {
	in := ioutil.NopCloser(strings.NewReader(`<h1>One</h1>
<h2>Two</h2>
<h3>Three</h3>`))

	in = HeadingAnchors(in)

	b, err := ioutil.ReadAll(in)
	assert.NoError(t, err, "reading")

	expected := `<h1 id="one">One</h1>
<h2 id="two">Two</h2>
<h3 id="three">Three</h3>`

	assert.Equal(t, expected, string(b))
}
