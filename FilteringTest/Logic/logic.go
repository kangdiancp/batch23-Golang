package main

import (
	"fmt"
	"sort"
	"strings"
)

func main() {
	//menjalankan fungsi no 10
	fmt.Println("No 10")
	array := []int{4, 7, 1, 20}
	element := 9
	result := SisipInteger(array, element)
	fmt.Println(result)

	//menjalankan fungsi no 9
	fmt.Println("\nNo 9")
	arr := []int{2, 3, 4, 5, 6, 7, 8, 9, 10}
	min, max := minMaxArray(arr)
	fmt.Printf("min=%d, max=%d\n", min, max)

	//menjalankan fungsi no 8
	fmt.Println("\nNo 8")
	arr1 := []string{"asep", "budi", "-", "budi", "asep"}
	fmt.Println(isPalindrome(arr1)) //output true

	// arr2 := []string{"Tom", "Tim", "tom", "tim"}
	// fmt.Println(isPalindrome(arr2)) //output true

	// arr3 := []string{"tik", "tok", "toko", "tik"}
	// fmt.Println(isPalindrome(arr3)) //output false

	// //menjalankan fungsi no 7
	fmt.Println("\nNo 7")
	fmt.Println(compareArrays([]string{"Mangga", "Apel", "Melon", "Pisang", "Sirsak", "Tomat", "Nanas", "Nangka", "Timun", "Mangga"}, []string{"Bayam", "Wortel", "Kangkung", "Mangga", "Tomat", "Kembang Kol", "Nangka", "Timun"}))
	// same, different := compareArrays(arr1, arr2)
	// fmt.Println("Same =", same)
	// fmt.Println("Different =", different)
	// arr1 :=[]string{"Mangga","Apel","Melon","Pisang","Sirsak","Tomat","Nanas","Nangka","Timun","Mangga"},
	// arr2 :=[]string{"Bayam","Wortel","Kangkung","Mangga","Tomat","Kembang Kol","Nangka","Timun"}

	// menjalankan fungsi no 6
	// fmt.Println("\nNo 6")
	// digits := []int{1, 2, 3}
	// result := plusOneDigit(digits)
	// fmt.Println(result)

	//menjalankan fungsi no 5
	fmt.Println("\nNo 5")
	fmt.Println(sumZero([]int{-1, 0, 1, 2, -1, 4}))
	// arr1 := []int{-1, 0, 1, 2, -1, 4}
	// result1 := sumZero(arr1)
	// fmt.Println(result1)

	// arr2 := []int{2, 3, 4, -1, -2}
	// result2 := sumZero(arr2)
	// fmt.Println(result2)

	// arr3 := []int{}
	// result3 := sumZero(arr3)
	// fmt.Println(result3)

	//menjalankan fungsi no 4
	fmt.Println("\nNo 4")
	fmt.Println(addSum([]int{2, 7, 11, 15}, 9))
	// 	nums := []int{2, 7, 11, 15}
	// 	target := 9
	// 	result := addSum(nums, target)
	// 	fmt.Println(result)

	//menjalankan fungsi no 3
	fmt.Println("\nNo 3")
	fmt.Println(findTarget([]int{1, 2, 4}, 4))
	fmt.Println(findTarget([]int{-1, 2, 5, 6, 7}, 6))

	//menjalankan fungsi no 2
	fmt.Println("\nNo2")
	printTriangle()
}

// no 10
func SisipInteger(arr []int, element int) []int {
	arr = append(arr, element)
	sort.Ints(arr)
	return arr

}

// n0 9
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

// no 8
func isPalindrome(arr []string) bool {
	for i := 0; i < len(arr)/2; i++ {
		if strings.ToLower(arr[i]) != strings.ToLower(arr[len(arr)-1-i]) {
			return false
		}
	}
	return true
}

// no 7
func compareArrays(arr1, arr2 []string) ([]string, []string) {
	same := make([]string, 0)
	different := make([]string, 0)
	set := make(map[string]bool)

	for _, str := range arr1 {
		set[str] = true
	}

	for _, str := range arr2 {
		if set[str] {
			same = append(same, str)
		} else {
			different = append(different, str)
		}
	}

	for _, str := range arr1 {
		if set[str] {
			same = append(same, str)
		} else {
			different = append(different, str)
		}
	}
	return same, different
}

// no 6
func plusOneDigit(digits []int) []int {
	n := len(digits)
	//untuk melakukan penambahan 1 pada digit terakhir
	digits[n-1] += 1
	//untuk melakukan propragasi carry jika diperlukan
	carry := 0
	for i := n - 1; i >= 0; i-- {
		digits[i] += carry
		carry = digits[i] / 10
		digits[i] %= 10
	}
	//jika carry masih ada setelah propagasi, tambahkan digit 1 pada array
	if carry > 0 {
		digits = append([]int{carry}, digits...)
	}
	return digits
}

// no 5
func sumZero(arr []int) []int {
	result := make([]int, 0)

	//menggunakan metode two-pointer untuk mencari pasangan elemen yang jumlahnya 0
	left := 0
	right := len(arr) - 1

	for left < right {
		sum := arr[left] + arr[right]

		if sum == 0 {
			result = append(result, arr[left], arr[right])
			left++
			right--
		} else if sum > 0 {
			right--
		} else {
			left++
		}
	}
	return result
}

// no 4
func addSum(nums []int, target int) []int {
	indices := make(map[int]int)

	for i, num := range nums {
		complement := target - num
		if idx, ok := indices[complement]; ok {
			return []int{idx, i}
		}
		indices[num] = i
	}
	return nil
}

// no 3
func findTarget(nums []int, target int) int {
	for i, num := range nums {
		if num == target {
			return i
		}
	}
	return -1
}

// no 2
func printTriangle() {
	n := 7
	index := 1

	for i := 0; i < n; i++ {
		count := index + i
		for j := 0; j <= i; j++ {
			fmt.Printf("%-3d", count)
			count++
		}
		fmt.Println()
	}
}
