package main

import (
	"fmt"
	"math"
)

func main() {

	soal1()
	soal2(30, 60)
	soal3(15, 15, 25)
	soal4(2012)

}

func soal1() {
	var jariJari float64
	fmt.Print("Masukkan jari-jari lingkaran: ")
	fmt.Scan(&jariJari)

	luas := math.Pi * math.Pow(jariJari, 2)
	fmt.Printf("Luas lingkaran: %.2f\n", luas)

}

func soal2(num1, num2 int) {
	jumlah := 180
	sudut3 := jumlah - num1 - num2

	fmt.Println("sudut 1 :", num1)
	fmt.Println("sudut 2 : ", num2)
	fmt.Println("sudut 3: ", sudut3)

}

func soal3(num1, num2, num3 float64) {
	x := 3
	jumlah := (num1 + num2 + num3) / float64(x)

	fmt.Println(jumlah)
}

func soal4(tahunKabisat int) {

	if tahunKabisat%4 == 0 && tahunKabisat%100 != 0 && tahunKabisat%400 != 0 {
		fmt.Printf("%v  adalah tahun Kabisat", tahunKabisat)
	} else {
		fmt.Printf("%v bukan tahun kabisat", tahunKabisat)
	}
}
