package main

import (
	"fmt"
	"strings"
	"time"
)

func main() {
	reverse("kasur")
}

func reverse(word string) {
	go func() {
		msgChar := strings.Split(word, "")
		result := ""

		for i := len(msgChar) - 1; i >= 0; i-- {
			result += msgChar[i]
			fmt.Println(result)
			time.Sleep(time.Second * 3)
		}
	}()

	time.Sleep(time.Duration(len(word)*3) * time.Second)
}
