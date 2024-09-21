package main

import (
	"fmt"
	"slices"
	"strings"
)

func main() {
	fmt.Println(pickVocals("alterra academy"))     // aea aae
	fmt.Println(pickVocals("sepulsa mantap jiwa")) // eua aa ia
	fmt.Println(pickVocals("kopi susu"))           // oi uu
}

func pickVocals(sentence string) string {
	vocals := []string{"a", "i", "u", "e", "o"}

	splittedStr := strings.Split(sentence, "")
	var result []string
	for _, v := range splittedStr {
		if slices.Contains(vocals, string(v)) || v == " " {
			result = append(result, v)
		}
	}

	return strings.Join(result, "")
}
