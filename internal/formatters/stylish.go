package formatters

import (
	"code/internal/differ"
	"fmt"
	"sort"
	"strings"
)

type StylishFormatter struct{}

func (s StylishFormatter) Format(diff *differ.Diff) (string, error) {
	return s.formatDiff(diff, 0)
}

func (s StylishFormatter) formatDiff(diff *differ.Diff, level int) (string, error) {
	var lines []string
	var line string

	keys := differ.SortedKeys(diff)

	for _, key := range keys {
		node := diff.GetNode(key)

		origVal, err := s.formatValue(node.OriginalValue, level)
		if err != nil {
			return "", err
		}
		newVal, err := s.formatValue(node.NewValue, level)
		if err != nil {
			return "", err
		}
		indent := strings.Repeat(" ", ((level+1)*4)-2)
		switch node.Status {
		case differ.SAME:
			line = fmt.Sprintf("%s  %s: %v", indent, node.FieldName, origVal)
		case differ.ADDED:
			line = fmt.Sprintf("%s+ %s: %v", indent, node.FieldName, newVal)
		case differ.REMOVED:
			line = fmt.Sprintf("%s- %s: %v", indent, node.FieldName, origVal)
		case differ.CHANGED:
			line = fmt.Sprintf("%s- %s: %v\n%s+ %s: %v", indent, node.FieldName, origVal, indent, node.FieldName, newVal)
		}
		lines = append(lines, line)

	}
	end := strings.Repeat("    ", level) + "}"
	res := fmt.Sprintf("{\n%s\n%s", strings.Join(lines, "\n"), end)
	return res, nil
}

func (s StylishFormatter) formatValue(value any, level int) (string, error) {
	switch val := value.(type) {
	case *differ.Diff:
		return s.formatDiff(val, level+1)
	case map[string]any:
		return s.formatMap(val, level+1)
	case nil:
		return "null", nil
	default:
		return fmt.Sprintf("%v", val), nil
	}
}

func (s StylishFormatter) formatMap(val map[string]any, depth int) (string, error) {
	keys := make([]string, 0)

	for key := range val {
		keys = append(keys, key)
	}

	sort.Slice(keys, func(i int, j int) bool {
		return keys[i] < keys[j]
	})

	var lines []string

	for _, key := range keys {
		value := val[key]
		formatted, err := s.formatValue(value, depth)
		if err != nil {
			return "", err
		}
		line := strings.Repeat("    ", depth+1) + fmt.Sprintf("%s: %s", key, formatted)
		lines = append(lines, line)
	}

	end := strings.Repeat("    ", depth) + "}"

	result := fmt.Sprintf("{\n%s\n%s", strings.Join(lines, "\n"), end)

	return result, nil
}
