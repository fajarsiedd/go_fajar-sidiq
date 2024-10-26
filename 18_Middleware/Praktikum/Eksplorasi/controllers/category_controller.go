package controllers

import (
	"net/http"
	"rest/models"
	"rest/services"

	"github.com/labstack/echo/v4"
)

type CategoryController struct {
	service services.CategoryService
}

func InitCategoryController() CategoryController {
	return CategoryController{
		service: services.InitCategoryService(),
	}
}

func (controller *CategoryController) GetCategories(c echo.Context) error {
	result, err := controller.service.GetCategories()
	if err != nil {
		return c.JSON(http.StatusBadRequest, models.BaseResponse{
			Status:  false,
			Message: err.Error(),
		})
	}

	return c.JSON(http.StatusOK, models.BaseResponse{
		Status:  true,
		Message: "all categories fetched",
		Data:    result,
	})
}

func (controller *CategoryController) AddCategory(c echo.Context) error {
	body := models.Categories{}

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

	result, err := controller.service.AddCategory(body)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, models.BaseResponse{
			Status:  false,
			Message: err.Error(),
		})
	}

	return c.JSON(http.StatusCreated, models.BaseResponse{
		Status:  true,
		Message: "category added",
		Data:    result,
	})
}

func (controller *CategoryController) DeleteCategory(c echo.Context) error {
	id := c.Param("id")

	err := controller.service.DeleteCategory(id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, models.BaseResponse{
			Status:  false,
			Message: err.Error(),
		})
	}

	return c.JSON(http.StatusOK, models.BaseResponse{
		Status:  true,
		Message: "category deleted",
	})
}
