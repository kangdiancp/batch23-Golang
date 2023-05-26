package main

import (
	"fmt"
)

// nomor1
func rightDiagonal(count int) {

	for baris := 0; baris < count; baris++ {

		//cetak baris ke samping
		for kolom := count - 1; 0 <= kolom; kolom-- {
			if baris == kolom {
				fmt.Printf("*")
			} else {
				fmt.Print("	")
			}

		}
		fmt.Println()

	}
}

//nomor 2

func segitigaSampingKanan(count int) {
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

// nomor3
func segitigaSampingKananValue(count int) {
	for baris := 0; baris < count; baris++ {

		//cetak baris ke samping
		for kolom := count - 1; 0 <= kolom; kolom-- {
			if baris == kolom {
				fmt.Printf("10 ")
			} else {
				if kolom >= baris {
					fmt.Printf("  ")
				} else {
					fmt.Printf("20")

				}
				fmt.Printf(" ")

			}

		}
		fmt.Println()

	}
}

// nomor4
func segitigaKananAtas(count int) {
	for baris := 0; baris < count; baris++ {

		//cetak baris ke samping
		for kolom := count - 1; 0 <= kolom; kolom-- {
			if baris == kolom {
				fmt.Printf("10")
			} else {
				if kolom >= baris {
					fmt.Printf("21")
				}
				fmt.Print(" ")

			}

		}
		fmt.Println()

	}
}

// nomor5
func kotakNilai(count int) {
	for baris := 0; baris < count; baris++ {

		//cetak baris ke samping
		for kolom := count - 1; 0 <= kolom; kolom-- {
			if baris == kolom {
				fmt.Printf("10 ")
			} else {
				if kolom >= baris {
					fmt.Printf("21")
				} else {
					fmt.Printf("20")

				}
				fmt.Printf(" ")

			}

		}
		fmt.Println()

	}
}

//nomor6

func angkaBarisKolom() {
	fmt.Println("Masukkan Angka: ")
	var angka int
	fmt.Scan(&angka)
	for row := 1; row <= angka; row++ {
		if row%2 == 1 {
			for col := angka; col >= 1; col-- {
				fmt.Print(col)
			}
		} else {
			for col := 1; col <= angka; col++ {
				fmt.Print(col)
			}
		}
		fmt.Println()
	}
}

// nomor 7
func angkaTengah() {
	fmt.Println("Masukkan Angka: ")
	var angka int
	fmt.Scan(&angka)

	for row := 1; row <= angka; row++ {

		for ulang := 0; ulang < angka/2; ulang++ {

			for col := 1; col <= angka; col++ {
				if col == angka {
					fmt.Print(row)
				} else {
					fmt.Print("")
				}
			}

			for col := angka; col >= 1; col-- {
				if col == 1 {
					fmt.Print(angka - row + 1)
				} else {
					fmt.Print("")
				}
			}
			//fmt.Println()

		}

		if angka%2 == 1 {
			for col := 1; col <= angka; col++ {
				if col == angka {
					fmt.Print(row)
				} else {
					fmt.Print("")
				}
			}
		}

		fmt.Println()

	}

}

// nomor8
func minusNumber() {
	fmt.Println("Masukkan Angka: ")
	var angka int
	fmt.Scan(&angka)
	for row := 0; row < angka; row++ {
		if row%2 == 0 {
			for col := 1; col <= angka; col++ {
				if col%2 == 0 {
					print(col)
				} else {
					print("-")
				}
			}
		} else {
			for col := 1; col <= angka; col++ {
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
	//rightDiagonal(5)
	//segitigaSampingKanan(5)
	//segitigaSampingKananValue(5)
	//segitigaKananAtas(5)
	//kotakNilai(5)
	//angkaBarisKolom()
	angkaTengah()
	//minusNumber()
}
