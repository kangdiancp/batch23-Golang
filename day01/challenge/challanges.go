package main

import (
	"fmt"
	"math/rand"
	"strconv"
)

func luasLingkaran(jari int) {
	const phi = 3.14159
	var jarinya = float64(jari)
	var luas float64 = jarinya * jarinya * phi
	format := strconv.FormatFloat(luas, 'f', 3, 64)
	fmt.Println("luas Lingkaran: ", format)

}

func sudutSegitiga(sudut1 int, sudut2 int) {

	var sudut3 int

	sudut3 = 180 - (sudut1 + sudut2)
	fmt.Printf("sudut1: %d \n", sudut1)
	fmt.Printf("sudut2: %d \n", sudut2)
	fmt.Printf("sudut3: %d \n", sudut3)

}

func computeAverege(n1 float64, n2 float64, n3 float64) {
	var average float64
	average = (n1 + n2 + n3) / 3
	fmt.Printf("Average: %8.3f", average)
}

func tahunKabisat(th int) {
	if th%4 == 0 && (th%400 == 0 || th%100 != 0) {
		fmt.Printf("Tahun %d tahun Kabisat", th)
	} else {
		fmt.Printf("Tahun %d bukan tahun Kabisat", th)

	}
}

func displayTime(detik int) {
	var menit int
	menit = detik / 60
	newdetik := detik % 60
	fmt.Printf("Result: %d minutes %d second", menit, newdetik)
}

func secondsToDays(seconds int) {
	hari := seconds / 86400
	jam := (seconds % 86400) / 3600
	menit := (seconds % 3600) / 60
	detik := seconds % 60
	fmt.Printf("Total Detik: %d", seconds)
	fmt.Printf("Total Konversi: %d hari, %d jam, %d menit, %d detik", hari, jam, menit, detik)

}

func reverseNumber(num int) {
	var a1, a2, a3, a4 int
	if num >= 1000 || num < 10000 {
		a1 = num / 1000
		sisa := num % 1000
		a2 = sisa / 100
		sisa = sisa % 100

		a3 = sisa / 10
		a4 = sisa % 10

		reverse := strconv.Itoa(a4) + strconv.Itoa(a3) + strconv.Itoa(a2) + strconv.Itoa(a1)
		fmt.Printf("Reverse Number: %d -> %s", num, reverse)

	} else {
		fmt.Println("input number harus dalam range 1_000 - 10_0000")
	}

}

// /LANJUTAN//
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
