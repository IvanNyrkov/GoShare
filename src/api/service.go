package api

import (
	"mime/multipart"
	"path/filepath"
)

// Service is an interface that defines API actions
type Service interface {
	UploadFile(fileName string, file multipart.File) (string, error)
}
type serviceImpl struct {
	randSentenceService randSentenceService
	fileStorageService  fileStorageService
}

// NewService creates new service
func NewService(randSentenceService randSentenceService, fileStorageService fileStorageService) *serviceImpl {
	return &serviceImpl{
		randSentenceService: randSentenceService,
		fileStorageService:  fileStorageService,
	}
}

// UploadFile saves file with randomly generated name
func (s *serviceImpl) UploadFile(fileName string, file multipart.File) (string, error) {
	randCode := s.randSentenceService.RandomSentence("-")
	if err := s.fileStorageService.SaveFile(randCode+filepath.Ext(fileName), file); err != nil {
		return "", err
	}
	return randCode, nil
}
