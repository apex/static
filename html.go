package static

import (
	"io"

	"github.com/PuerkitoBio/goquery"
	"github.com/segmentio/go-snakecase"
)

// HeadingAnchors returns a reader with
// HTML heading ids derived from their text.
func HeadingAnchors(r io.Reader) io.ReadCloser {
	pr, pw := io.Pipe()

	go func() {
		doc, err := goquery.NewDocumentFromReader(r)
		if err != nil {
			pw.CloseWithError(err)
			return
		}

		doc.Find("h1, h2, h3, h4, h5, h6").Each(func(i int, s *goquery.Selection) {
			s.SetAttr("id", snakecase.Snakecase(s.Text()))
		})

		html, err := doc.Html()
		if err != nil {
			pw.CloseWithError(err)
			return
		}

		pw.Write([]byte(html))
		pw.Close()
	}()

	return pr
}
