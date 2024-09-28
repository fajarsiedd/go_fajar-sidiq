package main

import "fmt"

func main() {
	fmt.Println(generateRows(5))
}

func generateRows(n int) [][]int {
	var arr [][]int

	for i := 0; i < n; i++ {
		row := []int{}
		for j := 0; j <= i; j++ {
			if j == 0 || j == i {
				row = append(row, 1)
			} else {
				row = append(row, arr[i-1][j]+arr[i-1][j-1])
			}
		}

		arr = append(arr, row)
	}

	return arr
}
