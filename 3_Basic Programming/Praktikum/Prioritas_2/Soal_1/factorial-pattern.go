package main

import (
	"fmt"
	"strings"
)

func main() {
	var number int

	fmt.Print("Masukkan angka: ")
	fmt.Scan(&number)

	fmt.Println(getFactorialPattern(number))
}

func getFactorialPattern(number int) string {
	var result []string

	for i := 1; i <= number; i++ {
		if number%i == 0 {
			if i%2 == 0 {
				result = append(result, "I")
			} else {
				result = append(result, "O")
			}
		}
	}

	return strings.Join(result, "")
}
