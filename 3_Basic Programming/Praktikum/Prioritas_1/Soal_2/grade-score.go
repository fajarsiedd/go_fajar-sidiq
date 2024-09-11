package main

import "fmt"

func main() {
	var score int32

	fmt.Print("Masukkan nilai: ")
	fmt.Scan(&score)

	fmt.Println(getGradeScore(score))
}

func getGradeScore(score int32) string {
	if score >= 85 && score <= 100 {
		return "A"
	} else if score >= 70 && score <= 84 {
		return "B"
	} else if score >= 55 && score <= 69 {
		return "C"
	} else if score >= 40 && score <= 54 {
		return "D"
	} else if score >= 0 && score <= 39 {
		return "E"
	} else {
		return "Nilai Invalid"
	}
}
