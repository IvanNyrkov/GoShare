package api

import "github.com/gorilla/mux"

// InitRoutes initializes api routes
func (m *moduleImpl) InitRoutes(router *mux.Router) {
	api := router.PathPrefix("/api").Subrouter()
	api.HandleFunc("/files", m.controller.UploadFile).Methods("POST")
}
