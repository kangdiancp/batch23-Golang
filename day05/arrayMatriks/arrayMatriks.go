package main

import "fmt"

func initBaris() {
	matriks := [1][5]int{
		{1, 2, 3, 4, 5},
	}

	fmt.Println(matriks)

}

func initKolom() {
	matriks := [5][1]int{
		{1},
		{2},
		{3},
		{4},
		{5},
	}

	fmt.Println(matriks)

}

func lowerTriangle() [3][4]int {
	matriks := [3][4]int{
		{1, 0, 0, 0},
		{1, 1, 0, 0},
		{1, 1, 1, 0},
	}
	return matriks

}

//print matriks
func printMatriks(matriks [5][5]int) {
	for i := 0; i < len(matriks); i++ {
		for j := 0; j < len(matriks[j]); j++ {
			fmt.Printf("%d ", matriks[i][j])
		}
		fmt.Println()
	}
}

//isi value matriks

func initArrayValue() [5][5]int {
	var matriks [5][5]int
	for i := 0; i < len(matriks); i++ {
		for j := 0; j < len(matriks); j++ {
			matriks[i][j] = i + j
		}
	}
	return matriks

}

func initArrayDiagonalValue() [5][5]int {
	var matriks [5][5]int
	for i := 0; i < len(matriks); i++ {
		for j := 0; j < len(matriks); j++ {
			//matriks[i][j] = i + j
			if i == j {
				matriks[i][j] = 10
			} else {
				if i < j {
					matriks[i][j] = 20
				} else {
					matriks[i][j] = 21

				}
			}
		}
	}
	return matriks

}

func main() {
	//initBaris()
	//initKolom()
	//lowerTriangle()
	//printMatriks(lowerTriangle())
	//printMatriks(initArrayValue())
	printMatriks(initArrayDiagonalValue())

}
