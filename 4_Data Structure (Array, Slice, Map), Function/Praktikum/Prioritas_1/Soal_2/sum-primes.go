package main

import (
	"fmt"
	"slices"
)

func main() {
	fmt.Println(primeSum([]int{1, 2, 4, 5, 12, 19, 30})) // 26
	fmt.Println(primeSum([]int{45, 17, 44, 67, 11}))     // 95
	fmt.Println(primeSum([]int{15, 12, 9, 10}))          // 0
}

func primeSum(numbers []int) int {
	numbers = slices.DeleteFunc(numbers, func(n int) bool {
		return !isPrime(n)
	})

	result := 0
	for _, v := range numbers {
		result += v
	}

	return result
}

func isPrime(number int) bool {
	if number <= 1 {
		return false
	}

	for i := 2; i <= number-1; i++ {
		if number%i == 0 {
			return false
		}
	}

	return true
}
