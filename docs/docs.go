package docs

import (
	"bytes"
	"html/template"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"

	"github.com/apex/log"
	"github.com/pkg/errors"
	"github.com/segmentio/go-snakecase"

	"github.com/apex/static"
	"github.com/apex/static/docs/themes/apex"
	"github.com/apex/static/inject"
)

// Config options.
type Config struct {
	// Src dir of markdown documentation files.
	Src string

	// Dst dir of the static site.
	Dst string

	// Title of the site.
	Title string

	// Subtitle of the site.
	Subtitle string

	// Theme name.
	Theme string

	// Segment write key.
	Segment string

	// Google Analytics tracking id.
	Google string
}

// Page model.
type Page struct {
	// Title of the page.
	Title string `yml:"title"`

	// Slug of the page.
	Slug string `yml:"slug"`

	// Skip the page.
	Skip bool `yml:"skip"`

	// Content of the page.
	Content template.HTML
}

// Compile docs site.
func Compile(c *Config) error {
	log.Infof("compiling %s to %s", c.Src, c.Dst)

	if err := os.MkdirAll(c.Dst, 0755); err != nil {
		return errors.Wrap(err, "mkdir")
	}

	if err := initTheme(c); err != nil {
		return errors.Wrap(err, "initializing theme")
	}

	files, err := ioutil.ReadDir(c.Src)
	if err != nil {
		return errors.Wrap(err, "reading dir")
	}

	var pages []*Page

	for _, f := range files {
		path := filepath.Join(c.Src, f.Name())

		log.Infof("compiling %q", path)
		p, err := compile(c, path)
		if err != nil {
			return errors.Wrapf(err, "compiling %q", path)
		}

		if p == nil {
			log.Infof("skipping %q", path)
			continue
		}

		pages = append(pages, p)
	}

	var buf bytes.Buffer

	if err := render(&buf, c, pages); err != nil {
		return errors.Wrap(err, "rendering")
	}

	html := buf.String()

	if c.Segment != "" {
		html = inject.Head(html, inject.Segment(c.Segment))
	}

	if c.Google != "" {
		html = inject.Head(html, inject.GoogleAnalytics(c.Google))
	}

	out := filepath.Join(c.Dst, "index.html")
	if err := ioutil.WriteFile(out, []byte(html), 0755); err != nil {
		return errors.Wrap(err, "writing")
	}

	return nil
}

// render to writer w.
func render(w io.Writer, c *Config, pages []*Page) error {
	path := filepath.Join(c.Dst, "theme", c.Theme, "views", "*.html")

	views, err := template.ParseGlob(path)
	if err != nil {
		return errors.Wrap(err, "parsing templates")
	}

	err = views.ExecuteTemplate(w, "index.html", struct {
		*Config
		Pages []*Page
	}{
		Config: c,
		Pages:  pages,
	})

	if err != nil {
		return errors.Wrap(err, "rendering")
	}

	return nil
}

// compile file.
func compile(c *Config, path string) (*Page, error) {
	// open
	f, err := os.Open(path)
	if err != nil {
		return nil, errors.Wrap(err, "open")
	}
	defer f.Close()

	// meta-data
	var page Page
	rc := static.Markdown(static.Frontmatter(f, &page))

	// contents
	b, err := ioutil.ReadAll(rc)
	if err != nil {
		rc.Close()
		return nil, errors.Wrap(err, "reading")
	}

	if err := rc.Close(); err != nil {
		return nil, errors.Wrap(err, "closing")
	}

	// populate
	if page.Slug == "" {
		page.Slug = snakecase.Snakecase(page.Title)
	}
	page.Content = template.HTML(b)

	// skip
	if page.Skip {
		return nil, nil
	}

	return &page, nil
}

// initTheme populates the theme directory unless present.
func initTheme(c *Config) error {
	dir := filepath.Join(c.Dst, "theme", c.Theme)
	_, err := os.Stat(dir)

	if os.IsNotExist(err) {
		return apex.RestoreAssets(dir, "")
	}

	return nil
}

// stripExt returns the path sans-extname.
func stripExt(s string) string {
	return strings.Replace(s, filepath.Ext(s), "", 1)
}
