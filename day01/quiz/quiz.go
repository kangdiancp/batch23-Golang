package main

import (
	"fmt"
	"log"
	"math/rand"
	"strconv"
)

// nomor 1
func totalLompat(x, y, k int) {
	//x = posisi awal, y = target posisi, k = jarak tempuh satu kali lompat
	sisa := (y - x) % k
	total := (y - x) / k
	if sisa != 0 {
		fmt.Println(total + 1)
	} else {
		fmt.Println(total)
	}
}

// nomor 2
func menitToTahun(menit int) {
	tahun := menit / 525600
	hari := (menit % 525600) / 1440
	fmt.Printf(" %d menit = %d tahun,  %d hari", menit, tahun, hari)
}

// nomor 3
func totalJarak() {
	var jarak, bensin, harga float64

	fmt.Println("Jarak yang ditempuh: ")
	fmt.Scan(&jarak)
	fmt.Println("1 Kilo Konsumsi Bensin: ")
	fmt.Scan(&bensin)
	fmt.Println("Harga per liter: ")
	fmt.Scan(&harga)

	total := harga / bensin * jarak

	format := strconv.FormatFloat(total, 'f', 2, 64)
	fmt.Println("Uang yang harus anda keluarkan adalah: ", format)
}

// nomor 4
func shioCalender() {
	fmt.Println("Enter tahun lahir : ")
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
	fmt.Print("Masukkan angka : ")
	_, err := fmt.Scan(&angka)
	if err != nil {
		log.Fatal(err)
	}

	angka1 := angka / 10
	sisa := angka % 10
	angka2 := sisa % 10

	random := rand.Intn(100) + 11
	random1 := random / 10
	sisa_random := random % 10
	random2 := sisa_random % 10

	angka_satu := strconv.Itoa(angka1)
	angka_dua := strconv.Itoa(angka2)
	angka_tiga := angka_satu + angka_dua
	randomm := strconv.Itoa(random)
	randomm1 := strconv.Itoa(random1)
	randomm2 := strconv.Itoa(random2)

	if randomm == angka_tiga {
		println("User input : ", angka)
		println("Komputer random : ", randomm)
		println("Match all digit : you win Rp.100.000")
	} else if randomm1 == angka_satu {
		println("User input : ", angka)
		println("Komputer random : ", randomm)
		println("Exact match : you win Rp.50.000")
	} else if randomm2 == angka_dua {
		println("User input : ", angka)
		println("Komputer random : ", randomm)
		println("Exact match : you win Rp.50.000")
	} else {
		println("User input : ", angka)
		println("Komputer random : ", randomm)
		println("You Lose!")
	}
}

func main() {
	totalLompat(10, 85, 30)
	menitToTahun(1_000_000_000)
	totalJarak()
	shioCalender()
	simulasiAtm()
	tebakHadiah()
}
