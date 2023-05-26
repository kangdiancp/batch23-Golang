package main

import "fmt"

func basic() {
	for col := 0; col < 5; col++ {
		fmt.Print("* ")
	}
	for row := 0; row < 5; row++ {
		fmt.Println("* ")
	}
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
			fmt.Printf("[%d-%d]", row, col)
		}
		fmt.Println()
	}

	//cetak baris ke bawah
	// 	for row := 0; row < count; row++ {
	// 		//cetak kolom
	// 		for col := 0; col < count; col++ {
	// 			fmt.Printf("[%d-%d]", row, col)
	// 		}
	// 		fmt.Println()
	// 	}
}

func leftDiagonal(count int) {
	for row := 0; row < count; row++ {
		for col := 0; col < count; col++ {
			if row == col {
				fmt.Printf("*[%d-%d]\n", row, col)
			} else {
				fmt.Print("	")
			}
		}
		fmt.Println()
	}
}

func middleEmptyBox(count int) {
	for row := 0; row < count; row++ {
		for col := 0; col < count; col++ {
			if col == 0 || col == count-1 || row == 0 || row == count-1 {
				fmt.Printf("*[%d-%d] ", row, col)
			} else {
				fmt.Print("       ")
			}
		}
		fmt.Println()
	}
}

func leftFullDiagonal(count int) {

	for row := 0; row < count; row++ {
		for col := 0; col <= row; col++ {
			fmt.Printf("*[%d-%d] ", row, col)
		}
		fmt.Println()
	}
}

func coba(count int) {
	for row := 0; row < count; row++ {
		for col := 0; col < count; col++ {
			if col >= count-row-1 {
				fmt.Print("*")
			} else {
				fmt.Print(" ")
			}
		}
		fmt.Println()
	}
}

// 	for row := 0; row < count; row++ {
// 		for col := 0; col < count; col++ {
// 			if col >= count-row-1 {
// 				fmt.Print("*")
// 			} else {
// 				fmt.Print("1")
// 			}
// 		}
// 		fmt.Println("*")
// 	}
// }

// 	for col := count - 1; 0 <= col; col--{
// 		for row := 0; row < count; row++ {
// 			if row >= count-col-1{
// 				fmt.Print("* ")
// 			} else {
// 				fmt.Print(" ")
// 			}
// 		}
// 		fmt.Println()
// 	}
// }

func main1() {
	var num int
	fmt.Print("Masukkan angka: ")
	fmt.Scan(&num)

	for i := 1; i <= num; i++ {
		for j := num; j >= 1; j-- {
			if i%2 == 0 {
				fmt.Print(j)
			} else {
				fmt.Print(j + (i-1)*num)
			}
		}
		fmt.Println()
	}
}

func main2() {
	var num int
	fmt.Print("Masukkan angka: ")
	fmt.Scan(&num)

	for i := 1; i <= num; i++ {
		for j := 1; j <= num; j++ {
			if i%2 == 0 {
				fmt.Print(num - j + 1)
			} else {
				fmt.Print(j)
			}
		}
		fmt.Println()
	}
}

func main() {
	// basic()
	// kotak(5)
	// kotakIndex(5)
	// leftDiagonal(5)
	// middleEmptyBox(5)
	// leftFullDiagonal(5)
	// coba(5)

	main1()
}
