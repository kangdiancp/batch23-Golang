package main

import "fmt"

func soal1(number float64) {
	phi := 3.14159
	luas := phi * number * number
	luasLingkaran := fmt.Sprintf("Luas : %.2f", luas)
	println(luasLingkaran)
}

func soal2(sudut1 int, sudut2 int) {
	jumlah_sudut := 180
	sudut3 := jumlah_sudut - (sudut1 + sudut2)
	println("Sudut 1 : ", sudut1)
	println("Sudut 2 : ", sudut2)
	println("Sudut 3 : ", sudut3)
}

func soal3(angka1, angka2, angka3 float64) {
	average := (angka1 + angka2 + angka3) / 3
	hasil := fmt.Sprintf("Average : %.3f", average)
	println(hasil)
}

func soal4(tahun int) {
	// kelipatan100 := tahun%100 == 0
	// println(kelipatan100)

	if tahun%400 != 0 && tahun%4 == 0 && tahun%100 != 0 {
		println()
		fmt.Printf("Tahun %d tahun kabisat", tahun)
	} else {
		fmt.Printf("Tahun %d bukan tahun kabisat", tahun)
	}
}

func main() {
	soal1(10)
	soal2(30, 60)
	soal3(15, 15, 25)
	soal3(12.5, 9.7, 6.98)
	soal4(2002)
	soal4(2012)
}
