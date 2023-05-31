package main

import "fmt"

func basic() {
	for col := 0; col < 5; col++ {
		fmt.Print("* ")
	}
	for row := 0; row < 5; row++ {
		fmt.Println("* ")
	}
}

func kotak(count int) {
	//cetak baris ke bawah
	for row := 0; row < count; row++ {
		//cetak kolom
		for col := 0; col < count; col++ {
			fmt.Print("* ")
		}
		fmt.Println()
	}
}

func kotakIndex(count int) {
	//cetak baris ke bawah
	for row := 0; row < count; row++ {
		//cetak kolom
		for col := 0; col < count; col++ {
			fmt.Printf("[%d-%d]", row, col)
		}
		fmt.Println()
	}

	//cetak baris ke bawah
	// 	for row := 0; row < count; row++ {
	//cetak kolom
	// 		for col := 0; col < count; col++ {
	// 			fmt.Printf("[%d-%d]", row, col)
	// 		}
	// 		fmt.Println()
	// 	}
}

func leftDiagonal(count int) {
	for row := 0; row < count; row++ {
		for col := 0; col < count; col++ {
			if row == col {
				fmt.Printf("*[%d-%d]\n", row, col)
			} else {
				fmt.Print("	")
			}
		}
		fmt.Println()
	}
}

func middleEmptyBox(count int) {
	for row := 0; row < count; row++ {
		for col := 0; col < count; col++ {
			if col == 0 || col == count-1 || row == 0 || row == count-1 {
				fmt.Printf("*[%d-%d] ", row, col)
			} else {
				fmt.Print("       ")
			}
		}
		fmt.Println()
	}
}

func leftFullDiagonal(count int) {

	for row := 0; row < count; row++ {
		for col := 0; col <= row; col++ {
			fmt.Printf("*[%d-%d] ", row, col)
		}
		fmt.Println()
	}
}

func printTriangle(rows int) {
	for i := 1; i <= rows; i++ {
		for j := 1; j <= i; j++ {
			fmt.Print("D ")
		}
		fmt.Println()
	}
}

func main() {
	// basic()
	// kotak(5)
	// kotakIndex(5)
	// leftDiagonal(5)
	// middleEmptyBox(5)
	// leftFullDiagonal(5)
	printTriangle(5)
}
