package app

import (
	"fmt"
	"io/fs"
	"path/filepath"

	"github.com/alfianchii/rs-nihongo-notes/internal/cli"
	"github.com/alfianchii/rs-nihongo-notes/internal/excalidraw"
	"github.com/alfianchii/rs-nihongo-notes/internal/excalidraw/services"
)

func Run(fsys fs.FS, opts cli.Options) error {
	file, err := excalidraw.Read(fsys, opts.Input)
	if err != nil { return err }

	exElements, err := services.RenumberDays(file, services.RenumberDayOptions{StartAt: opts.StartAt})
	if err != nil { return err }
	if len(exElements) == 0 { return fmt.Errorf("no matching elements") }

	if opts.DryRun {
		for _, el := range exElements {
			fmt.Printf("[%d] %q => %q\n", el.Idx, el.OldText, el.Text)
		}
	}

	if err := excalidraw.Write(opts.DocsRoot, opts.Output, file); err != nil {
		return err
	}
	fmt.Printf("Updated file written to %s\n", filepath.Join(opts.DocsRoot, opts.Output))
	return nil
}
