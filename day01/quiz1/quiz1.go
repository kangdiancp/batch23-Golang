package main

import (
	"fmt"
	"math/rand"
	"time"
)

func quiz1totalLompat(x, y, k int) {
	distance := y - x
	jumps := distance / k
	remainder := distance % k

	if remainder > 0 {
		jumps++
		fmt.Printf("Output : %d", jumps)
	} else if remainder == 0 {
		fmt.Printf("Output : %d", jumps)
	}
}

func quiz2TahunHari(menit int) {
	// 1 jam = 60 menit, 1 hari = 24 jam, 1 hari = 24 x 60 = 1440 menit, 1 tahun = 365 hari = 525600
	tahun := menit / 525600
	sisa := menit % 525600

	hari := sisa / 1440
	fmt.Printf("%d = %d Tahun, %d Hari", menit, tahun, hari)
}

func quiz3TotalJarak() {
	var jarak, bensin, harga, total_harga float64

	fmt.Println("Masukkan Jarak yang di tempuh : ")
	fmt.Scan(&jarak)
	fmt.Println("Masukkan Konsumsi bensin per kilo : ")
	fmt.Scan(&bensin)
	fmt.Println("Masukkan Harga bensin : ")
	fmt.Scan(&harga)

	total_harga = jarak / bensin * harga
	fmt.Printf("Uang Yang Harus Dikeluarkan Untuk Menempuh Jarak %8.1f = Rp.%8.2f", jarak, total_harga)
}

func quiz4ShioCalender() {
	var tahun_lahir int
	fmt.Println("Masukkan Tahun Lahir : ")
	fmt.Scan(&tahun_lahir)

	switch t := tahun_lahir; {
	case t%12 == 0:
		fmt.Println("Anda Lahir Di Tahun Monyet")
	case t%12 == 1:
		fmt.Println("Anda Lahir Di Tahun Ayam")
	case t%12 == 2:
		fmt.Println("Anda Lahir Di Tahun Anjing")
	case t%12 == 3:
		fmt.Println("Anda Lahir Di Tahun Babi")
	case t%12 == 4:
		fmt.Println("Anda Lahir Di Tahun Tikus")
	case t%12 == 5:
		fmt.Println("Anda Lahir Di Tahun Kerbau")
	case t%12 == 6:
		fmt.Println("Anda Lahir Di Tahun harimau")
	case t%12 == 7:
		fmt.Println("Anda Lahir Di Tahun Kelinci")
	case t%12 == 8:
		fmt.Println("Anda Lahir Di Tahun Naga")
	case t%12 == 9:
		fmt.Println("Anda Lahir Di Tahun Ular")
	case t%12 == 10:
		fmt.Println("Anda Lahir Di Tahun Kuda")
	case t%12 == 11:
		fmt.Println("Anda Lahir Di Tahun Kambing")
	}
}

func quiz5SimulasiATM() {
	var pin, jumlah_uang int

	fmt.Println("Inputkan pin Passsword Anda : ")
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
				fmt.Println("Pin Salah pin anda di lock : ")
			}
		}
	} else {
		fmt.Printf("Masukkan Nominal Uang yang akan diambil : ")
		fmt.Scan(&jumlah_uang)
		nominal1 := jumlah_uang / 10
		nominal2 := (jumlah_uang % 10) / 5
		nominal3 := jumlah_uang % 5

		fmt.Printf("$10 = %d Lembar, $5 = %d Lembar, $1 = %d Lembar", nominal1, nominal2, nominal3)
	}
}

func quiz6tebakhadiah() {
	// Menghasilkan seed acak berdasarkan waktu saat ini
	rand.Seed(time.Now().UnixNano())

	// Menghasilkan angka acak antara 10 hingga 99
	randomNumber := rand.Intn(90) + 10

	var userGuess int
	fmt.Print("Masukkan angka tebakan (10-99): ")
	fmt.Scanln(&userGuess)

	if userGuess < 10 || userGuess > 99 {
		fmt.Println("Angka yang dimasukkan tidak valid. Harap masukkan angka antara 10 dan 99.")
		return
	}

	fmt.Printf("Angka tebakan pengguna: %d\n", userGuess)
	fmt.Printf("Angka acak: %d\n", randomNumber)

	if userGuess == randomNumber {
		fmt.Println("You win Rp 100.000")
	} else {
		fmt.Println("You Lose")
	}

}

func main() {
	quiz1totalLompat(10, 85, 30)
	//quiz2TahunHari(1000000)
	//quiz3TotalJarak()
	//quiz4ShioCalender()
	//quiz5SimulasiATM()
	//quiz6tebakhadiah()

}
