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

func sumAllElement(arr []int) int {
	sum := 0
	for i := 0; i < len(arr); i++ {
		// sum = sum + arr[i]
		sum += arr[i]
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
	index := 0
	for i := 0; i < len(arr); i++ {
		if arr[i] > max {
			max = arr[i]
			index = i
		}
	}
	return index
}

func shifting(arr []int, rotate int) []int {
	var geser []int

	if rotate == 0 {
		return arr
	}

	for i := 0; i < rotate; i++ {
		geser = arr[1:len(arr)]
		geser = append(geser, arr[0])
		arr = geser
	}
	return geser
}

func linearSearch(n []int, cari int) int {
	var index int

	geser := n[1:len(n)]
	geser = append(geser, n[0])

	for i := 0; i < len(geser); i++ {
		if geser[i] == cari {
			index = i
		}
	}
	return index
}

func main() {
	// NOMOR 1
	nums := []int{1, 2, 3, 4, 5}
	fmt.Println(sumAllElement(nums))
	fmt.Println(sumAllElement([]int{10, 20, 30}))

	// NOMOR 2
	angka := []int{89, 76, 10, 99, 10}
	fmt.Println(maxElement(angka))

	// NOMOR 3
	number := []int{1, 1, 12, 10, 5, 12}
	fmt.Println(smallestIndex(number))

	// NOMOR 4
	arr := []int{2, 3, 5, 7, 11}
	fmt.Println(shifting(arr, 0))

	// NOMOR 5
	n := []int{2, 3, 5, 7, 11}
	fmt.Println(linearSearch(n, 7))

	list := initArray(5)
	fmt.Println(list)
	displayArray(initArray(10))
}
