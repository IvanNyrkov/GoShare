package api

import (
	"encoding/json"
	"log"
	"net/http"
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
	r.ParseMultipartForm(20000000)
	if r.MultipartForm == nil || len(r.MultipartForm.File["file"]) < 1 {
		log.Println("Can't parse file from form data")
		writeJSON(w, http.StatusBadRequest, responseStatus{
			StatusCode: http.StatusBadRequest,
			Message:    "Can't parse file from form data.",
		})
		return
	}
	fileHeader := r.MultipartForm.File["file"][0]
	file, err := fileHeader.Open()
	if err != nil {
		log.Println("Can't open file parsed from form data: Error: ", err.Error())
		writeJSON(w, http.StatusBadRequest, responseStatus{
			StatusCode: http.StatusBadRequest,
			Message:    "Can't open file parsed from form data.",
		})
		return
	}
	// Save file
	code, err := c.service.UploadFile(fileHeader.Filename, file)
	if err != nil {
		log.Println("Error while saving the file: Error: ", err.Error())
		writeJSON(w, http.StatusInternalServerError, responseStatus{
			StatusCode: http.StatusInternalServerError,
			Message:    "Error while saving the file.",
		})
		return
	}
	log.Println("File has been saved: Code has been generated: ", code)
	writeJSON(w, http.StatusOK, struct {
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

func writeJSON(w http.ResponseWriter, status int, data interface{}) {
	js, err := json.Marshal(data)
	if err != nil {
		http.Error(w, "Error while rendering response", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	w.Write(js)
}
