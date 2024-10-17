package main

import (
	"net/http"
	"slices"
	"strings"

	"github.com/labstack/echo/v4"
)

type WordReqBody struct {
	Word string `json:"word"`
}

type Words struct {
	Word       string `json:"word"`
	Length     int    `json:"length"`
	Vocals     int    `json:"num_of_vocals"`
	Palindrome bool   `json:"palindrome"`
}

type BaseResponse struct {
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

var words []Words

func main() {
	e := echo.New()

	e.GET("/words", GetWordsHandler)
	e.POST("/words", AddWordHandler)

	e.Start(":8000")
}

func GetWordsHandler(c echo.Context) error {
	return c.JSON(http.StatusOK, BaseResponse{Message: "all words", Data: words})
}

func AddWordHandler(c echo.Context) error {
	reqBody := WordReqBody{}
	c.Bind(&reqBody)

	if reqBody.Word == "" {
		return c.JSON(http.StatusBadRequest, BaseResponse{Message: "failed to add word data", Data: nil})
	}

	word := Words{
		Word:       reqBody.Word,
		Length:     len(reqBody.Word),
		Vocals:     countVocals(reqBody.Word),
		Palindrome: isPalindrome(reqBody.Word),
	}

	words = append(words, word)

	return c.JSON(http.StatusCreated, BaseResponse{Message: "word added", Data: word})
}

func countVocals(word string) int {
	vocals := []string{"a", "i", "u", "e", "o"}

	splittedStr := strings.Split(word, "")
	count := 0
	for _, v := range splittedStr {
		if slices.Contains(vocals, strings.ToLower(string(v))) {
			count++
		}
	}

	return count
}

func isPalindrome(word string) bool {
	splittedStr := strings.Split(word, "")
	i := 0
	j := len(splittedStr) - 1

	for i < j {
		if !strings.EqualFold(splittedStr[i], splittedStr[j]) {
			return false
		}

		i++
		j--
	}

	return true
}
