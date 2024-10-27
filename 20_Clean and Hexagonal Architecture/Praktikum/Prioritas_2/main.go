package main

import (
	"go-wishlist-api/database"
	"go-wishlist-api/middlewares"
	"go-wishlist-api/routes"
	"go-wishlist-api/utils"

	"github.com/labstack/echo/v4"
)

func main() {
	db, _ := database.InitDB()
	database.MigrateDB(db)

	e := echo.New()

	jwtConfig := middlewares.JWTConfig{
		SecretKey:       utils.GetConfig("JWT_SCRET_KEY"),
		ExpiresDuration: 1,
	}

	authMiddlewareConfig := jwtConfig.Init()

	routes.InitAuthRoute(e, db, &jwtConfig)
	routes.InitWishlistRoute(e, db, authMiddlewareConfig)

	e.Logger.Fatal(e.Start(":1323"))
}
