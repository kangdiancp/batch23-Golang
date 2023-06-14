package main

import "fmt"

func basic() {
	for col := 0; col < 5; col++ {
		fmt.Print("* ")
	}

	for row := 0; row < 5; row++ {
		fmt.Println("*")
	}
}

func kotak(count int) {
	// cetak baris ke bawah
	for row := 0; row < count; row++ {
		// cetak kolom
		for col := 0; col < count; col++ {
			fmt.Print("* ")
		}
		fmt.Println()
	}
}

func kotakIndex(count int) {
	// cetak baris ke bawah
	for row := 0; row < count; row++ {
		// cetak kolom
		for col := 0; col < count; col++ {
			fmt.Printf("[%d-%d]", row, col)
		}
		fmt.Println()
	}
}

func leftDiagonal(count int) {
	// cetak baris ke bawah
	for row := 0; row < count; row++ {
		// cetak kolom
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
	// cetak baris ke bawah
	for row := 0; row < count; row++ {
		// cetak kolom
		for col := 0; col < count; col++ {
			if row == 0 || row == count-1 {
				fmt.Print("* ")
			} else {
				if col == 0 || col == count-1 {
					fmt.Print("* ")
				} else {
					fmt.Print("_ ")
				}
			}
		}
		fmt.Println()
	}
}

func leftFullDiagonal(count int) {
	// cetak baris ke bawah
	for row := 0; row < count; row++ {
		// cetak kolom
		/* for col := 0; col < count; col++ {
			if col <= row {
				fmt.Print("* ")
			}
		} */

		for col := 0; col <= row; col++ {
			fmt.Print("* ")
		}
		fmt.Println()
	}
}

func main() {
	leftFullDiagonal(5)
	//emptyBox(5)
	//leftDiagonal(5)
	//kotakIndex(5)
	//kotak(5)
	//basic()
}
