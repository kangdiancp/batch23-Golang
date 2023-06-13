package main

import (
	"fmt"
	"sort"
	"strings"
)

func soal2() [7][7]int {
	var matrix [7][7]int
	var input int
	fmt.Print("Masukkan Inputan Angka :")
	fmt.Scan(&input)
	for row := 0; row < len(matrix); row++ {
		for col := 0; col < len(matrix); col++ {
			if row >= col {
				matrix[row][col] = (input + row) + col
			}
		}
	}

	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix); j++ {
			if i < j {
				fmt.Print(" ")
			} else {
				fmt.Printf("%d ", matrix[i][j])
			}
		}
		println()
	}
	println()
	return matrix
}

func soal3(index []int, find int) {
	for i := 0; i < len(index); i++ {
		if index[i] == find {
			fmt.Println("Index ke-", i)
		}
	}
	println()
}

func soal4(index []int, find int) {
	for i := 0; i < len(index); i++ {
		for j := i + 1; j < len(index); j++ {
			if index[i]+index[j] == find {
				fmt.Println("Index ke-", "[", i, ",", j, "]")
			}
		}
	}
	println()
}

func soal5(index []int) {
	for i := 0; i < len(index); i++ {
		for j := i + 1; j < len(index); j++ {
			for k := j + 1; k < len(index); k++ {
				if index[i]+index[j]+index[k] == 0 {
					fmt.Println("Nilai ke-", "[", index[i], ",", index[j], ",", index[k], "]")
				}
			}
		}
	}
	println()
}

func soal6(digit []int) []int {
	carry := 1

	for i := len(digit) - 1; i >= 0; i-- {
		sum := digit[i] + carry
		digit[i] = sum % 10
		carry = sum / 10
	}

	if carry > 0 {
		digit = append([]int{carry}, digit...)
	}
	fmt.Println("Result : ", digit)
	return digit
}

func soal7(arr1, arr2 []string) {
	var same, different []string
	for _, value := range arr2 {
		found := false
		for _, value1 := range arr1 {
			if value == value1 {
				found = true
				break
			}
		}
		if found {
			same = append(same, value)
		} else {
			different = append(different, value)
		}
	}
	for _, v := range arr1 {
		found := true
		for _, v1 := range arr2 {
			if v == v1 {
				found = false
				break
			}
		}
		if found {
			different = append(different, v)
		}
	}

	fmt.Println("Same = ", same)
	fmt.Println("Different = ", different)
}

func soal8(arr []string) bool {
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

func soal9(arr []int) (int, int) {
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
	fmt.Println("min =", min)
	fmt.Println("max =", max)
	return min, max
}

func soal10(arr []int, target int) []int {
	arr = append(arr, target)
	sort.Ints(arr)

	fmt.Println("Output: ", arr)
	return arr
}

func main() {
	soal2()
	soal3([]int{1, 2, 4}, 4)
	soal3([]int{-1, 2, 5, 6, 7}, 6)
	soal4([]int{3, 2, 3}, 6)
	soal5([]int{2, 3, 4, -1, -2})
	soal6([]int{1, 2, 3})

	arr1 := []string{"mangga", "apel", "melon", "pisang", "sirsak", "tomat", "nanas", "nangka", "timun", "mangga"}
	arr2 := []string{"bayam", "wortel", "kangkung", "mangga", "tomat", "kembang kol", "nangka", "timun"}
	soal7(arr1, arr2)

	arr := []string{"Tom", "Tim", "Tim", "Tom"}
	fmt.Println("Outputnya: ", soal8(arr))

	soal9([]int{2, 3, 4, 5, 6, 7, 8, 9, 1, 10})
	soal10([]int{2, 14, 10, 1, 11, 12, 3, 4}, 7)
}
