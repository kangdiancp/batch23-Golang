package main

import (
	"fmt"
	"sort"
	"strings"
)

// soal 2
func segiTiga() {
	num := 5
	count := 7
	// cetak baris ke bawah
	for row := 0; row < count; row++ {
		//cetak kolom
		for col := 0; col <= row; col++ {
			if col == 0 {
				fmt.Print(row+num, " ")
			} else {
				fmt.Print(col+row+num, " ")
			}
		}
		fmt.Println()
	}
}

// soal 3
func findTarget(slice []int, target int) int {
	index := 0
	for i, num := range slice {
		if num == target {
			index = i
		}
	}

	return index
}

// soal 4
func addSum(nums []int, target int) []int {
	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			if nums[i]+nums[j] == target {
				return []int{i, j}
			}
		}
	}

	return nil
}

// soal 5
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

	return []int{} // Jika nil
}

// soal 6
func plusOneDigit(arr []int) []int {
	fmt.Println("Input: digits = ", arr)
	n := len(arr)
	carry := 1
	for i := n - 1; i >= 0; i-- {
		digit := arr[i] + carry
		arr[i] = digit % 10
		carry = digit / 10
	}
	if carry > 0 {
		arr = append([]int{carry}, arr...)
	}
	return arr
}

// soal 7
func compareArrays(array1, array2 []string) ([]string, []string) {
	same := []string{}
	different := []string{}

	for _, item := range array1 {
		found := false
		for _, val := range array2 {
			if item == val {
				found = true
				break
			}
		}
		if found {
			same = append(same, item)
		} else {
			different = append(different, item)
		}
	}

	for _, item := range array2 {
		found := false
		for _, val := range array1 {
			if item == val {
				found = true
				break
			}
		}
		if !found {
			different = append(different, item)
		}
	}

	return same, different
}

// soal 8
func isPalindrome(arr []string) bool {
	for i := 0; i < len(arr); i++ {
		arr[i] = strings.ToLower(arr[i])
	}
	for i := 0; i < len(arr)/2; i++ {
		if arr[i] != arr[len(arr)-i-1] {
			return false
		}
	}

	return true
}

// soal 9
func minMaxArray(arr []int) {
	min := arr[0]
	max := arr[0]

	if len(arr) == 0 {
		fmt.Printf("min = %d, max = %d", min, max)
	}

	for _, num := range arr {

		if num < min {
			min = num
		}

		if num > max {
			max = num
		}
	}

	fmt.Printf("min = %d, max = %d", min, max)
}

// soal 10
func insertElement(arr []int, target int) []int {
	// Menambahkan target
	arr = append(arr, target)

	// Mengurutkan
	sort.Ints(arr)

	return arr
}

func main() {
	//soal 2
	// segiTiga()

	//soal 3
	// fmt.Println(findTarget([]int{1, 2, 4}, 4))        // output 2
	// fmt.Println(findTarget([]int{-1, 2, 5, 6, 7}, 6)) // output 3

	//soal 4
	// fmt.Println(addSum([]int{2, 7, 11, 15}, 9))
	// fmt.Println(addSum([]int{3, 2, 4}, 6))

	//soal 5
	// fmt.Println(sumZero([]int{2, 3, 4, -1, -2}))

	//soal 6
	// fmt.Println(plusOneDigit([]int{1, 2, 3}))

	//soal 7
	// array1 := []string{"Mangga", "Apel", "Melon", "Pisang", "Sirsak", "Tomat", "Nanas", "Nangka", "Timun", "Mangga"}
	// array2 := []string{"Bayam", "WorteI", "Kangkung", "Mangga", "Tomat", "Kembang Kol", "Nangka", "Timun"}

	// same, different := compareArrays(array1, array2)

	// fmt.Println("Same =", same)
	// fmt.Println("Different =", different)

	//soal 8
	// arr := []string{"asep", "budi", "-", "budi", "asep"}
	// arr2 := []string{"Tom", "Tim", "tim", "tom"}
	// arr3 := []string{"tik", "tok", "toko", "tik"}
	// fmt.Println(isPalindrome(arr))
	// fmt.Println(isPalindrome(arr2))
	// fmt.Println(isPalindrome(arr3))

	//soal 9
	// minMaxArray([]int{2, 3, 4, 5, 6, 7, 8, 9, 1, 10}) //output min 1 max 10

	//soal 10
	fmt.Println(insertElement([]int{4, 7, 1, 20}, 9))
}
