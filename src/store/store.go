package store

import (
	"github.com/IvanNyrkov/go-share/src/database"
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
func DaemonFileCleaner(db database.DBConnection) {
	// Delete old files from Database

	// Get list of files from db

	// Get list of files in folder

	// Remove all files from folder which not in db-files list
}
