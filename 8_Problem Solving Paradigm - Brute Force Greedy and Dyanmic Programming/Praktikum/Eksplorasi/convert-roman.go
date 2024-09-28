package main

import "fmt"

func ConvertToRoman(num int) string {
	vals := []int{1000, 900, 500, 400, 100, 90, 50, 40, 10, 9, 5, 4, 1}
	symbols := []string{"M", "CM", "D", "CD", "C", "XC", "L", "XL", "X", "IX", "V", "IV", "I"}

	roman := ""
	for i := 0; i < len(vals); i++ {
		for num >= vals[i] {
			roman += symbols[i]
			num -= vals[i]
		}
	}
	return roman
}

func main() {
	fmt.Println(ConvertToRoman(4))    // IV
	fmt.Println(ConvertToRoman(9))    // IX
	fmt.Println(ConvertToRoman(23))   // XXIII
	fmt.Println(ConvertToRoman(2021)) // MMXXI
	fmt.Println(ConvertToRoman(1646)) // MDCXLVI
}
