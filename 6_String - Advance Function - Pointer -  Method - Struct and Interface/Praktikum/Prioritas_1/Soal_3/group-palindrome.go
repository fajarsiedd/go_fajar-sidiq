package main

import (
	"fmt"
	"slices"
	"strings"
)

func main() {
	fmt.Println(groupPalindrome([]string{"katak", "civic", "kasur", "rusak"}))                  // [[katak, civic], kasur, rusak]
	fmt.Println(groupPalindrome([]string{"racecar", "seru", "kasur", "civic", "bilik", "kak"})) // [[racecar, civic, kak], seru, kasur, bilik]
	fmt.Println(groupPalindrome([]string{"masuk", "civic", "hahah", "garam"}))                  // [[civic, hahah], masuk, garam]
}

func groupPalindrome(words []string) []any {
   var palindromes, result []any
	for _, v := range words {
      splittedValue := strings.Split(v, "")
      slices.Reverse(splittedValue)
		reversed := strings.Join(splittedValue, "")

		if reversed == v {
			palindromes = append(palindromes, v)
		} else {
         result = append(result, v)
      }
	}

   if len(palindromes) > 0 {
      result = slices.Insert(result, 0, []any{palindromes}...)
   }

	return result
}
