package static

import (
	"io"
	"io/ioutil"

	"github.com/tj/front"
)

// Frontmatter returns a reader of contents and unmarshals frontmatter to v.
func Frontmatter(r io.Reader, v interface{}) io.ReadCloser {
	pr, pw := io.Pipe()

	go func() {
		b, err := ioutil.ReadAll(r)
		if err != nil {
			pw.CloseWithError(err)
			return
		}

		s, err := front.Unmarshal(b, v)
		if err != nil {
			pw.CloseWithError(err)
			return
		}

		pw.Write(s)
		pw.Close()
	}()

	return pr
}
