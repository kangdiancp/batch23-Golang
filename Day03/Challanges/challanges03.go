package main

import "fmt"

func soal1(count int) {
	for i := 0; i < count; i++ {
		for j := 0; j < count; j++ {
			if i+j == count-1 {
				fmt.Print("*")
			} else {
				fmt.Print("   ")
			}
		}
		fmt.Println()
	}
}

func soal2(count int) {
	for row := 0; row < count; row++ {
		for col := 0; col < count; col++ {
			if row == col {
				fmt.Print("* ")
			} else {
				fmt.Print("  ")
			}
		}
		println()
	}
}

// soal3
func soal3(count int) {
	for row := 0; row < count; row++ {
		for col := 0; col < count; col++ {
			if col == count-1-row {
				fmt.Print("10 ")
			} else if col > count-1-row {
				fmt.Print("20 ")
			} else {
				fmt.Print("   ")
			}
		}
		fmt.Println()
	}
}

// Soal 4
func soal4(count int) {
	for row := 0; row < count; row++ {
		for col := 0; col < count; col++ {
			if col == count-1-row {
				fmt.Print("10 ")
			} else if col > count-1-row {
				fmt.Print("	")
			} else {
				fmt.Print("21 ")
			}
		}
		fmt.Println()
	}
}

// Soal 5
func soal5(count int) {
	for row := 0; row < count; row++ {
		for col := 0; col < count; col++ {
			if col+row == count-1 {
				fmt.Printf("10")
			} else if col+row > count-1 {
				fmt.Printf("20")
			} else {
				fmt.Printf("21")
			}

		}
		fmt.Println()
	}
}

// Soal 6
func soal6(count int) {
	for i := 1; i <= count; i++ {
		if i%2 == 1 {
			// Cetak angka dari 987654321 hingga 123456789 untuk baris genap
			for j := count; j >= 1; j-- {
				fmt.Print(j)
			}
		} else {
			// Cetak angka dari 123456789 hingga 987654321 untuk baris ganjil
			for j := 1; j <= count; j++ {
				fmt.Print(j)
			}
		}
		fmt.Println()
	}
}

// Soal 7
func soal7(count int) {
	for row := 1; row <= count; row++ {
		for ulang := 0; ulang < count/2; ulang++ {

			for col := 1; col <= count; col++ {
				if col == count {
					fmt.Print(row)
				} else {
					fmt.Print("")
				}
			}

			for col := count; col >= 1; col-- {
				if col == 1 {
					fmt.Print(count - row + 1)
				} else {
					fmt.Print("")
				}
			}
			//fmt.Println()

		}

		if count%2 == 1 {
			for col := 1; col <= count; col++ {
				if col == count {
					fmt.Print(row)
				} else {
					fmt.Print("")
				}
			}
		}
		fmt.Println()
	}
}

// Soal 8
func soal8(count int) {
	for row := 0; row < count; row++ {
		if row%2 == 0 {
			for col := 1; col <= count; col++ {
				if col%2 == 0 {
					print(col)
				} else {
					print("-")
				}
			}
		} else {
			for col := 1; col <= count; col++ {
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
	// soal1(5)
	// soal2(5)
	// soal3(5)
	// soal4(5)
	soal5(5)
	// soal6(9)
	// soal7(5)
	// soal8(9)
}
