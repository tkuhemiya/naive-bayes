package main

import (
	"os"
	"path/filepath"
	"regexp"
	"strings"
)

func FileToString(path string) ([]string, error) {
	filePaths, err := filepath.Glob(path)
	if err != nil {
		return nil, err
	}
	fileContent := make([]string, len(filePaths))

	for i, filePath := range filePaths {
		file, err := os.ReadFile(filePath)
		if err != nil {
			return fileContent, err
		}
		content := string(file)
		fileContent[i] = content
	}

	return fileContent, nil
}

var nonWord = regexp.MustCompile(`[^A-Za-z0-9]+`)

func Tokenize(text string) []string {
	raw := strings.Fields(text)
	out := make([]string, 0, len(raw))
	for _, r := range raw {
		t := strings.ToUpper(r)
		t = nonWord.ReplaceAllString(t, "")
		if t != "" {
			out = append(out, t)
		}
	}
	return out
}
