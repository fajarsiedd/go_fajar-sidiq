package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(spinString("alta"))    // aalt
	fmt.Println(spinString("alterra")) // aalrtre
	fmt.Println(spinString("sepulsa")) // saesplu
}

func spinString(sentence string) string {
	low := 0
	high := len(sentence) - 1
	indexPosition := 0

	var result []string
	for len(result) < len(sentence) {
		if indexPosition == low {
			result = append(result, string(sentence[low]))
			low++
			indexPosition = high
		} else {
			result = append(result, string(sentence[high]))
			high--
			indexPosition = low
		}
	}

	return strings.Join(result, "")
}
