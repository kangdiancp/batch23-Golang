package main

import "fmt"

//nomor 1
func rightDiagonal(count int) {
	// cetak baris ke bawah
	for row := 0; row < count; row++ {
		//cetak kolom
		for col := count - 1; 0 <= col; col-- {
			if col == row {
				fmt.Print("*")
			} else {
				fmt.Print(" ")
			}
		}
		fmt.Println()
	}
}

//nomor 2
func segiTigaKanan(count int) {
	// cetak baris ke bawah
	for row := 0; row < count; row++ {
		//cetak kolom
		for col := 0; col < count; col++ {
			if col < count-row-1 {
				fmt.Print(" ")
			} else {
				fmt.Print("*")
			}
		}
		fmt.Println()
	}
}

//nomor 3
func segiTigaKananAngka(count int) {
	// cetak baris ke bawah
	for row := 0; row < count; row++ {
		for col := 0; col < count-row-1; col++ {
			fmt.Print("  ")
		}
		for col := 0; col <= row; col++ {
			if col == 0 {
				fmt.Print(10)
			} else {
				fmt.Print(20)
			}
		}

		fmt.Println()
	}

}

//nomor 4
func segiTigaKiriAngka(count int) {
	// cetak baris ke bawah
	for row := count; row >= 1; row-- {
		//cetak kolom
		for col := 1; col <= row; col++ {
			if col < row {
				fmt.Print("21")
			} else {
				fmt.Println("10")
			}

		}
		fmt.Println()
	}
}

// nomor 5
func kotakAngka(count int) {
	// cetak baris ke bawah
	for row := 0; row < count; row++ {
		for col := 0; col < count-row-1; col++ {
			fmt.Print(21)
		}
		for col := 0; col <= row; col++ {
			if col == 0 {
				fmt.Print(10)
			} else {
				fmt.Print(20)
			}
		}

		fmt.Println()
	}

}

//no 1
func angkaReverse(count int) {
	// cetak baris ke bawah
	for row := 0; row < count; row++ {
		//cetak kolom

		for col := 0; col < count; col++ {
			if row%2 == 0 {
				fmt.Print(col + 1)
			} else {
				fmt.Print(count - col)
			}
		}
		fmt.Println()
	}
}

//no 2
func angkaReverseTerbalik(count int) {
	// cetak baris ke bawah
	for row := 0; row < count; row++ {
		//cetak kolom

		for col := 0; col < count; col++ {
			if col%2 == 0 {
				fmt.Print(row + 1)
			} else {
				fmt.Print(count - row)
			}
		}
		fmt.Println()
	}
}

func ganjilGenap(count int) {
	// cetak baris ke bawah
	for row := 0; row < count; row++ {
		//cetak kolom

		for col := 0; col < count; col++ {
			if col%2 == 0 && row%2 != 0 {
				fmt.Print(col + 1)
			} else if col%2 != 0 && row%2 == 0 {
				fmt.Print(col + 1)
			} else {
				fmt.Print("_")
			}
		}
		fmt.Println()
	}
}

func main() {
	// rightDiagonal(5)
	// segiTigaKanan(5)
	// segiTigaKananAngka(5)
	// segiTigaKiriAngka(5)
	// kotakAngka(5)

	// angkaReverse(5)
	// angkaReverseTerbalik(5)
	ganjilGenap(5)
}
