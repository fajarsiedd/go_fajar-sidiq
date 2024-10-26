package controllers

import (
	"net/http"
	"rest/models"
	"rest/services"

	"github.com/labstack/echo/v4"
)

type PostController struct {
	service services.PostService
}

func InitPostController() PostController {
	return PostController{
		service: services.InitPostService(),
	}
}

func (controller *PostController) GetPosts(c echo.Context) error {
	result, err := controller.service.GetPosts()
	if err != nil {
		return c.JSON(http.StatusBadRequest, models.BaseResponse{
			Status:  false,
			Message: err.Error(),
		})
	}

	return c.JSON(http.StatusOK, models.BaseResponse{
		Status:  true,
		Message: "all posts fetched",
		Data:    result,
	})
}

func (controller *PostController) GetPostByID(c echo.Context) error {
	id := c.Param("id")

	result, err := controller.service.GetPostByID(id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, models.BaseResponse{
			Status:  false,
			Message: err.Error(),
		})
	}

	return c.JSON(http.StatusOK, models.BaseResponse{
		Status:  true,
		Message: "post fetched",
		Data:    result,
	})
}

func (controller *PostController) AddPost(c echo.Context) error {
	body := models.Posts{}

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

	result, err := controller.service.AddPost(body)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, models.BaseResponse{
			Status:  false,
			Message: err.Error(),
		})
	}

	return c.JSON(http.StatusCreated, models.BaseResponse{
		Status:  true,
		Message: "post added",
		Data:    result,
	})
}

func (controller *PostController) UpdatePost(c echo.Context) error {
	id := c.Param("id")
	body := models.Posts{Base: models.Base{ID: id}}

	if err := c.Bind(&body); err != nil {
		return c.JSON(http.StatusBadRequest, models.BaseResponse{
			Status:  false,
			Message: err.Error(),
		})
	}

	result, err := controller.service.UpdatePost(body)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, models.BaseResponse{
			Status:  false,
			Message: err.Error(),
		})
	}

	return c.JSON(http.StatusOK, models.BaseResponse{
		Status:  true,
		Message: "post updated",
		Data:    result,
	})
}

func (controller *PostController) DeletePost(c echo.Context) error {
	id := c.Param("id")

	err := controller.service.DeletePost(id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, models.BaseResponse{
			Status:  false,
			Message: err.Error(),
		})
	}

	return c.JSON(http.StatusOK, models.BaseResponse{
		Status:  true,
		Message: "post deleted",
	})
}
