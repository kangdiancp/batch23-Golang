package main

import (
	"fmt"
)

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
		{2, 1, 0},
		{2, 1, 1},
	}
	return matrix
}

func PrintMatrix(matrix [5][5]int) {
	for row := 0; row < len(matrix); row++ {
		for col := 0; col < len(matrix); col++ {
			fmt.Printf("%d ", matrix[row][col])
		}
		fmt.Println()
	}
}

func initArray() [3][3]int {
	//default zero value
	var matrix [3][3]int
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
			} else if col > row {
				matrix[row][col] = 20
			} else {
				matrix[row][col] = 21
			}
		}
	}
	return matrix
}

func main() {
	//initRow()
	//lowerTriangle()
	//PrintMatrix(lowerTriangle())
	// PrintMatrix(initArray())
	PrintMatrix(initDiagonal())

}
