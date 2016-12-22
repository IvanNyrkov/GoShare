package main

// InitRoutes sets up application routes
func (app *App) InitRoutes() error {
	// Handle static files
	app.Router.Static("/", "public")
	// API endpoints
	if app.APIModule != nil {
		app.APIModule.InitRoutes(app.Router)
	}
	return nil
}
