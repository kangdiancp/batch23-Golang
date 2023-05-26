package main

import "fmt"

func basic() {
	for i := 0; i < 5; i++ {
		fmt.Print(" * ")
	}

	for row := 0; row < 5; row++ {
		fmt.Println("* ")
	}
}

func kotakIndex(count int) {
	// cetak baris kebawah
	for row := 0; row < count; row++ {

		for col := 0; col < count; col++ {
			fmt.Printf("[%d- %d]", row, col)
		}
		fmt.Println()
	}
}

func leftdiagonal(count int) {
	//cetak baris kebawah

	for i := 0; i < count; i++ {
		// cetak kolom
		for j := 0; j < count; j++ {
			if i == 0 || i == count-1 || j == 0 || j == count-1 {
				fmt.Printf("[%d,%d]", i, j)
			} else {
				fmt.Print(" *** ")
			}
		}
		fmt.Println(" ")
	}
}

func segitiga(count int) {
	for i := 0; i < count; i++ {
		// cetak kolom
		for j := 0; j <= i; j++ {
			fmt.Printf("[%d,%d]", i, j)

		}
		fmt.Println(" ")
	}
}

func main() {
	//basic()
	//kotakIndex(5)
	leftdiagonal(5)
	//segitiga(5)

}
