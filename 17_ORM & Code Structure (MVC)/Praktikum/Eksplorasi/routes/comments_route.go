package routes

import (
	"rest/controllers"

	"github.com/labstack/echo/v4"
)

func InitCommentsRoute(e *echo.Echo) {
	e.POST("/api/v1/comments/:post_id", controllers.AddCommentHandler)
	e.PUT("/api/v1/comments/:id", controllers.UpdateCommentHandler)
	e.DELETE("/api/v1/comments/:id", controllers.DeleteCommentHandler)
}
