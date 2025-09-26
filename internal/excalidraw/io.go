package excalidraw

import (
	"encoding/json"
	"fmt"
	"io"
	"io/fs"
	"os"
	"path/filepath"
)

func Read(fsys fs.FS, path string) (*Doc, error) {
	rc, err := fsys.Open(path)
	if err != nil {
		return nil, fmt.Errorf("open input %q: %w", path, err)
	}
	defer rc.Close()

	data, err := io.ReadAll(rc)
	if err != nil {
		return nil, fmt.Errorf("read input %q: %w", path, err)
	}

	var f Doc
	if err := json.Unmarshal(data, &f); err != nil {
		return nil, fmt.Errorf("parse JSON %q: %w", path, err)
	}

	return &f, nil
}

func Write(base string, path string, f *Doc) error {
	if err := os.MkdirAll(base, 0o755); err != nil {
		return fmt.Errorf("ensure docs dir: %w", err)
	}
	clean := filepath.Clean(path)
	full := filepath.Join(base, clean)

	out, err := json.MarshalIndent(f, "", "  ")
	if err != nil {
		return fmt.Errorf("marshal output: %w", err)
	}

	if err := os.WriteFile(full, out, 0o644); err != nil {
		return fmt.Errorf("write %q: %w", full, err)
	}
	return nil
}
