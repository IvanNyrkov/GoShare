package store

import (
	"path/filepath"
	"runtime"
)

var (
	defaultUploadDir = filepath.Join(getCurrentDir(), "files/")
)

func getCurrentDir() string {
	_, currentPath, _, _ := runtime.Caller(1)
	return filepath.Dir(currentPath)
}
