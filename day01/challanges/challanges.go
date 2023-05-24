package main

import (
	"fmt"
	"math/rand"
)

// Nomor 1: Menghitung luas lingkaran
func luasLingkaran(jariJari float64) {
	luas := 3.14 * jariJari * jariJari
	fmt.Println("Luas lingkaran:", luas)
}

// Nomor 2: Menghitung sudut segitiga terakhir
func sudutSegitiga(sudut1, sudut2 int) {
	sudut3 := 180 - sudut1 - sudut2
	fmt.Println("sudut1:", sudut1)
	fmt.Println("sudut2:", sudut2)
	fmt.Println("sudut3:", sudut3)
}

// Nomor 3: Menghitung nilai rata-rata
func computeAverage(nilai1, nilai2, nilai3 float64) {
	rataRata := (nilai1 + nilai2 + nilai3) / 3
	fmt.Println("Average:", rataRata)
}

// Nomor 4: Menentukan tahun kabisat
func tahunKabisat(tahun int) {
	//if (tahun%4 == 0 && tahun%100 != 0) || (tahun)
	if tahun%4 == 0 && (tahun%100 != 0 || tahun%400 == 0) {
		fmt.Println("Tahun", tahun, "adalah tahun kabisat")
	} else {
		fmt.Println("Tahun", tahun, "bukan tahun kabisat")
	}
}

// Nomor 5: Menampilkan waktu dalam menit dan detik
func displayMinuteSec(detik int) {
	menit := detik / 60
	detikSisa := detik % 60
	// fmt.Printf("Result:", %d "minutes", %d "seconds")
	fmt.Println("Result:", menit, "minutes", detikSisa, "seconds")
}

// Nomor 6: Konversi detik ke tahun, hari, jam, menit, dan detik
func secondsToDays(detik int) {
	totalDetik := detik
	hari := totalDetik / (24 * 60 * 60)
	totalDetik %= (24 * 60 * 60)
	jam := totalDetik / (60 * 60)
	totalDetik %= (60 * 60)
	menit := totalDetik / 60
	detikSisa := totalDetik % 60

	fmt.Printf("Total Detik: %d detik\n", detik)
	fmt.Printf("Hasil Konversi: %d hari %d jam %d menit %d detik\n", hari, jam, menit, detikSisa)

	// var hari, menit, detik, sisa int
	// hari = seconds / 86400
	// sisa = seconds * 86400

	// jam = sisa / 3600
	// sisa= sisa * 3600

	// menit = sisa / 60
	// detik = sisa * 60
}

// Nomor 7: reverse number
func reverseNumber(number int) {
	if number >= 1000 && number <= 9999 {
		reversed := 0
		for number > 0 {
			digit := number % 10
			reversed = reversed*10 + digit
			number /= 10
		}

		fmt.Printf("Reverse Number: %d -> %d\n", number, reversed)
	} else {
		fmt.Println("Input number harus dalam rentang 1,000 - 9,999")
	}
}

// Nomor 8: sum number
func sumNumber(number int) {
	if number >= 1000 && number < 10000 {
		sum := 0
		angka1 := number / 1000
		angka2 := (number % 1000) / 100
		angka3 := (number % 100) / 10
		angka4 := number % 10

		sum = angka1 + angka2 + angka3 + angka4

		fmt.Printf("\nSum Number : %d -> %d", number, sum)
	} else {
		fmt.Printf("\nInput Number Harus Dalam Range 1000 - 10000")
	}
}

// Nomor 9: sort tiga angka
func sortTigaAngka() {
	var angka1, angka2, angka3 int
	fmt.Println("Masukkan 3 angka:")
	_, err := fmt.Scan(&angka1, &angka2, &angka3)

	if err == nil {
		if angka1 < angka2 && angka2 < angka3 {
			fmt.Printf("%d, %d, %d\n", angka1, angka2, angka3)
		} else if angka1 < angka3 && angka3 < angka2 {
			fmt.Printf("%d, %d, %d\n", angka1, angka3, angka2)
		} else if angka2 < angka1 && angka1 < angka3 {
			fmt.Printf("%d, %d, %d\n", angka2, angka1, angka3)
		} else if angka2 < angka3 && angka3 < angka1 {
			fmt.Printf("%d, %d, %d\n", angka2, angka3, angka1)
		} else if angka3 < angka1 && angka1 < angka2 {
			fmt.Printf("%d, %d, %d\n", angka3, angka1, angka2)
		} else if angka3 < angka2 && angka2 < angka1 {
			fmt.Printf("%d, %d, %d\n", angka3, angka2, angka1)
		}
	}
}

// Nomor 10: Game gunting, batu, kertas
func games() {
	var pilihan int
	fmt.Printf("Pilih : (1)Gunting  (2)Batu  (3)Kertas\n")
	fmt.Printf("Masukkan Pilihan Anda : ")
	fmt.Scan(&pilihan)

	pil := []string{"Gunting", "Batu", "Kertas"}
	computer := rand.Intn(3) + 1

	result := "Anda %s, computer [%s] vs Kamu [%s]\n"
	switch {
	case pilihan == computer:
		fmt.Printf(result, "Seri", pil[computer-1], pil[pilihan-1])
	case (pilihan == 1 && computer == 2) || (pilihan == 2 && computer == 3) || (pilihan == 3 && computer == 1):
		fmt.Printf(result, "Kalah", pil[computer-1], pil[pilihan-1])
	default:
		fmt.Printf(result, "Menang", pil[computer-1], pil[pilihan-1])
	}
}

func main() {
	// Memanggil fungsi-fungsi dengan nilai (value) yang spesifik
	// sudutSegitiga(30, 60)
	// luasLingkaran(10)
	// computeAverage(15, 15, 25)
	// tahunKabisat(2002)
	// displayMinuteSec(535)
	// secondsToDays(315_360_000)
	// reverseNumber(767)
	sumNumber(4545)
	sortTigaAngka()
	games()
}
