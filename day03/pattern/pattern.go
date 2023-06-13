package main

import "fmt"

func basic() {
	for col := 0; col < 5; col++ {
		fmt.Print("* ")
	}
	for row := 0; row < 5; row++ {
		fmt.Println("* ")
	}
	fmt.Println()
}

func kotak(count int) {
	//cetak baris ke bawah
	for row := 0; row < count; row++ {
		//cetak kolom
		for col := 0; col < count; col++ {
			fmt.Print("* ")
		}
		fmt.Println()
	}
}

func kotakIndex(count int) {
	//cetak baris ke bawah
	for row := 0; row < count; row++ {
		//cetak kolom
		for col := 0; col < count; col++ {
			fmt.Printf("[%d-%d] ", row, col)
		}
		fmt.Println()
	}
}

func diagonal(count int) {
	//cetak baris ke bawah
	for row := 0; row < count; row++ {
		//cetak kolom
		for col := 0; col < count; col++ {
			if row == col {
				fmt.Printf("[%d-%d] ", row, col)
			} else {
				fmt.Printf("     ")
			}

		}
		fmt.Println()
	}
}

func kotakKosong(count int) {
	//cetak baris ke bawah
	for row := 0; row < count; row++ {
		// cetak kolom
		for col := 0; col < count; col++ {
			if row == 0 || col == 0 || row == count-1 || col == count-1 {
				fmt.Printf("*[%d-%d] ", row, col)
			} else {
				fmt.Printf("       ")
			}
		}
		fmt.Println()
	}
}

func segitiga(count int) {
	// cetak baris ke bawah
	for row := 0; row < count; row++ {
		// cetak kolom
		for col := 0; col < count; col++ {
			if col > row {
				fmt.Printf("       ")
			} else {
				fmt.Printf("*[%d-%d] ", row, col)
			}
		}
		fmt.Println()
	}
}

func main() {
	basic()
	//kotak(5)
	//kotakIndex(5)
	//diagonal(5)
	//kotakKosong(5)
	//segitiga(5)
}
