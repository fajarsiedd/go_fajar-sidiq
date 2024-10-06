package main

import (
	"fmt"
	"math"
	"time"
)

func main() {
	compositeCh := make(chan int)
	powCh := make(chan float64)

	go compositeNums(compositeCh)

	go powComposite(powCh, compositeCh)

	go checkEvenOdd(powCh)

	<-time.After(time.Second * 5)
}

func checkEvenOdd(powCh chan float64) {
	for powVal := range powCh {
		if int(powVal)%2 == 0 {
			fmt.Println("Ping")
		} else {
			fmt.Println("Pong")
		}
	}
}

func powComposite(powCh chan float64, compositeCh chan int) {
	for composite := range compositeCh {
		powVal := math.Pow(float64(composite), 2)
		powCh <- powVal
	}
}

func compositeNums(compositeCh chan int) {
	for i := 1; i <= 100; i++ {
		if isComposite(i) {
			compositeCh <- i
		}
	}
}

func isComposite(number int) bool {
	if number <= 3 {
		return false
	}

	for i := 4; i <= int(math.Sqrt(float64(number))); i++ {
		if number%i == 0 {
			return true
		}
	}
	return false
}
