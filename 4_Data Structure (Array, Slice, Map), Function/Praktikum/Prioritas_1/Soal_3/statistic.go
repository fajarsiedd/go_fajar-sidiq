package main

import (
	"fmt"
	"slices"
)

func main() {
	data1 := []float64{1, 1, 2, 5, 6, 8, 12, 4, 5, 5, 5, 8, 9}

	fmt.Printf("sum: %.2f\n", sum(data1))       // 71.00
	fmt.Printf("mean: %.2f\n", mean(data1))     // 5.46
	fmt.Printf("median: %.2f\n", median(data1)) // 5.00
	fmt.Printf("mode: %.2f\n", mode(data1))     // 5.00
	fmt.Println()

	data2 := []float64{6, 7, 1, 11, 8, 12, 6, 12, 1, 7, 8, 10, 14}

	fmt.Printf("sum: %.2f\n", sum(data2))       // 103.00
	fmt.Printf("mean: %.2f\n", mean(data2))     // 7.92
	fmt.Printf("median: %.2f\n", median(data2)) // 8.00
	fmt.Printf("mode: %.2f\n", mode(data2))     // 1.00
}

func sum(data []float64) float64 {
	var result float64
	for _, v := range data {
		result += v
	}

	return result
}

func mean(data []float64) float64 {
	var sum float64
	for _, v := range data {
		sum += v
	}

	return sum / float64(len(data))
}

func median(data []float64) float64 {
	slices.Sort(data)

	if n := len(data); n%2 == 0 {
		return (data[n/2-1] + data[n/2]) / 2
	} else {
		return data[(n-1)/2]
	}
}

func mode(data []float64) (result float64) {
	frequency := map[float64]int{}
	for _, v := range data {
		_, existed := frequency[v]

		if existed {
			frequency[v] += 1
		} else {
			frequency[v] = 1
		}
	}

	var keys []float64
	for k := range frequency {
		keys = append(keys, k)
	}
	slices.Sort(keys)

	max := 0
	for _, k := range keys {
		if frequency[k] > max {
			max = frequency[k]
			result = k
		}
	}

	return result
}
