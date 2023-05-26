package main

import "fmt"

// Nomor 1 :
func rightDiagonal(count int) {
	for row := 0; row < count; row++ {
		for col := 0; col < count; col++ {
			if col == count-1-row {
				fmt.Print("*")
			} else {
				fmt.Print("		")
			}
		}
		fmt.Println()
	}
}

// Nomor 2 :
func rightFullDiagonal(count int) {
	for row := 0; row < count; row++ {
		for col := 0; col < count; col++ {
			if col >= count-row-1 {
				fmt.Print("*")
			} else {
				fmt.Print(" ")
			}
		}
		fmt.Println()
	}
}

// Nomor 3 :
func rightFullNumber(count int) {
	for row := 0; row < count; row++ {
		for col := 0; col < count; col++ {
			if col == count-1-row {
				fmt.Print("10 ")
			} else if col > count-1-row {
				fmt.Print("20 ")
			} else {
				fmt.Print(" ")
			}
		}
		fmt.Println()
	}
}

// Nomor 4 :
func leftNumberDiagonal(count int) {
	for row := 0; row < count; row++ {
		for col := 0; col < count; col++ {
			if col == count-1-row {
				fmt.Print("10 ")
			} else if col > count-1-row {
				fmt.Print(" ")
			} else {
				fmt.Print("21 ")
			}
		}
		fmt.Println()
	}
}

// Nomor 5 :
func numberBox(count int) {
	for row := 0; row < count; row++ {
		for col := 0; col < count; col++ {
			if col+row == count-1 {
				fmt.Printf("10 ")
			} else if col+row > count-1 {
				fmt.Printf("20 ")
			} else {
				fmt.Printf("21 ")
			}
		}
		fmt.Println()
	}
}

// Nomor 6 : Program untuk mencetak pola angka terbalik
func pola_1(count int) {
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

// Nomor 7 : Program untuk mencetak pola angka urut berurutan
func pola_2(count int) {
	for row := 1; row <= count; row++ {

		for loop := 0; loop < count/2; loop++ {

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

// Nomor 8 : Program untuk mencetak pola angka negatif dan positif bergantian
func pola_3(count int) {
	for row := 0; row < count; row++ {
		for col := 1; col <= count; col++ {
			if row%2 == 0 {
				if col%2 == 0 {
					fmt.Print(col)
				} else {
					fmt.Print("-")
				}
			} else {
				if col%2 == 0 {
					fmt.Print("-")
				} else {
					fmt.Print(col)
				}
			}
		}
		fmt.Println()
	}
}

func main() {
	rightDiagonal(5)
	rightFullDiagonal(5)
	rightFullNumber(5)
	leftNumberDiagonal(5)
	numberBox(5)
	pola_1(5)
	pola_2(5)
	pola_3(5)
}
