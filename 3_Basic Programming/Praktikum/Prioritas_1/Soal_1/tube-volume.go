package main

import "fmt"

func main() {
	var r, t float32

	fmt.Print("Masukkan jari-jari: ")
	fmt.Scan(&r)
	fmt.Print("Masukkan tinggi: ")
	fmt.Scan(&t)

	fmt.Println(calcTubeVolume(r, t))
}

func calcTubeVolume(r, t float32) float32 {
	return 3.14 * (r * r) * t
}
