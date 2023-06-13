package main

import (
	"fmt"
	"sort"
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
	arr := []int{}
	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			if nums[i]+nums[j] == target {
				return []int{i, j}
			}
		}
	}
	return arr
}

// Nomor 5 :
func sumZero(arr []int) []int {
	for i := 0; i < len(arr); i++ {
		for j := 0; j < len(arr); j++ {
			for k := 0; k < len(arr); k++ {
				if arr[i]+arr[j]+arr[k] == 0 {
					return []int{arr[i], arr[j], arr[k]}
				}
			}
		}
	}
	return []int{}
}

// Nomor 6 :
func plusOne(digits []int) []int {
	return increment(digits)
}

func increment(digits []int) []int {
	if len(digits) == 0 {
		return []int{1}
	}

	last_digit := digits[len(digits)-1]

	last_digit += 1
	if last_digit > 9 {
		last_digit = 0
		result := increment(digits[:len(digits)-1])
		return append(result, last_digit)
	} else {
		digits[len(digits)-1] = last_digit
		return digits
	}
}

// Nomor 7 :
func compareArrays(arr1, arr2 []string) ([]string, []string) {
	same := []string{}
	different := []string{}

	// Membuat map untuk menyimpan elemen-elemen dalam arr1
	arr1Map := make(map[string]bool)
	for _, elem := range arr1 {
		arr1Map[elem] = true
	}

	// Membandingkan elemen-elemen dalam arr2 dengan arr1Map
	for _, elem := range arr2 {
		if arr1Map[elem] {
			same = append(same, elem)
		} else {
			different = append(different, elem)
		}
	}

	return same, different
}

// Nomor 8 :
func isPalindrome(arr []string) bool {
	length := len(arr)

	// Iterasi melalui setengah panjang array
	for i := 0; i < length/2; i++ {
		// Membandingkan elemen-elemen pada posisi yang berlawanan
		if arr[i] != arr[length-1-i] {
			return false
		}
	}

	return true
}

// Nomor 9 :
func minMaxArray(arr []int) (int, int) {
	if len(arr) == 0 {
		return 0, 0
	}

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

// Nomor 10 :
func insertElement(arr []int, target int) []int {
	arr = append(arr, target)

	sort.Ints(arr)

	return arr
}

func main() {
	//Nomor 2 :
	// printPattern()

	// Nomor 3 :
	// fmt.Println(findTarget([]int{-1, 2, 5, 6, 7}, 6))

	// Nomor 4 :
	// fmt.Println(addSum([]int{2, 7, 11, 15}, 9)) // Output: [0, 1]
	// fmt.Println(addSum([]int{3, 2, 3}, 6))      // Output: [0, 2]
	// fmt.Println(addSum([]int{3, 2, 4}, 6))      // Output: [1, 2]
	// fmt.Println(addSum([]int{3, 3}, 6))         // Output: [0, 1]

	// Nomor 5:
	// fmt.Println(sumZero([]int{-1, 0, 1, 2, -1, 4}))

	// Nomor 6 :
	// fmt.Println(plusOne([]int{9}))

	// Nomor 7 :
	arr1 := []string{"Mangga", "Apel", "Melon", "Pisang", "Sirsak", "Tomat", "Nanas", "Nangka", "Timun", "Mangga"}
	arr2 := []string{"Bayam", "Wortel", "Kangkung", "Mangga", "Tomat", "Kembang Kol", "Nangka", "Timun"}

	same, different := compareArrays(arr1, arr2)
	fmt.Println("Same:", same)           // Output: Same: [Mangga Tomat Nangka Timun]
	fmt.Println("Different:", different) // Output: Different: [Bayam Wortel Kangkung Kembang Kol]

	// Nomor 8 :
	// fmt.Println(isPalindrome([]string{"asep", "budi", " - ", "budi", "asep"})) // Output: true
	// fmt.Println(isPalindrome([]string{" Tom", "Tim", "tim", "tom"})) // Output: true
	// fmt.Println(isPalindrome([]string{"tik", "tok", "toko", "tik"})) // Output: false

	// Nomor 9 :
	// arr := []int{2, 3, 4, 5, 6, 7, 8, 9, 1, 10}
	// min, max := minMaxArray(arr)
	// fmt.Println("Min:", min) // Output: Min: 1
	// fmt.Println("Max:", max) // Output: Max: 10

	// Nomor 10 :
	// fmt.Println(insertElement([]int{4, 7, 1, 20}, 9))
}
