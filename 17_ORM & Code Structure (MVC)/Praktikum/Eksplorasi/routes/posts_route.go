package routes

import (
	"rest/controllers"

	"github.com/labstack/echo/v4"
)

func InitPostsRoute(e *echo.Echo) {
	e.GET("/api/v1/posts", controllers.GetPostsHandler)
	e.GET("/api/v1/posts/:id", controllers.GetDetailPostHandler)
	e.POST("/api/v1/posts", controllers.AddPostHandler)
	e.PUT("/api/v1/posts/:id", controllers.UpdatePostHandler)
	e.DELETE("/api/v1/posts/:id", controllers.DeletePostHandler)
}
