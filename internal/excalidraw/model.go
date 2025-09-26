package excalidraw

import "encoding/json"

type Doc struct {
	Type     string                     `json:"type"`
	Version  int                        `json:"version"`
	Source   string                     `json:"source"`
	Elements []json.RawMessage          `json:"elements"`
	AppState json.RawMessage            `json:"appState"`
	Files    map[string]json.RawMessage `json:"files"`
}

type Element struct {
	Idx          int    `json:"-"`
	Day          int    `json:"-"`
	OldText      string `json:"-"`
	Type         string `json:"type"`
	Text         string `json:"text,omitempty"`
	OriginalText string `json:"originalText,omitempty"`
	Raw          map[string]any
}
