package main

import (
	"fmt"
)

func printMatrix(matrix [5][5]int) {
	for row := 0; row < len(matrix); row++ {
		for col := 0; col < len(matrix); col++ {
			fmt.Printf("%d ", matrix[row][col])
		}
		fmt.Println()
	}
}

// Nomor 1 :
func sumDiagonal1() [5][5]int {
	var matrix [5][5]int

	for row := 0; row < len(matrix); row++ {
		for col := 0; col < len(matrix); col++ {
			if row == col {
				matrix[row][col] = row + 1
			} else if row > col {
				matrix[row][col] = 10
			} else {
				matrix[row][col] = 20
			}
		}
	}
	return matrix
}

// Nomor 2 :
func sumDiagonal2() [5][5]int {
	var matrix [5][5]int

	for row := 0; row < len(matrix); row++ {
		for col := 0; col < len(matrix); col++ {
			if row == col {
				matrix[row][col] = 2 << col
			} else if row > col {
				matrix[row][col] = 22
			} else {
				matrix[row][col] = 88
			}
		}
	}
	return matrix
}

// Nomor 3 :
func keyAnswers() {
	kunciJawaban := []string{"D", "B", "D", "C", "C", "D", "A", "E", "A", "D"}
	jawabanSiswa := [][]string{
		{"A", "B", "A", "C", "C", "D", "E", "E", "A", "D"},
		{"D", "B", "A", "B", "C", "A", "E", "E", "A", "D"},
		{"E", "D", "D", "A", "C", "B", "E", "E", "A", "D"},
		{"A", "B", "D", "C", "C", "D", "E", "E", "A", "D"},
		{"B", "B", "E", "C", "C", "D", "E", "E", "A", "D"},
		{"B", "B", "A", "C", "C", "D", "E", "E", "A", "D"},
		{"E", "B", "E", "C", "C", "D", "E", "E", "A", "D"},
	}

	for siswa, jawaban := range jawabanSiswa {
		point := 0
		for i, jawab := range jawaban {
			if jawab == kunciJawaban[i] {
				point += 10
			}
		}
		fmt.Printf("siswa-%d : %d point\n", siswa, point)
	}
}

// Nomor 4 :
func diagonalNum1() [7][7]int {
	var matrix [7][7]int
	index0 := 17
	midValue := 3
	index1 := index0 + midValue
	index2 := index0 + (midValue * (len(matrix) + 1))

	for row := 0; row < len(matrix); row++ {
		for col := 0; col < len(matrix); col++ {
			if row+col == len(matrix)-1 && row != col {
				matrix[row][col] = index0
				index0 -= midValue
			} else if row == len(matrix)/2 && col == len(matrix)/2 {
				matrix[row][col] = midValue
			} else if row == 1 && col == len(matrix)-1 {
				matrix[row][col] = midValue * 3
			} else if row > 1 && col == len(matrix)-1 && row != len(matrix)-2 {
				matrix[row][col] = index1
				index1 += 3
			} else if row == len(matrix)-2 && col == len(matrix)-1 {
				matrix[row][col] = midValue * 9
			} else if col > 0 && row == len(matrix)-1 && col != len(matrix)/2 {
				matrix[row][col] = index2
				index2 -= midValue
			} else if row == len(matrix)-1 && col == len(matrix)-4 {
				matrix[row][col] = midValue * 27
			}
		}
	}
	for row := 0; row < len(matrix); row++ {
		for col := 0; col < len(matrix[row]); col++ {
			fmt.Printf("%d\t", matrix[row][col])
		}
		fmt.Println()
	}
	return matrix
}

// Nomor 5 :
func diagonalNum2() [7][7]int {
	var matrix [7][7]int
	n := 2

	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix); j++ {
			if i == 0 && j != len(matrix)/2 {
				matrix[i][j] = n
				n += 3
			} else if (i == 0 && j == len(matrix)/2) || (i == len(matrix)-1 && j == len(matrix)/2) || (j == 0 && (i == 1 || i == len(matrix)-2)) || (j == len(matrix)-1 && (i == 1 || i == len(matrix)-2)) {
				matrix[i][j] = 3
			} else if i > 1 && j == len(matrix)-1 && i != len(matrix)-2 {
				matrix[i][j] = n
				n += 3
			} else if i == 0 && j == len(matrix)-1 && i != len(matrix)-2 {
				matrix[i][j] = n
				n += 3
			}
		}
	}
	for row := len(matrix) - 1; row >= 0; row-- {
		for col := len(matrix) - 1; col >= 0; col-- {
			if (row == len(matrix)-1 && col < len(matrix)-1) && col != len(matrix)/2 {
				matrix[row][col] = n
				n += 3
			} else if (row > 1 && col == 0) && row != len(matrix)-2 {
				matrix[row][col] = n
				n += 3
			}
		}

	}

	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix); j++ {
			fmt.Printf("%d\t", matrix[i][j])
		}
		fmt.Println()
	}
	return matrix
}

// Nomor 6 :
func diagonalNum3(n int) {
	matrix := make([][]int, n)
	for i := 0; i < n; i++ {
		matrix[i] = make([]int, n)
	}

	for row := 0; row < n; row++ {
		for col := 0; col < n; col++ {
			if row == col || row+col == n-1 {
				matrix[row][col] = 0
			} else {
				matrix[row][col] = row + col
			}
		}
	}

	for row := 0; row < n; row++ {
		for col := 0; col < n; col++ {
			fmt.Printf("%d\t", matrix[row][col])
		}
		fmt.Println()
	}
}

// Nomor 7 :
func diagonalNum4(n int) {
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
		rowSum := 0
		// sum row
		for col := 0; col < n; col++ {
			rowSum += matrix[row][col]
			fmt.Printf("%d\t", matrix[row][col])
		}
		fmt.Println(rowSum)
	}
	diagonalSum := 0
	// sum row
	for col := 0; col < n; col++ {
		colSum := 0
		// sum column
		for row := 0; row < n; row++ {
			colSum += matrix[row][col]
			if row == col {
				diagonalSum += matrix[row][col]
			}
		}
		fmt.Printf("%d\t", colSum)
	}
	fmt.Print(diagonalSum)
	fmt.Println()
}

func main() {
	printMatrix(sumDiagonal1())
	printMatrix(sumDiagonal2())
	keyAnswers()
	diagonalNum1()
	diagonalNum2()
	diagonalNum3(7)
	diagonalNum4(7)
}
