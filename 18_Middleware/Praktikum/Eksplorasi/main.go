package main

import (
	"rest/database"
	"rest/routes"

	"github.com/labstack/echo/v4"
)

func main() {
	database.InitDB()

	database.MigrateDB()

	e := echo.New()

	routes.InitRoutes(e)

	e.Logger.Fatal(e.Start(":8000"))
}
