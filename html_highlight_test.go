package static

import (
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
	"testing"

	"github.com/tj/assert"
)

func fixture(t testing.TB, name string) io.ReadCloser {
	path := filepath.Join("testdata", name)
	f, err := os.Open(path)
	assert.NoError(t, err, "open")
	return f
}

func TestSyntaxHighlight(t *testing.T) {
	in := fixture(t, "highlight_input.html")
	out := fixture(t, "highlight_output.html")

	got, _ := ioutil.ReadAll(SyntaxHighlight(in))
	expect, _ := ioutil.ReadAll(out)

	if string(got) != string(expect) {
		t.Errorf("\nExpected:\n\n%s\n\nGot:\n\n%s", string(expect), string(got))
		// ioutil.WriteFile("testdata/highlight_output.html", got, 0644)
	}
}
