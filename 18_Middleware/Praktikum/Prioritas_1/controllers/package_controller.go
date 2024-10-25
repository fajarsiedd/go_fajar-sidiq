package controllers

import (
	"net/http"
	"rest/models"
	"rest/services"

	"github.com/labstack/echo/v4"
)

type PackageController struct {
	service services.PackageService
}

func InitPackageController() PackageController {
	return PackageController{
		service: services.InitPackageService(),
	}
}

func (controller *PackageController) GetPackages(c echo.Context) error {
	result, err := controller.service.GetPackages()
	if err != nil {
		return c.JSON(http.StatusBadRequest, models.BaseResponse{
			Status:  false,
			Message: err.Error(),
		})
	}

	return c.JSON(http.StatusOK, models.BaseResponse{
		Status:  true,
		Message: "all packages fetched",
		Data:    result,
	})
}

func (controller *PackageController) GetPackageByID(c echo.Context) error {
	id := c.Param("id")

	result, err := controller.service.GetPackageByID(id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, models.BaseResponse{
			Status:  false,
			Message: err.Error(),
		})
	}

	return c.JSON(http.StatusOK, models.BaseResponse{
		Status:  true,
		Message: "package fetched",
		Data:    result,
	})
}

func (controller *PackageController) AddPackage(c echo.Context) error {
	body := models.Packages{}

	if err := c.Bind(&body); err != nil {
		return c.JSON(http.StatusBadRequest, models.BaseResponse{
			Status:  false,
			Message: err.Error(),
		})
	}

	if err := c.Validate(&body); err != nil {
		return c.JSON(http.StatusBadRequest, models.BaseResponse{
			Status:  false,
			Message: err.Error(),
		})
	}

	result, err := controller.service.AddPackage(body)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, models.BaseResponse{
			Status:  false,
			Message: err.Error(),
		})
	}

	return c.JSON(http.StatusCreated, models.BaseResponse{
		Status:  true,
		Message: "package added",
		Data:    result,
	})
}

func (controller *PackageController) UpdatePackage(c echo.Context) error {
	id := c.Param("id")
	body := models.Packages{Base: models.Base{ID: id}}

	if err := c.Bind(&body); err != nil {
		return c.JSON(http.StatusBadRequest, models.BaseResponse{
			Status:  false,
			Message: err.Error(),
		})
	}

	result, err := controller.service.UpdatePackage(body)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, models.BaseResponse{
			Status:  false,
			Message: err.Error(),
		})
	}

	return c.JSON(http.StatusOK, models.BaseResponse{
		Status:  true,
		Message: "package updated",
		Data:    result,
	})
}

func (controller *PackageController) DeletePackage(c echo.Context) error {
	id := c.Param("id")

	err := controller.service.DeletePackage(id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, models.BaseResponse{
			Status:  false,
			Message: err.Error(),
		})
	}

	return c.JSON(http.StatusOK, models.BaseResponse{
		Status:  true,
		Message: "package deleted",
	})
}
