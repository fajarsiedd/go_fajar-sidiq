package routes

import (
	"rest/controllers"

	"github.com/labstack/echo/v4"
)

func InitPaymentsRoute(e *echo.Echo) {
	e.GET("/api/v1/payments", controllers.GetPaymentsHandler)
	e.GET("/api/v1/payments/:id", controllers.GetDetailPaymentHandler)
	e.POST("/api/v1/payments", controllers.AddPaymentHandler)
	e.PUT("/api/v1/payments/:id", controllers.UpdatePaymentHandler)
	e.DELETE("/api/v1/payments/:id", controllers.DeletePaymentHandler)
}
