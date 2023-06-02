package main

import "fmt"

func main() {
	initRow()
	initCol()
	lowerTriangel()
	// printMatrix(lowerTriangel())
	// printMatrix(initArray())
	printMatrix(initDiagonal())
}

func initRow() {
	matrix := [1][5]int{
		{1, 2, 3, 4, 5},
	}
	fmt.Println(matrix)
}

//bikin 5 baris 1 kolom
func initCol() {
	matrix := [5][1]int{
		{1},
		{3},
		{4},
		{7},
		{9},
	}
	fmt.Println(matrix)
}

func lowerTriangel() [3][3]int {
	matrix := [3][3]int{
		{1, 0, 0},
		{1, 1, 0},
		{2, 1, 1},
	}
	return matrix
}

func printMatrix(matrix [5][5]int) {
	for row := 0; row < len(matrix); row++ {
		for col := 0; col < len(matrix[row]); col++ {
			fmt.Printf("%d ", matrix[row][col])
		}
		fmt.Println()
	}
}

func initArray() [3][3]int {
	//default zero value
	var matrix [3][3]int

	for row := 0; row < len(matrix); row++ {
		for col := 0; col < len(matrix[row]); col++ {
			matrix[row][col] = row + col
		}
	}
	return matrix
}

func initDiagonal() [5][5]int {
	var matrix [5][5]int
	for row := 0; row < len(matrix); row++ {
		for col := 0; col < len(matrix[row]); col++ {
			if row == col {
				matrix[row][col] = 10
			} else if col > row {
				matrix[row][col] = 20
			} else {
				matrix[row][col] = 21
			}
		}
	}
	return matrix
}
