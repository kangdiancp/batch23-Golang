package main

import "fmt"

func kotak(count int) {
	//cetak baris ke bawah
	for baris := 0; baris < count; baris++ {

		//cetak baris ke samping
		for colom := 0; colom < count; colom++ {
			fmt.Print("* ")
		}
		fmt.Println()

	}
}

func leftDiagonal(count int) {

	for baris := 0; baris < count; baris++ {

		//cetak baris ke samping
		for kolom := 0; kolom < count; kolom++ {
			if baris == kolom {
				fmt.Printf("* (%d - %d)", baris, kolom)
			} else {
				fmt.Print("	")
			}

		}
		fmt.Println()

	}
}

func kotakTengah(count int) {

	for baris := 0; baris < count; baris++ {

		for kolom := 0; kolom < count; kolom++ {
			if baris == 0 || baris == count-1 || kolom == 0 || kolom == count-1 {

				fmt.Printf("	(%d - %d)", baris, kolom)

			} else {
				fmt.Print("	")
			}

		}
		fmt.Println()

	}
}

func segitigaSamping(count int) {

	for baris := 0; baris < count; baris++ {

		for kolom := 0; kolom < count; kolom++ {
			if baris >= count-1 || kolom <= baris {

				fmt.Printf("	(%d - %d)", baris, kolom)

			} else {
				fmt.Print("	")
			}

		}
		fmt.Println()

	}
}

func rightDiagonal(count int) {

	for baris := 0; baris < count; baris++ {

		//cetak baris ke samping
		for kolom := 0; kolom < count; kolom++ {
			if baris == kolom {
				fmt.Printf("* (%d - %d)", baris, kolom)
			} else {
				fmt.Print("	")
			}

		}
		fmt.Println()

	}
}

func main() {
	//kotak(5)
	//LeftDiagonal(5)
	//kotakTengah(10)
	//segitigaSamping(5)
	rightDiagonal(5)
}
