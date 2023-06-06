package main

import (
	"fmt"
	"sort"
	"strings"
)

// Soal 02 -> Segitiga siku"
func soal02(n int) {
	var matriks [7][7]int
	matriks[0][0] = n

	for row := 1; row < len(matriks); row++ {
		for col := 1; col < len(matriks); col++ {
			if row > col {
				matriks[row][col] = (n + row) + col
			} else if row == col {
				matriks[row][col] = matriks[row-1][col-1] + 2
			}

		}
		for k := 0; k < len(matriks); k++ {
			if k == 0 {
				matriks[row][k] = matriks[row-1][k] + 1
			}
		}

	}

	for row := 0; row < len(matriks); row++ {
		for col := 0; col < len(matriks); col++ {
			fmt.Print(matriks[row][col], "  ")
		}
		fmt.Println()
	}

}

// Soal 03 -> Find Target
func soal03(arr []int, n int) {

	for i := 0; i < len(arr); i++ {
		if arr[i] == n {
			fmt.Println(i)
		}
	}

}

// Soal 04 -> addSum
func soal04(nums []int, target int) []int {
	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			if nums[i]+nums[j] == target {
				return []int{i, j}
			}
		}
	}
	return nil
}

// Soal 05
func soal05(nums []int) []int {
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

// Soal 06
func soal06(digits []int) []int {
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

// Soal 07 -> Sama dan Beda
func soal07(array1, array2 []string) ([]string, []string) {
	same := []string{}
	different := []string{}

	seen := []string{}

	for _, str1 := range array1 {

		found := false
		for _, seenStr := range seen {
			if seenStr == str1 {
				found = true
				break
			}
		}

		if found {
			continue
		}

		isSame := false
		for _, str2 := range array2 {
			if str1 == str2 {
				isSame = true
				break
			}
		}

		if isSame {
			same = append(same, str1)
		} else {
			different = append(different, str1)
		}

		seen = append(seen, str1)
	}

	for _, str2 := range array2 {

		found := false
		for _, seenStr := range seen {
			if seenStr == str2 {
				found = true
				break
			}
		}

		if found {
			continue
		}

		isSame := false
		for _, str1 := range array1 {
			if str2 == str1 {
				isSame = true
				break
			}
		}

		if isSame {
			same = append(same, str2)
		} else {
			different = append(different, str2)
		}

		seen = append(seen, str2)
	}

	return same, different
}

// nomor 08 -> isPalindrome
func soal08(arr []string) bool {
	n := len(arr)
	for i := 0; i < len(arr); i++ {
		arr[i] = strings.ToLower(arr[i])
	}

	for i := 0; i < n/2; i++ {
		if arr[i] != arr[n-i-1] {
			return false
		}
	}

	return true
}

// nomor 9 -> minMaxArray
func soal09(arr []int) (int, int) {

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

// nomor 10 -> insertElement
func soal10(arr []int, target int) []int {
	arr = append(arr, target)
	sort.Ints(arr)
	return arr
}

func main() {
	soal02(1)
	soal03([]int{-1, 2, 5, 6, 7}, 6)

	// Soal 04
	counts := []int{2, 7, 11, 15}
	results := soal04(counts, 17)
	fmt.Println(results)

	// Soal 05
	result2 := soal05([]int{})
	fmt.Println(result2)

	// Soal 06
	digits := []int{1, 2, 3}
	fmt.Println(soal06(digits))

	// Soal 07
	array1 := []string{"Mangga", "Apel", "Melon", "Pisang", "Sirsak", "Tomat", "Nanas", "Nangka", "Timun", "Mangga"}
	array2 := []string{"Bayam", "Apel", "Wortel", "Kangkung", "Mangga", "Tomat", "Kembang Kol", "Nangka", "Timun"}

	same, different := soal07(array1, array2)
	fmt.Println("Same =", same)
	fmt.Println("Different =", different)

	// Soal 08
	// arr2 := []string{"asep", "budi", "-", "budi", "asep"}
	// result3 := soal08(arr2)
	fmt.Println(soal08([]string{"asep", "budi", "-", "budi", "asep"}))

	// Soal 09
	array3 := []int{2, 3, 4, 5, 6, 7, 8, 9, 1, 10}
	min1, max1 := soal09(array3)
	fmt.Println("min =", min1)
	fmt.Println("max =", max1)

	// Soal 10
	// input := []int{4, 7, 1, 20}
	// target := 9
	// output := soal10([]int{4, 7, 1, 20}, 9)
	fmt.Println(soal10([]int{4, 7, 1, 20}, 9))
}
