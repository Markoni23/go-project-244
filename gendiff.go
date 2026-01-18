package code

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"slices"
	"strings"
)

type DiffType uint8

const (
	SameValueType DiffType = iota
	ChangedValueType
	RemovedValueType
	AddedValueType
)

type FieldDiff struct {
	FieldName               string
	FirstValue, SecondValue any
	Type                    DiffType
}

func (fd FieldDiff) Print(level int) string {
	indent := strings.Repeat(" ", level)
	switch fd.Type {
	case SameValueType:
		return fmt.Sprintf("%s  %s: %v\n", indent, fd.FieldName, fd.FirstValue)
	case ChangedValueType:
		return fmt.Sprintf("%s- %s: %v\n%s+ %s: %v\n", indent, fd.FieldName, fd.FirstValue, indent, fd.FieldName, fd.SecondValue)
	case RemovedValueType:
		return fmt.Sprintf("%s- %s: %v\n", indent, fd.FieldName, fd.FirstValue)
	case AddedValueType:
		return fmt.Sprintf("%s+ %s: %v\n", indent, fd.FieldName, fd.SecondValue)
	default:
		return ""
	}
}

type FileDiff map[string]*FieldDiff

func ParseFile(fileName string) (map[string]any, error) {
	var res map[string]any
	path, err := filepath.Abs(fileName)

	if err != nil {
		return res, err
	}

	content, err := os.ReadFile(path)
	if err != nil {
		return res, err
	}

	err = json.Unmarshal(content, &res)
	if err != nil {
		return res, err
	}
	
	return res, nil
}

func GenDiff(firstFile, secondFile string) (string, error) {
	firstFileMap, err := ParseFile(firstFile)
	if err != nil {
		return "", err
	}

	secondFileMap, err := ParseFile(secondFile)
	if err != nil {
		return "", err
	}

	res := make(FileDiff, len(firstFileMap))
	keys := make([]string, 0, len(firstFileMap))

	for key, value := range firstFileMap {
		fieldDiff := FieldDiff{FieldName: key, FirstValue: value}
		if secondVal, exists := secondFileMap[key]; exists {
			fieldDiff.SecondValue = secondVal
			if secondVal == value {
				fieldDiff.Type = SameValueType
			} else {
				fieldDiff.Type = ChangedValueType
			}
		} else {
			fieldDiff.Type = RemovedValueType
		}

		res[key] = &fieldDiff
		keys = append(keys, key)
	}

	for key, value := range secondFileMap {
		if _, exists := res[key]; exists {
			continue
		}

		res[key] = &FieldDiff{FieldName: key, SecondValue: value, Type: AddedValueType}
		keys = append(keys, key)
	}

	var sb strings.Builder
	slices.Sort(keys)

	sb.WriteString("{\n")
	for _, key := range keys {
		sb.WriteString(res[key].Print(2))
	}
	sb.WriteString("}")

	return sb.String(), nil
}
