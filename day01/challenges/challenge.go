package main

import (
	"fmt"
)

func soal1(number float64) {
	var luas_lingkaran float64
	phi := 3.14159
    luas_lingkaran = phi * number * number
	fmt.Printf("Luas lingkaran =  %8.3f \n", luas_lingkaran)
}

func soal2(sudut1 int, sudut2 int) {
	var sudut3 int
	sudut3 = 180 - (sudut1 + sudut2)
	println("sudut 1 = ", sudut1)
	println("sudut 2 = ", sudut2)
	fmt.Println("Jumlah sudut segitiga  : ", sudut3)
}

func soal3(N1 float64, N2 float64, N3 float64){
	var Average float64
	Average = (N1 + N2 + N3) / 3
	fmt.Printf("Average : %8.3f \n", Average)
}

func soal4(tahun int){
	if tahun % 4 == 0 && (tahun % 400 == 0 || tahun % 100 != 0) {
		fmt.Printf("Tahun %d tahun kabisat \n", tahun)
	}else{
		fmt.Printf("Tahun %d bukan tahun kabisat \n", tahun)
	}
}

func main() {
	//soal1(10)
	//soal2(30, 60)
	//soal3(15, 15, 25)
	//soal3(12.5, 9.7, 6.98)
	soal4(2012)
	soal4(2002)
}