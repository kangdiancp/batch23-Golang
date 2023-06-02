package main

import "fmt"

// func printMatrix(matrix[5][5]int)  {

// }

func quis1() [5][5]int {
	var matrix [5][5]int
	var sum int
	for row := 0; row < len(matrix); row++ {
		for col := 0; col < len(matrix); col++ {
			if row == col {
				matrix[row][col] = row + 1
				sum += matrix[row][col]
			} else if row <= col {
				matrix[row][col] = 20
			} else {
				matrix[row][col] = 10
			}
			fmt.Printf("%d  ", matrix[row][col])
		}
		fmt.Println()
	}
	fmt.Println("Total Sum : ", sum)
	return matrix
}

func quis2() [5][5]int {
	var matrix [5][5]int
	var sum int
	for row := 0; row < len(matrix); row++ {
		for col := 0; col < len(matrix); col++ {
			if row == col {
				matrix[row][col] = 2 << col
				sum += 2 << col
			} else if row <= col {
				matrix[row][col] = 88
			} else {
				matrix[row][col] = 22
			}
			fmt.Printf("%d  ", matrix[row][col])
		}
		fmt.Println()
	}
	fmt.Println("Total Sum : ", sum)
	return matrix
}

func quis3() {
	data := [9][11]string{
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

	for row := 0; row < len(data)-1; row++ {
		points := 0
		for col := 1; col < len(data[row]); col++ {
			if data[row][col] == data[len(data)-1][col] {
				points += 10
			}
		}
		fmt.Printf("\n%s: %d point", data[row][0], points)
	}
}

func quis4() [7][7]int {
	var matrix [7][7]int
	index := 17
	index0 := 3
	index1 := index + index0
	index2 := index1 + (len(matrix) * index0)
	for row := 0; row < len(matrix); row++ {
		for col := 0; col < len(matrix); col++ {
			if row+col == len(matrix)-1 && row != col {
				matrix[row][col] = index
				index -= index0
			} else if row == len(matrix)/2 && col == len(matrix)/2 {
				matrix[row][col] = index0
			} else if row == 1 && col == len(matrix)-1 {
				matrix[row][col] = index0 * 3
			} else if row > 1 && col == len(matrix)-1 && row != len(matrix)-2 {
				matrix[row][col] = index1
				index1 += index0
			} else if row == len(matrix)-2 && col == len(matrix)-1 {
				matrix[row][col] = index0 * 9
			} else if row > 0 && row == len(matrix)-1 && col != len(matrix)-4 {
				matrix[row][col] = index2
				index2 -= index0
			} else if col == len(matrix)-4 && row == len(matrix)-1 {
				matrix[row][col] = index0 * 27
			}
			fmt.Printf(" %d    ", matrix[row][col])
		}
		fmt.Println()
	}
	return matrix
}

func quis5() [7][7]int {
	var matrix [7][7]int
	index := 3
	index1 := 2
	for row := 0; row < len(matrix); row++ {
		for col := 0; col < len(matrix); col++ {
			if col == 0 && row == 1 || row == 1 && col == len(matrix)-1 || col == 0 && row == len(matrix)-2 ||
				row == len(matrix)-2 && col == len(matrix)-1 || row == 0 && col == len(matrix)/2 ||
				row == len(matrix)-1 && col == len(matrix)/2 {
				matrix[row][col] = index
			} else if row == 0 {
				matrix[row][0] = index1
				matrix[0][col] = index1
				index1 += index
			} else if row > 1 && col == len(matrix)-1 {
				matrix[row][col] += index1
				index1 += index
			}
		}
	}

	for row1 := len(matrix) - 1; row1 >= 0; row1-- {
		for col1 := len(matrix) - 1; col1 >= 0; col1-- {
			if row1 == len(matrix)-1 && col1 < len(matrix)-1 && col1 != len(matrix)/2 {
				matrix[row1][col1] = index1
				index1 += index
			} else if row1 > 1 && col1 == 0 && row1 != len(matrix)-2 {
				matrix[row1][col1] = index1
				index1 += index
			}
		}
	}

	return matrix
}

func quis6() [7][7]int {
	var matrix [7][7]int
	for row := 0; row < len(matrix); row++ {
		for col := 0; col < len(matrix); col++ {
			if row == len(matrix)-1 || col == len(matrix)-1 {
				matrix[row][len(matrix)-1] = row + 6
				matrix[len(matrix)-1][col] = col + 6
			} else if row == 0 || col == 0 {
				matrix[row][0] = row
				matrix[0][col] = col
			}
			fmt.Printf("%d  ", matrix[row][col])
		}
		fmt.Println()
	}
	return matrix
}

func quis7(n int) {

	matrix := make([][]int, n+1)
	for row := 0; row < n+1; row++ {
		matrix[row] = make([]int, n+1)
	}

	for row := 0; row < n; row++ {
		for col := 0; col < n; col++ {
			matrix[row][col] = row + col
		}
	}

	var sumRow int
	for row := 0; row < n; row++ {
		for col := 0; col < n; col++ {
			fmt.Print(matrix[row][col], ",")
			sumRow += matrix[row][col]
		}
		fmt.Print(sumRow)
		sumRow = 0
		fmt.Println()
	}

	sumKol := 0
	sumDia := 0
	for row := 0; row < n; row++ {
		for col := 0; col < n; col++ {

			sumKol += matrix[row][col]
			if row == col {
				sumDia += matrix[row][col]
			}
		}
		fmt.Print("[", sumKol, "]")
		sumKol = 0
	}
	fmt.Print("[", sumDia, "]")
}

func printQuis(matrix [7][7]int) {
	for row := 0; row < len(matrix); row++ {
		for col := 0; col < len(matrix); col++ {
			fmt.Printf("%d  ", matrix[row][col])
		}
		fmt.Println()
	}
}

func main() {
	//quis1()
	// quis2()
	//quis3()
	//quis4()
	printQuis(quis5())
	//quis6()
	//quis7(7)
}
