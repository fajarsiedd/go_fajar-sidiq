package main

import "fmt"

func main() {
	primeFactors(20)   // 2, 2, 5
	primeFactors(75)   // 3, 5, 5
	primeFactors(1024) // 2, 2, 2, 2, 2, 2, 2, 2, 2, 2
}

func primeFactors(n int) {
	if n <= 1 {
		return
	}

	for i := 2; i*i <= n; i++ {
		if n%i == 0 {
			fmt.Print(i, ", ")
			primeFactors(n / i)
			return
		}
	}
	fmt.Print(n)
	fmt.Println()
}
