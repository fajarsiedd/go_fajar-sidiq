package services

import (
	"errors"
	"rest/database"
	"rest/middlewares"
	"rest/models"
	req "rest/models/request"
	res "rest/models/response"
	"rest/utils"
)

type AuthService struct {
	jwtAuth *middlewares.JWTConfig
}

func InitAuthService(jwtAuth *middlewares.JWTConfig) AuthService {
	return AuthService{
		jwtAuth: jwtAuth,
	}
}

func (service *AuthService) Register(body req.RegisterRequest) (res.RegisterResponse, error) {
	config := &utils.ArgonConfig{
		Memory:     64 * 1024,
		Iterations: 3,
		Pararelism: 2,
		SaltLength: 16,
		KeyLength:  32,
	}

	password, err := utils.CreatePassword(body.Password, config)
	if err != nil {
		return res.RegisterResponse{}, err
	}

	user := models.Users{
		Name:     body.Name,
		Email:    body.Email,
		Password: password,
	}

	result := database.DB.Create(&user)
	if err := result.Error; err != nil {
		return res.RegisterResponse{}, err
	}

	err = result.Last(&user).Error
	if err != nil {
		return res.RegisterResponse{}, err
	}

	return res.RegisterResponse{
		Base:  user.Base,
		Name:  user.Name,
		Email: user.Email,
	}, nil
}

func (service *AuthService) Login(body req.LoginRequest) (res.LoginResponse, error) {
	user := models.Users{}

	err := database.DB.First(&user, "email = ?", body.Email).Error
	if err != nil {
		return res.LoginResponse{}, errors.New("invalid email")
	}

	match, err := utils.ComparePassword(body.Password, user.Password)

	isFailed := err != nil || !match
	if isFailed {
		return res.LoginResponse{}, errors.New("invalid password")
	}

	token, err := service.jwtAuth.GenerateToken(user.ID)
	if err != nil {
		return res.LoginResponse{}, err
	}

	return res.LoginResponse{
		Base:        user.Base,
		Name:        user.Name,
		Email:       user.Email,
		AccessToken: token,
	}, nil
}
