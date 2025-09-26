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
	var elements []excalidraw.Element
	for idx, raw := range f.Elements {
		var element excalidraw.Element
		if err := json.Unmarshal(raw, &element); err != nil {
			return nil, fmt.Errorf("parse element %d: %w", idx, err)
		}

		if day, ok, err := GetDay(&element); err != nil {
			return nil, err
		} else if ok {
			element.Idx = idx
			element.Day = day
			elements = append(elements, element)
		}
	}

	sort.SliceStable(elements, func(i, j int) bool {
		if elements[i].Day == elements[j].Day {
			return elements[i].Idx < elements[j].Idx
		}

		return elements[i].Day < elements[j].Day
	})

	startAt := opt.StartAt
	for idx, element := range elements {
		var exElement map[string]any
		if err := json.Unmarshal(f.Elements[element.Idx], &exElement); err != nil {
			return nil, fmt.Errorf("reparse element %d: %w", elements[idx].Idx, err)
		}

		renumberedDay := reDay.ReplaceAllString(element.Text, fmt.Sprintf("Day %d", startAt))
		exElement["text"] = renumberedDay
		exElement["originalText"] = renumberedDay
		elements[idx].OldText = element.Text
		elements[idx].Text = renumberedDay

		byteExElement, err := json.Marshal(exElement)
		if err != nil {
			return nil, fmt.Errorf("remarshal element %d: %w", elements[idx].Idx, err)
		}
		f.Elements[element.Idx] = byteExElement

		startAt++
	}

	return elements, nil
}
