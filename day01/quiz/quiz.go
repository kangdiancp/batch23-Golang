package main

import (
	"fmt"
	"math/rand"
	"strconv"
)

func totalLompat(x, y, k int) {
	sisa := (y - x) % k
	total := (y - x) / k
	if sisa != 0 {
		fmt.Println(total + 1)
	} else {
		fmt.Println(total)
	}
}

func tahunhari(menit int) {
	// 1 hari = 1440, 365 hari = 525600
	tahun := menit / 525600
	sisa := menit % 525600

	hari := sisa / 1440
	fmt.Printf("%d = %d tahun, %d hari", menit, tahun, hari)
	println()
}

func totaljarak() {
	var jarak, bensin, harga, total_harga float64
	println()
	fmt.Println("Jarak yang di tempuh : ")
	fmt.Scan(&jarak)
	fmt.Println("Konsumsi bensin per kilo : ")
	fmt.Scan(&bensin)
	fmt.Println("Harga bensin : ")
	fmt.Scan(&harga)

	total_harga = jarak / bensin * harga
	println()
	fmt.Printf("Uang Yang Harus Dikeluarkan Untuk Menempuh Jarak %8.1f = Rp.%8.2f", jarak, total_harga)
}

func shioCalender() {
	var tahun_lahir int
	println()
	fmt.Println("input tahun lahir : ")
	fmt.Scan(&tahun_lahir)

	switch t := tahun_lahir; {
	case t%12 == 0:
		fmt.Println("Tahun Monyet")
	case t%12 == 1:
		fmt.Println("Tahun Ayam")
	case t%12 == 2:
		fmt.Println("Tahun Anjing")
	case t%12 == 3:
		fmt.Println("Tahun Babi")
	case t%12 == 4:
		fmt.Println("Tahun Tikus")
	case t%12 == 5:
		fmt.Println("Tahun Kerbau")
	case t%12 == 6:
		fmt.Println("Tahun harimau")
	case t%12 == 7:
		fmt.Println("Tahun Kelinci")
	case t%12 == 8:
		fmt.Println("Tahun Naga")
	case t%12 == 9:
		fmt.Println("Tahun Ular")
	case t%12 == 10:
		fmt.Println("Tahun Kuda")
	case t%12 == 11:
		fmt.Println("Tahun Kambing")
	}
}

func simulasiATM() {
	var pin, jumlah_uang int

	fmt.Println("\nInputkan pin Passsword Anda : ")
	fmt.Scan(&pin)
	if pin != 1234 {
		fmt.Println("Pin Salah ! Kesempatan 2 Kali Lagi Untuk Mencoba : ")
		fmt.Println("Inputkan pin Passsword Anda : ")
		fmt.Scan(&pin)
		if pin != 1234 {
			fmt.Println("Pin Salah ! Kesempatan 1 Kali Lagi Untuk Mencoba : ")
			fmt.Println("Inputkan pin Passsword Anda : ")
			fmt.Scan(&pin)
			if pin != 1234 {
				fmt.Println("Pin Salah pin anda di lock")
			} else {
				fmt.Printf("Masukkan Nominal Uang yang akan diambil : $")
				fmt.Scan(&jumlah_uang)
				nominal1 := jumlah_uang / 10
				nominal2 := (jumlah_uang % 10) / 5
				nominal3 := jumlah_uang % 5

				fmt.Printf("$10 = %d Lembar, $5 = %d Lembar, $1 = %d Lembar", nominal1, nominal2, nominal3)
			}
		} else {
			fmt.Printf("Masukkan Nominal Uang yang akan diambil : $")
			fmt.Scan(&jumlah_uang)
			nominal1 := jumlah_uang / 10
			nominal2 := (jumlah_uang % 10) / 5
			nominal3 := jumlah_uang % 5

			fmt.Printf("$10 = %d Lembar, $5 = %d Lembar, $1 = %d Lembar", nominal1, nominal2, nominal3)
		}
	} else {
		fmt.Printf("Masukkan Nominal Uang yang akan diambil : $")
		fmt.Scan(&jumlah_uang)
		nominal1 := jumlah_uang / 10
		nominal2 := (jumlah_uang % 10) / 5
		nominal3 := jumlah_uang % 5

		fmt.Printf("$10 = %d Lembar, $5 = %d Lembar, $1 = %d Lembar", nominal1, nominal2, nominal3)
	}
}

func tebakhadiah() {
	var angka int
	println()
	fmt.Println("Masukkan angka : ")
	fmt.Scan(&angka)

	num1 := angka / 10
	sisa := angka % 10
	num2 := sisa % 10

	random := rand.Intn(100) + 11
	comp_random := random / 10
	sisa_random := random % 10
	comp_random1 := sisa_random % 10

	number_satu := strconv.Itoa(num1)
	number_dua := strconv.Itoa(num2)
	number_tiga := number_satu + number_dua
	gacha1 := strconv.Itoa(random)
	gacha2 := strconv.Itoa(comp_random)
	gacha3 := strconv.Itoa(comp_random1)

	if gacha1 == number_satu {
		println("User input : ", angka)
		println("Komputer random : ", gacha1)
		println("Match all digit : you win Rp.100.000")
	} else if gacha2 == number_dua {
		println("User input : ", angka)
		println("Komputer random : ", gacha1)
		println("Exact match : you win Rp.50.000")
	} else if gacha3 == number_tiga {
		println("User input : ", angka)
		println("Komputer random : ", gacha1)
		println("Exact match : you win Rp.50.000")
	} else {
		println("User input : ", angka)
		println("Komputer random : ", gacha1)
		println("You Lose!")
	}
}

func main() {
	/*totalLompat(10, 85, 30)
	tahunhari(1_000_000_000)
	totaljarak()
	shioCalender()
	tebakhadiah()*/
	simulasiATM()
}