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

type DictionaryResponse []struct {
	Word     string `json:"word"`
	Meanings []struct {
		PartOfSpeech string `json:"partOfSpeech"`
		Definitions  []struct {
			Definition string `json:"definition"`
			Synonyms   []any  `json:"synonyms"`
			Antonyms   []any  `json:"antonyms"`
			Example    string `json:"example"`
		} `json:"definitions"`
		Synonyms []string `json:"synonyms"`
		Antonyms []string `json:"antonyms"`
	} `json:"meanings"`
}

type WordReqBody struct {
	Word string `json:"word"`
}

type Words struct {
	ID           string `json:"id" gorm:"primaryKey"`
	Word         string `json:"word"`
	PartOfSpeech string `json:"part_of_speech"`
	Definition   string `json:"definition"`
}

type BaseResponse struct {
	Status  bool        `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

var DB *gorm.DB

func initDB() {
	dsn := "root:@tcp(127.0.0.1:3306)/dictionary?charset=utf8mb4&parseTime=True&loc=Local"
	var err error
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("failed to connect database")
	}
}

func migrateDB() {
	DB.AutoMigrate(&Words{})

	// Add table suffix when creating tables
	DB.Set("gorm:table_options", "ENGINE=InnoDB").AutoMigrate(&Words{})
}

func main() {
	initDB()
	migrateDB()

	e := echo.New()

	e.GET("/api/v1/words", GetWordsHandler)
	e.GET("/api/v1/words/:id", GetWordDetailHandler)
	e.POST("/api/v1/words", AddWordHandler)
	e.DELETE("/api/v1/words/:id", DeleteWordHandler)

	e.Start(":8000")
}

func AddWordHandler(c echo.Context) error {
	wordReqBody := WordReqBody{}
	c.Bind(&wordReqBody)

	dictionary := getDict(wordReqBody.Word)

	word := Words{
		ID:           uuid.New().String(),
		Word:         dictionary[0].Word,
		PartOfSpeech: dictionary[0].Meanings[0].PartOfSpeech,
		Definition:   dictionary[0].Meanings[0].Definitions[0].Definition,
	}

	result := DB.Create(&word)

	if result.Error != nil {
		var errorMessage string
		if result.Error != nil {
			errorMessage = result.Error.Error()
		} else {
			errorMessage = "unknown error"
		}
		return c.JSON(http.StatusInternalServerError, BaseResponse{Status: false, Message: errorMessage, Data: nil})
	}

	return c.JSON(http.StatusCreated, BaseResponse{Status: true, Message: "data added successfully", Data: word})
}

func GetWordsHandler(c echo.Context) error {
	words := []Words{}
	result := DB.Find(&words)

	if result.Error != nil {
		var errorMessage string
		if result.Error != nil {
			errorMessage = result.Error.Error()
		} else {
			errorMessage = "unknown error"
		}
		return c.JSON(http.StatusInternalServerError, BaseResponse{Status: false, Message: errorMessage, Data: nil})
	}

	return c.JSON(http.StatusOK, BaseResponse{Status: true, Message: "data fetched successfully", Data: words})
}

func GetWordDetailHandler(c echo.Context) error {
	id := c.Param("id")
	word := Words{ID: id}
	result := DB.First(&word)

	if result.Error != nil {
		var errorMessage string
		if result.Error != nil {
			errorMessage = result.Error.Error()
		} else {
			errorMessage = "Unknown error"
		}
		return c.JSON(http.StatusInternalServerError, BaseResponse{Status: false, Message: errorMessage, Data: nil})
	}

	return c.JSON(http.StatusOK, BaseResponse{Status: true, Message: "data fetched successfully", Data: word})
}

func DeleteWordHandler(c echo.Context) error {
	id := c.Param("id")
	result := DB.Delete(&Words{ID: id})

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

func getDict(word string) DictionaryResponse {
	client := http.Client{}
	url := "https://api.dictionaryapi.dev/api/v2/entries/en/" + word

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		fmt.Print(err.Error())
		os.Exit(1)
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

	var responseObject DictionaryResponse
	json.Unmarshal(responseData, &responseObject)

	return responseObject
}
