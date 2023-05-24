package main

import (
	"fmt"
)

func soal1(number float64) {
	phi := 3.1416
	luas := phi * number * number
	fmt.Println("luas lingkaran :", luas)

}

func soal2(numb1, numb2 int) {
	jumlah := 180
	numb3 := jumlah - numb1 - numb2

	println("sudut1 :", numb1)
	println("sudut2 :", numb2)
	println("sudut3 :", numb3)

}

func soal3(angka1, angka2, angka3 float64) {
	bagi := 3
	average := (angka1 + angka2 + angka3) / float64(bagi)
	fmt.Println(average)

}

func soal4(kabisat int) {
	if kabisat%4 == 0 && kabisat%100 != 0 && kabisat%400 != 0 {
		println(kabisat, "Tahun Kabisat")
	} else {
		println(kabisat, "Bukan Tahun Kabisat")
	}
}

func main() {
	// soal1(10)
	// soal2(30, 60)
	// soal3(15, 15, 25)
	soal4(2012)
}
