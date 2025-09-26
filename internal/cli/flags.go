package cli

import (
	"flag"
	"fmt"
)

type Options struct {
	Input    string
	Output   string
	StartAt  int
	DocsRoot string
	DryRun   bool
}

func Parse() (Options, error) {
	var o Options
	flag.StringVar(&o.DocsRoot, "docs", "./docs", "Docs directory")
	flag.BoolVar(&o.DryRun, "dry-run", false, "Preview output to stdout without writing file")
	flag.StringVar(&o.Input, "f", "", "Input Excalidraw JSON file (relative to ./docs)")
	flag.IntVar(&o.StartAt, "s", 1, "Starting day count (>=1)")
	flag.StringVar(&o.Output, "o", "RSN.excalidraw", "Output filename (relative to ./docs)")

	flag.Parse()

	if o.Input == "" {
		return o, fmt.Errorf("-in is required")
	}
	if o.StartAt < 1 {
		return o, fmt.Errorf("-s must be >= 1")
	}
	return o, nil
}
