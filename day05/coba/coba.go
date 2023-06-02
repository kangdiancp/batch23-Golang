package main

import (
	"fmt"
)

func generateMatrix(n, o, m int) [][]int {
	matrix := make([][]int, n)
	for i := 0; i < n; i++ {
		matrix[i] = make([]int, n)
	}

	rowStart, rowEnd := 0, n-1
	colStart, colEnd := 0, n-1
	//direction := 0
	for o <= n*n {
		// Menambahkan elemen pada baris atas
		for i := colStart; i <= colEnd; i++ {

			matrix[rowStart][i] = o

			if (i+1)%(m+1) == 0 {
				matrix[rowStart][i] = m
				o -= m

			}

			o += m
		}
		sisa := (len(matrix)) % 4

		fmt.Println(sisa)
		rowStart++

		// Menambahkan elemen pada kolom kanan
		for i := rowStart; i <= rowEnd; i++ {
			matrix[i][colEnd] = o
			if (i+sisa)%(m+1) == 0 {
				matrix[i][colEnd] = m
				o -= m
			}

			o += m
		}
		colEnd--
		var sisa2 int
		if sisa == 0 {
			sisa2 = 3
			fmt.Println(sisa2)

		} else if sisa == 1 {
			sisa2 = 1
			fmt.Println(sisa2)

		} else if sisa == 2 {
			sisa2 = 3
			fmt.Println(sisa2)

		} else if sisa == 3 {
			sisa2 = 1
			fmt.Println(sisa2)

		}

		// Menambahkan elemen pada baris bawah
		for i := colEnd; i >= colStart; i-- {
			matrix[rowEnd][i] = o

			if (i+sisa)%(m+1) == 2 {

				matrix[rowEnd][i] = m
				o -= m
			}
			o += m

		}
		rowEnd--
		var sisa3 int
		if sisa2 == 0 {
			sisa3 = 3
			fmt.Println(sisa3)

		} else if sisa2 == 1 {
			sisa3 = 1
			fmt.Println(sisa3)

		} else if sisa2 == 2 {
			sisa3 = 3
			fmt.Println(sisa3)

		} else if sisa2 == 3 {
			sisa3 = 1
			fmt.Println(sisa3)

		}

		// Menambahkan elemen pada kolom kiri
		for i := rowEnd; i >= rowStart; i-- {
			matrix[i][colStart] = o
			if (i+sisa3)%(m+1) == 2 {

				matrix[i][colStart] = m
				o -= m
			}
			o += m
		}
		colStart++
	}

	return matrix
}

func main() {

	result := generateMatrix(7, 2, 3)

	for _, row := range result {
		for _, num := range row {
			fmt.Printf("%d,", num)
		}
		fmt.Println()
	}
}
