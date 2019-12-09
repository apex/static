package static

import (
	"bytes"
	"io"
	"strings"

	dom "github.com/PuerkitoBio/goquery"
	"github.com/alecthomas/chroma"
	"github.com/alecthomas/chroma/formatters/html"
	"github.com/alecthomas/chroma/lexers"
)

// blankStyle is a blank style.
var blankStyle = chroma.MustNewStyle("blank", chroma.StyleEntries{})

// SyntaxHighlight returns a reader with HTML code prettified with chroma.
func SyntaxHighlight(r io.Reader) io.ReadCloser {
	pr, pw := io.Pipe()

	go func() {
		doc, err := dom.NewDocumentFromReader(r)
		if err != nil {
			pw.CloseWithError(err)
			return
		}

		formatter := html.New(html.WithClasses(true))
		doc.Find("pre > code").Each(func(i int, s *dom.Selection) {
			lexer := detectLexer(s)
			code := s.Contents().Text()

			iterator, err := lexer.Tokenise(nil, code)
			if err != nil {
				pw.CloseWithError(err)
				return
			}

			var buf bytes.Buffer
			err = formatter.Format(&buf, blankStyle, iterator)
			if err != nil {
				pw.CloseWithError(err)
				return
			}

			// Parent() because chroma html is already in a <pre><code>.
			s.Parent().ReplaceWithHtml(buf.String())
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

// detectLexer returns a chroma lexer based on the classname.
func detectLexer(s *dom.Selection) chroma.Lexer {
	var lexer chroma.Lexer
	var lang string

	if classes, ok := s.Attr("class"); ok {
		for _, c := range strings.Split(classes, " ") {
			if strings.HasPrefix(c, "language-") {
				lang = strings.TrimPrefix(c, "language-")
			}
		}
	}

	if lang != "" {
		lexer = lexers.Get(lang)
	}

	if lexer == nil {
		lexer = lexers.Analyse(s.Contents().Text())
	}

	if lexer == nil {
		lexer = lexers.Fallback
	}

	return chroma.Coalesce(lexer)
}
