package main

import (
	"fmt"
	"strconv"
)

// function Two Sum (Sum 2 element array based on index)
func TwoSum(nums []int, target int) []int {
	var twosum []int
	for i := 0; i < len(nums); i++ {
		for j := 1; j < len(nums); j++ {
			if nums[i]+nums[j] == target && i!=j && i < j{
				twosum = append(twosum, i, j)
					
			}
		}
	}
	return twosum
}

// Function Palindrom number based on element of array in integer type
func IsPalindrome(x int)bool{
	str := strconv.Itoa(x)
	long := len(str)
	for i := 0; i < long/2; i++ {
		if str[i] != str[long-i-1]{
			return false
		}
	}
	return true
}

// Func Roman To Integer
// func RomanToInt(s string) int{

// 	var romawi 
// 	switch romawi := s; {
// 	case romawi == "I":
// 		return 1
// 	case romawi == "V":
// 		return 5
// 	case romawi == "X":
// 		return 10
// 	case romawi == "L":
// 		return 50
// 	case romawi == "C":
// 		return 100
// 	case romawi == "D":
// 		return 500
// 	case romawi == "M":
// 		return 1000
// 	return -1
// 	}


	
// }

func main() {
	// no 1 Two Sum
	// fmt.Println(TwoSum([]int {2, 7, 11, 15}, 9))
	// fmt.Println(TwoSum([]int {3, 2, 4}, 6))
	// fmt.Println(TwoSum([]int {3, 3}, 6))
	// fmt.Println(TwoSum([]int {3, 2, 3}, 6))

	// no 2 Palindrome Number 
	/*	Input: x = -121
		Output: false
		Explanation: From left to right, it reads -121. From right to left, it becomes 121-. Therefore it is not a palindrome.
		Input: x = 121
		Output: true
		Explanation: 121 reads as 121 from left to right and from right to left.*/
	// fmt.Println(IsPalindrome(121))
	// fmt.Println(IsPalindrome(-121))
	fmt.Println(IsPalindrome(10))


}