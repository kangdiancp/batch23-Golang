package main

import (
	"fmt"
	"sort"
	"strings"
)

// 2.
func segitigaSikuAngka(n int, num int) {
	for i := 1; i <= n; i++ {
		for j := 1; j <= i; j++ {
			fmt.Print(num, " ")
			num++
		}
		num -= i - 1
		fmt.Println()
	}
}

// 3.
func findTarget(nums []int, target int) int {
	for i := 0; i < len(nums); i++ {
		if nums[i] == target {
			return i
		}
	}

	return 0
}

// 4.
func addSum(nums []int, target int) []int {
	for i := 0; i < len(nums)-1; i++ {
		for j := i + 1; j < len(nums); j++ {
			if nums[i]+nums[j] == target {
				return []int{i, j}
			}
		}
	}

	return []int{}
}

// 5.
func sumZero(nums []int) []int {
	sort.Ints(nums)

	// left := 0
	// right := len(nums) - 1

	// for left < right {
	// 	sum := nums[left] + nums[right]

	// 	if sum == 0 {
	// 		return []int{nums[left], nums[nums[right]]}
	// 	} else if sum < 0 {
	// 		left++
	// 	} else {
	// 		right--
	// 	}
	// }

	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			for k := j + 1; k < len(nums); k++ {
				if nums[i]+nums[j]+nums[k] == 0 {
					return []int{nums[i], nums[j], nums[k]}
				}
			}
		}
	}

	return []int{}
}

// 6.
func plusOne(digits []int) []int {
	if len(digits) == 0 {
		return []int{1}
	}

	lastDigit := digits[len(digits)-1]

	lastDigit += 1

	if lastDigit > 9 {
		lastDigit = 0

		result := plusOne(digits[:len(digits)-1])

		return append(result, lastDigit)
	} else {
		digits[len(digits)-1] = lastDigit

		return digits
	}
}

// 7.
func duaBuahArray(word1, word2 []string) ([]string, []string) {
	sameStrings := []string{}
	differentStrings := []string{}

	for _, str1 := range word1 {
		found := false
		for _, str2 := range word2 {
			if str1 == str2 {
				sameStrings = append(sameStrings, str1)
				found = true
				break
			}
		}
		if !found {
			differentStrings = append(differentStrings, str1)
		}
	}

	for _, str2 := range word2 {
		found := false
		for _, str1 := range word1 {
			if str2 == str1 {
				found = true
				break
			}
		}
		if !found {
			differentStrings = append(differentStrings, str2)
		}
	}

	return sameStrings, differentStrings
}

// 8.
func isPalindrome(nums []string) bool {
	for i := 0; i < len(nums); i++ {
		nums[i] = strings.ToLower(nums[i])
	}

	for j := 0; j < len(nums)/2; j++ {
		if nums[j] != nums[len(nums)-j-1] {
			return false
		}
	}

	return true
}

// 9.
func minMaxArray(nums []int) (int, int) {
	min := nums[0]
	max := nums[0]

	for _, num := range nums {
		if num < min {
			min = num
		}
		if num > max {
			max = num
		}
	}

	return min, max
}

// 10.
func insertElement(arr []int, target int) []int {
	arr = append(arr, target)

	sort.Ints(arr)

	return arr
}

func main() {
	fmt.Println("(2.) ")
	segitigaSikuAngka(7, 5)

	fmt.Println()

	fmt.Println("(3.) ")
	nums3 := []int{-1, 2, 5, 6, 7}
	fmt.Println(findTarget(nums3, 6))

	fmt.Println()

	fmt.Println("(4.) ")
	nums4 := []int{3, 3}
	fmt.Println(addSum(nums4, 6))

	fmt.Println()

	fmt.Println("(5.) ")
	nums5 := []int{2, 3, 4, -1, -2}
	fmt.Println(sumZero(nums5))

	fmt.Println()

	fmt.Println("(6.) ")
	nums6 := []int{9}
	fmt.Println(plusOne(nums6))

	fmt.Println()

	fmt.Println("(7.) ")
	word1 := []string{"Mangga", "Apel", "Melon", "Pisang", "Sirsak", "Tomat", "Nanas", "Nangka", "Timun", "Mangga"}
	word2 := []string{"Bayam", "Wortel", "Kangkung", "Mangga", "Tomat", "Kembang Kol", "Nangka", "Timun"}
	fmt.Println(duaBuahArray(word1, word2))

	fmt.Println()

	fmt.Println("(8.) ")
	string8 := []string{"asep", "budi", "-", "budi", "asep"}
	fmt.Println(isPalindrome(string8))

	fmt.Println()

	fmt.Println("(9.) ")
	nums9 := []int{2, 3, 4, 5, 6, 7, 8, 9, 1, 10}
	min, max := minMaxArray(nums9)
	fmt.Println("min =", min)
	fmt.Println("max =", max)

	fmt.Println()

	fmt.Println("(10.) ")
	nums10 := []int{4, 7, 1, 20}
	fmt.Println(insertElement(nums10, 9))
}
