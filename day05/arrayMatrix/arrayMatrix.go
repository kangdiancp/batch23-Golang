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
		{2},
		{3},
		{4},
		{5},
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
	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix); j++ {
			fmt.Printf("%d ", matrix[i][j])
		}

		fmt.Println(" ")
	}
}

func initArray() [5][5]int {
	// default zero value
	var matrix [5][5]int

	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix); j++ {
			matrix[i][j] = i + j
		}
	}

	return matrix
}

func initDiagonalTriangle() [5][5]int {
	var matrix [5][5]int

	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix); j++ {
			if i == j {
				matrix[i][j] = 10
			} else {
				matrix[i][j] = 21
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

func main() {
	initRow()
	initCol()
	// printMatrix(lowerTriangle())
	printMatrix(initArray())
	printMatrix(initDiagonalTriangle())
}
