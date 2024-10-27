package auth

import (
	"go-wishlist-api/handlers/auth/request"
	"go-wishlist-api/models"
)

type AuthRepository interface {
	Login(input request.LoginInput) (models.Users, error)
	Register(input request.RegisterInput) (models.Users, error)
}
