package main

import (
	"log"

	"github.com/labstack/echo/engine/standard"
)

// Main function of application
func main() {
	config := NewConfig()
	app := &App{}
	router := app.NewRouter()

	log.Printf("Listening at port %s", config.ServerPort)
	router.Run(standard.New(config.ServerPort))
}
