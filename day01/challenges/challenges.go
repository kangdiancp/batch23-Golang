package main

import "fmt"

func soal1(number float64) {
	phii := 3.14159
	luasLingkaran := phii * (number * number)
	println(luasLingkaran)
}

func soal2(sudut1, sudut2 int) {
	sudutSegiTiga := 180 - (sudut1 + sudut2)
	println("sudut1: ", sudut1)
	println("sudut2: ", sudut2)
	println("sudut3: ", sudutSegiTiga)
}

func soal3(num1, num2, num3 float32) {
	computeAverage := (num1 + num2 + num3) / 3
	println("Average: ", computeAverage)
}

func soal4(tahun int) {
	if tahun%400 == 0 && tahun%100 == 0 && tahun%4 == 0 {
		fmt.Printf("Tahun %d tahun kabisat", tahun)
	} else {
		fmt.Printf("Tahun %d bukan tahun kabisat", tahun)
	}
}

func main() {
	soal1(10)
	soal2(30, 60)
	soal3(12.5, 9.7, 6.98)
	soal4(2002)
}
