package main

import (
	"fmt"
)

func nomor1(count int) {
	for row := 0; row < count; row++ {
		for col := 0; col < count; col++ {
			if row == count-1-col {
				fmt.Print("*  ")
			} else {
				print("   ")
			}
		}
		fmt.Println()
	}
}

func nomor2(count int) {
	for row := 0; row < count; row++ {
		for col := count - 1; 0 <= col; col-- {
			if row == col {
				fmt.Print("*  ")
			} else {
				if row <= col {
					fmt.Print("   ")
				} else {
					fmt.Print("*  ")
				}
			}
		}
		fmt.Println()
	}
}

func nomor3(count int) {
	for row := 0; row < count; row++ {
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

func nomor4(count int) {
	for row := 0; row < count; row++ {
		for col := count - 1; 0 <= col; col-- {
			if row == col {
				fmt.Print("10 ")
			} else {
				if row <= col {
					fmt.Print("21 ")
				} else {
					fmt.Print(" ")
				}
			}
		}
		fmt.Println()
	}
}

func nomor5(count int) {
	for row := 0; row < count; row++ {
		for col := count - 1; 0 <= col; col-- {
			if row == col {
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

func quis1() {
	var input int
	fmt.Print("Masukkan Inputan Angka : ")
	fmt.Scan(&input)
	for row := 1; row <= input; row++ {
		if row%2 == 0 {
			for col := input; 1 <= col; col-- {
				fmt.Print(col)
			}
		} else {
			for col := 1; col <= input; col++ {
				fmt.Print(col)
			}

		}
		fmt.Println()
	}
}

func quis2() {
	var input int
	fmt.Print("Masukkan Inputan Angka :")
	fmt.Scan(&input)
	for row := 1; row <= input; row++ {
		for ulang := 0; ulang < input/2; ulang++ {
			for col := 1; col <= input; col++ {
				if col == input {
					fmt.Print(row)
				} else {
					fmt.Print("")
				}
			}
			for col := input; col >= 1; col-- {
				if col == 1 {
					fmt.Print(input - row + 1)
				} else {
					fmt.Print("")
				}
			}
		}
		if input%2 == 1 {
			for col := 1; col <= input; col++ {
				if col == input {
					fmt.Print(row)
				} else {
					fmt.Print("")
				}
			}
		}

		fmt.Println()

	}
}

func quis3() {
	var input int
	fmt.Print("Masukkan Inputan Angka : ")
	fmt.Scan(&input)
	for row := 0; row < input; row++ {
		if row%2 == 1 {
			for col := 1; col <= input; col++ {
				if col%2 == 1 {
					fmt.Print(col)
				} else {
					fmt.Print("_")
				}
			}
		} else {
			for col := 1; col <= input; col++ {
				if col%2 == 1 {
					fmt.Print("_")
				} else {
					fmt.Print(col)
				}
			}
		}
		fmt.Println()
	}

}
func main() {
	nomor1(5)
	nomor2(5)
	nomor3(5)
	nomor4(5)
	nomor5(5)
	quis1()
	quis2()
	quis3()
}
