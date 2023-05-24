package main

import "fmt"

func soal1() {
	var luaslingkaran int = 10

	println("luaslingkaran :", 22/7*luaslingkaran)

}

func soal2() {
	var sudut1 int = 30
	sudut2 := 60
	sudut3 := sudut1 + sudut2

	println("sudut1 :", sudut1)
	println("sudut2 :", sudut2)
	println("sudut3 :", sudut3)

}

func soal3(f1 float64, f2 float64, f3 float64) {
	var avarage float64
	avarage = (f1 + f2 + f2) / 3
	fmt.Printf("avarage : %.3f\n", avarage)

}

func soal4(tahun int) {
	if tahun%4 == 0 && (tahun%100 != 0 || tahun%400 == 0) {
		fmt.Printf("Tahun %d tahun kabisat : \n", tahun)
	} else {
		fmt.Printf("Tahun %d tidak tahun kabisat : \n", tahun)
	}
}

func main() {
	soal1()
	soal2()
	soal3(15, 15, 25)
	soal4(2020)

}
