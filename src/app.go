package main

import (
	"log"

	"github.com/IvanNyrkov/GoShare/src/rand/sentence"
	"github.com/labstack/echo"
	"github.com/labstack/echo/engine/standard"
	"github.com/labstack/echo/middleware"
)

// App stores config, db connection and all injected modules
type App struct {
	Config             *Config
	Router             *echo.Echo
	RandSentenceModule sentence.Module
}

// NewApp creates default application structure
func NewApp() *App {
	app := new(App)
	// Create default config
	app.Config = &Config{
		Port: ":80",
	}
	// Create default router
	app.Router = echo.New()
	// Implementing C.O.R.S. headers enable pages within a modern web browser
	// to consume resources (such as REST APIs) from servers that are on a different domain.
	app.Router.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{echo.GET, echo.HEAD, echo.PUT, echo.PATCH, echo.POST, echo.DELETE},
	}))
	// Logger middleware
	app.Router.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: `${time_rfc3339} | ${method} | ${uri} | ${status} | ${latency_human}` + "\n",
	}))
	return app
}

// Run starts application
func (app *App) Run() error {
	log.Printf("Listening at port %s", app.Config.Port)
	return app.Router.Run(standard.New(app.Config.Port))
}
