package app

import (
	"net/http"
	_ "net/http/pprof"
)

// InitRoutes sets up application routes
func (app *App) InitRoutes() error {
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
