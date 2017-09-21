package static

import (
	"io"
	"strings"

	dom "github.com/PuerkitoBio/goquery"
	"github.com/segmentio/go-snakecase"
)

// anchorHTML is the anchor element and SVG icon.
var anchorHTML = `<a class="Anchor" aria-hidden="true">
	<svg xmlns="http://www.w3.org/2000/svg" aria-hidden="true" width="14" height="14" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round" class="feather feather-link"><path d="M10 13a5 5 0 0 0 7.54.54l3-3a5 5 0 0 0-7.07-7.07l-1.72 1.71"></path><path d="M14 11a5 5 0 0 0-7.54-.54l-3 3a5 5 0 0 0 7.07 7.07l1.71-1.71"></path></svg>
</a>`

// anchorEl is the instance of the anchor element which is cloned.
var anchorEl *dom.Selection

// initialize anchor element.
func init() {
	doc, err := dom.NewDocumentFromReader(strings.NewReader(anchorHTML))
	if err != nil {
		panic(err)
	}

	anchorEl = doc.Find("a")
}

// HeadingAnchors returns a reader with
// HTML heading ids derived from their text.
func HeadingAnchors(r io.Reader) io.ReadCloser {
	pr, pw := io.Pipe()

	go func() {
		doc, err := dom.NewDocumentFromReader(r)
		if err != nil {
			pw.CloseWithError(err)
			return
		}

		var section string
		doc.Find("h1, h2, h3, h4, h5, h6").Each(func(i int, s *dom.Selection) {
			id := snakecase.Snakecase(s.Text())

			if s.Is("h1") {
				section = id
			}

			id = section + "__" + id
			a := anchorEl.Clone()
			a.SetAttr("id", id)
			a.SetAttr("href", "#"+id)
			s.SetAttr("class", "Heading")
			s.PrependSelection(a)
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
