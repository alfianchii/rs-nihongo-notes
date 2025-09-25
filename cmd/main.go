package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"sort"
	"strconv"

	"github.com/alfianchii/rs-nihongo-notes/internal/models"
	u "github.com/alfianchii/rs-nihongo-notes/internal/utils"
)

const (
	docPath        = "./docs"
	defaultOutName = "RSN.excalidraw"
)

var (
	reDay = regexp.MustCompile(`^Day\s+(\d+)\b`)
)

func main() {
	var (
		targetedPath string
		startAt      int
	)

	flag.StringVar(&targetedPath, "f", "", "Input Excalidraw JSON file (.excalidraw) from ./docs")
	flag.IntVar(&startAt, "s", 1, "Starting day count (e.g., 9)")
	flag.Parse()

	if targetedPath == "" {
		u.Fatal("-f is required")
	}
	if startAt < 1 {
		u.Fatal("-s must be >= 1")
	}

	inPath := filepath.Join(docPath, targetedPath)
	outPath := filepath.Join(docPath, defaultOutName)

	data, err := os.ReadFile(inPath)
	u.Must(err, "read input")

	var excali models.Excalidraw
	u.Must(json.Unmarshal(data, &excali), "parse input JSON")

	var elements []models.ExcalidrawElement
	for idx, raw := range excali.Elements {
		var element models.ExcalidrawElement
		u.Must(json.Unmarshal(raw, &element), "parse element ")

		if element.Type != "text" {
			continue
		}
		match := reDay.FindStringSubmatch(element.Text)
		if len(match) != 2 {
			continue
		}

		day, err := strconv.Atoi(match[1])
		u.Must(err, "parse day number")

		element.Day = day
		element.Idx = idx
		elements = append(elements, element)
	}

	sort.SliceStable(elements, func(i, j int) bool {
		if elements[i].Day == elements[j].Day {
			return elements[i].Idx < elements[j].Idx
		}

		return elements[i].Day < elements[j].Day
	})

	for idx, element := range elements {
		elements[idx].Text = reDay.ReplaceAllString(element.Text, fmt.Sprintf("Day %d", startAt))
		startAt++
	}

	for _, element := range elements {
		var exElement map[string]any
		u.Must(json.Unmarshal(excali.Elements[element.Idx], &exElement), "re-parse element")

		exElement["text"] = element.Text
		exElement["originalText"] = element.OriginalText

		byteExElement, err := json.Marshal(exElement)
		u.Must(err, "re-marshal element")
		excali.Elements[element.Idx] = byteExElement
	}

	out, err := json.MarshalIndent(excali, "", "  ")
	u.Must(err, "marshal output")
	u.Must(os.WriteFile(outPath, out, 0o644), "write output")

	fmt.Printf("Updated file written to %s\n", outPath)
}
