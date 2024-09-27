package main

import "fmt"

func main() {
	fmt.Println(catalan(7))  // 429
	fmt.Println(catalan(10)) // 16796
}

func catalan(n int) int {
	// C(n) = (2n)! / ((n+1)! * n!)
	result := factorial(2*n) / (factorial(n+1) * factorial(n))

	return result
}

func factorial(n int) int {
	if n == 1 {
		return n
	}

	return n * factorial(n-1)
}
