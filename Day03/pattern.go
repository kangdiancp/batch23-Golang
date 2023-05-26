package main

import "fmt"

func basic() {
	for i := 0; i < 5; i++ {
		fmt.Print(" * ")
	}

	for row := 0; row < 5; row++ {
		fmt.Println("* ")
	}
}

func kotakIndex(count int) {
	//Cetak Garis Kebawah
	for row := 0; row < count; row++ {
		// Cetak Kolom
		for col := 0; col < count; col++ {
			fmt.Printf("[%d %d]", row, col)
		}
		fmt.Println()
	}
}

func leftDiagonal(count int) {
	// Baris
	for row := 0; row < count; row++ {
		// Kolom
		for col := 0; col < count; col++ {
			if row == col {
				fmt.Print("* ")
			} else {
				fmt.Print("_ ")
			}
		}
		fmt.Println()
	}
}

func emptyBox(count int) {
	// Baris
	for row := 0; row < count; row++ {
		for col := 0; col < count; col++ {
			if row == 0 || row == count-1 || col == 0 || col == count-1 {
				fmt.Printf("[%d %d]", row, col)
			} else {
				fmt.Print("     ")
			}
		}
		fmt.Println()
	}
}

func triangleT(count int) {
	// Baris
	for row := 0; row < count; row++ {
		for col := 0; col <= row; col++ {
			fmt.Printf("[%d %d]", row, col)
		}
		fmt.Println()
	}
}

func main() {
	// kotakIndex(5)
	leftDiagonal(5)
	// emptyBox(5)
	// triangleT(5)
}
