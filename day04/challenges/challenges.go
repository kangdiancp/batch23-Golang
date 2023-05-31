package main

import "fmt"

// Function initArray with 1 parameter & has return array
func initArray(n int) []int {
	arr := make([]int, n)

	for i := 0; i < n; i++ {
		arr[i] = i
	}

	return arr
}

// 1.
func sumAllElement(arr []int) int {
	var result int

	for i := 0; i < len(arr); i++ {
		// result = result + arr[i]
		result += arr[i]
	}

	return result
}

// 2.
func maxElement(arr []int) int {
	var result int

	for i := 0; i < len(arr); i++ {
		if arr[i] > result {
			result = arr[i]
		}
	}

	return result
}

// 3.
func smallestIndex(arr []int) int {
	var max int
	var result int

	for i := 0; i < len(arr); i++ {
		if arr[i] > max {
			max = arr[i]
			result = i
		}
	}

	return result
}

// 4.
func shiftingArray(arr []int, rotate int) []int {
	for i := 0; i < rotate; i++ {
		first := arr[0]

		for j := 0; j < (len(arr))-1; j++ {
			arr[j] = arr[j+1]
		}

		arr[(len(arr))-1] = first
	}

	return arr
}

// 5.
func linearSearch(n []int, target int) int {
	for i, v := range n {
		if v == target {
			return i
		}
	}

	return -1
}

func displayArray(arr []int) {
	for i := 0; i < len(arr)-1; i++ {
		fmt.Print(i, ", ")
	}
	fmt.Print(arr[len(arr)-1])
	fmt.Println(" ")
}

func main() {
	list := initArray(5)
	fmt.Println(list)
	displayArray(list)
	fmt.Println()
	displayArray(initArray(10))

	fmt.Println(" ")
	nums1 := []int{1, 2, 3, 4, 5}
	fmt.Println("(1.) ", sumAllElement(nums1))

	fmt.Println(" ")
	nums2 := []int{89, 76, 10, 99, 10}
	fmt.Println("(2.) ", maxElement(nums2))

	fmt.Println(" ")
	nums3 := []int{1, 1, 12, 10, 5, 12}
	fmt.Println("(3.) ", smallestIndex(nums3))

	fmt.Println(" ")
	nums4 := []int{2, 3, 5, 7, 11}
	rotate4 := 1
	fmt.Println("(4.) ", shiftingArray(nums4, rotate4))

	fmt.Println(" ")
	nums5 := []int{3, 5, 7, 11, 2}
	target5 := 7
	fmt.Println("(5.) ", linearSearch(nums5, target5))
}
