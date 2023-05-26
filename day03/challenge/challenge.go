package main

import (
	"fmt"
)

// soal no 1
func rightDiagonal(count int) {
	//cetak baris ke bawah
	for row := 0; row < count; row++ {
		//cetak kolom
		for col := 0; col < count; col++ {
			if row == count-1-col {
				fmt.Printf("*")
			} else {
				fmt.Print(" ")
			}
		}
		fmt.Println()
	}
}

// soal no 2
func rightDiagonal1(count int) {
	//cetak baris ke bawah
	for row := 0; row < count; row++ {
		//cetak kolom
		for col := count - 1; 0 <= col; col-- {
			if row == col {
				fmt.Print("* ")
			} else {
				if row <= col {
					fmt.Print("  ")
				} else {
					fmt.Print("* ")
				}
			}
		}
		fmt.Println()
	}
}

// soal no 3
func rightDiagonal2(count int) {
	//cetak baris ke bawah
	for row := 0; row < count; row++ {
		//cetak kolom
		for col := count - 1; 0 <= col; col-- {
			if row == col {
				fmt.Print("10 ")
			} else {
				if row <= col {
					fmt.Print("   ")
				} else {
					fmt.Print("20 ")
				}
			}
		}
		fmt.Println()
	}
}

// soal no 4
func rightDiagonal3(count int) {
	//cetak baris ke bawah
	for row := 0; row < count; row++ {
		//cetak kolom
		for col := count - 1; 0 <= col; col-- {
			if row == col {
				fmt.Print("10 ")
			} else {
				if row >= col {
					fmt.Print("   ")
				} else {
					fmt.Print("21 ")
				}
			}
		}
		fmt.Println()
	}
}

// soal no 5
func rightDiagonal4(count int) {
	//cetak baris ke bawah
	for row := 0; row < count; row++ {
		//cetak kolom
		for col := count - 1; 0 <= col; col-- {
			if col == row {
				fmt.Print("10 ")
			} else {
				if row <= col {
					fmt.Print("21 ")
				} else {
					fmt.Print("20 ")
				}
			}
		}
		fmt.Println()
	}
}

// quiz no 1
func quiz1(count int) {
	for row := 1; row <= count; row++ {
		if row%2 == 0 {
			for col := 1; col <= count; col++ {
				fmt.Print(col)
			}
		} else {
			for col := count; col >= 1; col-- {
				fmt.Print(col)
			}
		}
		fmt.Println()
	}
}

// quiz no 2
func quiz2(nilai int) {
	for row := 1; row <= nilai; row++ {
		for ulang := 0; ulang < nilai/2; ulang++ {
			for col := 1; col <= nilai; col++ {
				if col == nilai {
					fmt.Print(row)
				} else {
					fmt.Print("")
				}
			}

			for col := nilai; col >= 1; col-- {
				if col == 1 {
					fmt.Print(nilai - row + 1)
				} else {
					fmt.Print("")
				}
			}
		}
		if nilai%2 == 1 {
			for col := 1; col <= nilai; col++ {
				if col == nilai {
					fmt.Print(row)
				} else {
					fmt.Print("")
				}
			}
		}
		fmt.Println()

	}
}

// quiz no 3
func quiz3(input int) {
	for row := 0; row < input; row++ {
		if row%2 == 0 {
			for col := 1; col <= input; col++ {
				if col%2 == 0 {
					print(col)
				} else {
					print("-")
				}
			}
		} else {
			for col := 1; col <= input; col++ {
				if col%2 == 0 {
					print("-")
				} else {
					print(col)
				}
			}
		}
		println()
	}
}

func main() {
	// rightDiagonal(5)
	// rightDiagonal1(5)
	// rightDiagonal2(5)
	// rightDiagonal3(5)
	// rightDiagonal4(5)
	// quiz1(5)
	// quiz2(5)
	quiz3(5)

}
