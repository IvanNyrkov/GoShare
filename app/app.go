package app

import (
	"log"

	"net/http"

	"os"

	"github.com/gorilla/mux"
	"github.com/nrkv/GoShare/app/api"
	"github.com/nrkv/GoShare/app/rand/sentence"
	"github.com/nrkv/GoShare/app/store"
	"github.com/nrkv/snippers/middleware"
)

// App stores config, db connection and all injected modules
type App struct {
	Config             *Config
	Router             *mux.Router
	RandSentenceModule sentence.Module
	FileStorageModule  store.Module
	APIModule          api.Module
}

// New creates default application structure
func New() *App {
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
	return http.ListenAndServe(app.Config.Port, middleware.Logger(app.Router, os.Stdout, middleware.DefaultLoggerConfig))
}
