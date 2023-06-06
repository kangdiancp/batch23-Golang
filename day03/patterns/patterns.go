package main

import (
	"fmt"
)

func basic() {
	for col := 0; col < 5; col++ {
		fmt.Print("* ")
	}
	for row := 0; row < 5; row++ {
		fmt.Println("*")
	}
}

func kotak(count int) {
	//cetak baris ke bawah
	for row := 0; row < count; row++ {
		for col := 0; col < count; col++ {
			fmt.Printf("* ")
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

func kotakDiagonal(count int) {
	for row := 0; row < count; row++ {
		for col := 0; col < count; col++ {
			if row == col {
				fmt.Printf("  * [%d - %d]  ", row, col)
			} else {
				fmt.Print("  ")
			}
		}
		fmt.Println()
	}
}

func kotakMiddleBlank(count int) {
	for row := 0; row < count; row++ {
		for col := 0; col < count; col++ {
			if col == 0 || col == count-1 || row == 0 || row == count-1 {
				fmt.Printf("* [%d - %d]  ", row, col)
			} else {
				fmt.Print("           ")
			}
		}
		fmt.Println()
	}
}

func segitiga(count int) {
	for row := 0; row < count; row++ {
		for col := 0; col < count; col++ {
			if row >= col {
				fmt.Printf("* [%d - %d]  ", row, col)
			} else {
				fmt.Print("           ")
			}
		}
		fmt.Println()
	}
}

func segitigastar(count int) {
	for row := 0; row < count; row++ {
		for col := 0; col < count; col++ {
			if row <= col {
				fmt.Print("* ", row, col)
			} else {
				fmt.Print("           ")
			}
		}
		fmt.Println()
	}
}

//cara2 (validasi di looping)
/*func segitiga(count int) {
	for row := 0; row < count; row++ {
		for col := 0; col <= row; col++ {
				fmt.Print("           ")
			}
		}
		fmt.Println()
	}
}*/

func main() {
	//basic()
	//kotak(5)
	//kotakIndex(5)
	//kotakDiagonal(5)
	//kotakMiddleBlank(5)
	//segitiga(5)
	//segitigastar(5)
}
