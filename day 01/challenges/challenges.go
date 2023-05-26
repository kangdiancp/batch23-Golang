package main

import (
	"fmt"
	"math/rand"
	"strconv"
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

func displayTime() {
	display := 21
	minute := display / 60
	second := display % 60
	fmt.Printf("Output : %d minutes %d second", minute, second)
	println()
}

func yearTimed() {
	display := 1000000
	day := display / 86400
	hour := (display % 86400) / 3600
	minute := (display % 3600) / 60
	second := display % 60
	fmt.Printf("Output : %d day %d hour %d minutes %d second ", day, hour, minute, second)
	println()
}

func reverseNumber(number int) {
	var angka1, angka2, angka3, angka4, sisa int
	if number < 1000 || number >= 10000 {
		fmt.Println("Input Number Range >=1000 dari <10000")
	} else {
		angka1 = number / 1000
		sisa = number % 1000

		angka2 = sisa / 100
		sisa = sisa % 100

		angka3 = sisa / 10
		angka4 = sisa % 10

		reversed := strconv.Itoa(angka4) + strconv.Itoa(angka3) + strconv.Itoa(angka2) + strconv.Itoa(angka1)
		fmt.Printf("\nReverse : %d -> %s", number, reversed)
	}
}

func sumNumber(number int) {
	var angka1, angka2, angka3, angka4, sisa int
	if number < 1000 || number >= 10000 {
		fmt.Println("Input Number Range >=1000 dari <10000")
	} else {
		angka1 = number / 1000
		sisa = number % 1000

		angka2 = sisa / 100
		sisa = sisa % 100

		angka3 = sisa / 10
		angka4 = sisa % 10

		reversed := angka1 + angka2 + angka3 + angka4
		fmt.Printf("\nSum number : %d", reversed)
	}
}

func games() {
	fmt.Println("Pilih : (1)Gunting (2)Batu (3)Kertas")
	fmt.Println("Pilihan Anda: ")
	var input int
	fmt.Scan(&input)
	computer := rand.Intn(3) + 1
	fmt.Print(computer)
	if computer == input {
		println("Seri")
	} else if computer == 1 && input == 3 {
		println("Kalah")
	} else if computer == 2 && input == 1 {
		println("kalah")
	} else if computer == 3 && input == 2 {
		println("Kalah")
	} else {
		println("Menang")
	}
}

func sortTigaAngka() {
	fmt.Println("Input Tiga Angka: ")
	var angka1, angka2, angka3 int
	fmt.Scan(&angka1, &angka2, &angka3)
	if angka1 <= angka2 && angka1 <= angka3 {
		if angka2 <= angka3 {
			fmt.Printf("%d %d %d", angka1, angka2, angka3)
		} else {

			fmt.Printf("%d %d %d", angka1, angka3, angka2)
		}
	} else if angka2 <= angka1 && angka2 <= angka3 {
		if angka1 <= angka3 {
			fmt.Printf("%d %d %d", angka2, angka1, angka3)
		} else {
			fmt.Printf("%d %d %d", angka2, angka3, angka1)
		}
	} else {
		if angka1 <= angka2 {
			fmt.Printf("%d %d %d", angka3, angka1, angka2)
		} else {
			fmt.Printf("%d %d %d", angka3, angka2, angka1)
		}
	}
}

func main() {
	// var jari float32
	// var sudut1, sudut2 int
	// var av1, av2, av3 float32
	// var input int

	// //Nomor 1
	// fmt.Println("Masukkan Nilai lingkaran : ")
	// fmt.Scan(&jari)
	// luasLingkaran(jari)

	// //Nomor 2
	// fmt.Println("Masukkan Sudut Segitiga : ")
	// fmt.Scan(&sudut1, &sudut2)
	// sudutSegitiga(sudut1, sudut2)

	// //nomor3
	// fmt.Println("Masukkan inputan yang akan diratakan : ")
	// fmt.Scan(&av1, &av2, &av3)
	// average(av1, av2, av3)

	// //nomor 4
	// fmt.Println("Input Tahun: ")
	// fmt.Scan(&input)
	// tahunKabinet(input)
	// displayTime()
	// yearTimed()
	// reverseNumber(1234)
	// sumNumber(2424)
	// games()
	sortTigaAngka()
}
