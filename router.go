package main

import (
	"net/http"
	_ "net/http/pprof"
	"os"
	"log"
	"github.com/nrkv/snippers/middleware"
	"github.com/gorilla/mux"
)

// InitRouter sets up application router
func (app *App) InitRouter() error {
	// API endpoints
	if app.APIModule != nil {
		app.APIModule.InitRoutes(app.Router)
	}
	// Profiling
	app.Router.PathPrefix("/debug").Handler(http.DefaultServeMux)
	// Handle static files
	app.Router.PathPrefix("/").Handler(http.StripPrefix("/", http.FileServer(http.Dir("public"))))
	return nil
}

// ListenAndServe starts application server
func (app *App) ListenAndServe() error {
	port := app.Config.Port
	if app.Config.Port == "" {
		port = ":80"
	}
	router := app.Router
	if app.Router == nil {
		router = mux.NewRouter()
	}
	log.Printf("Listening at port %s", port)
	return http.ListenAndServe(port, middleware.Logger(router, os.Stdout, middleware.DefaultLoggerConfig))
}