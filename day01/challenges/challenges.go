package main

import (
	"fmt"
	"math/rand"
	"strconv"
)

func soal1(number float64) {
	phi := 3.1416
	luas := phi * number * number
	fmt.Println("luas lingkaran :", luas)

}

func soal2(num1, num2 int) {
	jumlah := 180
	num3 := jumlah - num1 - num2

	println("sudut1:", num1)
	println("sudut2:", num2)
	println("sudut3:", num3)
}

func soal3(num1, num2, num3 float64) {
	n := 3
	average := (num1 + num2 + num3) / float64(n)
	fmt.Println("Average:", average)

}

func soal4(tahunkabisat int) {
	if tahunkabisat%4 == 0 && tahunkabisat%100 != 0 && tahunkabisat%400 != 0 {
		fmt.Printf("Tahun %d : Tahun Kabisat", tahunkabisat)
	} else {
		fmt.Printf("Tahun %d : Bukan Tahun Kabisat", tahunkabisat)
	}

}

func soal5(second int) {
	minute := second / 60
	sisaSecond := second % 60
	fmt.Printf("\nhasil : %d minute %d second\n", minute, sisaSecond)

}

func soal6(second int) {
	var hari, jam, menit, detik, sisa int
	//1 hari=24 jam, 1jam=60 menit, 1menit=60detik
	//1 hari = 24jam * 3600 = 86_400 detik
	jam = second / 86400
	sisa = second % 86400

	//hitung jam, 1jam=3600 detik
	jam = sisa / 3600
	sisa = sisa % 3600

	//hitung menit, 1menit=60 detik
	menit = second / 60
	detik = second % 60
	fmt.Printf("Total Detik: %d", second)
	fmt.Printf("\nTotal Konversi : %d hari %d jam %d menit %d detik", hari, jam, menit, detik)
}

func soal7(number int) {
	var angka1, angka2, angka3, angka4, sisa int
	if number < 1_000 || number >= 10_000 {
		fmt.Println("Input number range >= 1000 dan < 10000")
	} else {
		angka1 = number / 1000
		sisa = number % 1000

		angka2 = sisa / 100
		sisa = sisa % 100

		angka3 = sisa / 10
		angka4 = sisa % 10

		reverse := strconv.Itoa(angka4) + strconv.Itoa(angka3) + strconv.Itoa(angka2) + strconv.Itoa(angka1)
		fmt.Printf("Reverse Number: %d -> %s", number, reverse)
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
	soal4(2012)
	soal5(12345)
	soal6(2_000_000)
	soal7(1234)
	sortTigaAngka()
	gameGunting()
	sumNumber(1234)
}
