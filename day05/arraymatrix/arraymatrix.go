package main

import "fmt"

// 1 Baris 5 Colom
func initRow() {
	metrix := [1][5]int{
		{1, 2, 3, 4, 5},
	}
	fmt.Println(metrix)
}

// 5 Baris 1 Kolom
func initCol(){
	metrix := [5][1]int{
		{1},
		{3},
		{4},
		{7},
		{9},
	}
	fmt.Println(metrix)
}

func lowerTriangle()[3][3]int{
	matrix := [3][3]int{
		{1, 0, 0},
		{1, 1, 0},
		{2, 1, 1},
	}
	return matrix
}

func printMatrix(matrix[5][5]int){
	for row := 0; row < len(matrix); row++ {
		for Colom := 0; Colom < len(matrix); Colom++ {
			fmt.Printf("%d ", matrix[row][Colom])
		}
		fmt.Println()
	}
}

func initArray()[5][5]int{
	// default zero value
	var matrix  [5][5]int
	
	for row := 0; row < len(matrix); row++ {
		for col := 0; col < len(matrix); col++ {
			matrix[row][col] = row + col
		}
	}
	return matrix
}

func initDiagonal()[5][5]int{
	var matrix [5][5] int

	for row := 0; row < len(matrix); row++ {
		for col := 0; col < len(matrix); col++ {
			if row == col{
				matrix[row][col] = 10
			}else if row > col{
				matrix[row][col] = 21
			}else{
				matrix[row][col] = 20
			}
		}
	}

	return matrix
}

func main() {
	//initRow()
	//initCol()
	// printMatrix(lowerTriangle())
	//printMatrix(initArray())
	//printMatrix(initDiagonal())
}