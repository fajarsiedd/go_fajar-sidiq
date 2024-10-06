package main

import (
	"go-math-api/routes"

	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()

	routes.CalculatorRoutes(e)

	e.Start(":1323")
}
