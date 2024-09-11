package main

import "fmt"

func main() {
	var budget, duration, difficulty int32

	fmt.Print("Masukkan Budget: ")
	fmt.Scan(&budget)
	fmt.Print("Masukkan Waktu Pengerjaan: ")
	fmt.Scan(&duration)
	fmt.Print("Masukkan Tingkat Kesulitan: ")
	fmt.Scan(&difficulty)

	fmt.Println(calcScore(budget, duration, difficulty))
}

func calcScore(budget, duration, difficulty int32) string {
	finalScore := getBudgetScore(budget) + getDurationScore(duration) + getDifficultyScore(difficulty)

	if finalScore >= 17 && finalScore <= 24 {
		return "High"
	} else if finalScore >= 10 && finalScore <= 16 {
		return "Medium"
	} else if finalScore >= 3 && finalScore <= 9 {
		return "Low"
	} else if finalScore <= 2 {
		return "Impossible"
	} else {
		return "Undefined"
	}
}

func getBudgetScore(budget int32) int32 {
	if budget > 100 {
		return 1
	} else if budget >= 81 && budget <= 100 {
		return 2
	} else if budget >= 51 && budget <= 80 {
		return 3
	} else if budget >= 0 && budget <= 50 {
		return 4
	} else {
		return 0
	}
}

func getDurationScore(duration int32) int32 {
	if duration > 50 {
		return 1
	} else if duration >= 31 && duration <= 50 {
		return 2
	} else if duration >= 21 && duration <= 30 {
		return 5
	} else if duration >= 0 && duration <= 20 {
		return 10
	} else {
		return 0
	}
}

func getDifficultyScore(difficulty int32) int32 {
	if difficulty > 10 {
		return 0
	} else if difficulty >= 8 && difficulty <= 10 {
		return 1
	} else if difficulty >= 4 && difficulty <= 6 {
		return 5
	} else if difficulty >= 0 && difficulty <= 3 {
		return 10
	} else {
		return 0
	}
}
