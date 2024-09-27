package main

import (
	"cmp"
	"fmt"
	"slices"
)

type Student struct {
	Name         string
	MathScore    int
	ScienceScore int
	EnglishScore int
}

func main() {
	students := []Student{
		{"jamie", 67, 88, 90},
		{"michael", 77, 77, 90},
		{"aziz", 100, 65, 88},
		{"jamal", 50, 80, 75},
		{"eric", 70, 80, 65},
		{"john", 80, 76, 68},
	}

	getInfo(students)
	// output:
	// best student in math: aziz with 100
	// best student in science: jamie with 88
	// best student in english: jamie with 90
	// best student in average: aziz with 84

}

func getInfo(students []Student) {
	// Menggunakan package Slices
	bestInMath := slices.MaxFunc(students, func(a, b Student) int {
		return cmp.Compare(a.MathScore, b.MathScore)
	})
	bestInEnglish := slices.MaxFunc(students, func(a, b Student) int {
		return cmp.Compare(a.EnglishScore, b.EnglishScore)
	})

	// Manual
	bestInScience := getBestIntScience(students)
	bestInAvgName, bestInAvgValue := getBestInAverage(students)

	// Output
	fmt.Printf("best student in math: %s with %d\n", bestInMath.Name, bestInMath.MathScore)
	fmt.Printf("best student in science: %s with %d\n", bestInScience.Name, bestInScience.ScienceScore)
	fmt.Printf("best student in english: %s with %d\n", bestInEnglish.Name, bestInEnglish.EnglishScore)
	fmt.Printf("best student in average: %s with %d", bestInAvgName, bestInAvgValue)
}

func getBestInAverage(students []Student) (string, int) {
	if len(students) == 0 {
		return "", 0
	}

	best := students[0]
	for _, v := range students {
		averageA := (v.EnglishScore + v.MathScore + v.ScienceScore) / 3
		averageB := (best.EnglishScore + best.MathScore + best.ScienceScore) / 3

		if averageA > averageB {
			best = v
		}
	}

	return best.Name, (best.EnglishScore + best.MathScore + best.ScienceScore) / 3
}

func getBestIntScience(students []Student) Student {
	if len(students) == 0 {
		return Student{}
	}

	best := students[0]
	for _, v := range students {
		if v.ScienceScore > best.ScienceScore {
			best = v
		}
	}

	return best
}
