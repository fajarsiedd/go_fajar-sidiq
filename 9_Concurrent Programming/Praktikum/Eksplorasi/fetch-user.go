package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"sync"
)

type User struct {
	ID       int    `json:"id"`
	Username string `json:"username"`
	Email    string `json:"email"`
	Name     struct {
		Firstname string `json:"firstname"`
		Lastname  string `json:"lastname"`
	} `json:"name"`
}

func main() {
	userCh := make(chan []User)

	go fetchUsers(userCh)

	users := <-userCh

	fmt.Println("Users:")
	for _, user := range users {
		fmt.Println("====")
		fmt.Println("ID:", user.ID)
		fmt.Println("Username:", user.Username)
		fmt.Println("Email:", user.Email)
		fmt.Println("First name:", user.Name.Firstname)
		fmt.Println("Last name:", user.Name.Lastname)
		fmt.Println("====")
	}
}

func fetchUsers(userCh chan []User) {
	var wg sync.WaitGroup

	url := "https://fakestoreapi.com/users"

	wg.Add(1)
	go func() {
		defer wg.Done()
		resp, err := http.Get(url)
		if err != nil {
			log.Fatalf("Failed to fetch data from API: %v", err)
		}
		defer resp.Body.Close()

		body, err := io.ReadAll(resp.Body)
		if err != nil {
			log.Fatalf("Failed to read response body: %v", err)
		}

		var users []User
		err = json.Unmarshal(body, &users)
		if err != nil {
			log.Fatalf("Failed to parse JSON: %v", err)
		}

		userCh <- users
	}()

	go func() {
		wg.Wait()
		close(userCh)
	}()
}
