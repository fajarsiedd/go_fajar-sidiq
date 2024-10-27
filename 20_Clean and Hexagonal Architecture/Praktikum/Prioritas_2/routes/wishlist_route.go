package routes

import (
	handler "go-wishlist-api/handlers/wishlist"
	repository "go-wishlist-api/repositories/wishlist"
	usecase "go-wishlist-api/usecases/wishlist"

	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func InitWishlistRoute(e *echo.Echo, db *gorm.DB, authMiddlewareConfig echojwt.Config) {
	wishlistRepo := repository.NewWishlistRepository(db)
	wishlistUsecase := usecase.NewWishlistUsecase(wishlistRepo)
	wishlistHandler := handler.NewWishlistHandler(wishlistUsecase)

	wishlist := e.Group("/api/v1/wishlist", echojwt.WithConfig(authMiddlewareConfig))
	wishlist.POST("", wishlistHandler.Create)
	wishlist.GET("", wishlistHandler.GetAll)
}
