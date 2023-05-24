package main

import (
	"fmt"
	"math/rand"
	"strconv"
)

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
	if (tahun%4 == 0 && tahun%100 != 0) || (tahun%400 == 0) {
		println()
		fmt.Printf("Tahun %d tahun kabisat", tahun)
	} else {
		fmt.Printf("Tahun %d bukan tahun kabisat", tahun)
	}
}

func soal5(input int) {
	x := 60
	menit := input / x
	detik := input % x
	println()
	fmt.Printf("Result : %d menit %d seconds", menit, detik)
}

func soal6(input int) {
	// CARA 1
	// secToDay := 86400
	// secToHours := 3600
	// secToMinute := 60

	// hari := input / secToDay
	// jam := (input % secToDay) / secToHours
	// menit := (input % secToHours) / secToMinute
	// detik := input % secToMinute

	// println()
	// fmt.Printf("%d Hari %d jam %d menit %d detik", hari, jam, menit, detik)

	//CARA 2
	var hari, jam, menit, detik, sisa int
	hari = input / 86400
	sisa = input % 86400

	jam = sisa / 3600
	sisa = sisa % 3600

	menit = sisa / 60
	detik = sisa % 60

	fmt.Printf("%d Hari %d jam %d menit %d detik", hari, jam, menit, detik)
}

func soal7(input int) {
	if input < 1_000 || input >= 10_000 {
		println("Input number harus dalam range >= 1000 dan < 10_000")
	} else {
		angka1 := input / 1000
		sisa := input % 1000

		angka2 := sisa / 100
		sisa = sisa % 100

		angka3 := sisa / 10
		angka4 := sisa % 10

		reverse := strconv.Itoa(angka4) + strconv.Itoa(angka3) + strconv.Itoa(angka2) + strconv.Itoa(angka1)
		println()
		fmt.Println(reverse)
	}
}

func gameGunting() {
	fmt.Println("Pilih: (1) Gunting, (2)Batu, (3)Kertas")
	fmt.Println("input: ")

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

func main() {
	soal1(10)
	soal2(30, 60)
	soal3(15, 15, 25)
	soal3(12.5, 9.7, 6.98)
	soal4(2002)
	soal4(2012)
	soal5(60)
	soal5(12345)
	soal6(1_000_000)
	soal7(1234)
	gameGunting()
	games()
	sortTigaAngka()
	sumNumber(1234)

}
