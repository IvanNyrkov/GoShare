package main

import (
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

// NewRouter creates echo router with all application endpoints
func (a *App) NewRouter() *echo.Echo {
	// Create router with default middleware
	router := echo.New()

	// Implementing C.O.R.S. headers enable pages within a modern web browser
	// to consume resources (such as REST APIs) from servers that are on a different domain.
	router.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{echo.GET, echo.HEAD, echo.PUT, echo.PATCH, echo.POST, echo.DELETE},
	}))

	// Logger middleware
	router.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: `${time_rfc3339} | ${method} | ${uri} | ${status} | ${latency_human}` + "\n",
	}))

	// Handle static files
	router.Static("/", "public")

	return router
}
