package main

import "fmt"

func main() {
	fmt.Println(spinSlice([]int{1, 2, 3, 4, 5}))    // [1,5,2,4,3]
	fmt.Println(spinSlice([]int{6, 7, 8}))          // [6,8,7]
	fmt.Println(spinSlice([]int{8, 5, 4, 3, 2, 1})) // [8,1,5,2,4,3]
}

func spinSlice(numbers []int) (result []int) {
	low := 0
	high := len(numbers) - 1
	indexPosition := 0

	for len(result) < len(numbers) {
		if indexPosition == low {
			result = append(result, numbers[low])
			low++
			indexPosition = high
		} else {
			result = append(result, numbers[high])
			high--
			indexPosition = low
		}
	}

	return result
}
