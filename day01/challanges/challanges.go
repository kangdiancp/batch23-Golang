package main

import (
	"fmt"
	"strconv"
)

func tahunKabisat(year int) {
	if (year%4 == 0 && year%100 != 0) || (year%400 == 0) {

	}
}

func displayTime(seconds int) {
	minute := seconds / 60
	remaining := seconds % 60
	fmt.Printf("Result : %d minutes %d seconds", minute, remaining)
}

func secondsToDays(seconds int) {
	var hari, jam, menit, detik, sisa int
	//1 hari = 24 jam, 1 jam = 60 menit, 1 menit =60detik
	// 1 hari = 24 jam * 3600 = 86_400detik
	hari = seconds / 86400
	sisa = seconds % 86400

	//hitung jam, 1jam = 3600detik
	jam = sisa / 3600
	sisa = sisa % 3600

	//hitung menit, 1menit = 60detik
	menit = sisa / 60
	detik = sisa % 60

	//display
	fmt.Printf("\nKonversi : %d hari %d jam %d menit %d detik", hari, jam, menit, detik)
}

/*
case input : 1234
*/
func reverseNumber(number int) {

	var angka1, angka2, angka3, angka4, sisa int
	if number < 1000 || number >= 10000 {
		fmt.Println("Input number range >= 1000 dan < 10000")
	} else {
		angka1 = number / 1000
		sisa = number % 1000

		angka2 = sisa / 100
		sisa = sisa % 100

		angka3 = sisa / 10
		angka4 = sisa % 10

		//Itoa -> convert integer to string
		reverse := strconv.Itoa(angka4) + strconv.Itoa(angka3) + strconv.Itoa(angka2) + strconv.Itoa(angka1)
		fmt.Printf("\n Reverse : %d -> %s", number, reverse)

	}
}

func main() {
	displayTime(12345)
	secondsToDays(1_000_000)
	reverseNumber(1234)
}
