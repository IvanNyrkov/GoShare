package main

import (
	"log"

	"github.com/gorilla/mux"
	"github.com/nrkv/GoShare/api"
	"github.com/nrkv/GoShare/rand/sentence"
	"github.com/nrkv/GoShare/store"
	"os"
)

// App stores config, db connection and all injected modules
type App struct {
	Config             *Config
	Router             *mux.Router
	RandSentenceModule sentence.Module
	FileStorageModule  store.Module
	APIModule          api.Module
}

// Main function of application
func main() {
	// Create application with default settings
	app := &App{
		Router: mux.NewRouter(),
	}
	// Init application config
	if err := app.InitConfig(os.Getenv("CONFIG_FILE")); err != nil {
		log.Fatal("Error while parsing config file: " + err.Error())
	}
	// Init application modules
	if err := app.InitModules(); err != nil {
		log.Fatal("Error while injecting application modules: " + err.Error())
	}
	// Init application routes
	if err := app.InitRoutes(); err != nil {
		log.Fatal("Error while setting up application routes: " + err.Error())
	}
	// Run application
	if err := app.ListenAndServe(); err != nil {
		log.Fatal("Server is stopped. Error: " + err.Error())
	}
}
