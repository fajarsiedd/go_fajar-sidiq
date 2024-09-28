package main

import "fmt"

func SimpleEquations(a, b, c int) {
	for x := 1; x <= 100; x++ {
		for y := x; y <= 100; y++ {
			for z := y; z <= 100; z++ {
				if (x+y+z == a) && (x*y*z == b) && (x*x+y*y+z*z == c) {
					fmt.Println(x, y, z)
					return
				}
			}
		}
	}

	fmt.Println("no solution")
}

func main() {
	SimpleEquations(1, 2, 3)  // no solution
	SimpleEquations(6, 6, 14) // 1 2 3
}
