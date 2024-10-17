package main

import (
	"net/http"
	"slices"
	"strings"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

type Foods struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Price       int    `json:"price"`
	Description string `json:"description"`
}

type BaseResponse struct {
	Status  bool        `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

var foods []Foods

func main() {
	e := echo.New()

	e.GET("/api/v1/foods", GetFoodsHandler)
	e.GET("/api/v1/foods/:id", GetFoodDetailHandler)
	e.POST("/api/v1/foods", AddFoodHandler)
	e.PUT("/api/v1/foods/:id", UpdateFoodHandler)
	e.DELETE("/api/v1/foods/:id", DeleteFoodHandler)

	e.Start(":8000")
}

func GetFoodsHandler(c echo.Context) error {
	return c.JSON(http.StatusOK, BaseResponse{Status: true, Message: "Success", Data: foods})
}

func GetFoodDetailHandler(c echo.Context) error {
	id := c.Param("id")

	index, found := findFoods(id)

	if !found {
		return c.JSON(http.StatusBadRequest, BaseResponse{Status: false, Message: "Data not found", Data: nil})
	}

	return c.JSON(http.StatusOK, BaseResponse{Status: true, Message: "Success", Data: foods[index]})
}

func AddFoodHandler(c echo.Context) error {
	food := Foods{ID: uuid.New().String()}
	c.Bind(&food)

	foods = append(foods, food)

	return c.JSON(http.StatusCreated, BaseResponse{Status: true, Message: "Success", Data: food})
}

func UpdateFoodHandler(c echo.Context) error {
	id := c.Param("id")
	food := Foods{ID: id}
	c.Bind(&food)

	index, found := findFoods(id)

	if !found {
		return c.JSON(http.StatusBadRequest, BaseResponse{Status: false, Message: "Data not found", Data: nil})
	}

	foods[index] = food

	return c.JSON(http.StatusOK, BaseResponse{Status: true, Message: "Success", Data: food})
}

func DeleteFoodHandler(c echo.Context) error {
	id := c.Param("id")

	index, found := findFoods(id)

	if !found {
		return c.JSON(http.StatusBadRequest, BaseResponse{Status: false, Message: "Data not found", Data: nil})
	}

	foods = append(foods[:index], foods[index+1:]...)

	return c.JSON(http.StatusNoContent, nil)
}

func findFoods(id string) (foodIndex int, isFound bool) {
	slices.SortFunc(foods, func(a, b Foods) int {
		return strings.Compare(a.ID, b.ID)
	})

	index, found := slices.BinarySearchFunc(foods, Foods{ID: id}, func(a, b Foods) int {
		return strings.Compare(a.ID, b.ID)
	})

	return index, found
}
