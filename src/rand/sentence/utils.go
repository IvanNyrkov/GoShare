package sentence

import (
	"path/filepath"
	"runtime"
)

const (
	defaultNounsFile      = "nouns.txt"
	defaultAdjectivesFile = "adjectives.txt"
)

func getCurrentDirectory() string {
	_, currentPath, _, _ := runtime.Caller(1)
	return filepath.Dir(currentPath)
}

func getNounsFilename(filename string) string {
	if filename == "" {
		dir := getCurrentDirectory()
		return filepath.Join(dir, "files", defaultNounsFile)
	}
	return filename
}

func getAdjectivesFilename(filename string) string {
	if filename == "" {
		dir := getCurrentDirectory()
		return filepath.Join(dir, "files", defaultAdjectivesFile)
	}
	return filename
}
