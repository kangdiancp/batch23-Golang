package main

import "fmt"

// No 1

func soalno1(count int) {
	//cetak baris kebawah
	for baris := 0; baris < count; baris++ {
		for kolom := count - 1; 0 <= kolom; kolom-- {
			if baris == kolom {
				fmt.Printf("*")
			} else {
				fmt.Print(" ")
			}
		}
		fmt.Println()
	}
}

//NO2
func soalno2(count int) {
	for baris := 0; baris < count; baris++ {
		for kolom := count - 1; 0 <= kolom; kolom-- {
			if baris == kolom {
				fmt.Printf("* ")
			} else {
				if kolom >= baris {
					fmt.Printf(" ")
				} else {
					fmt.Printf("*")
				}
				fmt.Printf(" ")
			}
		}
		fmt.Println()
	}
}

//NO 3
func soalno3() {
	for row := 0; row < 5; row++ {
		for col := 5 - 1; col >= 0; col-- {
			if row == col {
				fmt.Printf("10")
			} else if row <= col {
				fmt.Printf("-")

			} else {
				fmt.Printf("20")

			}
		}
		fmt.Println()
	}
}

//no 4

func soalno4() {
	for row := 0; row <= 5; row++ {
		for col := 5 - 1; col >= 0; col-- {
			if row == col {
				fmt.Print("10")
			} else if row <= col {
				fmt.Print("21")
			} else {
				fmt.Print(" ")
			}
		}
		fmt.Println()
	}
}

//NO 5

func soalnomor5(count int) {
	for row := 0; row < count; row++ {
		//cetak kolom
		for col := 0; col < count; col++ {
			if col+row > count-1 {
				fmt.Printf(" 20")
			} else if col+row == count-1 {
				fmt.Printf(" 10")
			} else {
				fmt.Printf(" 21")
			}
		}
		fmt.Println()
	}
}

// NO 6

func soalnomor6(nilai int) {
	for row := 0; row < nilai; row++ {
		if row%2 == 0 {
			for col := nilai; col > 0; col-- {
				print(col)
			}
		} else {
			for col := 1; col <= nilai; col++ {
				print(col)
			}
		}
		println()
	}
}

//No7
func soalnomor7(nilai int) {

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
			//fmt.Println()

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

func soalnomer8(nilai int) {
	for row := 0; row < nilai; row++ {
		if row%2 == 0 {
			for col := 1; col <= nilai; col++ {
				if col%2 == 0 {
					print(col)
				} else {
					print("-")
				}
			}
		} else {
			for col := 1; col <= nilai; col++ {
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
	//soalnomor1()
	//alnomor2()
	//soalnomor3()
	//soalnomor4()
	//soalnomor5()
	//soalnomor6()
	//soalnomor7(9)
	soalnomer8(8)

}
