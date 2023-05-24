package main

import "fmt"

//hitung luas lingkaran dg satu inputan
//contoh : luaslingkaran (10)
//output : luas lingkaran : 314.16
func soal1(number float64) {
	phi := 3.1416
	luas := phi * number * number
	fmt.Println("luas lingkaran :", luas)

}

//soal no 2 hitung sudut segitiga terakhir :
// contoh sudut segitiga (30,60)
//output : sudut 1 : 30, sudut 2: 60, sudut3 : 90

func soal2(num1, num2 int) {
	jumlah := 180
	num3 := jumlah - num1 - num2

	println("sudut1:", num1)
	println("sudut2:", num2)
	println("sudut3:", num3)
}

//soal no 3, hitung niai rata2 dari 3 inputan untu function computerAverage
//contoh : computerAverage(15,15,25)
//output : avergae : 18.333
// contoh : computerAverage(12.5, 9.7, 6.98)
//output : avergae : 16.667
func soal3(num1, num2, num3 float64) {
	n := 3
	average := (num1 + num2 + num3) / float64(n)
	fmt.Println("Average:", average)

}

//soal 4, buat function tahun kabisat dengan inputan integer
//contoh : tahun kabisat(2002)
//output : tahun 2002 bukan tahun kabisat
// contoh : tahunkabisat(2012)
//output: tahun 2012 tahun kabisat
func soal4(tahunkabisat int) {
	if tahunkabisat%4 == 0 && tahunkabisat%100 != 0 && tahunkabisat%400 != 0 {
		fmt.Printf("Tahun %d : Tahun Kabisat", tahunkabisat)
	} else {
		fmt.Printf("Tahun %d : Bukan Tahun Kabisat", tahunkabisat)
	}

}

func main() {
	soal1(10)
	soal2(30, 60)
	soal3(15, 15, 25)
	soal4(2012)
}
