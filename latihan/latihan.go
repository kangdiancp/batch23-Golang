package main

import (
	"fmt"
)

func main() {
	// TwoSum
	fmt.Println(twoSum([]int{2, 7, 11, 15}, 9))
	fmt.Println(twoSum([]int{3, 2, 4}, 6))
	fmt.Println(twoSum([]int{3, 3}, 6))

	//Longest Substring Without Repeating Characters
	fmt.Println(longestSubstring("abcabcbb"))
	fmt.Println(longestSubstring("pwwkew"))
	fmt.Println(longestSubstring("bbbbb"))
}

func twoSum(nums []int, target int) []int {
	for i := 0; i < len(nums); i++ {
		for j := 0; j < len(nums); j++ {
			if nums[i]+nums[j] == target && (i > j || i != j) {
				return []int{i, j}
			}
		}
	}
	return []int{}
}

func longestSubstring(s string) int {
	count := 0
	// for _, value := range s {
	// 	for _, nums := range s {
	// 		if value == nums {
	// 			count += 0
	// 			continue
	// 		} else if value != nums {
	// 			count++
	// 			break
	// 		}
	// 	}
	// }
	for i := 0; i < len(s); i++ {
		for j := 0; j < len(s); j++ {
			if s[i] == s[j] && s[i] > s[j] {
				count -= count
			} else if s[i] != s[j] {
				count++
			}

		}
		break
	}
	return count
}
