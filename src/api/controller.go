package api

import (
	"log"
	"net/http"

	"github.com/IvanNyrkov/GoShare/src/store"
	"github.com/labstack/echo"
)

// Controller is an interface that defines handlers
type Controller interface {
	UploadFile(ctx echo.Context) error
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
func (c *controllerImpl) UploadFile(ctx echo.Context) error {
	fileHeader, err := ctx.FormFile("file")
	if err != nil {
		log.Println("Can't parse file from form data: Error: ", err.Error())
		return ctx.JSON(http.StatusBadRequest, responseStatus{
			StatusCode: http.StatusBadRequest,
			Message:    "Can't parse file from form data.",
		})
	}
	file, err := fileHeader.Open()
	if err != nil {
		log.Println("Can't open file parsed from form data: Error: ", err.Error())
		return ctx.JSON(http.StatusBadRequest, responseStatus{
			StatusCode: http.StatusBadRequest,
			Message:    "Can't open file parsed from form data.",
		})
	}
	if err := store.SaveFile(fileHeader.Filename, file); err != nil {
		log.Println("Error while saving the file: Error: ", err.Error())
		return ctx.JSON(http.StatusInternalServerError, responseStatus{
			StatusCode: http.StatusInternalServerError,
			Message:    "Error while saving the file.",
		})
	}
	code := "GENERATED-CODE"
	log.Println("File has been saved: Code has been generated: ", code)
	return ctx.JSON(http.StatusOK, struct {
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
