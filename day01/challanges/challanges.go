package main

import (
	"fmt"
	"math/rand"
	"strconv"
)

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

func soal5(second int) {
	menit := second / 60
	detik := second % 60

	fmt.Printf("Result : %d minutes %d seconds \n", menit, detik)

}

func soal6(detik int) {
	// 1 hari 24 jam, 1 jam= 60 detik, 1 menit=60 detik
	hari := detik / 86400
	jam := (detik % 86400) / 3600
	menit := ((detik % 86400) % 3600) / 60
	second := detik % 60
	fmt.Printf("%d hari %d jam %d menit %d detik\n", hari, jam, menit, second)
}

func soal7(number int) {
	var angka1, angka2, angka3, angka4, sisa int
	if number < 1000 || number >= 10000 {
		fmt.Printf("Input Number Harus dalem Range 1000-10000")
	} else {
		angka1 = number / 1000
		sisa = number % 1000

		angka2 = sisa / 1000
		sisa = sisa / 100

		angka3 = sisa / 10
		sisa = sisa % 10

		reverse := strconv.Itoa(angka4) + strconv.Itoa(angka3) + strconv.Itoa(angka2) + strconv.Itoa(angka1)
		fmt.Printf("\n Reverse : %d -> %s", number, reverse)
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
	//luasLingkaran(10)
	//sudutSegitiga(30, 60)
	//computeAverege(15, 15, 25)
	//computeAverege(12.5, 9.7, 6.98)
	//tahunKabisat(2200)
	//displayTime(12345)
	//secondsToDays(2_000_000)
	//reverseNumber(1234)
	sortTigaAngka()
	gameGunting()
	sumNumber(1234)

}
