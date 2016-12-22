package main

// InitRoutes sets up application routes
func (app *App) InitRoutes() error {
	// Handle static files
	app.Router.Static("/", "public")
	return nil
}
