package main

import (
	"fmt"
	"strings"
)

func main() {
	var words1, words2 string

	fmt.Print("Kata Pertama: ")
	fmt.Scan(&words1)
	fmt.Print("Kata Kedua: ")
	fmt.Scan(&words2)

	fmt.Println(checkAnagram(calcFrequency(words1), calcFrequency(words2)))
}

func checkAnagram(words1, words2 map[string]int32) string {
	if len(words1) != len(words2) {
		return "Bukan Anagram"
	}

	for k := range words1 {
		if words1[k] != words2[k] {
			return "Bukan Anagram"
		}
	}

	return "Anagram"
}

func calcFrequency(words string) map[string]int32 {
	arr := strings.Split(strings.ToLower(words), "")
	result := map[string]int32{}

	for _, v := range arr {
		_, existed := result[v]

		if existed {
			result[v] += 1
		} else {
			result[v] = 1
		}
	}

	return result
}
