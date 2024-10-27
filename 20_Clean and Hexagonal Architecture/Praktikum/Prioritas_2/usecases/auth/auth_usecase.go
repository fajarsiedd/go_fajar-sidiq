package auth

import (
	"go-wishlist-api/handlers/auth/request"
	"go-wishlist-api/handlers/auth/response"
)

type AuthUsecase interface {
	Login(input request.LoginInput) (response.LoginResponse, error)
	Register(input request.RegisterInput) (response.RegisterResponse, error)
}
