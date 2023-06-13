package main

import (
	"fmt"
	"sort"
)

// nomor 2
func segitiga1() {
	rows := 7
	start := 1
	end := 13
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

func segitga2() {
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

// nomor 3
func findTarget(nums []int, target int) int {
	for i, num := range nums {
		if num == target {
			return i
		}
	}
	return -1 // Jika target tidak ditemukan
}

// nomor 4
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

// nomor 5
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

	return []int{} // Jika tidak ditemukan pasangan elemen
}

// nomor 6
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

// nomor 7
func compareSlices(slice1 []string, slice2 []string) ([]string, []string) {
	var same []string
	var different []string

	for _, item1 := range slice1 {
		found := false

		for _, item2 := range slice2 {
			if item1 == item2 {
				found = true
				break
			}
		}

		if found {
			same = append(same, item1)
		} else {
			different = append(different, item1)
		}
	}

	for _, item2 := range slice2 {
		found := false

		for _, item1 := range slice1 {
			if item2 == item1 {
				found = true
				break
			}
		}

		if !found {
			different = append(different, item2)
		}
	}

	return same, different
}

// nomor 8
func isPalindrome(words []string) bool {
	length := len(words)
	for i := 0; i < length/2; i++ {
		if words[i] != words[length-1-i] {
			return false
		}
	}
	return true
}

// nomor 9
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

// nomor 10
func menyisipkan(arr []int, target int) []int {
	// Menambahkan elemen target ke dalam array
	arr = append(arr, target)

	// Mengurutkan array secara ascending
	sort.Ints(arr)

	return arr
}

func main() {
	// no.2
	//segitiga1()
	//segitga2()

	//no.3
	//result := findTarget([]int{1, 2, 4}, 4)
	//result := findTarget([]int{-1, 2, 5, 6, 7}, 6)
	//fmt.Println(result)

	// no.4
	fmt.Println(addSum([]int{3, 2, 3}, 6))
	fmt.Println(addSum([]int{3, 2, 4}, 6))
	fmt.Println(addSum([]int{3, 3}, 6))

	// no.5
	//result := sumZero([]int{-1, 0, 1, 2, -1, -4})
	//fmt.Println(result)

	// no.6
	//digits := []int{1, 2, 3}
	//fmt.Println(plusOneDigit(digits))

	// no.7
	// slice1 := []string{"mangga", "apel", "melon", "pisang", "sirsak", "tomat", "nanas", "nangka", "timun", "mangga"}
	// slice2 := []string{"bayam", "wortel", "kangkung", "mangga", "tomat", "kembang kol", "nangka", "timun"}
	// same, different := compareSlices(slice1, slice2)
	// fmt.Printf("Same: %v\n", same)
	// fmt.Printf("Different: %v\n", different)

	// no.8
	//fmt.Println(isPalindrome([]string{"asep", "budi", "-", "budi", "asep"}))
	//fmt.Println(isPalindrome([]string{"tom", "tim", "tim", "tom"}))
	//fmt.Println(isPalindrome([]string{"tik", "tok", "toko", "tik"}))

	// no.9
	//fmt.Println(minMaxArray([]int{2, 3, 4, 5, 6, 7, 8, 9, 1, 10}))

	//no.10
	//array := []int{4, 7, 1, 20}
	//target := 9
	//result := menyisipkan(array, target)
	//fmt.Println(result)

}
