package main

import (
	"fmt"
	"math/rand"
)

func main() {
	//Nomor 1
	totalLompat(10, 85, 30)
	//Nomor 2
	konversiTahunHari(1000000000)
	//Nomor 3
	var jarak, bensin, harga float64
	fmt.Print("Masukkan Jarak yang di tempuh : ")
	fmt.Scan(&jarak)
	fmt.Print("Masukkan Konsumsi bensin per kilo : ")
	fmt.Scan(&bensin)
	fmt.Print("Masukkan Harga bensin : ")
	fmt.Scan(&harga)
	totalJarak(jarak, bensin, harga)

	//Nomor 4
	var inputThnLahir int
	fmt.Print("Masukkan Tahun Lahir : ")
	fmt.Scan(&inputThnLahir)
	shioCalender(inputThnLahir)

	//Nomor 5
	simulasiATM(10, 5, 1)

	//Nomor 6
	var userInput int
	fmt.Print("User Input : ")
	fmt.Scan(&userInput)
	tebakhadiah(userInput)

}

// Nomor 1
func totalLompat(x, y, k int) {
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

// Nomor 2
func konversiTahunHari(menit int) {
	// 1 jam = 60 menit, 1 hari = 24 jam, 1 hari = 24 x 60 = 1440 menit, 1 tahun = 365 hari = 525600
	tahun := menit / 525600
	sisa := menit % 525600

	hari := sisa / 1440
	fmt.Printf("%d = %d Tahun, %d Hari", menit, tahun, hari)
}

// Nomor 3
func totalJarak(jarak, bensin, harga float64) {
	var total_harga float64

	total_harga = jarak / bensin * harga
	fmt.Printf("Uang Yang Harus Dikeluarkan Untuk Menempuh Jarak %2.1f = Rp.%3.2f", jarak, total_harga)
}

// Nomor 4
func shioCalender(tahunLahir int) {
	switch tL := tahunLahir; {
	case tL%12 == 0:
		println("Anda Lahir Di Tahun Monyet")
	case tL%12 == 1:
		println("Anda Lahir Di Tahun Ayam")
	case tL%12 == 2:
		println("Anda Lahir Di Tahun Anjing")
	case tL%12 == 3:
		println("Anda Lahir Di Tahun Babi")
	case tL%12 == 4:
		println("Anda Lahir Di Tahun Tikus")
	case tL%12 == 5:
		println("Anda Lahir Di Tahun Kerbau")
	case tL%12 == 6:
		println("Anda Lahir Di Tahun harimau")
	case tL%12 == 7:
		println("Anda Lahir Di Tahun Kelinci")
	case tL%12 == 8:
		println("Anda Lahir Di Tahun Naga")
	case tL%12 == 9:
		println("Anda Lahir Di Tahun Ular")
	case tL%12 == 10:
		println("Anda Lahir Di Tahun Kuda")
	case tL%12 == 11:
		println("Anda Lahir Di Tahun Kambing")
	}
}

// Nomor5
func simulasiATM(sepuluh, lima, satu int) {
	var jmlhUang int
	pin := 328339
	var inputpin int
	fmt.Println("Masukkan Pin ATM Anda : ")
	fmt.Scan(&inputpin)
	if pin == inputpin {
		fmt.Println("Masukkan Jumlah Uang yg ingin anda ambil : [Pecahan : $10, $5, $1]")
		fmt.Scan(&jmlhUang)
		pecahanSepuluh := jmlhUang / sepuluh
		sisa := jmlhUang % 10
		pecahanLima := sisa / lima
		sisa = jmlhUang % 5
		pecahanSatu := sisa / satu

		fmt.Printf("Pecahan : $10 = %d lembar, $5 = %d lembar, $1 = %d lembar", pecahanSepuluh, pecahanLima, pecahanSatu)
	} else {
		println("Kesempatan anda hanya 2 kali untuk mencoba: ")
		fmt.Println("Masukkan Pin ATM Anda : ")
		fmt.Scan(&inputpin)
		if pin == inputpin {
			fmt.Println("Masukkan Jumlah Uang yg ingin anda ambil : [Pecahan : $10, $5, $1]")
			fmt.Scan(&jmlhUang)
			pecahanSepuluh := jmlhUang / sepuluh
			sisa := jmlhUang % 10
			pecahanLima := sisa / lima
			sisa = jmlhUang % 5
			pecahanSatu := sisa / satu

			fmt.Printf("Pecahan : $10 = %d lembar, $5 = %d lembar, $1 = %d lembar", pecahanSepuluh, pecahanLima, pecahanSatu)
		} else {
			println("Kesempatan anda hanya 1 kali untuk mencoba: ")
			fmt.Println("Masukkan Pin ATM Anda : ")
			fmt.Scan(&inputpin)
			if pin == inputpin {
				fmt.Println("Masukkan Jumlah Uang yg ingin anda ambil : [Pecahan : $10, $5, $1]")
				fmt.Scan(&jmlhUang)
				pecahanSepuluh := jmlhUang / sepuluh
				sisa := jmlhUang % 10
				pecahanLima := sisa / lima
				sisa = jmlhUang % 5
				pecahanSatu := sisa / satu

				fmt.Printf("Pecahan : $10 = %d lembar, $5 = %d lembar, $1 = %d lembar", pecahanSepuluh, pecahanLima, pecahanSatu)
			} else {
				println("Pin Anda Di Lock")
			}
		}
	}
}

// Nomor 6
func tebakhadiah(userinput int) {
	var angkaUser, sisaUser, angkaCom, sisaCom int

	angkaUser = userinput / 10
	sisaUser = userinput % 10

	min := 10
	max := 99
	nilaiRandom := rand.Intn(max-min) + min
	angkaCom = nilaiRandom / 10
	sisaCom = nilaiRandom % 10

	if userinput < 10 || userinput > 99 {
		fmt.Printf("Input Harus berada pada range 10 - 99")
	} else if userinput == nilaiRandom {
		fmt.Println("User Input :", userinput)
		fmt.Println("Komputer random :", nilaiRandom)
		fmt.Printf("Match All Digit : You Win Rp 100.000")
	} else if angkaUser == angkaCom || angkaUser == sisaCom || sisaUser == angkaCom || sisaUser == sisaCom {
		fmt.Println("User Input :", userinput)
		fmt.Println("Komputer random :", nilaiRandom)
		fmt.Printf("Exact match : You win Rp 50.000")
	} else {
		fmt.Println("User Input :", userinput)
		fmt.Println("Komputer random :", nilaiRandom)
		fmt.Println("You Lose !")
	}
}
