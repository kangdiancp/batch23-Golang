package main

import (
	"fmt"
	"log"
	"math"
	"math/rand"
)

func main() {

	soal1()
	soal2(30, 60)
	soal3(15, 15, 25)
	soal4(2012)
	soal5(12345)
	soal6(1_000_000)
	soal7(1234)
	soal8(1234)
	soal9()
	soal10()

}

func soal1() {
	var jariJari float64
	fmt.Print("Masukkan jari-jari lingkaran: ")
	fmt.Scan(&jariJari)

	luas := math.Pi * math.Pow(jariJari, 2)
	fmt.Printf("Luas lingkaran: %.2f\n", luas)

}

func soal2(num1, num2 int) {
	jumlah := 180
	sudut3 := jumlah - num1 - num2

	fmt.Println("sudut 1 :", num1)
	fmt.Println("sudut 2 : ", num2)
	fmt.Println("sudut 3: ", sudut3)

}

func soal3(num1, num2, num3 float64) {
	x := 3
	jumlah := (num1 + num2 + num3) / float64(x)

	fmt.Println(jumlah)
}

func soal4(tahunKabisat int) {

	if tahunKabisat%4 == 0 && tahunKabisat%100 != 0 && tahunKabisat%400 != 0 {
		fmt.Printf("%v  adalah tahun Kabisat", tahunKabisat)
	} else {
		fmt.Printf("%v bukan tahun kabisat", tahunKabisat)
	}
}

func soal5(second int) {

	menit := second / 60
	remainingSecond := second % 60

	fmt.Printf("Result : %d menit %d detik  ", menit, remainingSecond)
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
	if number >= 1_000 && number < 10_000 {
		digit1 := number % 10
		digit2 := (number / 10) % 10
		digit3 := (number / 100) % 10
		digit4 := number / 1000

		reversed := (digit1 * 1000) + (digit2 * 100) + (digit3 * 10) + digit4

		fmt.Printf("\n Reverse Number : %d --> %d", number, reversed)
	} else {
		fmt.Printf("\n Input Number harus dalam Range 1_000 - 10_000")
	}
}

func soal8(number int) {
	var angka1, angka2, angka3, angka4, sisa int
	if number < 1000 || number >= 10000 {
		fmt.Printf("\n input number harus dalam range 1000 - 10000")
	} else {
		angka1 = number / 1000
		sisa = number % 1000

		angka2 = sisa / 100
		sisa = sisa % 100

		angka3 = sisa / 10
		angka4 = sisa % 10

		sum := angka1 + angka2 + angka3 + angka4

		fmt.Printf("\n Sum Number : %d > %d ", number, sum)
	}
}

func soal9() {
	var angka1, angka2, angka3 int
	fmt.Printf("\n Masukkan 3 angka : \n")
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

func soal10() {
	var pilihan int
	fmt.Printf("\n Pilih : (1)Gunting  (2)Batu  (3)Kertas\n")
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
