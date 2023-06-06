package main

import "fmt"

func soal01() {
	var matrix [5][5]int

	sumDiagonal := 0

	// Isi pada Matrix
	for row := 0; row < len(matrix); row++ {
		for col := 0; col < len(matrix); col++ {

			if row == col {
				matrix[row][col] = col + 1
				sumDiagonal = sumDiagonal + matrix[row][col]
			} else if row > col {
				matrix[row][col] = 10
			} else {
				matrix[row][col] = 20
			}
		}
	}

	// Menampilkan isi Nilai ke dalam Matrix
	for row := 0; row < len(matrix); row++ {
		for col := 0; col < len(matrix); col++ {
			fmt.Printf("%d ", matrix[row][col])
		}
		fmt.Println()
	}
	fmt.Println("Total Sum :", sumDiagonal)
}

func soal02() {
	var matrix [5][5]int
	nDiagonal := make([]int, len(matrix))
	nDiagonal[0] = 2
	sumDiagonal := 0

	// Menampung nilai nDiagonal
	for i := 1; i < len(matrix); i++ {
		nDiagonal[i] = nDiagonal[i-1] * 2
	}

	// Mengisi nilai Matrix
	for row := 0; row < len(matrix); row++ {
		for col := 0; col < len(matrix); col++ {

			// Kondisi
			if row == col {
				matrix[row][col] = nDiagonal[row]
				sumDiagonal = sumDiagonal + matrix[row][col]
			} else if row > col {
				matrix[row][col] = 22
			} else {
				matrix[row][col] = 88
			}
		}
	}

	// Menampilkan isi Nilai ke dalam Matrix
	for row := 0; row < len(matrix); row++ {
		for col := 0; col < len(matrix); col++ {
			fmt.Printf("%d ", matrix[row][col])
		}
		fmt.Println()
	}
	fmt.Println("Total Sum :", sumDiagonal)
}

func soal03() {
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
		fmt.Printf("\n%s: %d point", matriks[row][0], point)
	}
}

// soal 04
func soal04() [7][7]int {
	var matrix [7][7]int
	index0 := 17
	midValue := 3
	index1 := index0 + midValue // 20
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

func main() {
	// soal01()
	// soal02()
	// soal03()
	soal04()
}
