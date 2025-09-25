package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"

	"github.com/alfianchii/rs-nihongo-notes/internal/models"
)

const DOC_PATH = "./docs/"
const OUTPUT_NAME = "new.excalidraw"

func main() {
	var (
		inPath       string
		outPath      string
		startedCount int
	)

	flag.StringVar(&inPath, "in", "", "Input Excalidraw JSON file (.excalidraw) from ./docs")
	flag.IntVar(&startedCount, "start", 1, "Starting day count (e.g., 9)")
	flag.Parse()

	if inPath == "" {
		fmt.Fprintln(os.Stderr, "error: -in is required")
		os.Exit(2)
	}
	inPath = DOC_PATH + inPath
	outPath = DOC_PATH + OUTPUT_NAME

	data, err := os.ReadFile(inPath)
	if err != nil {
		fmt.Fprintf(os.Stderr, "read error: %v\n", err)
		os.Exit(1)
	}

	var excali models.Excalidraw
	json.Unmarshal(data, &excali)

	var filteredElements []models.ExcalidrawElement
	for idx, raw := range excali.Elements {
		var el models.ExcalidrawElement
		_ = json.Unmarshal(raw, &el)

		if el.Type == "text" && strings.HasPrefix(el.Text, "Day ") {
			el.Idx = idx
			filteredElements = append(filteredElements, el)
		}
	}

	sort.Slice(filteredElements, func(i, j int) bool {
		currEl := strings.Fields(filteredElements[i].Text)[1]
		nextEl := strings.Fields(filteredElements[j].Text)[1]

		currElDay, _ := strconv.Atoi(currEl)
		nextElDay, _ := strconv.Atoi(nextEl)

		return currElDay < nextElDay
	})

	for idx, element := range filteredElements {
		item := strings.Fields(element.Text)
		if len(item) >= 2 {
			item[1] = fmt.Sprintf("%d", startedCount)
			element.Text = strings.Join(item, " ")

			filteredElements[idx] = element
			startedCount++
		}
	}

	for i, excaliItem := range excali.Elements {
		for _, filteredEl := range filteredElements {
			if i == filteredEl.Idx {
				var m map[string]any
				json.Unmarshal(excaliItem, &m)

				m["text"] = filteredEl.Text
				m["originalText"] = filteredEl.OriginalText

				b, _ := json.Marshal(m)
				excali.Elements[i] = b
			}
		}
	}

	out, _ := json.MarshalIndent(excali, "", "  ")
	os.WriteFile(outPath, out, 0644)
}
