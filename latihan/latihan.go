package main

import (
	"fmt"
	"strconv"
)

func twoSum(nums []int, target int) []int {
	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			if nums[i]+nums[j] == target {
				return []int{i, j}
			}
		}
	}
	return []int{}
}

func isPalindrome(x int) bool {
	var reservedNum int
	tmp := x
	for tmp > 0 {
		reservedNum = reservedNum*10 + tmp%10
		tmp = tmp / 10
	}
	return x == reservedNum
}

func containsDuplicate(nums []int) bool {
	for i := 0; i < len(nums); i++ {
		for j := 0; j < len(nums); j++ {
			if nums[i] == nums[j] {
				return true
			}
		}
	}
	return false
}

func summaryRanges(nums []int) []string {
	if len(nums) == 0 {
		return []string{}
	}

	ranges := []string{}
	start := nums[0]

	for i := 0; i < len(nums); i++ {
		// Jika mencapai akhir array atau menemukan celah
		if i == len(nums)-1 || nums[i]+1 != nums[i+1] {
			//Jika start sama dengan nums[i], tidak ada rentang, hanya satu angka
			if start == nums[i] {
				ranges = append(ranges, strconv.Itoa(start))
			} else {
				ranges = append(ranges, strconv.Itoa(start)+"->"+strconv.Itoa(nums[i]))
			}
			// Jika belum mencapai akhir array, mulai rentang baru
			if i != len(nums)-1 {
				start = nums[i+1]
			}
		}
	}
	return ranges
}

// type ListNode struct {
// 	Val  int
// 	Next *ListNode
// }

// func mergeTwoLists(lists1 *ListNode, lists2 *ListNode) *ListNode {
// 	dummy := &ListNode{}
// 	current:=
// }

func main() {
	// fmt.Println(twoSum([]int{2, 7, 11, 15}, 9)) // Outputs: [0 1]
	// fmt.Println(twoSum([]int{3, 2, 4}, 6))      // Outputs: [1 2]
	// fmt.Println(twoSum([]int{3, 3}, 6))         // Outputs: [0 1]

	// fmt.Println(isPalindrome(121))	// Output: true
	// fmt.Println(isPalindrome(-121))	// Output: false
	// fmt.Println(isPalindrome(10))	// Output: false

	// fmt.Println(containsDuplicate([]int{1, 2, 3, 1}))                   // Output: true
	// fmt.Println(containsDuplicate([]int{1, 2, 3, 4}))                   // Output: false
	// fmt.Println(containsDuplicate([]int{1, 1, 1, 3, 3, 4, 3, 2, 4, 2})) // Output: true

	// fmt.Println(summaryRanges([]int{0, 1, 2, 4, 5, 7}))    // Output: ["0->2","4->5","7"]
	fmt.Println(summaryRanges([]int{0, 2, 3, 4, 6, 8, 9})) // Output: ["0","2->4","6","8->9"]

}
