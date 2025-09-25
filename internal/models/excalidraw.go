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
	Idx          int    `json:"idx,omitempty"`
	Day          int    `json:"day,omitempty"`
	Type         string `json:"type"`
	Text         string `json:"text,omitempty"`
	OriginalText string `json:"originalText,omitempty"`
}
