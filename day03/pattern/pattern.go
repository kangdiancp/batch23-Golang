package main

import "fmt"

func basic() {
	for i := 0; i < 5; i++ {
		fmt.Print("* ")
	}

	for i := 0; i < 5; i++ {
		fmt.Println("* ")
	}
}

func kotak(count int) {
	// Cetak baris ke bawah
	for i := 0; i < count; i++ {
		// Cetak Kolom
		for j := 0; j < count; j++ {
			fmt.Print("* ")
		}
		fmt.Println("* ")
	}
}

func kotakIndex(count int) {
	// Cetak baris ke bawah
	for i := 0; i < count; i++ {
		// Cetak Kolom
		for j := 0; j < count; j++ {
			fmt.Printf("[%d, %d] ", i, j)
		}
		fmt.Println(" ")
	}
}

func diagonal(count int) {
	// Cetak baris ke bawah
	for i := 0; i < count; i++ {
		// Cetak Kolom
		for j := 0; j < count; j++ {
			if i == j {
				fmt.Printf("[%d, %d] ", i, j)
			} else {
				fmt.Print("* ")
			}
		}
		fmt.Println(" ")
	}
}

func kotakBolong(count int) {
	// Cetak baris ke bawah
	for i := 0; i < count; i++ {
		// Cetak Kolom
		for j := 0; j < count; j++ {
			if i == 0 || i == count-1 || j == 0 || j == count-1 {
				fmt.Printf("[%d, %d] ", i, j)
			} else {
				fmt.Print("****** ")
			}
		}
		fmt.Println(" ")
	}
}

func segitigaSiku(count int) {
	// Cetak baris ke bawah
	for i := 0; i < count; i++ {
		// Cetak Kolom

		for j := 0; j <= i; j++ {
			fmt.Printf("*[%d, %d] ", i, j)
		}
		fmt.Println(" ")
	}
}

func main() {
	// basic()
	// kotak(5)
	// kotakIndex(5)
	// diagonal(5)
	// kotakBolong(5)
	segitigaSiku(5)
}
