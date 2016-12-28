package main

import (
	"github.com/gorilla/mux"
	"github.com/nrkv/snippers/middleware"
	"log"
	"net/http"
	_ "net/http/pprof"
	"github.com/justinas/alice"
	"github.com/rs/cors"
)

// InitRoutes sets up application router
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
	chain := alice.New(
		cors.Default().Handler,
		middleware.Recover,
		middleware.Logger,
	).Then(router)
	return http.ListenAndServe(port, chain)
}