package controllers

import (
	"net/http"
	"rest/config"
	"rest/models"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

func AddPaymentHandler(c echo.Context) error {
	payment := models.Payments{ID: uuid.New().String()}
	c.Bind(&payment)
	result := config.DB.Create(&payment)

	if result.Error != nil {
		var errorMessage string
		if result.Error != nil {
			errorMessage = result.Error.Error()
		} else {
			errorMessage = "unknown error"
		}
		return c.JSON(http.StatusInternalServerError, models.BaseResponse{Status: false, Message: errorMessage, Data: nil})
	}

	return c.JSON(http.StatusCreated, models.BaseResponse{Status: true, Message: "Success", Data: payment})
}

func GetDetailPaymentHandler(c echo.Context) error {
	id := c.Param("id")
	payment := models.Payments{ID: id}
	result := config.DB.First(&payment)

	if result.Error != nil {
		var errorMessage string
		if result.Error != nil {
			errorMessage = result.Error.Error()
		} else {
			errorMessage = "unknown error"
		}
		return c.JSON(http.StatusInternalServerError, models.BaseResponse{Status: false, Message: errorMessage, Data: nil})
	}

	return c.JSON(http.StatusOK, models.BaseResponse{Status: true, Message: "Success", Data: payment})
}

func GetPaymentsHandler(c echo.Context) error {
	payments := []models.Payments{}
	result := config.DB.Find(&payments)

	if result.Error != nil {
		var errorMessage string
		if result.Error != nil {
			errorMessage = result.Error.Error()
		} else {
			errorMessage = "unknown error"
		}
		return c.JSON(http.StatusInternalServerError, models.BaseResponse{Status: false, Message: errorMessage, Data: nil})
	}

	return c.JSON(http.StatusOK, models.BaseResponse{Status: true, Message: "Success", Data: payments})
}

func UpdatePaymentHandler(c echo.Context) error {
	id := c.Param("id")
	payment := models.Payments{ID: id}
	c.Bind(&payment)

	result := config.DB.Save(&payment)

	if result.Error != nil {
		var errorMessage string
		if result.Error != nil {
			errorMessage = result.Error.Error()
		} else {
			errorMessage = "unknown error"
		}
		return c.JSON(http.StatusInternalServerError, models.BaseResponse{Status: false, Message: errorMessage, Data: nil})
	}

	return c.JSON(http.StatusOK, models.BaseResponse{Status: true, Message: "Success", Data: payment})
}

func DeletePaymentHandler(c echo.Context) error {
	id := c.Param("id")
	result := config.DB.Delete(&models.Payments{ID: id})

	if result.Error != nil {
		var errorMessage string
		if result.Error != nil {
			errorMessage = result.Error.Error()
		} else {
			errorMessage = "unknown error"
		}
		return c.JSON(http.StatusInternalServerError, models.BaseResponse{Status: false, Message: errorMessage, Data: nil})
	}

	return c.JSON(http.StatusNoContent, nil)
}
