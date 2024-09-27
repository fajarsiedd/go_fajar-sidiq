package main

import "fmt"

func main() {
	fmt.Println(fibX(5))  // 12
	fmt.Println(fibX(10)) // 143
}

func fibX(n int) int {
	if n == 0 {
		return n
	}

	return fibonacci(n) + fibX(n-1)
}

func fibonacci(n int) int {
	if n <= 1 {
		return n
	}

	return fibonacci(n-2) + fibonacci(n-1)
}
