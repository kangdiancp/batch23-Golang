package main

import (
	"fmt"
	"log"
	"math/rand"
	"strconv"
)

func soal1(number float64) {
	const phi = 314.16

	var luas = phi * number * number
	println("luas lingkaran :", luas)
}

func soal2(sudut1 int, sudut2 int) {
	jumlah_sudut := 180
	sudut3 := jumlah_sudut - sudut1 - sudut2
	println("sudut 1:", sudut1)
	println("sudut 2:", sudut2)
	println("sudut 3:", sudut3)
}

func soal3(rata1, rata2, rata3 float64) {
	average1 := (rata1 + rata2 + rata3) / 3
	jumlah := fmt.Sprintf("average : %8.3f", average1)
	println(jumlah)

}

func soal4(tahun int) {
	if tahun%400 == 0 {
		println("bukan tahun kabisat")
	} else if tahun%4 == 0 {
		println("tahun kabisat")
	} else if tahun%100 == 0 {
		println("bukan tahun kabisat")
	} else {
		println("bukan tahun kabisat")
	}

}

func soal5(input int) {
	menit := input / 60
	detik := input % 60
	println()
	fmt.Printf("result : %d menit %d detik", menit, detik)
	println()
}

//cara 1
/*func soal6(input int) {
	detik := input % 60
	menit := input % 3600 / 60
	jam := input % 86400 / 3600
	hari := input / 86400
	println()
	fmt.Printf("%d hari %d jam %d menit %d detik", hari, jam, menit, detik)

}*/

// cara 2
func soal6(input int) {
	var hari, jam, menit, detik, sisa int
	hari = input / 86400
	sisa = input % 86400

	jam = sisa / 3600
	sisa = sisa % 3600

	menit = sisa / 60
	detik = sisa % 60

	fmt.Printf("%d Hari %d jam %d menit %d detik", hari, jam, menit, detik)
}

func soal7(number int) {
	var n1, n2, n3, n4, sisa int
	if number < 1_000 && number >= 10_000 {
		fmt.Println("input number harus dalam range >= 1_000 dan < 10_000")
	} else {
		n1 = number / 1000
		sisa = number % 1000

		n2 = sisa / 100
		sisa = number % 100

		n3 = sisa / 10
		n4 = sisa % 10

		reverse := strconv.Itoa(n4) + strconv.Itoa(n3) + strconv.Itoa(n2) + strconv.Itoa(n1)
		fmt.Printf("\n reverse : %d -> %s", number, reverse)

	}

}

func soal8(number int) {
	var n1, n2, n3, n4, sisa int
	if number < 1_000 || number >= 10000 {
		fmt.Println("input number harus dalam range 1_000 - 10_000")
	} else {
		n1 = number / 1000
		sisa = number % 1000

		n2 = sisa / 100
		sisa = number % 100

		n3 = sisa / 10
		n4 = sisa % 10

		sum := n1 + n2 + n3 + n4
		fmt.Printf("\n sum number : %d -> %d", number, sum)

	}

}

func soal9(int) {
	var pilihan int
	fmt.Printf("pilih : (1)Gunting  (2)Batu  (3)Kertas\n")
	fmt.Printf("masukan pilihan: ")
	p, err := fmt.Scan(&pilihan)

	plh1 := "Gunting"
	plh2 := "Batu"
	plh3 := "Kertas"
	min := 1
	max := 3
	computer := rand.Intn(max-min) + min
	if p == 1 && computer == 1 {
		fmt.Printf("Anda Seri, computer [%s] vs Kamu [%s]\n", plh1, plh1)
	} else if p == 1 && computer == 2 {
		fmt.Printf("Anda Kalah, computer [%s] vs Kamu [%s]\n", plh1, plh2)
	} else if p == 1 && computer == 3 {
		fmt.Printf("Anda Menang, computer [%s] vs Kamu [%s]\n", plh1, plh3)
	} else if p == 2 && computer == 1 {
		fmt.Printf("Anda Menang, computer [%s] vs Kamu [%s]\n", plh2, plh1)
	} else if p == 2 && computer == 2 {
		fmt.Printf("Anda Seri, computer [%s] vs Kamu [%s]\n", plh2, plh2)
	} else if p == 2 && computer == 3 {
		fmt.Printf("Anda Kalah, computer [%s] vs Kamu [%s]\n", plh2, plh3)
	} else if p == 3 && computer == 1 {
		fmt.Printf("Anda Kalah, computer [%s] vs Kamu [%s]\n", plh3, plh1)
	} else if p == 3 && computer == 2 {
		fmt.Printf("Anda Menang, computer [%s] vs Kamu [%s]\n", plh3, plh2)
	} else if p == 3 && computer == 3 {
		fmt.Printf("Anda Seri, computer [%s] vs Kamu [%s]\n", plh3, plh3)
	} else if err != nil {
		log.Fatal(err)
	}
}

func soal10(number int) {
	var n1, n2, n3 int
	fmt.Printf("input 3 angka : \n")
	_, err := fmt.Scan(&n1, &n2, &n3)

	if n1 < n2 && n1 < n3 && n2 < n3 {
		fmt.Printf("%d , %d, %d", n1, n2, n3)
	} else if n1 > n2 && n1 < n3 && n2 < n3 {
		fmt.Printf("%d, %d, %d", n2, n1, n3)
	} else if n1 > n2 && n1 > n3 && n2 < n3 {
		fmt.Printf("%d, %d, %d", n2, n3, n1)
	} else if n1 < n2 && n1 > n3 && n2 > n3 {
		fmt.Printf("%d, %d, %d", n3, n1, n2)
	} else if n1 > n2 && n1 > n3 && n2 > n3 {
		fmt.Printf("%d, %d, %d", n3, n2, n1)
	} else if n1 < n2 && n1 < n3 && n2 > n3 {
		fmt.Printf("%d, %d, %d", n1, n3, n2)
	} else if err != nil {
		log.Fatal(err)
	}
}

func main() {
	/*soal1(10)
	soal2(30, 60)
	soal3(15, 15, 25)
	soal3(12.5, 9.7, 6.98)
	soal4(2012)
	soal4(2002)
	soal5(60)
	soal5(12345)
	soal6(1_000_000)
	soal7(1234)
	soal8(1234)*/
	soal8(4545)
	soal9(3)
	soal10(3)

}
