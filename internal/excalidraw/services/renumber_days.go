package services

import (
	"encoding/json"
	"fmt"
	"regexp"
	"sort"
	"strconv"

	"github.com/alfianchii/rs-nihongo-notes/internal/excalidraw"
)

var reDay = regexp.MustCompile(`^Day\s+(\d+)\b`)

type RenumberDayOptions struct {
	StartAt int
}

func GetDay(e *excalidraw.Element) (int, bool, error) {
	if e.Type != "text" {
		return 0, false, nil
	}

	match := reDay.FindStringSubmatch(e.Text)
	if len(match) != 2 {
		return 0, false, nil
	}

	day, err := strconv.Atoi(match[1])
	if err != nil {
		return 0, false, fmt.Errorf("parse day number %q: %w", match[1], err)
	}
	return day, true, nil
}

func RenumberDays(f *excalidraw.Doc, opt RenumberDayOptions) ([]excalidraw.Element, error) {
	elements := make([]excalidraw.Element, len(f.Elements))
	filteredElements := make([]excalidraw.Element, 0, len(f.Elements))

	for idx, raw := range f.Elements {
		var item map[string]any
		if err := json.Unmarshal(raw, &item); err != nil {
			return nil, fmt.Errorf("parse item %d: %w", idx, err)
		}

		var element excalidraw.Element = excalidraw.Element{
			Raw: item,
		}
		if v, ok := item["type"].(string); ok {
			element.Type = v
		}
		if v, ok := item["text"].(string); ok {
			element.Text = v
		}
		if v, ok := item["originalText"].(string); ok {
			element.OriginalText = v
		}
		elements[idx] = element

		if day, ok, err := GetDay(&element); err != nil {
			return nil, err
		} else if ok {
			filteredElements = append(filteredElements, excalidraw.Element{
				Idx:  idx,
				Day:  day,
				Type: element.Type,
				Text: element.Text,
			})
		}
	}

	sort.SliceStable(filteredElements, func(i, j int) bool {
		if filteredElements[i].Day == filteredElements[j].Day {
			return filteredElements[i].Idx < filteredElements[j].Idx
		}

		return filteredElements[i].Day < filteredElements[j].Day
	})

	startAt := opt.StartAt
	for idx, element := range filteredElements {
		exElement := elements[element.Idx].Raw
		oldText := exElement["text"].(string)
		newText := reDay.ReplaceAllString(oldText, fmt.Sprintf("Day %d", startAt))

		filteredElements[idx].OldText = oldText
		filteredElements[idx].Text = newText
		exElement["text"] = newText
		exElement["originalText"] = newText

		byteExElement, err := json.Marshal(exElement)
		if err != nil {
			return nil, fmt.Errorf("remarshal element %d: %w", element.Idx, err)
		}
		f.Elements[element.Idx] = byteExElement

		startAt++
	}

	return filteredElements, nil
}
