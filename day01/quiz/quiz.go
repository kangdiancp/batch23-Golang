package main

import (
	"fmt"
	"math/rand"
)

// 1. Quiz
func totalLompat(x, y, k int) {
	distance := y - x
	jumps := distance / k
	remainder := distance % k

	if remainder > 0 {
		jumps++
		fmt.Printf("(1.) Output : %d", jumps)
	} else if remainder == 0 {
		fmt.Printf("(1.) Output : %d", jumps)
	}
}

// 2. Quiz
func tahunHari(menit int) {
	// 1 jam = 60 menit, 1 hari = 24 jam, 1 hari = 24 x 60 = 1440 menit, 1 tahun = 365 hari = 525600
	tahun := menit / 525600
	sisa := menit % 525600

	hari := sisa / 1440
	fmt.Printf("\n(2.) %d Menit = %d Tahun, %d Hari", menit, tahun, hari)
}

// 3. Quiz
func totalJarak() {
	var jarak, bensin, harga, total_harga float64

	fmt.Println("\n(3.) Masukkan Jarak yang di tempuh : ")
	fmt.Scan(&jarak)
	fmt.Println("Masukkan Konsumsi bensin per kilo : ")
	fmt.Scan(&bensin)
	fmt.Println("Masukkan Harga bensin : ")
	fmt.Scan(&harga)

	total_harga = jarak / bensin * harga
	fmt.Printf("Uang Yang Harus Dikeluarkan Untuk Menempuh Jarak %8.1f = Rp.%8.2f", jarak, total_harga)
}

// 4. Quiz
func shioCalender() {
	var tahun_lahir int
	fmt.Println("\n(4.) Masukkan Tahun Lahir : ")
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
		fmt.Println("Anda Lahir Di Tahun Kambing")
	}
}

// 5. Quiz
func simulasiATM() {
	var pin, jumlah_uang int

	fmt.Println("\n(5.) Inputkan pin Passsword Anda : ")
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

// 6. Quiz
func tebakHadiah() {
	var nomor_togel, angka1, angka2, angka3, angka4 int
	fmt.Println("$$$ ==> SELAMAT DATANG DI PERMAINAN TEBAK BERHADIAH <=== $$$")
	fmt.Println("")
	fmt.Println("Silahkan Masukkan 2 digit Nomor Keberuntungan anda antara 10 - 99")
	fmt.Scan(&nomor_togel)
	angka1 = nomor_togel / 10
	angka2 = nomor_togel % 10

	min := 10
	max := 99
	komputer_random := rand.Intn(max-min) + min
	angka3 = komputer_random / 10
	angka4 = komputer_random % 10

	if nomor_togel < 10 || nomor_togel > 99 {
		fmt.Printf("Input Harus berada pada range 10 - 99")
	} else if nomor_togel == komputer_random {
		fmt.Println("User Input :", nomor_togel)
		fmt.Println("Komputer random :", komputer_random)
		fmt.Printf("Match All Digit : You Win Rp 100.000")
	} else if angka1 == angka3 || angka1 == angka4 || angka2 == angka3 || angka2 == angka4 {
		fmt.Println("User Input :", nomor_togel)
		fmt.Println("Komputer random :", komputer_random)
		fmt.Printf("Exact match : You win Rp 50.000")
	} else {
		fmt.Println("User Input :", nomor_togel)
		fmt.Println("Komputer random :", komputer_random)
		fmt.Println("You Lose !")
	}
}

func main() {
	totalLompat(10, 100, 30)
	tahunHari(1_000_000_000)
	totalJarak()
	shioCalender()
	simulasiATM()
	tebakHadiah()
}
