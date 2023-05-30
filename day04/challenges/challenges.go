package main

import (
	"fmt"
)

// func init array with 1 parameter & has return array
func initArray(n int) []int {
	arr := make([]int, n)
	for i := 0; i < n; i++ {
		arr[i] = i
	}
	return arr
}

func displayArray(arr []int) {
	for i := 0; i < len(arr); i++ {
		fmt.Print(i, ", ")
	}
}

func sumAllElement(arr []int) int {
	sum := 0
	for i := 0; i < len(arr); i++ {
		sum += arr[i] //sum = sum +arr[1]
	}
	return sum
}

func maxElement(arr []int) int {
	max := 0
	/*	for i := 0; i < len(arr); i++ {
			if arr[i] > max {
				max = arr[i]
			}
		}
		return max */
	for _, v := range arr {
		if v > max {
			max = v
		}
	}
	return max
}

func smallestIndex(arr []int) int {
	index := 0
	small := 0
	/*	for i, v := range arr {
		if v > index {
			index = v
			small = i
		}
	} */
	for i := 0; i < len(arr); i++ {
		if arr[i] > index {
			index = arr[i]
			small = i
		}
	}
	return small
}

func shifting(arr []int, rotate int) []int {

	for i := 0; i < rotate; i++ {
		index := arr[0]
		indexTerakhir := len(arr) - 1
		arr[0] = arr[1]
		arr[1] = arr[2]
		arr[2] = arr[3]
		arr[3] = arr[indexTerakhir]

		arr[indexTerakhir] = index

	}
	return arr
}

func linearSearch(arr []int, search int) int {
	index := arr[0]
	indexTerakhir := len(arr) - 1
	arr[0] = arr[1]
	arr[1] = arr[2]
	arr[2] = arr[3]
	arr[3] = arr[indexTerakhir]

	arr[indexTerakhir] = index
	for i, value := range arr {
		if value == search {
			index = i
		}
	}
	return index
}

func main() {
	// list := initArray(5)
	// displayArray(list)
	// displayArray(initArray(10))
	// num := []int{1, 2, 3, 4, 5}
	// fmt.Println(sumAllElement(num))
	// fmt.Println(sumAllElement([]int{12, 13, 14, 45, 32}))
	// fmt.Println(maxElement([]int{89, 76, 10, 99, 10}))
	// sfmt.Println(smallestIndex([]int{1, 1, 12, 10, 5, 12}))
	fmt.Println(shifting([]int{2, 3, 5, 7, 11}, 1))
	fmt.Println(linearSearch([]int{2, 3, 5, 7, 11}, 7))
}
