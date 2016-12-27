package main

import (
	"log"

	"github.com/nrkv/GoShare/app"
	"os"
)

// Main function of application
func main() {
	// Create app with default settings
	app := app.New()
	// Init config
	if err := app.InitConfig(os.Getenv("CONFIG_FILE")); err != nil {
		log.Fatal("Error while parsing config file: " + err.Error())
	}
	// Init modules
	if err := app.InitModules(); err != nil {
		log.Fatal("Error while setting up application modules: " + err.Error())
	}
	// Init routes
	if err := app.InitRoutes(); err != nil {
		log.Fatal("Error while injecting application routes: " + err.Error())
	}
	// Run application
	if err := app.Run(); err != nil {
		log.Fatal("Server is stopped. Error: " + err.Error())
	}
}
