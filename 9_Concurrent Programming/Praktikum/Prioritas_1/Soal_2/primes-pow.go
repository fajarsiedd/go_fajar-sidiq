package main

import (
	"fmt"
	"math"
	"time"
)

func main() {
	primeCh := make(chan int)

	go primeNumbers(primeCh)

	go powPrimes(primeCh)

	<-time.After(time.Millisecond * 100)
}

func powPrimes(primeCh chan int) {
	for prime := range primeCh {
		powVal := math.Pow(float64(prime), 2)
		fmt.Println(powVal)
	}
}

func primeNumbers(primeCh chan int) {
	for i := 1; i <= 100; i++ {
		if isPrime(i) {
			primeCh <- i
		}
	}
}

func isPrime(number int) bool {
	if number <= 1 {
		return false
	}
	if number <= 3 {
		return true
	}
	if number%2 == 0 {
		return false
	}

	for i := 3; i <= int(math.Sqrt(float64(number))); i += 2 {
		if number%i == 0 {
			return false
		}
	}
	return true
}
