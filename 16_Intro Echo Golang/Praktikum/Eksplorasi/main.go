package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type NutritionsResponse []struct {
	Name                string      `json:"name"`
	Calories            interface{} `json:"calories"`
	ServingSizeG        interface{} `json:"serving_size_g"`
	FatTotalG           float64     `json:"fat_total_g"`
	FatSaturatedG       float64     `json:"fat_saturated_g"`
	ProteinG            interface{} `json:"protein_g"`
	SodiumMg            float64     `json:"sodium_mg"`
	PotassiumMg         float64     `json:"potassium_mg"`
	CholesterolMg       float64     `json:"cholesterol_mg"`
	CarbohydratesTotalG float64     `json:"carbohydrates_total_g"`
	FiberG              float64     `json:"fiber_g"`
	SugarG              float64     `json:"sugar_g"`
}

type FruitReqBody struct {
	Name  string `json:"name"`
	Price int    `json:"price"`
}

type Fruits struct {
	ID        string     `json:"id"`
	Name      string     `json:"name"`
	Price     int        `json:"price"`
	Nutrition Nutritions `json:"nutrition"`
}

type Nutritions struct {
	Calories      interface{} `json:"calories"`
	Fat           float64     `json:"fat"`
	Sugar         float64     `json:"sugar"`
	Carbohydrates float64     `json:"carbohydrates"`
	Protein       interface{} `json:"protein"`
}

type BaseResponse struct {
	Status  bool        `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

var DB *gorm.DB

func initDB() {
	dsn := "root:@tcp(127.0.0.1:3306)/fruits?charset=utf8mb4&parseTime=True&loc=Local"
	var err error
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("failed to connect database")
	}
}

func main() {
	initDB()

	e := echo.New()

	e.GET("/api/v1/fruits", GetFruitsHandler)
	e.GET("/api/v1/fruits/:id", GetFruitDetailHandler)
	e.POST("/api/v1/fruits", AddFruitHandler)
	e.DELETE("/api/v1/fruits/:id", DeleteFruitHandler)

	e.Start(":8000")
}

func AddFruitHandler(c echo.Context) error {
	fruitReqBody := FruitReqBody{}
	c.Bind(&fruitReqBody)

	nutritions := getNutritions(fruitReqBody.Name)

	fruit := Fruits{
		ID:        uuid.New().String(),
		Name:      fruitReqBody.Name,
		Price:     fruitReqBody.Price,
		Nutrition: nutritions,
	}

	result := DB.Create(&fruit)

	if result.Error != nil {
		var errorMessage string
		if result.Error != nil {
			errorMessage = result.Error.Error()
		} else {
			errorMessage = "unknown error"
		}
		return c.JSON(http.StatusInternalServerError, BaseResponse{Status: false, Message: errorMessage, Data: nil})
	}

	return c.JSON(http.StatusCreated, BaseResponse{Status: true, Message: "data added successfully", Data: fruit})
}

func GetFruitsHandler(c echo.Context) error {
	fruits := []Fruits{}
	result := DB.Find(&fruits)

	if result.Error != nil {
		var errorMessage string
		if result.Error != nil {
			errorMessage = result.Error.Error()
		} else {
			errorMessage = "unknown error"
		}
		return c.JSON(http.StatusInternalServerError, BaseResponse{Status: false, Message: errorMessage, Data: nil})
	}

	return c.JSON(http.StatusOK, BaseResponse{Status: true, Message: "data fetched successfully", Data: fruits})
}

func GetFruitDetailHandler(c echo.Context) error {
	id := c.Param("id")
	fruit := Fruits{}
	result := DB.First(&fruit, id)

	if result.Error != nil {
		var errorMessage string
		if result.Error != nil {
			errorMessage = result.Error.Error()
		} else {
			errorMessage = "Unknown error"
		}
		return c.JSON(http.StatusInternalServerError, BaseResponse{Status: false, Message: errorMessage, Data: nil})
	}

	return c.JSON(http.StatusOK, BaseResponse{Status: true, Message: "data fetched successfully", Data: fruit})
}

func DeleteFruitHandler(c echo.Context) error {
	id := c.Param("id")
	result := DB.Delete(&Fruits{}, id)

	if result.Error != nil {
		var errorMessage string
		if result.Error != nil {
			errorMessage = result.Error.Error()
		} else {
			errorMessage = "unknown error"
		}
		return c.JSON(http.StatusInternalServerError, BaseResponse{Status: false, Message: errorMessage, Data: nil})
	}

	return c.JSON(http.StatusNoContent, nil)
}

// Hanya untuk ilustrasi, karena API berbayar
func getNutritions(fruit string) Nutritions {
	client := http.Client{}
	url := "https://api.api-ninjas.com/v1/nutrition?query=" + fruit

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		fmt.Print(err.Error())
		os.Exit(1)
	}

	req.Header = http.Header{
		"X-Api-Key": {"YOUR_API_KEY"},
	}

	res, err := client.Do(req)
	if err != nil {
		fmt.Print(err.Error())
		os.Exit(1)
	}

	responseData, err := io.ReadAll(res.Body)
	if err != nil {
		log.Fatal(err)
	}

	var responseObject NutritionsResponse
	json.Unmarshal(responseData, &responseObject)

	nutritions := Nutritions{}

	if len(responseObject) != 0 {
		data := responseObject[0]

		nutritions.Calories = data.Calories
		nutritions.Carbohydrates = data.CarbohydratesTotalG
		nutritions.Fat = data.FatTotalG
		nutritions.Protein = data.ProteinG
		nutritions.Sugar = data.SugarG
	}

	return nutritions
}
