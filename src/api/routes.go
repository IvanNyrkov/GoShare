package api

import "github.com/labstack/echo"

// InitRoutes initializes api routes
func (m *moduleImpl) InitRoutes(router *echo.Echo) {
	api := router.Group("/api")
	api.POST("/files", m.controller.UploadFile)
}
