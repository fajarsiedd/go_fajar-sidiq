package main

import "fmt"

func MinCost(jumps []int, n int) int {
	if n == 1 {
		return 0
	}

	if n == 2 {
		result := Abs(jumps[n-1], jumps[n-2])

		return result
	}

	step1 := MinCost(jumps, n-1) + Abs(jumps[n-1], jumps[n-2])
	step2 := MinCost(jumps, n-2) + Abs(jumps[n-1], jumps[n-3])

	if step1 < step2 {
		return step1
	} else {
		return step2
	}
}

func Frog(jumps []int) int {
	return MinCost(jumps, len(jumps))
}

func Abs(a, b int) int {
	if a > b {
		return a - b
	} else {
		return b - a
	}
}

func main() {
	fmt.Println(Frog([]int{10, 30, 40, 20}))         // 30
	fmt.Println(Frog([]int{30, 10, 60, 10, 60, 50})) // 40
}
