package main

import (
	"fmt"
)

func main() {
	// basic()
	// kotak(5)
	// basic2(5)
	// emptyBox(5)
	// bintangSegitigaSikufromleft(5)
	bintangSegitigaSikufromright(5)
}

func basic() {
	// for col := 0; col < 5; col++ {
	// 	for row := col; row > 1; row-- {
	// 		print("*")
	// 	}
	// }

	for col := 0; col < 5; col++ {
		print("*")
	}
	for col := 0; col < 5; col++ {
		println("*")
	}
}

func kotak(count int) {
	for row := 0; row < count; row++ {
		for col := 0; col < count; col++ {
			print("* ")
		}
		println()
	}
}

func basic2(new int) {
	for row := 0; row < new; row++ {
		for col := 0; col < new; col++ {
			fmt.Printf("[%d - %d]", row, col)
		}
		println()
	}

	for row := 0; row < new; row++ {
		for col := 0; col < new; col++ {
			if row == col {
				fmt.Printf("[%d - %d]", row, col)
			} else {
				fmt.Print("       ")
			}
		}
		println()
	}
}
func emptyBox(nilai int) {
	for row := 0; row < nilai; row++ {
		for col := 0; col < nilai; col++ {
			// fmt.Printf("[%d - %d]", row, col)

			if row == 0 || row == nilai-1 || col == 0 || col == nilai-1 {
				fmt.Printf("[%d - %d]", row, col)
			} else {
				fmt.Print("       ")
			}
		}
		println()
	}
}
func bintangSegitigaSikufromleft(nilai int) {
	for row := 0; row < nilai; row++ {
		for col := 0; col <= row; col++ {
			fmt.Printf("[%d - %d]", row, col)
		}
		println()
	}

}
func bintangSegitigaSikufromright(nilai int) {
	for row := nilai - 1; row > 0; row-- {
		for col := 0; col < row; col++ {
			fmt.Printf("[%d - %d]", row, col)
		}
		println()
	}

}
