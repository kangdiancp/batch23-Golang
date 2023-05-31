package main

import "fmt"

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

// Nomor 1 :
func sumAllElement(arr []int) int {
	sum := 0
	for i := 0; i < len(arr); i++ {
		// sum = sum + arr[i]
		sum += arr[i]
	}
	return sum
}

// Nomor 2 :
func maxElement(arr []int) int {
	max := arr[0]
	for i := 1; i < len(arr); i++ {
		if arr[i] > max {
			max = arr[i]
		}
	}
	return max
}

// Nomor 3 :
func smallestIndex(arr []int) int {
	smallest := arr[0]
	smallestIndex := 0
	for i := 1; i < len(arr); i++ {
		if arr[i] > smallest {
			smallest = arr[i]
			smallestIndex = i
		}
	}
	return smallestIndex
}

// Nomor 4 :
func shifting(arr []int, rotate int) []int {
	length := len(arr)
	shiftArr := make([]int, length)

	for i := 0; i < length; i++ {
		shiftIndex := (i + rotate) % length
		shiftArr[i] = arr[shiftIndex]
	}
	return shiftArr
}

// Nomor 5 :
func linearSearch(arr []int, search int) int {
	searchIndex := 0
	for i, num := range arr {
		if num == search {
			searchIndex = i
		}
	}
	return searchIndex
}

func main() {

	/*
		list := initArray(5)
		displayArray(list)
		fmt.Println()
		displayArray(initArray(10))
		fmt.Println(list)
	*/

	// Nomor 1:
	fmt.Println(sumAllElement([]int{11, 32, 54, 35, 55}))

	// Nomor 2:
	fmt.Println(maxElement([]int{89, 76, 10, 99, 10}))

	// Nomor 3:
	fmt.Println(smallestIndex([]int{1, 1, 12, 10, 5, 12}))

	// Nomor 4:
	fmt.Println(shifting([]int{2, 3, 5, 7, 11}, 1))

	// Nomor 5 :
	fmt.Println(linearSearch([]int{2, 3, 5, 7, 11}, 7))

}
