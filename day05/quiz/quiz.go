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

func nomor1() {
	var matrix [5][5]int
	var sum int

	for row := 0; row < len(matrix); row++ {
		for col := 0; col < len(matrix); col++ {
			if row == col {
				for i := 1; i <= row+1; i++ {
					matrix[row][col] = i
				}
				sum += matrix[row][col]
			} else if row > col {
				matrix[row][col] = 10
			} else {
				matrix[row][col] = 20
			}
		}
	}

	for row := 0; row < len(matrix); row++ {
		for col := 0; col < len(matrix); col++ {
			fmt.Printf("%d ", matrix[row][col])
		}
		fmt.Println()
	}

	fmt.Println("Total Sum : ", sum)

}

func nomor2() {
	var matrix [5][5]int
	var sum int
	diagonal := make([]int, len(matrix))
	diagonal[0] = 2

	for i := 1; i < len(matrix); i++ {
		diagonal[i] = diagonal[i-1] * 2
	}

	for row := 0; row < len(matrix); row++ {
		for col := 0; col < len(matrix); col++ {
			if row == col {
				matrix[row][col] = diagonal[row]
				sum += matrix[row][col]
			} else if row > col {
				matrix[row][col] = 22
			} else {
				matrix[row][col] = 88
			}
		}
	}

	for row := 0; row < len(matrix); row++ {
		for col := 0; col < len(matrix); col++ {
			fmt.Printf("%d ", matrix[row][col])
		}
		fmt.Println()
	}

	fmt.Println("Total Sum : ", sum)
}

func nomor3() {
	matriks := [9][11]string{
		{"student-0", "A", "B", "A", "C", "C", "D", "E", "E", "A", "D"},
		{"student-1", "D", "B", "A", "B", "C", "A", "E", "E", "A", "D"},
		{"student-2", "E", "D", "D", "A", "C", "B", "E", "E", "A", "D"},
		{"student-3", "C", "B", "A", "E", "D", "C", "E", "E", "A", "D"},
		{"student-4", "A", "B", "D", "C", "C", "D", "E", "E", "A", "D"},
		{"student-5", "B", "B", "E", "C", "C", "D", "E", "E", "A", "D"},
		{"student-6", "B", "B", "A", "C", "C", "D", "E", "E", "A", "D"},
		{"student-7", "E", "B", "E", "C", "C", "D", "E", "E", "A", "D"},
		{"Kunci Jawaban", "D", "B", "D", "C", "C", "D", "A", "E", "A", "D"},
	}

	for row := 0; row < len(matriks)-1; row++ {
		point := 0
		for col := 1; col < len(matriks[row]); col++ {
			if matriks[row][col] == matriks[len(matriks)-1][col] {
				point += 10
			}
		}
		fmt.Printf("\n%s: %d point ", matriks[row][0], point)
	}
}

func nomor4() [7][7]int {
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
	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[i]); j++ {
			fmt.Printf("%d ", matrix[i][j])
		}
		fmt.Println()
	}
	return matrix
}

func nomor5() [7][7]int {
	var matriks [7][7]int
	n := 2

	for i := 0; i < len(matriks); i++ {
		for j := 0; j < len(matriks); j++ {
			if i == 0 && j != len(matriks)/2 {
				matriks[i][j] = n
				n += 3
			} else if (i == 0 && j == len(matriks)/2) || (i == len(matriks)-1 && j == len(matriks)/2) || (j == 0 && (i == 1 || i == len(matriks)-2)) || (j == len(matriks)-1 && (i == 1 || i == len(matriks)-2)) {
				matriks[i][j] = 3
			} else if i > 1 && j == len(matriks)-1 && i != len(matriks)-2 {
				matriks[i][j] = n
				n += 3
			} else if i == 0 && j == len(matriks)-1 && i != len(matriks)-2 {
				matriks[i][j] = n
				n += 3
			}
		}
	}
	for row := len(matriks) - 1; row >= 0; row-- {
		for col := len(matriks) - 1; col >= 0; col-- {
			if (row == len(matriks)-1 && col < len(matriks)-1) && col != len(matriks)/2 {
				matriks[row][col] = n
				n += 3
			} else if (row > 1 && col == 0) && row != len(matriks)-2 {
				matriks[row][col] = n
				n += 3
			}
		}

	}
	for i := 0; i < len(matriks); i++ {
		for j := 0; j < len(matriks); j++ {
			fmt.Printf("%d ", matriks[i][j])
		}
		fmt.Println()
	}
	return matriks
}

func nomor6(length int) {
	matrix := make([][]int, length+1)
	for i := 0; i <= length; i++ {
		matrix[i] = make([]int, length+1)
	}

	for row := 0; row < length; row++ {
		for col := 0; col < length; col++ {
			matrix[0][col] = col
			matrix[row][0] = row
			matrix[length-1][col] = col + length - 1
			matrix[row][length-1] = row + length - 1
		}
	}

	for i := 0; i < length; i++ {
		for j := 0; j < length; j++ {
			fmt.Printf("%d ", matrix[i][j])
		}
		fmt.Println()
	}
}

func nomor7(length int) {
	matrix := make([][]int, length+1)
	for i := 0; i < length+1; i++ {
		matrix[i] = make([]int, length+1)
	}
	for i := 0; i < length; i++ {
		for j := 0; j < length; j++ {
			matrix[i][j] = i + j
		}
	}
	var sumBaris int
	for i := 0; i < length; i++ {
		for j := 0; j < length; j++ {
			fmt.Print(matrix[i][j], " ")
			sumBaris += matrix[i][j]

		}
		fmt.Print("[", sumBaris, "]")

		sumBaris = 0

		fmt.Println()
	}

	sumKolom := 0
	sumDiagonal := 0
	for i := 0; i < length; i++ {
		for j := 0; j < length; j++ {
			sumKolom += matrix[i][j]
			if i == j {
				sumDiagonal += matrix[i][j]
			}
		}
		fmt.Print("[", sumKolom, "]")
		sumKolom = 0
	}
	fmt.Print("[", sumDiagonal, "]")
}

func main() {
	nomor1()
	nomor2()
	nomor3()
	nomor4()
	nomor5()
	nomor6(7)
	nomor7(7)
}
