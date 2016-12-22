package sentence

import (
	"path/filepath"
	"runtime"
)

const (
	nounsFileName      = "default_nouns.txt"
	adjectivesFileName = "default_adjectives.txt"
)

var (
	defaultNounsFilePath      = filepath.Join(getCurrentDir(), nounsFileName)
	defaultAdjectivesFilePath = filepath.Join(getCurrentDir(), adjectivesFileName)
)

func getCurrentDir() string {
	_, currentPath, _, _ := runtime.Caller(1)
	return filepath.Dir(currentPath)
}
