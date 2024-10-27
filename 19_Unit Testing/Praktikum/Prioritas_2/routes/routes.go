package routes

import (
	"go-todos-api/controllers"

	"github.com/labstack/echo/v4"
)

func SetupRoutes(e *echo.Echo) {
	todoController := controllers.InitTodoController()

	todoRoutes := e.Group("/api/v1")

	todoRoutes.GET("/todos", todoController.GetAll)
	todoRoutes.GET("/todos/:id", todoController.GetByID)
	todoRoutes.POST("/todos", todoController.Create)
	todoRoutes.PUT("/todos/:id", todoController.Update)
	todoRoutes.DELETE("/todos/:id", todoController.Delete)
}
