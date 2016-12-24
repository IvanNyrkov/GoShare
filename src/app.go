package main

import (
	"log"

	"net/http"

	"time"

	"github.com/IvanNyrkov/GoShare/src/api"
	"github.com/IvanNyrkov/GoShare/src/rand/sentence"
	"github.com/IvanNyrkov/GoShare/src/store"
	"github.com/gorilla/mux"
)

// App stores config, db connection and all injected modules
type App struct {
	Config             *Config
	Router             *mux.Router
	RandSentenceModule sentence.Module
	FileStorageModule  store.Module
	APIModule          api.Module
}

// NewApp creates default application structure
func NewApp() *App {
	app := new(App)
	// Create default config
	app.Config = &Config{
		Port: ":80",
	}
	// Create default router
	app.Router = mux.NewRouter()
	return app
}

// Run starts application
func (app *App) Run() error {
	log.Printf("Listening at port %s", app.Config.Port)
	return http.ListenAndServe(app.Config.Port, func(inner http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			start := time.Now()
			inner.ServeHTTP(w, r)
			log.Printf(
				"[%s] | %s | %s",
				r.Method,
				r.RequestURI,
				time.Since(start),
			)
		})
	}(app.Router))
}
