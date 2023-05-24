package main

import "fmt"

func soal1(number float64) {
	const phi = 314.16

	var luas = phi * number * number
	println("luas lingkaran :", luas)
}

func soal2(sudut1 int, sudut2 int) {
	jumlah_sudut := 180
	sudut3 := jumlah_sudut - sudut1 - sudut2
	println("sudut 1:", sudut1)
	println("sudut 2:", sudut2)
	println("sudut 3:", sudut3)
}

func soal3(rata1, rata2, rata3 float64) {
	average1 := (rata1 + rata2 + rata3) / 3
	jumlah := fmt.Sprintf("average : %8.3f", average1)
	println(jumlah)

}

func soal4(tahun int) {
	if tahun%400 == 0 {
		println("bukan tahun kabisat")
	} else if tahun%4 == 0 {
		println("tahun kabisat")
	} else if tahun%100 == 0 {
		println("bukan tahun kabisat")
	} else {
		println("bukan tahun kabisat")
	}

}

func main() {
	soal1(10)
	soal2(30, 60)
	soal3(15, 15, 25)
	soal3(12.5, 9.7, 6.98)
	soal4(2012)
	soal4(2002)

}
