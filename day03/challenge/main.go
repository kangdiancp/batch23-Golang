package main

import "fmt"

func main() {

	soal1(5)
	soal2(5)
	soal3(5)
	soal4(5)
	soal5(5)
	soal6()
	soal7()
	//soal7v2(5)
	soal8()
}

func soal1(count int) {

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

func soal2(count int) {

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

func soal3(count int) {
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

func soal4(count int) {
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

func soal5(count int) {
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

func soal6() {

	var angka int
	fmt.Print("Masukkan Inputan : ")
	fmt.Println()
	fmt.Scanf("%d", &angka)

	for row := 1; row <= angka; row++ {
		if row%2 == 0 {
			for col := 1; col <= angka; col++ {
				fmt.Print(col)
			}
		} else {
			for col := angka; col >= 1; col-- {
				fmt.Print(col)
			}
		}
		fmt.Println()
	}
}

func soal7() {

	var angka int
	fmt.Print("Masukkan Inputan : ")
	fmt.Println()
	fmt.Scanf("%d", &angka)

	for row := 0; row <= angka; row++ {
		if row%2 == 0 {
			for col := 1; col <= angka; col++ {
				fmt.Print(col)
			}
		} else {
			for col := 1; col >= 1; col++ {

				fmt.Print(col)
			}
		}
		fmt.Println()
	}
}

func soal8() {
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

//func soal7v2(angka int) {
//
//	fmt.Print("masukkan angka : ")
//	fmt.Scan(&angka)
//	for i := 1; i <= angka; i++ {
//		for j := 1; j <= angka; j++ {
//			if j%2 == 0 {
//				fmt.Println(angka + 1 - i)
//			} else {
//				fmt.Print(i)
//			}
//		}
//	}
//}
