package main

import (
	"rest/config"
	"rest/routes"

	"github.com/labstack/echo/v4"
)

func main() {
	config.ConnectDatabase()
	config.MigrateDB()

	e := echo.New()

	routes.InitPaymentsRoute(e)

	e.Start(":8000")
}
