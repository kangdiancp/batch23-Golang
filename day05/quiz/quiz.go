package main

import "fmt"

func main() {
	printMatrix(initDiagonal())
	// printMatrix(initDiagonal2())
	// (kunciJawabanSoal())
	// segitigaMatrix()
	// printMatrix(kotakMatrix())
	// (kotakMatrix2(7))
	// (kotakMatrix3(7))
}

// no 1
// program array 2d baris diagonal
func initDiagonal() [5][5]int {
	// sumDiagonal := 0
	var matrix [5][5]int
	for row := 0; row < len(matrix); row++ {
		for col := 0; col < len(matrix[row]); col++ {
			if row == col {
				matrix[row][col] = row + (1 * 1)
			} else if col < row {
				matrix[row][col] = 10
			} else {
				matrix[row][col] = 20
			}
		}
	}
	return matrix
}

func printMatrix(matrix [7][7]int) {
	for row := 0; row < len(matrix); row++ {
		for col := 0; col < len(matrix[row]); col++ {
			fmt.Printf("%d ", matrix[row][col])
		}
		fmt.Println()
	}
}

//no 2
//program array 2d jumlahkan diagonalnya
func initDiagonal2() [5][5]int {
	// sumDiagonal := 0
	var matrix [5][5]int
	for row := 0; row < len(matrix); row++ {
		for col := 0; col < len(matrix[row]); col++ {
			if row == col {
				matrix[row][col] = 2 << row
			} else if col < row {
				matrix[row][col] = 88
			} else {
				matrix[row][col] = 22
			}
		}
	}
	return matrix
}

func printMatrix2(matrix [5][5]int) {
	for row := 0; row < len(matrix); row++ {
		for col := 0; col < len(matrix[row]); col++ {
			fmt.Printf("%d ", matrix[row][col])
		}
		fmt.Println()
	}
}

//no 3
//tampilkan point yang didapat dari 8 orang siswa jika jawabannya benar
func kunciJawabanSoal() {
	kunciJawaban := []string{"D", "B", "D", "C", "C", "D", "A", "E", "A", "D"}
	jawabanSiswa := [][]string{
		{"A", "B", "A", "C", "C", "D", "E", "E", "A", "D"},
		{"D", "B", "A", "B", "C", "A", "E", "E", "A", "D"},
		{"E", "D", "D", "A", "C", "B", "E", "E", "A", "D"},
		{"C", "B", "A", "E", "D", "C", "E", "E", "A", "D"},
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

//no 4
//buat kotak array segitiga n=7 m=3 o=2
func segitigaMatrix() [7][7]int {
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
				index2 = midValue
			} else if i == len(matrix)-1 && j == len(matrix)-4 {
				matrix[i][j] = midValue * 27
			}
		}
	}
	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[i]); j++ {
			fmt.Printf("%d", matrix[i][j])
		}
		fmt.Println()
	}
	return matrix
}

//no 5
//buat kotak array sama sisi n=7 m=3 o=2
func kotakMatrix() [7][7]int {
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
		fmt.Println()
	}
	return matrix
}

//no 6
//buat kotak array sama n=7
func kotakMatrix2(length int) {
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
			fmt.Printf("%d", matrix[i][j])
		}
		fmt.Println()
	}
}

//no 7
//kotak array n=7
func kotakMatrix3(length int) {
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
