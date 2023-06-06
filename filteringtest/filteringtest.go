package main

import (
	"fmt"
)

// Nomor 2 :
func printPattern() {
	rows := 7
	start := 1

	for i := 0; i < rows; i++ {
		count := start + i
		for j := 0; j <= i; j++ {
			fmt.Printf("%-3d", count)
			count++
		}
		fmt.Println()
	}
}

// Nomor 3 :
func findTarget(arr []int, target int) int {
	index := -1
	for i := 0; i < len(arr); i++ {
		if arr[i] == target {
			return i
		}
	}
	return index
}

// Nomor 4 :
func addSum(nums []int, target int) []int {
	slice := []int{}
	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			if nums[i]+nums[j] == target {
				return []int{i, j}
			}
		}
	}
	return slice
}

// Nomor 5 :
func sumZero(arr []int) []int {
	slice := []int{}

	for i := 0; i < len(arr); i++ {
		sum := arr[i]
		if sum == 0 {
			slice = append(slice, arr[i])
		}
		for j := i + 1; j < len(arr); j++ {
			sum += arr[j]
			if sum == 0 {
				slice = append(slice, arr[i:j+1]...)
				break
			}
		}
	}
	return slice
}

func main() {
	//Nomor 2 :
	// printPattern()

	// Nomor 3 :
	// arr := []int{-1, 2, 5, 6, 7}
	// targetIndex := findTarget(arr, 6)
	// fmt.Println(targetIndex)

	// Nomor 4 :
	fmt.Println(addSum([]int{2, 7, 11, 15}, 9)) // Output: [0, 1]
	fmt.Println(addSum([]int{3, 2, 3}, 6))      // Output: [0, 2]
	fmt.Println(addSum([]int{3, 2, 4}, 6))      // Output: [1, 2]
	fmt.Println(addSum([]int{3, 3}, 6))         // Output: [0, 1]

	// Nomor 5:

}
