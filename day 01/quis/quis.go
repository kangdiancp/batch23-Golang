package main

import (
	"fmt"
	"log"
	"math/rand"
	"strconv"
)

// nomor 1
func totalLompat(x, y, k int) {
	sisa := (y - x) % k
	total := (y - x) / k
	if sisa != 0 {
		fmt.Println(total + 1)
	} else {
		fmt.Println(total)
	}
}

// nomor 2//
func menitToTahun(menit int) {
	tahun := menit / 525600
	hari := (menit % 525600) / 1440
	fmt.Printf("%d menit = %d tahun,  %d hari", menit, tahun, hari)
	println()
}

// Nomor 3//
func totalJarak() {
	var jarak, bensin, harga float64

	fmt.Print("Jarak yang ditempuh: ")
	fmt.Scan(&jarak)
	fmt.Print("1 Kilo Konsumsi Bensin: ")
	fmt.Scan(&bensin)
	fmt.Print("Harga per liter: ")
	fmt.Scan(&harga)

	total := harga / bensin * jarak

	format := strconv.FormatFloat(total, 'f', 2, 64)
	fmt.Println("Uang yang harus anda keluarkan adalah: ", format)
}

//Nomor 4//

func shioCalender() {
	fmt.Print("Enter tahun lahir : ")
	var tahun int
	_, err := fmt.Scan(&tahun)
	if err != nil {
		log.Fatal(err)
	}

	switch shio := tahun % 12; {
	case shio == 0:
		fmt.Println("Anda lahir di tahun monyet")
	case shio == 1:
		fmt.Println("Anda lahir di tahun ayam")
	case shio == 2:
		fmt.Println("Anda lahir di tahun anjing")
	case shio == 3:
		fmt.Println("Anda lahir di tahun babi")
	case shio == 4:
		fmt.Println("Anda lahir di tahun tikus")
	case shio == 5:
		fmt.Println("Anda lahir di tahun kerbau")
	case shio == 6:
		fmt.Println("Anda lahir di tahun harimau")
	case shio == 7:
		fmt.Println("Anda lahir di tahun kelinci")
	case shio == 8:
		fmt.Println("Anda lahir di tahun naga")
	case shio == 9:
		fmt.Println("Anda lahir di tahun ular")
	case shio == 10:
		fmt.Println("Anda lahir di tahun kuda")
	case shio == 11:
		fmt.Println("Anda lahir di tahun kambing")

	}
}

// nomor 5
func simulasiAtm() {
	fmt.Print("Enter pin : ")
	var pin int
	_, err := fmt.Scan(&pin)
	if err != nil {
		log.Fatal(err)
	}
	pinBenar := 1234
	if pin == pinBenar {
		fmt.Print("Enter jumlah uang : ")
		var jumlah int
		_, err := fmt.Scan(&jumlah)
		if err != nil {
			log.Fatal(err)
		}
		pecahan10 := jumlah / 10
		sisaUang := jumlah % 10
		pecahan5 := sisaUang / 5
		sisaUang = sisaUang % 5
		pecahan1 := sisaUang / 1
		fmt.Printf("$10 = %d lembar, $5 = %d lembar, $1 = %d", pecahan10, pecahan5, pecahan1)
	} else {
		fmt.Print("pin salah (kesempatan 2 kali lagi untuk mencoba) : ")
		var pin int
		_, err := fmt.Scan(&pin)
		if err != nil {
			log.Fatal(err)
		}
		if pin == pinBenar {
			fmt.Print("Enter jumlah uang : ")
			var jumlah int
			_, err := fmt.Scan(&jumlah)
			if err != nil {
				log.Fatal(err)
			}
			pecahan10 := jumlah / 10
			sisaUang := jumlah % 10
			pecahan5 := sisaUang / 5
			sisaUang = sisaUang % 5
			pecahan1 := sisaUang / 1

			fmt.Printf("$10 = %d lembar, $5 = %d lembar, $1 = %d", pecahan10, pecahan5, pecahan1)
		} else {
			fmt.Print("pin salah (kesempatan 3 kali lagi untuk mencoba) : ")
			var pin int
			_, err := fmt.Scan(&pin)
			if err != nil {
				log.Fatal(err)
			}
			if pin == pinBenar {
				fmt.Print("Enter jumlah uang : ")
				var jumlah int
				_, err := fmt.Scan(&jumlah)
				if err != nil {
					log.Fatal(err)
				}
				pecahan10 := jumlah / 10
				sisaUang := jumlah % 10
				pecahan5 := sisaUang / 5
				sisaUang = sisaUang % 5
				pecahan1 := sisaUang / 1

				fmt.Printf("$10 = %d lembar, $5 = %d lembar, $1 = %d", pecahan10, pecahan5, pecahan1)
			} else {
				fmt.Println("pin anda di lock")
			}
		}
	}
}

// nomor 6
func tebakHadiah() {
	var angka int
	fmt.Print("Inputkan Angka: ")
	fmt.Scan(&angka)

	angka1 := angka / 10
	sisa := angka % 10
	angka2 := sisa % 10

	computer := rand.Intn(100) + 11
	fmt.Println("Jackpot: ", computer)
	computer1 := computer / 10
	sisacom := computer % 10
	computer2 := sisacom % 10

	angkaa1 := strconv.Itoa(angka1)
	angkaa2 := strconv.Itoa(angka2)
	angkatotal := angkaa1 + angkaa2
	computerr := strconv.Itoa(computer)
	computerr1 := strconv.Itoa(computer1)
	computerr2 := strconv.Itoa(computer2)

	if computerr == angkatotal {
		println("Match All Digit : You Win Rp. 100.000")
	} else if computerr1 == angkaa1 {
		println("Exact Match : You Win Rp. 50.000")
	} else if computerr2 == angkaa2 {
		println("Exact Match : You Win Rp. 50.000")
	} else {
		println("You Lose, Coba Lagi Ya")
	}
}

func main() {
	// totalLompat(0, 120, 30)
	// menitToTahun(1000000000)
	// totalJarak()
	// shioCalender()
	simulasiAtm()
	// tebakHadiah()

}
