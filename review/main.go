package main

import "fmt"

func addSum(nums []int, target int) []int {
	out := []int{}
	for i := 0; i < len(nums); i++ {
		for j := len(nums) - 1; j > i; j-- {
			if nums[i]+nums[j] == target {
				out = append(out, i, j)
			}
		}

	}

	return out
}

func main() {
	//   fmt.Println(addaddSum(2,))
	fmt.Println(addSum([]int{2, 7, 11, 15}, 9))
}
