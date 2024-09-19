package main

import (
	"fmt"
	"slices"
)

func main() {
	fmt.Println(groupPrime([]int{2, 3, 4, 5, 6, 7, 8})) // [[2,3,5,7],4,6,8]
	fmt.Println(groupPrime([]int{15, 24, 3, 9, 5}))     // [[3,5],15,24,9]
	fmt.Println(groupPrime([]int{4, 8, 9, 12}))         // [4,8,9,12]
}

func groupPrime(numbers []int) (result []any) {
	var primes []int
	for _, v := range numbers {
		if isPrime(v) {
			primes = append(primes, v)
		} else {
			result = append(result, v)
		}
	}

	if len(primes) > 0 {
		result = slices.Insert(result, 0, []any{primes}...)
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
