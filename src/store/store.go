package store

import (
	"io"
	"mime/multipart"
	"os"
	"path/filepath"
	"runtime"
)

var (
	uploadDir      = "files/"
	fullUploadPath string
)

func init() {
	_, currentPath, _, _ := runtime.Caller(1)
	fullUploadPath = filepath.Join(filepath.Dir(currentPath), uploadDir)
}

// SaveFile saves file in path
func SaveFile(fileName string, file multipart.File) (err error) {
	filePath := filepath.Join(fullUploadPath, fileName)
	f, err := os.OpenFile(filePath, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		return err
	}
	defer f.Close()
	io.Copy(f, file)
	return
}

// DaemonFileCleaner deletes all old files
func DaemonFileCleaner() {
	// Get list of files market Deleted from db

	// Get list of files in folder

	// Remove all files from folder which is in db-list
}
