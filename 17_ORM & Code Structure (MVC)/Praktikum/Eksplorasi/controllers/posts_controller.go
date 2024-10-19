package controllers

import (
	"net/http"
	"rest/config"
	"rest/models"

	"github.com/labstack/echo/v4"
)

func AddPostHandler(c echo.Context) error {
	post := models.Posts{}
	c.Bind(&post)
	result := config.DB.Create(&post)

	if result.Error != nil {
		var errorMessage string
		if result.Error != nil {
			errorMessage = result.Error.Error()
		} else {
			errorMessage = "unknown error"
		}
		return c.JSON(http.StatusInternalServerError, models.BaseResponse{Status: false, Message: errorMessage, Data: nil})
	}

	return c.JSON(http.StatusCreated, models.BaseResponse{Status: true, Message: "Success", Data: post})
}

func GetDetailPostHandler(c echo.Context) error {
	id := c.Param("id")
	post := models.Posts{Base: models.Base{ID: id}}
	result := config.DB.Model(&models.Posts{}).Preload("Comments").First(&post)

	if result.Error != nil {
		var errorMessage string
		if result.Error != nil {
			errorMessage = result.Error.Error()
		} else {
			errorMessage = "unknown error"
		}
		return c.JSON(http.StatusInternalServerError, models.BaseResponse{Status: false, Message: errorMessage, Data: nil})
	}

	return c.JSON(http.StatusOK, models.BaseResponse{Status: true, Message: "Success", Data: post})
}

func GetPostsHandler(c echo.Context) error {
	posts := []models.Posts{}
	result := config.DB.Find(&posts)

	if result.Error != nil {
		var errorMessage string
		if result.Error != nil {
			errorMessage = result.Error.Error()
		} else {
			errorMessage = "unknown error"
		}
		return c.JSON(http.StatusInternalServerError, models.BaseResponse{Status: false, Message: errorMessage, Data: nil})
	}

	return c.JSON(http.StatusOK, models.BaseResponse{Status: true, Message: "Success", Data: posts})
}

func UpdatePostHandler(c echo.Context) error {
	id := c.Param("id")
	post := models.Posts{Base: models.Base{ID: id}}
	c.Bind(&post)

	result := config.DB.Updates(&post)

	if result.Error != nil {
		var errorMessage string
		if result.Error != nil {
			errorMessage = result.Error.Error()
		} else {
			errorMessage = "unknown error"
		}
		return c.JSON(http.StatusInternalServerError, models.BaseResponse{Status: false, Message: errorMessage, Data: nil})
	}

	return c.JSON(http.StatusOK, models.BaseResponse{Status: true, Message: "Success", Data: post})
}

func DeletePostHandler(c echo.Context) error {
	id := c.Param("id")
	result := config.DB.Select("Comments").Delete(&models.Posts{Base: models.Base{ID: id}})

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
