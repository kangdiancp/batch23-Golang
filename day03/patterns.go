package main

import (
	"fmt"
)

//print : cetak ke kanan/baris
//printl : cetak ke bawah/kolom

func basic() {
	for col := 0; col < 5; col++ {
		fmt.Print("*")
	}
	for row := 0; row < 5; row++ {
		fmt.Println("*")
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
}

func kotakIndex1(count int) {
	//cetak baris ke bawah
	for row := 0; row < count; row++ {
		//cetak kolom
		for col := 0; col < count; col++ {
			fmt.Printf("* [%d-%d]", row, col)
		}
		fmt.Println()
	}
}

func leftDiagonal(count int) {
	//cetak baris ke bawah
	for row := 0; row < count; row++ {
		//cetak kolom
		for col := 0; col < count; col++ {
			if row == col {
				fmt.Printf("* [%d-%d]", row, col)
			} else {
				fmt.Print("  ")
			}
		}
		fmt.Println()
	}
}

func emptyBox(count int) {
	//cetak baris ke bawah
	for row := 0; row < count; row++ {
		//cetak kolom
		for col := 0; col < count; col++ {
			if col == 0 || col == 4 || row == 4 || row == 0 {
				fmt.Printf("* [%d-%d]", row, col)
			} else {
				fmt.Print("       ")
			}
		}
		fmt.Println()
	}
}

func leftFullDiagonal(count int) {
	//cetak baris ke bawah
	for row := 0; row < count; row++ {
		//cetak kolom
		for col := 0; col < count; col++ {
			if row >= col {
				fmt.Printf("* [%d-%d]", row, col)
			} else {
				fmt.Print("  ")
			}
		}
		fmt.Println()
	}
}

func cobaPiramid(count int) { //masih gagal
	//cetak baris ke bawah
	for col := 0; col < count; col++ {
		//cetak kolom
		for row := 0; row < count; row++ {
			if col > row {
				fmt.Printf("* [%d-%d]", col, row)
			} else {
				fmt.Print("  ")
			}
		}
		fmt.Println()
	}
}

func main() {
	// cobaPiramid(5)
	leftFullDiagonal(5)
	// emptyBox(5)
	// leftDiagonal(5)
	// kotakIndex1(5)
	// kotakIndex(5)
	// kotak(5)
	// basic()
}
