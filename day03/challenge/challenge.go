package main

import "fmt"

func diagonal(count int) {
	//cetak baris ke bawah
	for row := 0; row < count; row++ {
		//cetak kolom
		for col := 0; col < count; col++ {
			if row+col == count-1 {
				fmt.Print("* ")
			} else {
				fmt.Print("0 ")
			}

		}
		fmt.Println()
	}
}

func segitiga(count int) {
	//cetak baris ke bawah
	for row := 0; row < count; row++ {
		//cetak kolom
		for col := 0; col < count; col++ {
			if col+row >= count-1 {
				fmt.Print("* ")
			} else {
				fmt.Print("  ")
			}
		}
		fmt.Println()
	}
}

func segitigaAngka(count int) {
	//cetak baris ke bawah
	for row := 0; row < count; row++ {
		//cetak kolom
		for col := 0; col < count; col++ {
			if col+row >= count-1 {
				if col+row == count-1 {
					fmt.Print("10 ")
				} else {
					fmt.Print("20 ")
				}
			} else {
				fmt.Print("   ")
			}
		}
		fmt.Println()
	}
}

func segitigaAngka_(count int) {
	//cetak baris ke bawah
	for row := 0; row < count; row++ {
		//cetak kolom
		for col := 0; col < count; col++ {
			if col+row == count-1 {
				fmt.Printf("10 ")
			} else if col+row >= count-1 {
				fmt.Print("  ")
			} else {
				fmt.Print("21 ")
			}
		}
		fmt.Println()
	}
}

func segitigaPenuh(count int) {
	//cetak baris ke bawah
	for row := 0; row < count; row++ {
		//cetak kolom
		for col := 0; col < count; col++ {
			if col+row == count-1 {
				fmt.Printf("10 ")
			} else if col+row >= count-1 {
				fmt.Print("20 ")
			} else {
				fmt.Print("21 ")
			}
		}
		fmt.Println()
	}
}

func nomor6(input int) {
	for a := 0; a < input; a++ {
		if a%2 == 0 {
			for i := input; i >= 1; i-- {
				fmt.Print(i)
			}
		} else {
			for j := 1; j <= input; j++ {
				fmt.Print(j)
			}
		}
		println()
	}
}

func nomor7(nilai int) {
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

func nomor8(nilai int) {
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
	diagonal(5)
	segitiga(5)
	segitigaAngka(5)
	segitigaAngka_(5)
	segitigaPenuh(5)
	nomor6(9)
	nomor7(5)
	nomor8(5)
}
