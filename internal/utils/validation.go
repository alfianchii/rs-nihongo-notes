package utils

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

func Must(err error, context string) {
	if err != nil {
		Fatal("%s: %v", context, err)
	}
}

func Fatal(format string, args ...any) {
	fmt.Fprintf(os.Stderr, "error: "+format+"\n", args...)
	os.Exit(1)
}

func AssertExcalidrawExt(filename *string) (error) {
	if *filename == "" {
		*filename = "RSN.excalidraw"
		return nil
	}

	if filepath.Ext(*filename) == "" {
		*filename = *filename + ".excalidraw"
		return nil
	}

	ext := strings.ToLower(filepath.Ext(*filename))
	if ext != ".excalidraw" {
		return fmt.Errorf("-o must end with .excalidraw (got %q)", *filename)
	}

	return nil
}