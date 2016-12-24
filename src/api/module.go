package api

import "github.com/gorilla/mux"

// Module defines module interface
type Module interface {
	GetService() Service
	InitRoutes(router *mux.Router)
}
type moduleImpl struct {
	controller Controller
	service    Service
}

// ModuleConfig contains required configs
type ModuleConfig struct {
	RandSentenceService randSentenceService
	FileStorageService  fileStorageService
}

// NewModule creates struct that encapsulates the module
func NewModule(config ModuleConfig) *moduleImpl {
	service := NewService(config.RandSentenceService, config.FileStorageService)
	controller := NewController(service)
	return &moduleImpl{
		controller: controller,
		service:    service,
	}
}

// GetService returns module service
func (m *moduleImpl) GetService() Service {
	return m.service
}
