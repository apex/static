package main

import (
	"flag"
	"time"

	"github.com/apex/log"
	"github.com/apex/log/handlers/cli"
	"github.com/tj/go/flag/usage"

	"github.com/apex/static/docs"
)

func init() {
	log.SetHandler(cli.Default)
}

func main() {
	flag.Usage = usage.Output(&usage.Config{
		Examples: []usage.Example{
			{
				Help:    "Generate site in ./build from ./docs/*.md.",
				Command: "static-docs -title Up -in ./docs",
			},
			{
				Help:    "Generate a site in ./ from ../up/docs/*.md",
				Command: "static-docs -title Up -in ../up/docs -out .",
			},
		},
	})

	title := flag.String("title", "", "Site title.")
	subtitle := flag.String("subtitle", "", "Site subtitle or slogan.")
	theme := flag.String("theme", "apex", "Theme name.")
	src := flag.String("in", ".", "Source directory for markdown files.")
	dst := flag.String("out", "build", "Output directory for the static site.")
	segment := flag.String("segment", "", "Segment write key.")
	google := flag.String("google", "", "Google Analytics tracking id.")
	flag.Parse()

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
