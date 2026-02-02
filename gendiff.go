package code

import (
	"code/internal/differ"
	"code/internal/formatters"
	"code/internal/parsers"
	"os"
	"path/filepath"
)

func GenDiff(firstFile, secondFile string, format string) (string, error) {
	firstFileMap, err := ParseFile(firstFile)
	if err != nil {
		return "", err
	}

	secondFileMap, err := ParseFile(secondFile)
	if err != nil {
		return "", err
	}

	diff, err := differ.CalculateDiffTree(firstFileMap, secondFileMap)
	if err != nil {
		return "", err
	}

	res, err := FormatDiff(diff, format)
	if err != nil {
		return "", err
	}
	return res, nil
}

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

	extension := filepath.Ext(path)[1:]

	parser, err := parsers.ParserFabric{}.CreateParser(extension)

	if err != nil {
		return res, err
	}

	res, err = parser.Parse(content)

	if err != nil {
		return res, err
	}

	return res, nil
}

func FormatDiff(diff *differ.Diff, format string) (string, error) {
	formatter, err := formatters.FormatterFabric{}.CreateFormatter(format)
	if err != nil {
		return "", err
	}

	res, err := formatter.Format(diff)

	if err != nil {
		return "", err
	}

	return res, nil
}
