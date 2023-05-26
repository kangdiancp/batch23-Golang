package main

import "fmt"

// 1.
func diagonalKanan(count int) {
	for i := 0; i < count; i++ {
		for j := 0; j < count; j++ {
			// [0, 4], [1, 3], [2, 2], [3, 1], [4, 0]
			if i+j == count-1 {
				fmt.Printf("[%d, %d] ", i, j)
			} else {
				fmt.Print("* ")
			}
		}
		fmt.Println(" ")
	}
}

// 2.
func segitigaKanan(count int) {
	for i := 0; i < count; i++ {
		for j := 0; j < count; j++ {
			if j >= (count-i)-1 {
				fmt.Print("*")
			} else {
				fmt.Print(" ")
			}
		}
		fmt.Println()
	}
}

// 3.
func segitigaKananAngkaBeda(count int) {
	for i := 0; i < count; i++ {
		for j := 0; j < count; j++ {
			if i+j == count-1 {
				fmt.Print("10")
			} else if j >= (count-i)-1 {
				fmt.Print("20")
			} else {
				fmt.Print("  ")
			}
		}
		fmt.Println(" ")
	}
}

// 4.
func segitigaKebalikBedaAngka(count int) {
	for i := 0; i < count; i++ {
		for j := count - 1; j > i; j-- {
			fmt.Print("21 ")
		}

		for k := 0; k < count; k++ {
			if i+k == count-1 {
				fmt.Print("10 ")
			}
		}
		fmt.Println(" ")
	}
}

// 5.
func kotakBedaAngka(count int) {
	for i := 0; i < count; i++ {
		for j := count - 1; j > i; j-- {
			fmt.Print("21 ")
		}

		for k := 0; k < count; k++ {
			if i+k == count-1 {
				fmt.Print("10 ")
			} else if k >= (count-i)-1 {
				fmt.Print("20 ")
			}
		}
		fmt.Println(" ")
	}
}

// 6.
func inputAngkaHitungMundur() {
	var number int
	fmt.Print("Input (Masukkan Angka 1 - 9) : ")
	fmt.Println(" ")
	fmt.Scanf("%d ", &number)
	fmt.Println("Output : ")

	for i := 1; i <= number; i++ {
		if i%2 == 0 {
			for j := 1; j <= number; j++ {
				fmt.Printf("%d ", j)
			}
		} else {
			for k := number; k >= 1; k-- {
				fmt.Printf("%d ", k)
			}
		}
		fmt.Println(" ")
	}
}

// 7.
func inputAngkaHitung2Pola() {
	var number int
	fmt.Print("Input (Masukkan Angka 1 - 9) : ")
	fmt.Println(" ")
	fmt.Scanf("%d ", &number)
	fmt.Println("Output : ")

	for i := 1; i <= number; i++ {
		for j := 0; j < number/2; j++ {
			for k := 1; k <= number; k++ {
				if k == number {
					fmt.Printf("%d ", i)
				} else {
					fmt.Print(" ")
				}
			}

			for l := number; l >= 1; l-- {
				if l == 1 {
					fmt.Printf("%d ", (number-i)+1)
				} else {
					fmt.Print(" ")
				}
			}
		}
		if number%2 == 1 {
			for m := 1; m <= number; m++ {
				if m == number {
					fmt.Printf("%d ", i)
				} else {
					fmt.Print(" ")
				}
			}
		}
		fmt.Println(" ")
	}
}

// 8.
func inputAngkaHitungAngkaDanSimbol() {
	var number int
	fmt.Print("Input (Masukkan Angka 1 - 9) : ")
	fmt.Println(" ")
	fmt.Scanf("%d ", &number)
	fmt.Println("Output : ")

	for i := 0; i < number; i++ {
		if i%2 == 0 {
			for j := 1; j <= number; j++ {
				if j%2 == 0 {
					fmt.Printf("%d ", j)
				} else {
					fmt.Print("- ")
				}
			}
		} else {
			for k := 1; k <= number; k++ {
				if k%2 == 0 {
					fmt.Print("- ")
				} else {
					fmt.Printf("%d ", k)
				}
			}
		}
		fmt.Println(" ")
	}
}

func main() {
	diagonalKanan(5)
	segitigaKanan(5)
	segitigaKananAngkaBeda(5)
	segitigaKebalikBedaAngka(5)
	kotakBedaAngka(5)
	inputAngkaHitungMundur()
	inputAngkaHitung2Pola()
	inputAngkaHitungAngkaDanSimbol()
}
