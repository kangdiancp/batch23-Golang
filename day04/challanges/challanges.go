package main

import (
	"fmt"
)

// functioniniArray with 1 parameter & has return array

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
		//sum = sum + arr[i]
		sum += arr[i]

	}
	return sum
}

func maxElement(arr []int) int {
	hasil := 0

	for _, value := range arr {
		if hasil < value {
			hasil = value
		}
	}
	fmt.Println(hasil)
	return hasil
}

func smallesIndex(arr []int) int {
	small := 0
	max := 0

	for i := 0; i < len(arr); i++ {
		if max < arr[i] {
			max = arr[i]
			small = i
		}
	}
	return small
}
func shifting(arr []int, rotate int) []int {
	temp := 0
	panjangArr := len(arr)

	for i := 0; i < rotate; i++ {
		temp = arr[panjangArr-1]

		arr[panjangArr-1] = arr[0]

		arr[0] = arr[1]
		arr[1] = arr[2]
		arr[2] = arr[3]
		arr[3] = temp
	}

	return arr

}

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
	//nums := []int{1, 2, 3, 4, 5}
	//fmt.Println(sumAllElement(nums))
	//fmt.Println(sumAllElement([]int{12, 13, 14, 45, 32}))
	//list := initArray(5)
	//isplayArray(list)
	//fmt.Println()
	// displayArray(initArray(10))
	//fmt.Println(list)
	//fmt.Println(smallesIndex([]int{1, 1, 12, 10, 5, 12}))
	//fmt.Println(maxElement([]int{89, 76, 10, 99, 10}))
	//fmt.Println(shifting([]int{2, 3, 5, 7, 11}, 3))
	fmt.Println(shifting([]int{2, 3, 5, 7, 11}, 7))

}
