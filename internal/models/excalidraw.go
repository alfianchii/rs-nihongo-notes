package models

import "encoding/json"

type Excalidraw struct {
	Type     string                     `json:"type"`
	Version  int                        `json:"version"`
	Source   string                     `json:"source"`
	Elements []json.RawMessage          `json:"elements"`
	AppState json.RawMessage            `json:"appState"`
	Files    map[string]json.RawMessage `json:"files"`
}

type ExcalidrawElement struct {
	ID           string   `json:"id"`
	Idx          int      `json:"idx"`
	Type         string   `json:"type"`
	Text         string   `json:"text,omitempty"`
	OriginalText string   `json:"originalText,omitempty"`
	X            *float64 `json:"x"`
	Y            *float64 `json:"y"`
}
