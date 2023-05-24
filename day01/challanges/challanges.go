package main

import "fmt"

func soal1(number float64) {
	var luas_lingkaran float64
	phi := 3.14159
	luas_lingkaran = phi * number * number
	fmt.Printf("luas_lingkaran = %8.3f\n", luas_lingkaran)

}

func soal2(num1, num2 int) {
	jumlah := 180
	sudut3 := jumlah - num1 - num2

	fmt.Println("sudut 1 :", num1)
	fmt.Println("sudut 2 : ", num2)
	fmt.Println("sudut 3: ", sudut3)

}

func soal3(f1 float64, f2 float64, f3 float64) {
	var avarage float64
	avarage = (f1 + f2 + f2) / 3
	fmt.Printf("avarage : %.3f\n", avarage)

}

func soal4(tahun int) {
	if tahun%4 == 0 && (tahun%400 == 0 || tahun%100 != 0) {
		fmt.Printf("tahun %d tahun kabisat\n", tahun)

	} else {
		fmt.Printf("tahun %d bukan tahun kabisat \n", tahun)
	}
}

func main() {
	soal1(10)
	//soal2(30, 60)
	//soal3(15, 15, 25)
	//soal3(12, 5, 9, 7, 6.98)
	soal4(2012)
	soal4(2002)

}
