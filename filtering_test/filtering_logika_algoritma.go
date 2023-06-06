package main

import (
	"fmt"
	"sort"
	"strings"
)

func nomor2() {

	var matrix [7][7]int
	nilai := 1
	matrix[0][0] = nilai
	for row := 1; row < len(matrix); row++ {
		for col := 0; col < len(matrix); col++ {
			if col == row {
				matrix[row][col] = matrix[row-1][col-1] + 2
			} else if row > col {
				matrix[row][col] = (nilai + row) + col
			} else if col == 0 && row != 0 {
				matrix[col][row] = matrix[row-1][col-1] + 1
			}

		}
	}
	for row := 0; row < len(matrix); row++ {
		for Colom := 0; Colom < len(matrix); Colom++ {
			if Colom > row {
				fmt.Print(" ")
			} else {
				fmt.Printf("%d ", matrix[row][Colom])
			}
		}
		fmt.Println()
	}

}

func nomor3(arr []int, cari int) {
	for i, value := range arr {
		if value == cari {
			fmt.Printf("output : %d", i)
		}
	}
	fmt.Println()
}

func nomor4(arr []int, cari int) []int {
	for i := 0; i < len(arr); i++ {
		for j := 0; j < len(arr); j++ {
			if arr[i]+arr[j] == cari {
				// fmt.Printf("%d, %d", i, j)
				// fmt.Println()
				return []int{i, j}
			}
		}
	}
	// fmt.Println()
	return []int{}
}

func nomor5(arr []int) []int {
	for i := 0; i < len(arr); i++ {
		for j := 0; j < len(arr); j++ {
			for k := 0; k < len(arr); k++ {
				if arr[i]+arr[j]+arr[k] == 0 {
					// fmt.Printf("%d, %d, %d", arr[i], arr[j], arr[k])
					// fmt.Println()
					// break
					return []int{arr[i], arr[j], arr[k]}
				}
			}
		}
	}
	return []int{}
}

func nomor6(arr []int) {
	signAwal := 1

	simpanData := arr[len(arr)-1] + 1
	replaceElement := arr[0 : len(arr)-1]
	copy(arr, replaceElement)

	nilaiModulus := simpanData % 10
	masukkan := append(arr, signAwal, nilaiModulus)

	fmt.Println(masukkan)
}

func nomor7() {
	//Belum dapat
}

func nomor8(arr []string) bool {
	nilai := len(arr)
	for i := 0; i < len(arr); i++ {
		arr[i] = strings.ToLower(arr[i])
	}

	for i := 0; i < nilai/2; i++ {
		if arr[i] != arr[nilai-i-1] {
			return false
		}
	}

	return true
}
func nomor9(arr []int) {
	nilaiMin := arr[0]
	nilaiMax := arr[0]

	for _, num := range arr {
		if num < nilaiMin {
			nilaiMin = num
		}
		if num > nilaiMax {
			nilaiMax = num
		}
	}
	fmt.Printf("Output : min = %d, max = %d", nilaiMin, nilaiMax)
	println()
}
func nomor10(arr []int, sisip int) {
	arr = append(arr, sisip)

	sort.Ints(arr)

	fmt.Println(arr)
}

func main() {
	println("Nomor 02---------")
	nomor2()
	println()
	println("Nomor 03---------")
	nomor3([]int{1, 2, 4}, 4)
	println()
	println("Nomor 04---------")
	fmt.Println(nomor4([]int{2, 7, 11, 15}, 9))
	println()
	println("Nomor 05---------")
	fmt.Println(nomor5([]int{-1, 0, 1, 2, -1, 4}))
	println()
	println("Nomor 06---------")
	nomor6([]int{1, 2, 9})
	println()
	println("Nomor 08---------")
	fmt.Println(nomor8([]string{"asep", "budi", "-", "budi", "asep"}))
	println()
	println("Nomor 09---------")
	nomor9([]int{2, 3, 4, 5, 6, 7, 8, 9, 1, 10})
	println()
	println("Nomor 10---------")
	nomor10([]int{4, 7, 1, 20}, 9)
	println()
}
