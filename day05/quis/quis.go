package main

import (
	"fmt"
	"strconv"
)

// print matriks
// func printMatriks(matriks [5][5]int) {
// 	for i := 0; i < len(matriks); i++ {
// 		for j := 0; j < len(matriks[j]); j++ {
// 			fmt.Printf("%d ", matriks[i][j])
// 		}
// 		fmt.Println()
// 	}
// }

//nomor1

func initArrayDiagonalSum() {
	var matriks [5][5]int
	sum := 0
	for i := 0; i < len(matriks); i++ {
		for j := 0; j < len(matriks); j++ {

			if i == j {
				matriks[i][j] = j + 1
				sum += j + 1
			} else {
				if i < j {
					matriks[i][j] = 20
				} else {
					matriks[i][j] = 10

				}
			}
		}
	}
	for i := 0; i < len(matriks); i++ {
		for j := 0; j < len(matriks[j]); j++ {
			fmt.Printf("%d ", matriks[i][j])
		}
		fmt.Println()
	}
	fmt.Printf("Total Sum: %d", sum)

}

//nomor2

func initArrayDiagonalKaliDanSum() {
	var matriks [5][5]int
	sum := 0
	//matriks[0][0] = 2
	for i := 0; i < len(matriks); i++ {
		//matriks[i][0] = 22

		for j := 0; j < len(matriks); j++ {
			//matriks[0][j] = 88

			if i == j {
				matriks[i][j] = 2 << j
				sum += 2 << j
			} else {
				if i < j {
					matriks[i][j] = 88
				} else {
					matriks[i][j] = 22
				}
			}

		}
	}
	//sum += matriks[0][0]
	for i := 0; i < len(matriks); i++ {
		for j := 0; j < len(matriks[j]); j++ {
			fmt.Printf("%d ", matriks[i][j])
		}
		fmt.Println()
	}
	fmt.Printf("Total Sum: %d", sum)

}

//nomor3

func NilaiUlangan() {
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

	for i := 0; i < len(matrix)-1; i++ {
		points := 0
		for j := 1; j < len(matrix[i]); j++ {
			if matrix[i][j] == matrix[len(matrix)-1][j] {
				points = points + 10
			}
		}
		fmt.Printf("\n%s: %d point", matrix[i][0], points)
	}
}

// nomor4
func segitigaKanan() {
	var matrix [7][7]int
	simpanDiagonal := make([]int, len(matrix))
	simpanRight := make([]int, len(matrix))
	simpanBottom := make([]int, len(matrix))
	varInti := len(matrix) - 1

	simpanDiagonal[0] = 2
	increment := 3

	for i := 1; i < len(matrix); i++ {
		if i == varInti/2 {
			simpanDiagonal[i] = increment
			simpanDiagonal[i+1] = simpanDiagonal[i-1] + simpanDiagonal[i]
			i++
		} else if i%3 == 0 || i%3 != 0 {
			simpanDiagonal[i] = simpanDiagonal[i-1] + increment
		}
	}

	simpanRight[0] = simpanDiagonal[varInti]
	simpanRight[1] = simpanDiagonal[3] * increment
	for i := 2; i < len(matrix); i++ {
		if simpanRight[0] != 0 {
			if i == 2 {
				simpanRight[i] = simpanDiagonal[varInti] + increment
			} else if i == 5 {
				simpanRight[i] = simpanRight[1] * increment
				simpanRight[i+1] = simpanRight[i-1] + increment
				i++
			} else {
				simpanRight[i] = simpanRight[i-1] + increment
			}

		}
	}

	simpanBottom[0] = 2
	simpanBottom[1] = 41
	simpanBottom[varInti] = simpanRight[varInti]
	for i := 2; i < len(matrix)-1; i++ {
		if simpanBottom[0] != 0 {
			if i == 3 {
				simpanBottom[i] = simpanRight[varInti-1] * increment
				simpanBottom[i+1] = simpanBottom[i-1] - increment
				i++
			} else {
				simpanBottom[i] = simpanBottom[i-1] - increment
			}
		}
	}

	for row := 0; row < len(matrix); row++ {
		for col := 0; col < len(matrix); col++ {
			if row+col == varInti {
				matrix[row][col] = simpanDiagonal[col]
			} else if col == varInti {
				matrix[row][col] = simpanRight[row]
			} else if row == varInti {
				matrix[row][col] = simpanBottom[col]
			} else {
				matrix[row][col] = 0
			}

			fmt.Printf("%d  ", matrix[row][col])
		}
		fmt.Println()
	}

}

// nomor 5
func kotak(n, o, m int) {

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

		} else if sisa == 1 {
			sisa2 = 1

		} else if sisa == 2 {
			sisa2 = 3

		} else if sisa == 3 {
			sisa2 = 1

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

		} else if sisa2 == 1 {
			sisa3 = 1

		} else if sisa2 == 2 {
			sisa3 = 3

		} else if sisa2 == 3 {
			sisa3 = 1

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

	//print matrix
	stringMatrix := make([][]string, len(matrix))
	for i := 0; i < len(matrix); i++ {
		stringMatrix[i] = make([]string, len(matrix[i]))
		for j := 0; j < len(matrix); j++ {
			stringMatrix[i][j] = strconv.Itoa(matrix[i][j])
			if stringMatrix[i][j] == "0" {

				stringMatrix[i][j] = " "

			}
		}
	}

	for i := 0; i < len(stringMatrix); i++ {
		for j := 0; j < len(stringMatrix); j++ {
			fmt.Print(stringMatrix[i][j], "	")
		}
		fmt.Println()
	}
}

//nomor6

func kotakTengah() [7][7]int {
	var matriks [7][7]int

	for baris := 0; baris < len(matriks); baris++ {

		for kolom := 0; kolom < len(matriks[baris]); kolom++ {
			if baris == 0 || baris == len(matriks)-1 || kolom == 0 || kolom == len(matriks)-1 {

				matriks[0][kolom] = kolom
				matriks[baris][0] = baris
				matriks[baris][len(matriks)-1] = baris + matriks[0][len(matriks)-1]
				matriks[len(matriks)-1][kolom] = kolom + matriks[len(matriks)-1][0]

			} else {
				matriks[baris][kolom] = 0
			}

		}

	}

	return matriks
}

// print nomor 6
func printMatriks(matriks [7][7]int) {
	stringMatrix := make([][]string, len(matriks))
	for i := 0; i < len(matriks); i++ {
		stringMatrix[i] = make([]string, len(matriks[i]))
		for j := 0; j < len(matriks[i]); j++ {
			stringMatrix[i][j] = strconv.Itoa(matriks[i][j])
			if stringMatrix[i][j] == "0" {
				if i == 0 && j == 0 {
					stringMatrix[i][j] = "0"
				} else {
					stringMatrix[i][j] = " "

				}
			}
		}
	}

	for i := 0; i < len(stringMatrix); i++ {
		for j := 0; j < len(stringMatrix[i]); j++ {
			fmt.Print(stringMatrix[i][j], " ")
		}
		fmt.Println()
	}
}

//nomor7

func matriksUrutSum(n int) {

	matrix := make([][]int, n+1)
	for i := 0; i < n+1; i++ {
		matrix[i] = make([]int, n+1)
	}

	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			matrix[i][j] = i + j

		}

	}
	var sumBaris int
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			fmt.Print(matrix[i][j], ", ")
			sumBaris += matrix[i][j]

		}
		fmt.Print("[", sumBaris, "]")

		sumBaris = 0

		fmt.Println()
	}

	sumKolom := 0
	sumDiagonal := 0
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
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

	initArrayDiagonalSum()
	initArrayDiagonalKaliDanSum()
	NilaiUlangan()
	segitigaKanan()
	kotak(7, 2, 3)
	printMatriks(kotakTengah())
	matriksUrutSum(7)

}
