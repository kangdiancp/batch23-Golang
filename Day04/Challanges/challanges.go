package main

import "fmt"

// function iniArray witch 1 parameter & has return array
// func initArray(count int) []int {
// 	arr := make([]int, count)
// 	for i := 1; i < count; i++ {
// 		arr[i] = i
// 	}
// 	return arr
// }

// func displayArray(arr []int) {
// 	for i := 0; i < len(arr); i++ {
// 		fmt.Print(i, ", ")
// 	}
// }

// Soal1
func sumAllElement(arr []int) int {
	sum := 0
	for i := 0; i < len(arr); i++ {
		sum += arr[i]
	}
	return sum
}

// Soal2 = {89, 76, 10, 99, 10} --> Max Element = 99
func maxElement(arr []int) int {
	max := arr[0]
	for i := 0; i < len(arr); i++ {
		if arr[i] > max {
			max = arr[i]
		}
	}
	return max
}

// Soal03 = {1, 1, 12, 10, 5, 12} --> Small Index = 2
func smallIndex(arr []int) int {
	max := 0
	min := 0

	for i := 0; i < len(arr); i++ {
		if max < arr[i] {
			max = arr[i]
			min = i
		}
	}
	return min
}

// Soal04 {2, 3, 5, 7, 11} - Rotate -> {3, 5, 7, 11, 2}
func shifting(arr []int, rotate int) []int {
	temp := 0
	panjangArr := len(arr)

	for i := 0; i < rotate; i++ {
		temp = arr[panjangArr-1]

		arr[temp] = arr[0]

		arr[0] = arr[1]
		arr[1] = arr[2]
		arr[2] = arr[3]
		arr[3] = temp
	}
	return arr
}

func soal4(arr []int, rotate int) []int {
	temp := 0
	lengArr := len(arr)

	for i := 0; i < rotate; i++ {
		temp = arr[lengArr-1]

		arr[lengArr-1] = arr[0]

		arr[0] = arr[1]
		arr[1] = arr[2]
		arr[2] = arr[3]
		arr[3] = temp
	}
	return arr
}

// Soal05 {2, 3, 5, 7, 11} --> {3, 5, 7, 11, 2}
func linierSearch(arr []int, search int) int {
	temp := 0
	indexSearch := 0

	panjangArr := len(arr)

	temp = arr[panjangArr-1]

	arr[panjangArr-1] = arr[0]

	arr[0] = arr[1]
	arr[1] = arr[2]
	arr[2] = arr[3]
	arr[3] = temp

	for i := 0; i < len(arr); i++ {
		if arr[i] == search {
			indexSearch = i
		}
	}
	return indexSearch
}

func main() {
	// list := initArray(5)
	// displayArray(list)
	// fmt.Println()
	// displayArray(initArray(5))
	// fmt.Println(list)

	// Soal01
	// list := []int{1, 2, 3, 4, 5}
	// fmt.Println(sumAllElement(list))
	// fmt.Println(sumAllElement([]int{1, 2, 3, 4, 5}))

	// Soal02
	// fmt.Println(maxElement([]int{89, 76, 10, 99, 10}))

	// Soal03
	// fmt.Println(smallIndex([]int{1, 1, 12, 10, 5, 12}))

	// Soal04
	// count4 := []int{2, 3, 5, 7, 11}
	// fmt.Println(shifting(count4[0:], 2))

	// Soal05
	count5 := []int{2, 3, 5, 7, 11}
	fmt.Println(linierSearch(count5[0:], 3))

}
