package main

import "fmt"

func printMatrix1(matrix [5][5]int) {
	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix); j++ {
			fmt.Printf("%d ", matrix[i][j])
		}

		fmt.Println(" ")
	}
}

// 1.
func array2D() [5][5]int {
	var matrix [5][5]int

	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix); j++ {
			if i == j {
				matrix[i][j] = i + 1
			} else {
				matrix[i][j] = 10
			}
		}

		for k := len(matrix) - 1; k > i; k-- {
			if k < (len(matrix)+i)+1 {
				matrix[i][k] = 20
			}
		}
	}

	return matrix
}

// 2.
func array2D2() [5][5]int {
	var matrix [5][5]int
	pow := 2

	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix); j++ {
			if i == j {
				matrix[i][j] = pow
				pow *= 2
			} else {
				matrix[i][j] = 22
			}
		}

		for k := len(matrix) - 1; k > i; k-- {
			if k < (len(matrix)+i)+1 {
				matrix[i][k] = 88
			}
		}
	}

	return matrix
}

// // 3.
func pointStudents() {
	answers := [][]string{
		{"A", "B", "A", "C", "C", "D", "E", "E", "A", "D"},
		{"D", "B", "A", "B", "C", "A", "E", "E", "A", "D"},
		{"E", "D", "D", "A", "C", "B", "E", "E", "A", "D"},
		{"C", "B", "A", "E", "D", "C", "E", "E", "A", "D"},
		{"A", "B", "D", "C", "C", "D", "E", "E", "A", "D"},
		{"B", "B", "E", "C", "C", "D", "E", "E", "A", "D"},
		{"B", "B", "A", "C", "C", "D", "E", "E", "A", "D"},
		{"E", "B", "E", "C", "C", "D", "E", "E", "A", "D"},
	}

	answerKey := []string{"D", "B", "D", "C", "C", "D", "A", "E", "A", "D"}

	points := make([]int, len(answers))

	for i, answer := range answers {
		for j, choice := range answer {
			if choice == answerKey[j] {
				points[i]++
			}
		}
	}

	for k, point := range points {
		fmt.Printf("Student-%d: %d\n", k+1, point*10)
	}
}

func printMatrix2(matrix [7][7]int) {
	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix); j++ {
			fmt.Printf("%d ", matrix[i][j])
		}

		fmt.Println(" ")
	}
}

// 4.
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
				index2 -= midValue
			} else if i == len(matrix)-1 && j == len(matrix)-4 {
				matrix[i][j] = midValue * 27
			}
		}
	}

	return matrix
}

// 5.
func kotakBolongMatrix() [7][7]int {
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

// 6.
func kotakBolongMatrix2() [7][7]int {
	var matrix [7][7]int
	result1 := 0
	result2 := 1

	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix); j++ {
			if i == 0 || j == len(matrix)-1 {
				matrix[i][j] = result1
				result1++
			} else if j == 0 || i == len(matrix)-1 {
				matrix[i][j] = result2
				result2++
			}
		}
	}

	return matrix
}

// 7.
func kotakIsiMatrix(n int) {
	matrix := make([][]int, n+1)
	sumRow := 0
	sumCol := 0
	sumDiag := 0

	for i := 0; i < n+1; i++ {
		matrix[i] = make([]int, n+1)
	}

	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			matrix[i][j] = i + j
		}
	}

	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			fmt.Print(matrix[i][j], ", ")

			sumRow += matrix[i][j]
		}

		fmt.Print("[", sumRow, "]")

		sumRow = 0

		fmt.Println()
	}

	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			sumCol += matrix[i][j]

			if i == j {
				sumDiag += matrix[i][j]
			}
		}

		fmt.Print("[", sumCol, "]")

		sumCol = 0
	}

	fmt.Print("[", sumDiag, "]")
}

func main() {
	fmt.Println("(1.) ")
	printMatrix1(array2D())
	fmt.Println(" ")

	fmt.Println("(2.) ")
	printMatrix1(array2D2())
	fmt.Println(" ")

	fmt.Println("(3.) ")
	pointStudents()
	fmt.Println(" ")

	fmt.Println("(4.) ")
	printMatrix2(segitigaMatrix())
	fmt.Println(" ")

	fmt.Println("(5.) ")
	printMatrix2(kotakBolongMatrix())
	fmt.Println(" ")

	fmt.Println("(6.) ")
	printMatrix2(kotakBolongMatrix2())
	fmt.Println(" ")

	fmt.Println("(7.) ")
	kotakIsiMatrix(7)
	fmt.Println(" ")
}
