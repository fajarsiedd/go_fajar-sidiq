package controllers

import (
	"net/http"
	"rest/config"
	"rest/models"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

func AddPackageHandler(c echo.Context) error {
	pkg := models.Packages{ID: uuid.New().String()}
	c.Bind(&pkg)
	result := config.DB.Create(&pkg)

	if result.Error != nil {
		var errorMessage string
		if result.Error != nil {
			errorMessage = result.Error.Error()
		} else {
			errorMessage = "unknown error"
		}
		return c.JSON(http.StatusInternalServerError, models.BaseResponse{Status: false, Message: errorMessage, Data: nil})
	}

	return c.JSON(http.StatusCreated, models.BaseResponse{Status: true, Message: "Success", Data: pkg})
}

func GetDetailPackageHandler(c echo.Context) error {
	id := c.Param("id")
	pkg := models.Packages{ID: id}
	result := config.DB.First(&pkg)

	if result.Error != nil {
		var errorMessage string
		if result.Error != nil {
			errorMessage = result.Error.Error()
		} else {
			errorMessage = "unknown error"
		}
		return c.JSON(http.StatusInternalServerError, models.BaseResponse{Status: false, Message: errorMessage, Data: nil})
	}

	return c.JSON(http.StatusOK, models.BaseResponse{Status: true, Message: "Success", Data: pkg})
}

func GetPackagesHandler(c echo.Context) error {
	pkgs := []models.Packages{}
	result := config.DB.Find(&pkgs)

	if result.Error != nil {
		var errorMessage string
		if result.Error != nil {
			errorMessage = result.Error.Error()
		} else {
			errorMessage = "unknown error"
		}
		return c.JSON(http.StatusInternalServerError, models.BaseResponse{Status: false, Message: errorMessage, Data: nil})
	}

	return c.JSON(http.StatusOK, models.BaseResponse{Status: true, Message: "Success", Data: pkgs})
}

func UpdatePackageHandler(c echo.Context) error {
	id := c.Param("id")
	pkg := models.Packages{ID: id}
	c.Bind(&pkg)

	result := config.DB.Save(&pkg)

	if result.Error != nil {
		var errorMessage string
		if result.Error != nil {
			errorMessage = result.Error.Error()
		} else {
			errorMessage = "unknown error"
		}
		return c.JSON(http.StatusInternalServerError, models.BaseResponse{Status: false, Message: errorMessage, Data: nil})
	}

	return c.JSON(http.StatusOK, models.BaseResponse{Status: true, Message: "Success", Data: pkg})
}

func DeletePackageHandler(c echo.Context) error {
	id := c.Param("id")
	result := config.DB.Delete(&models.Packages{ID: id})

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
