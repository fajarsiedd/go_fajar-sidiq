package routes

import (
	"rest/controllers"

	"github.com/labstack/echo/v4"
)

func InitPackagesRoute(e *echo.Echo) {
	e.GET("/api/v1/packages", controllers.GetPackagesHandler)
	e.GET("/api/v1/packages/:id", controllers.GetDetailPackageHandler)
	e.POST("/api/v1/packages", controllers.AddPackageHandler)
	e.PUT("/api/v1/packages/:id", controllers.UpdatePackageHandler)
	e.DELETE("/api/v1/packages/:id", controllers.DeletePackageHandler)
}
