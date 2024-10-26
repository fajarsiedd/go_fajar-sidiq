package controllers

import (
	"net/http"
	"rest/middlewares"
	"rest/models"
	req "rest/models/request"
	"rest/services"

	"github.com/labstack/echo/v4"
)

type AuthController struct {
	service services.AuthService
}

func InitAuthController(jwtAuth *middlewares.JWTConfig) AuthController {
	return AuthController{
		service: services.InitAuthService(jwtAuth),
	}
}

func (controller *AuthController) Register(c echo.Context) error {
	body := req.RegisterRequest{}

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

	result, err := controller.service.Register(body)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, models.BaseResponse{
			Status:  false,
			Message: err.Error(),
		})
	}

	return c.JSON(http.StatusCreated, models.BaseResponse{
		Status:  true,
		Message: "user registered",
		Data:    result,
	})
}

func (controller *AuthController) Login(c echo.Context) error {
	body := req.LoginRequest{}

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

	result, err := controller.service.Login(body)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, models.BaseResponse{
			Status:  false,
			Message: err.Error(),
		})
	}

	return c.JSON(http.StatusOK, models.BaseResponse{
		Status:  true,
		Message: "login success",
		Data:    result,
	})
}
