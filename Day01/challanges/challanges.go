package main

import (
	"fmt"
	"math/rand"
)

func main() {
	// soal1(10)
	// soal2(30, 60)
	// soal3(15, 15, 25)
	// soal3(12, 5, 6.98)
	// soal4(2012)
	// soal4(2002)
	// soal5(3620)
	// secondToDays(1000000)
	// reverseNumber(1234)
	// gameGunting()
	// sortTigaAngka()
	// sumNumber(1234)
}

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

func soal3(f1, f2, f3 float64) {
	avarage := (f1 + f2 + f2) / 3
	fmt.Printf("avarage : %.3f\n", avarage)
}

func soal4(tahun int) {
	if tahun%4 == 0 && (tahun%400 == 0 || tahun%100 != 0) {
		fmt.Printf("tahun %d tahun kabisat\n", tahun)
	} else {
		fmt.Printf("tahun %d bukan tahun kabisat \n", tahun)
	}
}

// Soal 5 = displayime
func soal5(second int) {
	menit := second / 60
	detik := second % 60
	fmt.Printf("Result : %d minutes %d seconds", menit, detik)
}

// Soal 6
// func secondToDays(second int) {
// 	var hari, jam, menit, detik, sisa int
// 1 hari = 24 jam, 1 jam = 60 menit, 1 menit = 60 detik
// 1 hari = 24 jam * 3600 = 86_400 detk
// 	hari = second / 86400
// 	sisa = second % 86400

// hitung jam, 1 jam = 3600 detik
// 	jam = sisa / 3600
// 	sisa = sisa % 3600

// hitung menit, 1 menit = 60 detik
// 	menit = sisa / 60
// 	detik = sisa % 60

// display
// 	fmt.Printf("\nKonverisi : %dhari %d jam %d menit %d detik", hari, jam, menit, detik)
// }

// Soal 7 = reverseNumber
func soal7(number int) {

	var angka1, angka2, angka3, angka4, sisa int
	if number < 1000 || number >= 10000 {
		fmt.Println("Input number range >= 1000 - < 10000")
	} else {
		angka1 = number / 1000
		sisa = number % 1000

		angka2 = number / 100
		sisa = number % 100

		angka3 = sisa / 10
		angka4 = sisa % 10

// Itoa -> Convert integer to string
		reverse := strconv.Itoa(angka4) + strconv.Itoa(angka3) + strconv.Itoa(angka2) + strconv.Itoa(angka1)
		fmt.Printf("\n Reverse : %d -> %s", number, reverse)
	}
}

func sumNumber(num int) {
	var
}

// Soal 8
func gameGunting() {
	fmt.Println("Pilih: (1) Gunting, (2)Batu, (3)Kertas")
	fmt.Print("input: ")

	var pilihankamu int
	fmt.Scan(&pilihankamu)

	komputer := rand.Intn(3) + 1

	if komputer == 1 && pilihankamu == 1 {
		fmt.Printf("Seri, Komputer [Gunting] VS Kamu [Gunting]")

	} else if komputer == 1 && pilihankamu == 2 {
		fmt.Printf("Anda Menang, Komputer [Gunting] VS Kamu [Batu]")

	} else if komputer == 1 && pilihankamu == 3 {
		fmt.Printf("Anda Kalah, Komputer [Gunting] VS Kamu [Kertas]")

	} else if komputer == 2 && pilihankamu == 1 {
		fmt.Printf("Anda Kalah, Komputer [Batu] VS Kamu [Gunting]")

	} else if komputer == 2 && pilihankamu == 2 {
		fmt.Printf("Draw, Komputer [Batu] VS Kamu [Batu]")

	} else if komputer == 2 && pilihankamu == 3 {
		fmt.Printf("Anda Menang, Komputer [Batu] VS Kamu [Kertas]")

	} else if komputer == 3 && pilihankamu == 1 {
		fmt.Printf("Anda Menang, Komputer [Kertas] VS Kamu [Gunting]")
	} else if komputer == 3 && pilihankamu == 2 {
		fmt.Printf("Anda Kalah, Komputer [Kertas] VS Kamu [Batu]")
	} else {
		fmt.Printf("Draw, Komputer [Kertas] VS Kamu [Kertas]")
	}

}

// Soal 9
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

// Soal 9
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

// Soal 10
func sumNumber(number int) {

	num1 := number % 1000
	num2 := num1 % 100
	satuan := num2 % 10
	puluhan := num2 / 10
	ratusan := num1 / 100
	ribuan := number / 1000
	sum := ribuan + ratusan + puluhan + satuan

	if number < 1_000 || number > 9999 {
		println("input number harus dalam range 1_000 - 9999")
	} else {
		fmt.Printf("Sum Number: %d -> %d ", number, sum)
	}
}
