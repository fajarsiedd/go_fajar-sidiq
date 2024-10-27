package routes

import (
	handler "go-wishlist-api/handlers/auth"
	"go-wishlist-api/middlewares"
	repository "go-wishlist-api/repositories/auth"
	usecase "go-wishlist-api/usecases/auth"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func InitAuthRoute(e *echo.Echo, db *gorm.DB, jwtConfig *middlewares.JWTConfig) {
	authRepo := repository.NewAuthRepository(db)
	authUsecase := usecase.NewAuthUsecase(authRepo, jwtConfig)
	authHandler := handler.NewAuthHandler(authUsecase)

	wishlist := e.Group("/api/v1")
	wishlist.POST("/login", authHandler.Login)
	wishlist.POST("/register", authHandler.Register)
}
