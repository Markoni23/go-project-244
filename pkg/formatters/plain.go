package formatters

import (
	"code/pkg/differ"
	"fmt"
	"strings"
)

type PlainFormatter struct{}

func (p PlainFormatter) Format(diff *differ.Diff) (string, error) {
	return p.format("", diff)
}

func (p PlainFormatter) format(path string, diff *differ.Diff) (string, error) {
	var lines []string
	var currentPath, nodeRes string
	var err error

	for _, key := range diff.Keys() {
		node := diff.GetNode(key)

		if path != "" {
			currentPath = path + "." + key
		} else {
			currentPath = key
		}

		if node.Type == differ.OBJECT_TYPE && isDiff(node.OriginalValue) {
			nestedDiff := node.OriginalValue.(*differ.Diff)
			nodeRes, err = p.format(currentPath, nestedDiff)

			if err != nil {
				return "", err
			}

		} else {
			switch node.Status {
			case differ.ADDED:
				nodeRes = fmt.Sprintf("Property '%s' was added with value: %s", currentPath, p.formatValue(node.NewValue))
			case differ.REMOVED:
				nodeRes = fmt.Sprintf("Property '%s' was removed", currentPath)
			case differ.CHANGED:
				nodeRes = fmt.Sprintf("Property '%s' was updated. From %s to %s", currentPath, p.formatValue(node.OriginalValue), p.formatValue(node.NewValue))
			case differ.SAME:
				continue
			}
		}

		if nodeRes != "" {
			lines = append(lines, nodeRes)
		}
	}

	return strings.Join(lines, "\n"), nil
}

func (p PlainFormatter) formatValue(value any) string {
	switch val := value.(type) {
	case map[string]any:
		return "[complex value]"
	case nil:
		return "null"
	case string:
		return fmt.Sprintf("'%s'", val)
	default:
		return fmt.Sprintf("%v", val)
	}
}
