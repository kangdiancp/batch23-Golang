package main

import (
	"fmt"
)

const phi = 3.14

func luasLingkaran(nilai float32) {
	hasil := (nilai * nilai) * phi
	fmt.Printf("Output Luas Lingkaran : %8.2f", hasil)
}

func sudutSegitiga(nilai1, nilai2 int) {
	output := nilai1 + nilai2
	println("Maka nilai ketiga sudut : ", output)
}

func average(average1, average2, average3 float32) {
	output := (average1 + average2 + average3) / 3
	fmt.Printf("Average : %8.3f", output)
}

func tahunKabinet(input int) {
	if input%4 == 0 && input%100 != 0 && input%400 != 0 {
		println("Tahun Kabisat")
	} else {
		println("Bukan Tahun Kabisat")
	}
}

func main() {
	var jari float32
	var sudut1, sudut2 int
	var av1, av2, av3 float32
	var input int

	//Nomor 1
	fmt.Println("Masukkan Nilai lingkaran : ")
	fmt.Scan(&jari)
	luasLingkaran(jari)

	//Nomor 2
	fmt.Println("Masukkan Sudut Segitiga : ")
	fmt.Scan(&sudut1, &sudut2)
	sudutSegitiga(sudut1, sudut2)

	//nomor3
	fmt.Println("Masukkan inputan yang akan diratakan : ")
	fmt.Scan(&av1, &av2, &av3)
	average(av1, av2, av3)

	//nomor 4
	fmt.Println("Input Tahun: ")
	fmt.Scan(&input)
	tahunKabinet(input)
}
