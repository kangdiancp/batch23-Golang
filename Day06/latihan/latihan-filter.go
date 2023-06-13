package main

import (
	"fmt"
	"strconv"
)

// Soal 02
func soal02(n int) {
	var martix [7][7]int
	martix[0][0] = n

	for row := 1; row < len(martix); row++ {
		for col := 1; col < len(martix); col++ {
			if row > col {
				martix[row][col] = (n + row) + col
			} else if row == col {
				martix[row][col] = martix[row-1][col-1] + 2
			}
		}

		for k := 0; k < len(martix); k++ {
			if k == 0 {
				martix[row][k] = martix[row-1][k] + 1
			}
		}
	}

	for row := 0; row < len(martix); row++ {
		for col := 0; col < len(martix); col++ {
			fmt.Print(martix[row][col], " ")
		}
		fmt.Println()
	}
}

// Soal 03
func findTarget(arr []int, target int) {
	for i := 0; i < len(arr); i++ {
		if arr[i] == target {
			fmt.Println(i)
		}
	}
}

// Soal 04 -> []int{2, 7, 11, 15}
func addSum(arr []int, target int) []int {
	for i := 0; i < len(arr); i++ {
		for j := i + 1; j < len(arr); j++ {
			if arr[i]+arr[j] == target {
				return []int{i, j}
			}
		}
	}
	return nil
}

// Soal 05 -> []]int{-1, 0, 3, 2, -1, 4} mendapat 3 index
func sumZero(arr []int) []int {
	for i := 0; i < len(arr); i++ {
		for j := i + 1; j < len(arr); j++ {
			for k := j + 1; k < len(arr); k++ {
				if arr[i]+arr[j]+arr[k] == 0 {
					return []int{arr[i], arr[j], arr[k]}
				}
			}
		}
	}
	return nil
}

func summaryRanges(nums []int) []string {
	result := []string{}

	if len(nums) == 0 {
		return result
	}

	start := nums[0]
	end := nums[0]

	for i := 1; i < len(nums); i++ {
		if nums[i] == end+1 {
			end = nums[i]
		} else {
			if start == end {
				result = append(result, strconv.Itoa(start))
			} else {
				result = append(result, strconv.Itoa(start)+"->"+strconv.Itoa(end))
			}
			start = nums[i]
			end = nums[i]
		}
	}

	if start == end {
		result = append(result, strconv.Itoa(start))
	} else {
		result = append(result, strconv.Itoa(start)+"->"+strconv.Itoa(end))
	}

	return result
}

func main() {

	nums := []int{0, 1, 2, 4, 5, 7}
	output := summaryRanges(nums)
	fmt.Println(output)

	// soal02(1)
	// findTarget([]int{-1, 2, 5, 6, 7}, 6)

	// Soal 04
	// arr := []int{2, 7, 11, 15}
	// counts := addSum(arr, 9)
	// fmt.Println(counts)

	// Soal 05
	// arr2 := sumZero([]int{-1, 0, 3, 2, -1, 4})
	// fmt.Println(arr2)
}
