package main

import (
	"fmt"
	"sort"
	"strings"
)

// nomor 2
func segitiga(n int) {
	var matriks [7][7]int
	matriks[0][0] = n

	for i := 1; i < len(matriks); i++ {
		for j := 1; j < len(matriks); j++ {
			if i > j {
				matriks[i][j] = (n + i) + j
			} else if i == j {
				matriks[i][j] = matriks[i-1][j-1] + 2
			}

		}
		for k := 0; k < len(matriks); k++ {
			if k == 0 {
				matriks[i][k] = matriks[i-1][k] + 1
			}
		}

	}

	for i := 0; i < len(matriks); i++ {
		for j := 0; j < len(matriks); j++ {
			fmt.Print(matriks[i][j], "	")
		}
		fmt.Println()
	}

}

// nomor 3
func findTarget(arr []int, n int) {

	for i := 0; i < len(arr); i++ {
		if arr[i] == n {
			fmt.Println(i)
		}
	}

}

// nomor 4
func addSum(nums []int, target int) []int {
	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			if nums[i]+nums[j] == target {
				return []int{i, j}
			}
		}
	}

	return nil // Jika tidak ditemukan pasangan elemen
}

// nomor5
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

// nomor6
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
func samaDanBeda(array1, array2 []string) ([]string, []string) {
	same := []string{}
	different := []string{}

	// Membuat slice kosong untuk menyimpan string yang sudah ditemukan
	seen := []string{}

	// Mengecek setiap string pada array1
	for _, str1 := range array1 {
		// Mengecek apakah string sudah pernah ditemukan sebelumnya
		found := false
		for _, seenStr := range seen {
			if seenStr == str1 {
				found = true
				break
			}
		}

		// Jika string sudah pernah ditemukan, lanjut ke string berikutnya
		if found {
			continue
		}

		// Mengecek apakah string ada di array2
		isSame := false
		for _, str2 := range array2 {
			if str1 == str2 {
				isSame = true
				break
			}
		}

		// Memasukkan string ke dalam slice yang sesuai
		if isSame {
			same = append(same, str1)
		} else {
			different = append(different, str1)
		}

		// Menambahkan string ke dalam slice seen
		seen = append(seen, str1)
	}

	// Mengecek setiap string pada array2
	for _, str2 := range array2 {
		// Mengecek apakah string sudah pernah ditemukan sebelumnya
		found := false
		for _, seenStr := range seen {
			if seenStr == str2 {
				found = true
				break
			}
		}

		// Jika string sudah pernah ditemukan, lanjut ke string berikutnya
		if found {
			continue
		}

		// Mengecek apakah string ada di array1
		isSame := false
		for _, str1 := range array1 {
			if str2 == str1 {
				isSame = true
				break
			}
		}

		// Memasukkan string ke dalam slice yang sesuai
		if isSame {
			same = append(same, str2)
		} else {
			different = append(different, str2)
		}

		// Menambahkan string ke dalam slice seen
		seen = append(seen, str2)
	}

	return same, different
}

// nomor 8
func isPalindrome(arr []string) bool {
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

// nomor 9
func minMaxArray(arr []int) (int, int) {
	// Inisialisasi min dan max dengan nilai pertama dalam array
	min := arr[0]
	max := arr[0]

	// Mencari nilai minimum dan maksimum dalam array
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
func insertElement(arr []int, target int) []int {
	// Menambahkan elemen target ke dalam array
	arr = append(arr, target)

	// Mengurutkan array secara ascending
	sort.Ints(arr)

	return arr
}

func main() {
	//nomor 2
	segitiga(1)

	//nomor3
	//arr := []int{-1, 2, 5, 6, 7}
	//findTarget(arr, 6)

	//nomor4
	//num := []int{2, 7, 11, 15}
	//result := addSum(num, 17)
	//fmt.Println(result)

	//nomor 5
	// result := sumZero([]int{})
	// fmt.Println(result)

	//nomor 6
	// digits := []int{1, 2, 3}
	// fmt.Println(plusOneDigit(digits))

	//nomor 7
	// array1 := []string{"Mangga", "Apel", "Melon", "Pisang", "Sirsak", "Tomat", "Nanas", "Nangka", "Timun", "Mangga"}
	// array2 := []string{"Bayam", "Apel", "Wortel", "Kangkung", "Mangga", "Tomat", "Kembang Kol", "Nangka", "Timun"}

	// same, different := samaDanBeda(array1, array2)
	// fmt.Println("Same =", same)
	// fmt.Println("Different =", different)

	//nomor 8
	// arr := []string{"asep", "budi", "-", "budi", "asep"}
	// result := isPalindrome(arr)
	// fmt.Println(result)

	//nomor 9
	// array1 := []int{2, 3, 4, 5, 6, 7, 8, 9, 1, 10}
	// min1, max1 := minMaxArray(array1)
	// fmt.Println("min =", min1)
	// fmt.Println("max =", max1)

	//nomor 10
	//array := []int{4, 7, 1, 20}
	//target := 9
	//result := insertElement(array, target)
	//fmt.Println(result)

}
