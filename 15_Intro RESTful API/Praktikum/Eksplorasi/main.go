package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
)

type Response struct {
	Page       int `json:"page"`
	PerPage    int `json:"per_page"`
	Total      int `json:"total"`
	TotalPages int `json:"total_pages"`
	Data       []struct {
		ID        int    `json:"id"`
		Email     string `json:"email"`
		FirstName string `json:"first_name"`
		LastName  string `json:"last_name"`
		Avatar    string `json:"avatar"`
	} `json:"data"`
	Support struct {
		URL  string `json:"url"`
		Text string `json:"text"`
	} `json:"support"`
}

func main() {
	client := http.Client{}
	url := "https://reqres.in/api/users"

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

	var responseObject Response
	json.Unmarshal(responseData, &responseObject)

	fmt.Printf("Page: %d, Total Users: %d\n", responseObject.Page, responseObject.Total)

	for _, user := range responseObject.Data {
		fmt.Println("=======")
		fmt.Println("ID:", user.ID)
		fmt.Println("Name:", user.FirstName, user.LastName)
		fmt.Println("Email:", user.Email)
		fmt.Println("=======")
	}
}
