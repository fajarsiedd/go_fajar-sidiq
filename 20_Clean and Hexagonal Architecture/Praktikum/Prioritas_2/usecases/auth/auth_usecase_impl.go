package auth

import (
	"errors"
	"go-wishlist-api/handlers/auth/request"
	"go-wishlist-api/handlers/auth/response"
	"go-wishlist-api/middlewares"
	"go-wishlist-api/repositories/auth"
	"go-wishlist-api/utils"
)

type authUsecase struct {
	repository auth.AuthRepository
	jwtAuth    *middlewares.JWTConfig
}

func NewAuthUsecase(r auth.AuthRepository, jwtAuth *middlewares.JWTConfig) *authUsecase {
	return &authUsecase{
		repository: r,
		jwtAuth:    jwtAuth,
	}
}

func (usecase authUsecase) Login(input request.LoginInput) (response.LoginResponse, error) {
	result, err := usecase.repository.Login(input)

	if err != nil {
		return response.LoginResponse{}, err
	}

	match, err := utils.ComparePassword(input.Password, result.Password)

	isFailed := err != nil || !match
	if isFailed {
		return response.LoginResponse{}, errors.New("invalid password")
	}

	token, err := usecase.jwtAuth.GenerateToken(result.ID)
	if err != nil {
		return response.LoginResponse{}, err
	}

	return response.Login_FromModel(result, token), nil
}

func (usecase authUsecase) Register(input request.RegisterInput) (response.RegisterResponse, error) {
	config := &utils.ArgonConfig{
		Memory:     64 * 1024,
		Iterations: 3,
		Pararelism: 2,
		SaltLength: 16,
		KeyLength:  32,
	}

	var err error
	input.Password, err = utils.CreatePassword(input.Password, config)
	if err != nil {
		return response.RegisterResponse{}, err
	}

	result, err := usecase.repository.Register(input)

	if err != nil {
		return response.RegisterResponse{}, err
	}

	return response.Register_FromModel(result), nil
}
