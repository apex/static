package main

import (
	"os"
	"time"

	"github.com/apex/log"
	"github.com/apex/log/handlers/cli"
	"github.com/tj/kingpin"

	"github.com/apex/static/docs"
)

func init() {
	log.SetHandler(cli.Default)
}

func main() {
	app := kingpin.New("static-docs", "Generate static documentation sites")
	app.Example("static-docs --title Up --in ./docs", "Generate site in ./build from ./docs/*.md")
	app.Example("static-docs --title Up --in ../up/docs -out .", "Generate a site in ./ from ../up/docs/*.md")
	title := app.Flag("title", "Site title.").String()
	subtitle := app.Flag("subtitle", "Site subtitle or slogan.").String()
	theme := app.Flag("theme", "Theme name.").Default("apex").String()
	src := app.Flag("in", "Source directory for markdown files.").Default(".").String()
	dst := app.Flag("out", "Output directory for the static site.").Default("build").String()
	segment := app.Flag("segment", "Segment write key.").String()
	google := app.Flag("google", "Google Analytics tracking id.").String()
	kingpin.MustParse(app.Parse(os.Args[1:]))

	println()
	defer println()

	start := time.Now()

	c := &docs.Config{
		Src:      *src,
		Dst:      *dst,
		Title:    *title,
		Subtitle: *subtitle,
		Theme:    *theme,
		Segment:  *segment,
		Google:   *google,
	}

	if err := docs.Compile(c); err != nil {
		log.Fatalf("error: %s", err)
	}

	log.Infof("compiled in %s\n", time.Since(start).Round(time.Millisecond))
}
