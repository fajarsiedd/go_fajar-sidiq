package controllers

import (
	"go-math-api/models"
	"net/http"

	"github.com/labstack/echo/v4"
)

func Add(c echo.Context) error {
	request := new(models.Calculator)
	if err := c.Bind(request); err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": "Invalid input"})
	}
	result := request.Add()
	return c.JSON(http.StatusOK, echo.Map{"result": result})
}

func Subtract(c echo.Context) error {
	request := new(models.Calculator)
	if err := c.Bind(request); err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": "Invalid input"})
	}
	result := request.Subtract()
	return c.JSON(http.StatusOK, echo.Map{"result": result})
}

func Multiply(c echo.Context) error {
	request := new(models.Calculator)
	if err := c.Bind(request); err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": "Invalid input"})
	}
	result := request.Multiply()
	return c.JSON(http.StatusOK, echo.Map{"result": result})
}

func Divide(c echo.Context) error {
	request := new(models.Calculator)
	if err := c.Bind(request); err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": "Invalid input"})
	}
	result := request.Divide()
	return c.JSON(http.StatusOK, echo.Map{"result": result})
}
