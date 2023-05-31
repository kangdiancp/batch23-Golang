package main

import (
	"fmt"
)

func initArray(n int) []int {
	arr := make([]int, n)
	for i := 0; i < n; i++ {
		arr[i] = i

	}

	return arr
}

func displayArray(arr []int) {
	for i := 0; i < len(arr); i++ {
		fmt.Println(i, ",")
	}
}

func sumAllElement(arr []int) int {
	sum := 0
	for i := 0; i < len(arr); i++ {
		sum = sum + arr[i]

	}
	return sum
}

func maxElement(arr []int) int {
	max := arr[0]
	for i := 0; i < len(arr); i++ {
		if arr[i] > max {
			max = arr[i]
		}
	}
	return max
}

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

func shifting(arr []int, rotate int) []int {
	index := 0
	indexAkhir := len(arr)

	for i := 0; i < rotate; i++ {
		index = arr[indexAkhir-1]

		arr[indexAkhir-1] = arr[0]

		arr[0] = arr[1]
		arr[1] = arr[2]
		arr[2] = arr[3]
		arr[3] = index
	}
	return arr
}

func linierSearch(arr []int, search int) int {
	index := 0
	indexsearch := 0
	indexAkhir := len(arr)

	arr[indexAkhir-1] = arr[0]

	arr[0] = arr[1]
	arr[1] = arr[2]
	arr[2] = arr[3]
	arr[3] = index
	for i := 0; i < len(arr); i++ {
		if arr[i] == search {
			indexsearch = i
		}
	}
	return indexsearch
}


func main() {
	//fmt.Println(sumAllElement([]int{10, 20, 30, 40, 50}))
	//fmt.Println(maxElement([]int{89, 76, 10, 99}))
	//fmt.Println(smallestIndex([]int{1, 1, 12, 10, 5, 12}))
	//fmt.Println(shifting([]int{2, 3, 5, 7, 11}, 3))
	fmt.Println(linierSearch([]int{2, 3, 5, 7, 11}, 3))

}
