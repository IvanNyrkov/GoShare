package store

import (
	"io"
	"mime/multipart"
	"os"
	"path/filepath"
)

// Service is an interface that defines actions for storing files
type Service interface {
	SaveFile(fileName string, file multipart.File) error
}
type serviceImpl struct {
	uploadDir string
}

// NewService creates new service for storing files
func NewService(uploadDir string) *serviceImpl {
	return &serviceImpl{
		uploadDir: uploadDir,
	}
}

// SaveFile saves file in path
func (s *serviceImpl) SaveFile(fileName string, file multipart.File) error {
	filePath := filepath.Join(s.uploadDir, fileName)
	f, err := os.OpenFile(filePath, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		return err
	}
	defer f.Close()
	io.Copy(f, file)
	return nil
}
