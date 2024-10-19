package controllers

import (
	"net/http"
	"rest/config"
	"rest/models"

	"github.com/labstack/echo/v4"
)

func AddCommentHandler(c echo.Context) error {
	postID := c.Param("post_id")
	comment := models.Comments{PostID: postID}
	c.Bind(&comment)
	result := config.DB.Create(&comment)

	if result.Error != nil {
		var errorMessage string
		if result.Error != nil {
			errorMessage = result.Error.Error()
		} else {
			errorMessage = "unknown error"
		}
		return c.JSON(http.StatusInternalServerError, models.BaseResponse{Status: false, Message: errorMessage, Data: nil})
	}

	return c.JSON(http.StatusCreated, models.BaseResponse{Status: true, Message: "Success", Data: comment})
}

func UpdateCommentHandler(c echo.Context) error {
	id := c.Param("id")
	comment := models.Comments{Base: models.Base{ID: id}}
	c.Bind(&comment)

	result := config.DB.Updates(&comment)

	if result.Error != nil {
		var errorMessage string
		if result.Error != nil {
			errorMessage = result.Error.Error()
		} else {
			errorMessage = "unknown error"
		}
		return c.JSON(http.StatusInternalServerError, models.BaseResponse{Status: false, Message: errorMessage, Data: nil})
	}

	return c.JSON(http.StatusOK, models.BaseResponse{Status: true, Message: "Success", Data: comment})
}

func DeleteCommentHandler(c echo.Context) error {
	id := c.Param("id")
	result := config.DB.Delete(&models.Comments{Base: models.Base{ID: id}})

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
