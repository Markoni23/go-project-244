package fileparser

import (
	"fmt"
	"os"
	"path/filepath"
)

func printFile(path string) error {
	content, err := os.ReadFile(path)
	if err != nil {
		return err
	}

	fmt.Printf("%s", content)
	return nil
}

func ParseFiles(firstFile, secondFile string) error {
	path, err := filepath.Abs(firstFile)
	if err != nil {
		return err
	}

	if err := printFile(path); err != nil {
		return err
	}
	fmt.Println()

	path, err = filepath.Abs(secondFile)
	if err != nil {
		return err
	}
	if err := printFile(path); err != nil {
		return err
	}
	return nil
}
