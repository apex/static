package static

import (
	"bytes"
	"io"
	"strconv"
	"strings"

	dom "github.com/PuerkitoBio/goquery"
	"github.com/alecthomas/chroma"
	"github.com/alecthomas/chroma/formatters/html"
	"github.com/alecthomas/chroma/lexers"
	"github.com/alecthomas/chroma/styles"
	snakecase "github.com/segmentio/go-snakecase"
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

		scope := []string{""}
		prev := 1

		doc.Find("h1, h2, h3, h4, h5, h6").Each(func(i int, s *dom.Selection) {
			curr, _ := strconv.Atoi(string(s.Get(0).Data[1]))
			change := curr - prev

			if change <= 0 {
				scope = scope[:len(scope)+(change-1)]
			}

			scope = append(scope, snakecase.Snakecase(s.Text()))
			prev = curr

			id := strings.Join(scope, ".")
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

// SyntaxHighlight returns a reader with HTML code prettified with chroma
func SyntaxHighlight(r io.Reader) io.ReadCloser {
	pr, pw := io.Pipe()
	go func() {
		doc, err := dom.NewDocumentFromReader(r)
		if err != nil {
			pw.CloseWithError(err)
			return
		}
		// TODO: make this configurable?
		style := styles.Get("github")
		if style == nil {
			style = styles.Fallback
		}
		formatter := html.New(html.Standalone())
		doc.Find("pre>code").Each(func(i int, s *dom.Selection) {
			lexer := detectLexer(s)
			code := s.Contents().Text()
			iterator, err := lexer.Tokenise(nil, code)
			if err != nil {
				pw.CloseWithError(err)
				return
			}
			var html bytes.Buffer
			err = formatter.Format(&html, style, iterator)
			if err != nil {
				pw.CloseWithError(err)
				return
			}

			// Parent() because chroma html is already in a <pre><code>.
			s.Parent().ReplaceWithHtml(html.String())
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

func detectLexer(s *dom.Selection) chroma.Lexer {
	var lang string
	classes, ok := s.Attr("class")
	if ok {
		for _, class := range strings.Split(classes, " ") {
			if strings.HasPrefix(class, "language-") {
				lang = strings.TrimPrefix(class, "language-")
			}
		}
	}
	var lexer chroma.Lexer
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
