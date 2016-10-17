package dictionary

import (
	"path/filepath"
	"runtime"
)

const (
	defaultNounsFile      = "nouns.txt"
	defaultAdjectivesFile = "adjectives.txt"
)

// GetNounsFilename returns path to the specified file or to the default nouns file
func GetNounsFilename(filename string) string {
	if filename == "" {
		dir := getCurrentDirectory()
		return filepath.Join(dir, defaultNounsFile)
	}
	return filename
}

// GetAdjectivesFilename returns path to the specified file or to the default adjectives file
func GetAdjectivesFilename(filename string) string {
	if filename == "" {
		dir := getCurrentDirectory()
		return filepath.Join(dir, defaultAdjectivesFile)
	}
	return filename
}

func getCurrentDirectory() string {
	_, currentPath, _, _ := runtime.Caller(1)
	return filepath.Dir(currentPath)
}
