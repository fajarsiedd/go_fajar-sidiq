package routes

import (
	"go-math-api/controllers"

	"github.com/labstack/echo/v4"
)

func CalculatorRoutes(e *echo.Echo) {
	e.POST("/api/add", controllers.Add)
	e.POST("/api/subtract", controllers.Subtract)
	e.POST("/api/multiply", controllers.Multiply)
	e.POST("/api/div", controllers.Divide)
}
