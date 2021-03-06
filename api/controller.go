package api

import (
	"log"
	"net/http"

	"github.com/nrkv/snippers/response"
)

// Controller is an interface that defines handlers
type Controller interface {
	UploadFile(w http.ResponseWriter, r *http.Request)
}
type controllerImpl struct {
	service Service
}

type responseStatus struct {
	StatusCode int    `json:"status_code"`
	Message    string `json:"message"`
}

// NewController creates new controller struct and links it with required services
func NewController(service Service) *controllerImpl {
	return &controllerImpl{
		service: service,
	}
}

// UploadFile reads file from multipart form data, saves it and responds with randomly generated sentence
func (c *controllerImpl) UploadFile(w http.ResponseWriter, r *http.Request) {
	// Parse file from form
	file, fileHeader, err := r.FormFile("file")
	if err != nil {
		log.Println("Can't parse file from form data: Error: ", err.Error())
		response.JSON(w, http.StatusBadRequest, responseStatus{
			StatusCode: http.StatusBadRequest,
			Message:    "Can't parse file from form data.",
		})
		return
	}
	// Save file
	code, err := c.service.UploadFile(fileHeader.Filename, file)
	if err != nil {
		log.Println("Error while saving the file: Error: ", err.Error())
		response.JSON(w, http.StatusInternalServerError, responseStatus{
			StatusCode: http.StatusInternalServerError,
			Message:    "Error while saving the file.",
		})
		return
	}
	log.Println("File has been saved: Code has been generated: ", code)
	response.JSON(w, http.StatusOK, struct {
		responseStatus
		Code string `json:"code"`
	}{
		responseStatus: responseStatus{
			StatusCode: http.StatusOK,
			Message:    "File has been saved. Code has been generated.",
		},
		Code: code,
	})
}
