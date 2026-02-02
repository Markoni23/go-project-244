package formatters

import (
	"code/internal/differ"
	"encoding/json"
	"fmt"
	"strings"
)

type JSONFormatter struct{}

func (j JSONFormatter) Format(diff *differ.Diff) (string, error) {
	var lines []string
	var line string

	keys := differ.SortedKeys(diff)

	for _, key := range keys {
		node := diff.GetNode(key)
		if node.Type == differ.OBJECT_TYPE && isDiff(node.OriginalValue) {
			nestedRes, err := j.Format(node.OriginalValue.(*differ.Diff))
			if err != nil {
				return "", err
			}

			line := fmt.Sprintf("\"%s\": %s", node.FieldName, nestedRes)
			lines = append(lines, line)
		}

		origVal, err := json.Marshal(node.OriginalValue)
		if err != nil {
			return "", err
		}

		newVal, err := json.Marshal(node.NewValue)
		if err != nil {
			return "", err
		}

		switch node.Status {
		case differ.SAME:
			line = fmt.Sprintf("\"  %s\": %v", node.FieldName, string(origVal))
		case differ.ADDED:
			line = fmt.Sprintf("\"+ %s\": %v", node.FieldName, string(newVal))
		case differ.REMOVED:
			line = fmt.Sprintf("\"- %s\": %v", node.FieldName, string(origVal))
		case differ.CHANGED:
			line = fmt.Sprintf("\"- %s\": %v,\n\"+ %s\": %v", node.FieldName, string(origVal), node.FieldName, string(newVal))
		}
		lines = append(lines, line)

	}

	res := fmt.Sprintf("{\n%s\n}", strings.Join(lines, ",\n"))
	return res, nil
}
