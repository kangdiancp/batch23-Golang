package main

import "fmt"

// 1.
func luasLingkaran(number float64) {
	phi := 3.1416
	luas := phi * number * number
	fmt.Println("(1.) Luas Lingkaran = ", luas)
}

// 2.
func sudutSegitiga(num1, num2 int) {
	jumlah := 180
	num3 := jumlah - num1 - num2
	println("\n(2.) Sudut 1 = ", num1)
	println("Sudut 2 = ", num2)
	println("Sudut 3 = ", num3)
}

// 3.
func computeAverage(num1, num2, num3 float64) {
	n := 3
	rataRata := (num1 + num2 + num3) / float64(n)
	// println("Average = ", rataRata)
	fmt.Printf("\n(3.) Average = %.3f", rataRata)
}

// 4.
func tahunKabisat(number int) {
	if number%4 == 0 && number%100 != 0 && number%400 != 0 {
		fmt.Printf("\n(4.) Tahun %d Tahun Kabisat", number)
	} else {
		fmt.Printf("\n(4.) Tahun %d Bukan Tahun Kabisat", number)
	}
}

func main() {
	luasLingkaran(10)
	sudutSegitiga(30, 60)
	computeAverage(15, 15, 25)
	tahunKabisat(2012)
}
