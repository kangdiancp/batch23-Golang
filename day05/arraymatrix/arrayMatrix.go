package main

import "fmt"

func initRow() {
	matrix := [1][5]int{
		{1, 2, 3, 4, 5},
	}
	fmt.Println(matrix)
}

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

func lowerTriangle() [3][3]int {
	matrix := [3][3]int{
		{1, 0, 0},
		{1, 1, 0},
		{2, 1, 1},
	}
	return matrix
}

func printMatrix(matrix [5][5]int) {
	for row := 0; row < len(matrix); row++ {
		for col := 0; col < len(matrix); col++ {
			fmt.Printf("%d ", matrix[row][col])
		}
		fmt.Println()
	}
}

func initArray() [5][5]int {
	//default zero value
	var matrix [5][5]int

	for row := 0; row < len(matrix); row++ {
		for col := 0; col < len(matrix); col++ {
			matrix[row][col] = row + col
		}
	}
	return matrix
}

func initDiagonal() [5][5]int {
	var matrix [5][5]int
	for row := 0; row < len(matrix); row++ {
		for col := 0; col < len(matrix); col++ {
			if row == col {
				matrix[row][col] = 10
			} else if row > col {
				matrix[row][col] = 21
			} else {
				matrix[row][col] = 20
			}
		}
	}
	return matrix
}

func main() {
	// initRow()
	// initCol()
	// printMatrix(lowerTriangle())
	// printMatrix(initArray())
	printMatrix(initDiagonal())

}
