package main

import "fmt"

func barisDiagonal() {
	var matrix [5][5]int
	sumdiagonal := 0
	for row := 0; row < len(matrix); row++ {
		for col := 0; col < len(matrix); col++ {
			if row == col {
				matrix[row][col] = row + 1
				sumdiagonal += matrix[row][col]
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
	fmt.Print("   Total Sum : ", sumdiagonal)
}

func kaliDiagonal() {
	var matrix [5][5]int
	kalidua := make([]int, len(matrix))
	kalidua[0] = 2
	for i := 1; i < len(matrix); i++ {
		kalidua[i] = kalidua[i-1] * 2
	}

	sumdiagonal := 0
	for row := 0; row < len(matrix); row++ {
		for col := 0; col < len(matrix); col++ {
			if row == col {
				matrix[row][col] = kalidua[row]
				sumdiagonal += matrix[row][col]
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
	fmt.Print("   Total Sum : ", sumdiagonal)
}

func ujianSiswa() {
	answerKey := []string{"D", "B", "D", "C", "C", "D", "A", "E", "A", "D"}

	studentAnswers := [][]string{
		{"A", "B", "A", "C", "C", "D", "E", "E", "A", "D"},
		{"D", "B", "A", "B", "C", "A", "E", "E", "A", "D"},
		{"E", "D", "D", "A", "C", "B", "E", "E", "A", "D"},
		{"C", "B", "A", "E", "D", "C", "E", "E", "A", "D"},
		{"A", "B", "D", "C", "C", "D", "E", "E", "A", "D"},
		{"B", "B", "E", "C", "C", "D", "E", "E", "A", "D"},
		{"B", "B", "A", "C", "C", "D", "E", "E", "A", "D"},
		{"E", "B", "E", "C", "C", "D", "E", "E", "A", "D"},
	}
	for i, answers := range studentAnswers {
		points := 0
		for j, answer := range answers {
			if answer == answerKey[j] {
				points += 10
			}
		}
		fmt.Printf("Siswa-%d: %d point\n", i, points)
	}
}

func kotakenam() {
	N := 7

	kotakArray := make([][]int, N)
	for i := 0; i < N; i++ {
		kotakArray[i] = make([]int, N)
	}
	for i := 0; i < N; i++ {
		for j := 0; j < N; j++ {
			if i == 0 || j == 0 || i == N-1 || j == N-1 {
				kotakArray[i][j] = i + j
			}
		}
	}
	for i := 0; i < N; i++ {
		for j := 0; j < N; j++ {
			if kotakArray[i][j] != 0 {
				fmt.Printf("%2d ", kotakArray[i][j])
			} else {
				fmt.Printf("    ")
			}
		}
		fmt.Println()
	}

}

func kotaklima() [7][7]int {
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
			if matriks[i][j] != 0 {
				fmt.Printf("%d ", matriks[i][j])
			} else {
				fmt.Printf("   ")
			}
		}
		fmt.Println()
	}
	return matriks
}

func kotakempat() [7][7]int {
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
			if matrix[i][j] != 0 {
				fmt.Printf("%d ", matrix[i][j])
			} else {
				fmt.Printf("   ")
			}
		}
		fmt.Println()
	}
	return matrix
}

func kotaktujuh() {
	N := 8

	// Membuat matriks dengan ukuran N x N
	matrix := make([][]int, N)
	for i := 0; i < N; i++ {
		matrix[i] = make([]int, N)
	}

	// Mengisi matriks dengan pola yang diinginkan
	for i := 0; i < N-1; i++ {
		for j := 0; j < N-1; j++ {
			matrix[i][j] = i + j
		}
	}
	for i := 0; i < N-1; i++ {
		matrix[i][N-1] = (i + 3) * 7
		matrix[N-1][i] = (i + 3) * 7
	}
	matrix[N-1][N-1] = 42

	// Menampilkan matriks
	for i := 0; i < N; i++ {
		for j := 0; j < N; j++ {
			fmt.Printf("%2d ", matrix[i][j])
		}
		fmt.Println()
	}
}

func main() {
	//barisDiagonal()
	//kaliDiagonal()
	//ujianSiswa()
	//kotakenam()
	//kotaklima()
	//kotakempat()
	kotaktujuh()

}
