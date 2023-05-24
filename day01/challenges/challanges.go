package main

import (
	"fmt"
	"log"
	"math/rand"
)

// 1.
func luasLingkaran(number float64) {
	phi := 3.1416
	luas := phi * number * number
	fmt.Printf("(1.) Luas Lingkaran = %.2f", luas)
}

// 2.
func sudutSegitiga(num1, num2 int) {
	jumlah := 180
	num3 := jumlah - num1 - num2
	fmt.Printf("\n(2.) Sudut 1 = %d", num1)
	fmt.Printf("\nSudut 2 = %d", num2)
	fmt.Printf("\nSudut 3 = %d", num3)
}

// 3.
func computeAverage(num1, num2, num3 float64) {
	n := 3
	rataRata := (num1 + num2 + num3) / float64(n)
	// println("Average = ", rataRata)
	fmt.Printf("\n(3.) Average = %.3f", rataRata)
}

// 4.
func tahunKabisat(number int) {
	if (number%4 == 0 && number%100 != 0) || (number%400 != 0) {
		fmt.Printf("\n(4.) Tahun %d Tahun Kabisat", number)
	} else {
		fmt.Printf("\n(4.) Tahun %d Bukan Tahun Kabisat", number)
	}
}

// 5.
func displayMinutesSec(seconds int) {
	minutes := seconds / 60
	seconds2 := seconds % 60
	fmt.Printf("\n(5.) Result : %d minutes %d seconds", minutes, seconds2)
}

// 6.
func secondsToDays(seconds int) {
	hari := seconds / 86400
	sisaHari := seconds % 86400
	jam := sisaHari / 3600
	sisaJam := sisaHari % 3600
	menit := sisaJam / 60
	detik := sisaJam % 60

	fmt.Printf("\n(6.) Konversi : %d hari %d jam %d menit %d detik", hari, jam, menit, detik)
}

// 7.
func reverseNumber(number int) {
	if number >= 1_000 && number < 10_000 {
		digit1 := number % 10
		digit2 := (number / 10) % 10
		digit3 := (number / 100) % 10
		digit4 := number / 1000

		reversed := (digit1 * 1000) + (digit2 * 100) + (digit3 * 10) + digit4

		fmt.Printf("\n(7.) Reverse Number : %d --> %d", number, reversed)
	} else {
		fmt.Printf("\n(7.) Input Number harus dalam Range 1_000 - 10_000")
	}
}

// 8.
func sumNumber(number int) {
	var angka1, angka2, angka3, angka4, sisa int
	if number < 1000 || number >= 10000 {
		fmt.Printf("\n(8.) Input Number Harus Dalam Range 1000 - 10000")
	} else {
		angka1 = number / 1000
		sisa = number % 1000

		angka2 = sisa / 100
		sisa = sisa % 100

		angka3 = sisa / 10
		angka4 = sisa % 10

		sum := angka1 + angka2 + angka3 + angka4
		fmt.Printf("\n Sum Number : %d -> %d", number, sum)
	}
}

// 9.
func sortTigaAngka() {
	var angka1, angka2, angka3 int
	fmt.Printf("\n(9.) Masukkan 3 angka : \n")
	_, err := fmt.Scan(&angka1, &angka2, &angka3)

	if angka1 < angka2 && angka1 < angka3 && angka2 < angka3 {
		fmt.Printf("%d , %d, %d", angka1, angka2, angka3)
	} else if angka1 > angka2 && angka1 < angka3 && angka2 < angka3 {
		fmt.Printf("%d, %d, %d", angka2, angka1, angka3)
	} else if angka1 > angka2 && angka1 > angka3 && angka2 < angka3 {
		fmt.Printf("%d, %d, %d", angka2, angka3, angka1)
	} else if angka1 < angka2 && angka1 > angka3 && angka2 > angka3 {
		fmt.Printf("%d, %d, %d", angka3, angka1, angka2)
	} else if angka1 > angka2 && angka1 > angka3 && angka2 > angka3 {
		fmt.Printf("%d, %d, %d", angka3, angka2, angka1)
	} else if angka1 < angka2 && angka1 < angka3 && angka2 > angka3 {
		fmt.Printf("%d, %d, %d", angka1, angka3, angka2)
	} else if err != nil {
		log.Fatal(err)
	}
}

// 10.
func gameGunting() {
	var pilihan int
	fmt.Printf("\n(10.) Pilih : (1)Gunting  (2)Batu  (3)Kertas\n")
	fmt.Printf("Masukkan Pilihan Anda : ")
	p, err := fmt.Scan(&pilihan)

	pil1 := "Gunting"
	pil2 := "Batu"
	pil3 := "Kertas"
	min := 1
	max := 3
	computer := rand.Intn(max-min) + min
	if p == 1 && computer == 1 {
		fmt.Printf("Anda Seri, computer [%s] vs Kamu [%s]\n", pil1, pil1)
	} else if p == 1 && computer == 2 {
		fmt.Printf("Anda Kalah, computer [%s] vs Kamu [%s]\n", pil1, pil2)
	} else if p == 1 && computer == 3 {
		fmt.Printf("Anda Menang, computer [%s] vs Kamu [%s]\n", pil1, pil3)
	} else if p == 2 && computer == 1 {
		fmt.Printf("Anda Menang, computer [%s] vs Kamu [%s]\n", pil2, pil1)
	} else if p == 2 && computer == 2 {
		fmt.Printf("Anda Seri, computer [%s] vs Kamu [%s]\n", pil2, pil2)
	} else if p == 2 && computer == 3 {
		fmt.Printf("Anda Kalah, computer [%s] vs Kamu [%s]\n", pil2, pil3)
	} else if p == 3 && computer == 1 {
		fmt.Printf("Anda Kalah, computer [%s] vs Kamu [%s]\n", pil3, pil1)
	} else if p == 3 && computer == 2 {
		fmt.Printf("Anda Menang, computer [%s] vs Kamu [%s]\n", pil3, pil2)
	} else if p == 3 && computer == 3 {
		fmt.Printf("Anda Seri, computer [%s] vs Kamu [%s]\n", pil3, pil3)
	} else if err != nil {
		log.Fatal(err)
	}
}

func main() {
	luasLingkaran(10)
	sudutSegitiga(30, 60)
	computeAverage(15, 15, 25)
	tahunKabisat(2012)
	displayMinutesSec(12345)
	secondsToDays(1_000_000)
	reverseNumber(4352)
	sumNumber(55_000)
	sortTigaAngka()
	gameGunting()
}
