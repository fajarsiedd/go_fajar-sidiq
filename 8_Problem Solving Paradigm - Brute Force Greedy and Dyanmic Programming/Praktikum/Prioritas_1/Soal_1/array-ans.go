package main

import (
	"fmt"
	"slices"
	"strconv"
	"strings"
)

func getAns(n int) []string {
	ans := []string{}
	for i := 0; i <= n; i++ {
		ans = append(ans, decimalToBinary(i))
	}
	return ans
}

func main() {
	fmt.Println(getAns(2))
	fmt.Println(getAns(3))
}

func decimalToBinary(n int) string {
	if n == 0 {
		return strconv.Itoa(n)
	}

	result := []string{}
	for n > 0 {
		result = append(result, strconv.Itoa(n%2))
		n /= 2
	}

	slices.Reverse(result)

	return strings.Join(result, "")
}
