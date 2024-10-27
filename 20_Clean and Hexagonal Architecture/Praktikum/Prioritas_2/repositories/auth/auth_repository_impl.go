package auth

import (
	"errors"
	"go-wishlist-api/handlers/auth/request"
	"go-wishlist-api/models"

	"gorm.io/gorm"
)

type authRepositoryImpl struct {
	db *gorm.DB
}

func NewAuthRepository(db *gorm.DB) *authRepositoryImpl {
	return &authRepositoryImpl{
		db: db,
	}
}

func (repository authRepositoryImpl) Login(input request.LoginInput) (models.Users, error) {
	user := input.ToModel()

	err := repository.db.First(&user, "email = ?", user.Email).Error
	if err != nil {
		return models.Users{}, errors.New("invalid email")
	}

	return user, nil
}

func (repository authRepositoryImpl) Register(input request.RegisterInput) (models.Users, error) {
	user := input.ToModel()

	result := repository.db.Create(&user)
	if err := result.Error; err != nil {
		return models.Users{}, err
	}

	return user, nil
}
