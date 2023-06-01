package main

import "fmt"

// Soal01
func evenOddNumber(count int) {
	Odd := []int{}

	for i := 0; i < count; i++ {
		counts := (i * 2) + 1
		Odd = append(Odd, counts)
	}
	fmt.Println(Odd)
}

// Soal02
func powNumber(count int) {
	arr := make([]int, count)
	arr[0] = 1

	for i := 1; i < count; i++ {
		arr[i] = arr[i-1] * 2
	}
	fmt.Println(arr)
}

// Soal03
func fibonnaci(count int) {

}

func main() {
	// evenOddNumber(5)
	powNumber(5)
}
