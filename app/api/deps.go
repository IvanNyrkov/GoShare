package api

import "mime/multipart"

type randSentenceService interface {
	RandomSentence(sep string) string
}

type fileStorageService interface {
	SaveFile(fileName string, file multipart.File) error
}
