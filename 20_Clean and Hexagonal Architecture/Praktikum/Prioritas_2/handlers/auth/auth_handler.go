package auth

import (
	"go-wishlist-api/handlers/auth/request"
	"go-wishlist-api/handlers/base"
	"go-wishlist-api/usecases/auth"

	"github.com/labstack/echo/v4"
)

type authHandler struct {
	usecase auth.AuthUsecase
}

func NewAuthHandler(u auth.AuthUsecase) *authHandler {
	return &authHandler{
		usecase: u,
	}
}

func (handler authHandler) Login(c echo.Context) error {
	input := request.LoginInput{}

	if err := c.Bind(&input); err != nil {
		return base.BadRequestResponse(c, err)
	}

	result, err := handler.usecase.Login(input)

	if err != nil {
		return base.ErrorResponse(c, err)
	}

	return base.SuccesResponse(c, result)
}

func (handler authHandler) Register(c echo.Context) error {
	input := request.RegisterInput{}

	if err := c.Bind(&input); err != nil {
		return base.BadRequestResponse(c, err)
	}

	result, err := handler.usecase.Register(input)

	if err != nil {
		return base.ErrorResponse(c, err)
	}

	return base.SuccesResponse(c, result)
}
