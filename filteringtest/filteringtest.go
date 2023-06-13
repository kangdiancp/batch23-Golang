package main

import (
	"fmt"
	"sort"
	"strings"
)

// No.2
func segitiga() {
	rows := 7
	start := 5
	end := 17
	count := start

	for i := 0; i < rows; i++ {
		for j := 0; j <= i; j++ {
			if count > end {
				count = start
			}
			fmt.Printf("%-3d", count)
			count++
		}
		fmt.Println()
		count = start + i + 1
	}
}

// No.3
func findTarget(arr []int, n int) {

	for i := 0; i < len(arr); i++ {
		if arr[i] == n {
			fmt.Println(i)
		}
	}

}

// No.4
func addSum(nums []int, target int) []int {
	indexs := make([]int, 0)

	for i := 0; i < len(nums)-1; i++ {
		for j := i + 1; j < len(nums); j++ {
			if nums[i]+nums[j] == target {
				indexs = append(indexs, i, j)
			}
		}
	}
	return indexs
}

// No.5
func sumZero(nums []int) []int {
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

// No.6
func plusOneDigit(digits []int) []int {
	n := len(digits)
	carry := 1

	for i := n - 1; i >= 0; i-- {
		sum := digits[i] + carry
		digits[i] = sum % 10
		carry = sum / 10
	}

	if carry > 0 {
		digits = append([]int{carry}, digits...)
	}

	return digits
}

// No.7
func sameDifferent(array1, array2 []string) ([]string, []string) {
	sameStrings := make([]string, 0)
	differentStrings := make([]string, 0)

	for i := 0; i < len(array1); i++ {
		found := false
		for j := 0; j < len(array2); j++ {
			if array1[i] == array2[j] {
				sameStrings = append(sameStrings, array1[i])
				found = true
			}
		}
		if !found {
			differentStrings = append(differentStrings, array1[i])
		}
	}

	for i := 0; i < len(array2); i++ {
		found := false
		for j := 0; j < len(array1); j++ {
			if array2[i] == array1[j] {
				found = true
			}
		}
		if !found {
			differentStrings = append(differentStrings, array2[i])
		}
	}

	return sameStrings, differentStrings
}

// No.8
func isPalindrome(arr []string) bool {
	for i := 0; i < len(arr)/2; i++ {
		if strings.ToLower(arr[i]) != strings.ToLower(arr[len(arr)-1-i]) {
			return false
		}
	}
	return true
}

// No.9
func minMaxArray(arr []int) (int, int) {
	min := arr[0]
	max := arr[0]

	for _, num := range arr {
		if num < min {
			min = num
		}
		if num > max {
			max = num
		}
	}
	return min, max
}

// No.10
func insertElement(arr []int, target int) []int {
	// menambahkan elemen target ke dalam array
	arr = append(arr, target)

	// mengurutkan array (ascending)
	sort.Ints(arr)

	return arr
}

func main() {
	//segitiga()
	//(findTarget([]int{1, 2, 4}, 4))
	//fmt.Println(addSum([]int{3, 2, 3}, 6))
	//fmt.Println(sumZero([]int{2, 3, 4, -1, -2}))
	// fmt.Println(plusOneDigit([]int{1, 2, 9}))
	fmt.Println(sameDifferent([]string{"mangga", "apel", "melon", "pisang", "sirsak", "tomat", "nanas", "nangka", "timun"}, []string{"bayam", "wortel", "kangkung", "mangga", "tomat", "kembang kol", "nangka", "timun"}))
	//fmt.Println(isPalindrome([]string{"tik", "tok", "toko", "tik"}))
	//fmt.Println(minMaxArray([]int{2, 3, 4, 5, 6, 7, 8, 9, 1, 10}))
	//fmt.Println(insertElement([]int{4, 7, 1, 20}, 9))
}
