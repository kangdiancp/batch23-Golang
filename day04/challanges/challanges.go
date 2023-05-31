package main

import (
	"fmt"
)

// function initArray with 1 parameter & has return array
func initArray(n int) []int {
	arr := make([]int, n)
	for i := 0; i < n; i++ {
		arr[i] = i
	}
	return arr
}

func displayArray(arr []int) {
	for i := 0; i < len(arr); i++ {
		fmt.Print(i, ",")
	}
}

// no 1
func sumAllElement(arr []int) int {
	sum := 0
	for i := 0; i < len(arr); i++ {
		// atau bisa juga tulis fungsi seperti ini (sum = sum+arr[i])
		sum += arr[i]
	}
	return sum
}

// no 2
func maxElement(arr []int) int {
	max := arr[0]
	// max := 0
	for i := 0; i < len(arr); i++ {
		if arr[i] > max {
			max = arr[i]
		}
	}
	return max
}

// no 3
func smallestIndex(arr []int) int {
	max := arr[0]
	index := arr[0]
	for i := 0; i < len(arr); i++ {
		if arr[i] > max {
			max = arr[i]
			index = i
		}
	}
	return index
}

// no 4
func shifting(arr []int, rotate int) {
	length := len(arr)
	temp := make([]int, rotate)

	if length == 0 || rotate == 0 {
		return
	}

	rotate = rotate % length
	if rotate < 0 {
		rotate = rotate + length
	}

	copy(temp, arr[length-rotate:])
	copy(arr[rotate:], arr[:length-rotate])
	copy(arr, temp)
}

//no 5
// func linearSearch(n []int, 7){
// 	n := 0
// 	for i := 0; i < len(n); i++ {
// 		if n[i] == 7 {
// 			return i
// 			}
// 		}
// 			return -1
// }

// no 5
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
	// displayArray(initArray(10))
	// //fmt.Println(list)

	// untuk menjalankan fungsi no 1
	// nums := []int{1, 2, 3, 4, 5}
	// fmt.Println(sumAllElement(nums))
	// fmt.Println(sumAllElement([]int{12, 13, 14, 45, 32}))

	//untuk menjalankan fungsi no 2
	// nums := []int{1, 2, 3, 4, 5}
	// maxElement(nums)
	// fmt.Println(maxElement(nums))
	// fmt.Println(maxElement([]int{89, 76, 10, 99}))

	//untuk menjalankan fungsi no 3
	// fmt.Println(smallestIndex([]int{1, 1, 12, 10, 5, 12}))

	//untuk menjalankan fungsi no 4
	// arr := []int{2, 3, 5, 7, 11}
	// rotate := 4
	// fmt.Println("sebelum digeser :", arr)
	// shifting(arr, rotate)
	// fmt.Println("setelah digeser :", arr)

	//untuk menjalankan fungsi no 5
	// fmt.Println(linearSearch([]int{2, 3, 5, 7, 11}))
	nums := []int{2, 3, 5, 7, 11}
	fmt.Println(linierSearch(nums[0:], 7))
}
