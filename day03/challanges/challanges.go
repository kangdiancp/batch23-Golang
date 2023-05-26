package main

import (
	"fmt"
)

func main() {
	var inputNilai int
	println("Nomor 1 -------------------------")
	nomor1Diagonal(5)
	println("Nomor 2 -------------------------")
	nomor2(5)
	println("Nomor 3 -------------------------")
	nomor3(5)
	println("Nomor 4 -------------------------")
	nomor4(5)
	println("Nomor 5 -------------------------")
	nomor5(5)

	//Nomor 6
	println("Nomor 6 -------------------------")
	fmt.Print("Masukkan Angka : ")
	fmt.Scan(&inputNilai)
	nomor6(inputNilai)

	//Nomor 7
	println("Nomor 7 -------------------------")
	fmt.Print("Masukkan Angka : ")
	fmt.Scan(&inputNilai)
	nomor7(inputNilai)

	//Nomor 8
	println("Nomor 8 -------------------------")
	fmt.Print("Masukkan Angka : ")
	fmt.Scan(&inputNilai)
	nomor8(inputNilai)
}

//Nomor 1

func nomor1Diagonal(nilai int) {
	// for row := 0; row < nilai; row++ {
	// 	for col := 0; col < nilai; col++ {
	// 		if (row == 0 && col == nilai-1) || (row == nilai-1 && col == 0) || (row == 1 && col == nilai-2) || (row == nilai-2 && col == 1) || (row == 2 && col == 2) {
	// 			fmt.Printf("[%d - %d]", row, col)
	// 		} else {
	// 			fmt.Print("")
	// 		}
	// 		println()
	// 	}
	// 	println()
	// }
	for row := 0; row < nilai; row++ {
		for col := nilai - 1; col >= 0; col-- {
			if row == col {
				fmt.Print("* ")
			} else {
				if row <= col {
					fmt.Print("  ")
				} else {
					fmt.Print("  ")
				}
			}
		}
		fmt.Println()
	}
	println()
}

func nomor2(nilai int) {
	for row := 0; row < nilai; row++ {
		for col := nilai - 1; col >= 0; col-- {
			if row == col {
				fmt.Print("* ")
			} else {
				if row <= col {
					fmt.Print("  ")
				} else {
					fmt.Print("* ")
				}
			}
		}
		fmt.Println()
	}
	println()
}
func nomor3(nilai int) {

	// for row := nilai - 1; row <= 0; row-- {
	// 	for col := 0; col < nilai-row-1; col++ {
	// 		print("21 ")
	// 	}
	// 	if row <= nilai {
	// 		print("10 ")
	// 	}

	// 	println()
	// }
	// println()

	for row := 0; row < nilai; row++ {
		for col := nilai - 1; col >= 0; col-- {
			if row == col {
				fmt.Print("10 ")
			} else {
				if row <= col {
					fmt.Print("   ")
				} else {
					fmt.Print("20 ")
				}
			}
		}
		fmt.Println()
	}
	println()
}

func nomor4(nilai int) {
	for row := 0; row < nilai; row++ {
		for col := 0; col < nilai-row-1; col++ {
			print("21 ")
		}
		if row <= nilai {
			print("10 ")
		}

		println()
	}
	println()
}

func nomor5(nilai int) {
	for row := 0; row < nilai; row++ {
		for col := nilai - 1; col >= 0; col-- {
			if row == col {
				fmt.Print("10 ")
			} else {
				if row <= col {
					fmt.Print("21 ")
				} else {
					fmt.Print("20 ")
				}
			}
		}
		fmt.Println()
	}
	println()
}

func nomor6(nilai int) {
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
	println()
}

func nomor7(nilai int) {
	// for col := 0; col < nilai; col++ {
	// 	var output1, output2 string
	// 	if col%2 == 0 {
	// 		for row := 1; row <= nilai; row++ {
	// 			output1 = string(row)
	// 		}
	// 	} else {
	// 		for row := nilai; row > 0; row-- {
	// 			output2 = string(row)
	// 		}
	// 	}
	// 	print()
	// }

	for row := 1; row <= nilai; row++ {

		for loop := 0; loop < nilai/2; loop++ { //for ini digunakan untuk mengulang nilai column nya, dikarenakan row nya sudah sesuai dengan jumlah inputan nilai tetapi untuk columnnya tidak sesuai. Setelah jumlah col nya sudah sesuai, tetapi untuk inputan bernilai 5 tidak sesuai, yang dihasilkan nilai 4, maka dari itu diberi validasi dibawah perulangan ini.

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
	println()
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
