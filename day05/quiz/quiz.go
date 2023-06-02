package main

import "fmt"

func printMatrix(matrix [5][5]int, sum int) {
	for row := 0; row < len(matrix); row++ {
		for col := 0; col < len(matrix); col++ {
			fmt.Printf("%d ", matrix[row][col])
		}
		fmt.Println()
	}
	fmt.Println("total sum:", sum)
}

func printMatrix2(matrix [7][7]int) {
	for row := 0; row < len(matrix); row++ {
		for col := 0; col < len(matrix); col++ {
			fmt.Printf("%d ", matrix[row][col])
		}
		fmt.Println()
	}
}

func sumDiagonal() ([5][5]int, int) {
	var matrix [5][5]int
	sum := 0
	for row := 0; row < len(matrix); row++ {
		for col := 0; col < len(matrix); col++ {
			if row == col {
				matrix[row][col] = col + 1
				sum += col + 1
			} else if row > col {
				matrix[row][col] = 10
			} else {
				matrix[row][col] = 20
			}
		}
	}
	return matrix, sum
}

func sumDiagonal2() ([5][5]int, int) {
	var matrix [5][5]int
	sum := 0
	for row := 0; row < len(matrix); row++ {
		for col := 0; col < len(matrix); col++ {
			if row == col {
				matrix[row][col] = 2 << col
				sum += 2 << col
			} else if row > col {
				matrix[row][col] = 22
			} else {
				matrix[row][col] = 88
			}
		}
	}
	return matrix, sum
}

func pointUjian() {
	matrix := [9][11]string{
		{"Student-0", "A", "B", "A", "C", "C", "D", "E", "E", "A", "D"},
		{"Student-1", "D", "B", "A", "B", "C", "A", "E", "E", "A", "D"},
		{"Student-2", "E", "D", "D", "A", "C", "B", "E", "E", "A", "D"},
		{"Student-3", "C", "B", "A", "E", "D", "C", "E", "E", "A", "D"},
		{"Student-4", "A", "B", "D", "C", "C", "D", "E", "E", "A", "D"},
		{"Student-5", "B", "B", "E", "C", "C", "D", "E", "E", "A", "D"},
		{"Student-6", "B", "B", "A", "C", "C", "D", "E", "E", "A", "D"},
		{"Student-7", "E", "B", "E", "C", "C", "D", "E", "E", "A", "D"},
		{"Kunci Jawaban", "D", "B", "D", "C", "C", "D", "A", "E", "A", "D"},
	}
	// for row := 0; row < len(matrix); row++ {
	// 	for col := 0; col < len(matrix); col++ {
	// 		fmt.Printf("%s ", matrix[row][col])
	// 	}
	// 	fmt.Println()
	// }

	for row := 0; row < len(matrix)-1; row++ {
		points := 0
		for col := 1; col < len(matrix[row]); col++ {
			if matrix[row][col] == matrix[len(matrix)-1][col] {
				points += 10
			}
		}
		fmt.Printf("\n%s: %d point", matrix[row][0], points)
	}
}

func kotak4() [7][7]int {
	var matrix [7][7]int
	index0 := 17
	midValue := 3
	index1 := index0 + midValue
	index2 := index0 + (midValue * (len(matrix) + 1))

	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix); j++ {
			if i+j == len(matrix)-1 && i != j {
				matrix[i][j] = index0
				index0 -= midValue
			} else if i == len(matrix)/2 && j == len(matrix)/2 {
				matrix[i][j] = midValue
			} else if i == 1 && j == len(matrix)-1 {
				matrix[i][j] = midValue * 3
			} else if i > 1 && j == len(matrix)-1 && i != len(matrix)-2 {
				matrix[i][j] = index1
				index1 += 3
			} else if i == len(matrix)-2 && j == len(matrix)-1 {
				matrix[i][j] = midValue * 9
			} else if j > 0 && i == len(matrix)-1 && j != len(matrix)/2 {
				matrix[i][j] = index2
				index2 -= midValue
			} else if i == len(matrix)-1 && j == len(matrix)-4 {
				matrix[i][j] = midValue * 27
			}
		}
	}

	return matrix
}

func kotak5() [7][7]int {
	var matrix [7][7]int
	index0 := 2

	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix); j++ {
			if i == 0 && j != len(matrix)/2 {
				matrix[i][j] = index0
				index0 += 3
			} else if i > 1 && j == len(matrix)-1 && i != len(matrix)-2 {
				matrix[i][j] = index0
				index0 += 3
			} else if (i == 0 && j == len(matrix)/2) || (i == len(matrix)-1 && j == len(matrix)/2) || (j == 0 && (i == 1 || i == len(matrix)-2)) || (j == len(matrix)-1 && (i == 1 || i == len(matrix)-2)) {
				matrix[i][j] = 3
			}
		}
	}

	for k := len(matrix) - 1; k >= 0; k-- {
		for l := len(matrix) - 1; l >= 0; l-- {
			if (k == len(matrix)-1 && l < len(matrix)-1) && l != len(matrix)/2 {
				matrix[k][l] = index0
				index0 += 3
			} else if (k > 1 && l == 0) && k != len(matrix)-2 {
				matrix[k][l] = index0
				index0 += 3
			}
		}
	}

	return matrix
}

// func kotak5(n, m, o int) [][]int {
// 	matrix := make([][]int, n)

// 	for i := 0; i < n; i++ {
// 		matrix[i] = make([]int, n)
// 	}

// 	currentV := o
// 	printCount := 0

// 	for row := 0; row < n; row++ {
// 		for col := 0; col < n; col++ {
// 			matrix[row][col] = currentV
// 			currentV += 3

// 			if printCount < 3 {
// 				printCount++
// 			} else {
// 				matrix[row][col] = 3
// 				printCount = 0
// 			}
// 		}
// 		fmt.Println()
// 	}
// 	return matrix
// }

func kotak6(n int) {
	// var matrix [7][7]int

	matrix := make([][]int, n)
	for i := 0; i < n; i++ {
		matrix[i] = make([]int, n)
	}

	for row := 0; row < n; row++ {
		for col := 0; col < n; col++ {
			matrix[0][col] = col
			matrix[row][0] = row

			matrix[n-1][col] = col + n - 1
			matrix[row][n-1] = row + n - 1
		}
	}

	for row := 0; row < n; row++ {
		for col := 0; col < n; col++ {
			fmt.Printf("%d ", matrix[row][col])
		}
		fmt.Println()
	}
}

func kotak7(n int) {
	matrix := make([][]int, n)
	for i := 0; i < n; i++ {
		matrix[i] = make([]int, n)
	}

	for row := 0; row < n; row++ {
		rowSum := 0
		for col := 0; col < n; col++ {
			matrix[row][col] = row + col
			rowSum += matrix[row][col]
		}
	}

	for row := 0; row < n; row++ {
		rowSum := 0 // Jumlah baris
		for col := 0; col < n; col++ {
			rowSum += matrix[row][col]
			fmt.Printf("%d ", matrix[row][col])
		}
		fmt.Println(rowSum)
	}
	diagonalSum := 0 // Jumlah diagonal
	for col := 0; col < n; col++ {
		colSum := 0 // Jumlah kolom

		for row := 0; row < n; row++ {
			colSum += matrix[row][col]
			if row == col {
				diagonalSum += matrix[row][col]
			}
		}
		fmt.Printf("%d ", colSum)
	}
	fmt.Print(diagonalSum)
	fmt.Println()
}

func main() {
	// printMatrix(sumDiagonal())
	// printMatrix(sumDiagonal2())
	// pointUjian()

	// printMatrix2(kotak4())
	// printMatrix2(kotak5())

	// kotak6(7)
	kotak7(7)
}
