package main

import "fmt"

func basic() {
	for col := 0; col < 5; col++ {
		fmt.Print("* ")
	}
	for row := 0; row < 5; row++ {
		fmt.Println("*")
	}
}

func kotak(count int) {
	// catak baris ke bawah
	for row := 0; row < count; row++ {
		//cetak kolom
		for col := 0; col < count; col++ {
			fmt.Print("* ")
		}
		fmt.Println()
	}
}

func kotakIndex(count int) {
	// catak baris ke bawah
	for row := 0; row < count; row++ {
		//cetak kolom
		for col := 0; col < count; col++ {
			fmt.Printf("[%d-%d]", col, row)
		}
		fmt.Println()
	}
}

func kotakIndexDia(count int) {
	// catak baris ke bawah
	for row := 0; row < count; row++ {
		//cetak kolom
		for col := 0; col < count; col++ {
			if row == col {
				fmt.Printf("*")
			} else {
				fmt.Print("-")
			}
		}
		fmt.Println()
	}
}

func kotakIndexMid(count int) {
	// catak baris ke bawah
	for row := 0; row < count; row++ {
		//cetak kolom
		for col := 0; col < count; col++ {
			if col == 0 || col == count-1 || row == count-1 || row == 0 {
				fmt.Printf("* [%d-%d]", col, row)
			} else {
				fmt.Print("       ")
			}
		}
		fmt.Println()
	}
}

func kotakIndexLeft(count int) {
	// catak baris ke bawah
	for row := 0; row < count; row++ {
		//cetak kolom
		for col := 0; col < count; col++ {
			if row <= col {
				fmt.Printf("*")
			}
		}
		fmt.Println()
	}
}

func kotakIndexRight(count int) {
	// catak baris ke bawah
	for row := 0; row < count; row++ {
		//cetak kolom
		for col := 0; col < count; col++ {
			if row <= col {
				fmt.Printf("* [%d-%d]", col, row)
			} else {
				fmt.Print("-")
			}
		}
		fmt.Println()
	}
}

func main() {
	basic()
	kotak(5)
	kotakIndex(5)
	kotakIndexMid(5)
	kotakIndexDia(5)
	kotakIndexLeft(5)
	kotakIndexRight(5)
}
