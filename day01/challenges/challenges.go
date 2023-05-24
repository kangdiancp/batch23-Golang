package main

import (
	"fmt"
)

const phi = 3.14159

func main() {
	//Soal1
	var inputLuasLingkaran float32
	//Soal2
	var sudut1, sudut2 int
	//Soal3
	var av1, av2, av3 float32
	//Soal4
	var tahun int

	//Soal1
	fmt.Println("Masukkan Nilai lingkaran : ")
	fmt.Scan(&inputLuasLingkaran)
	luasLingkaran(inputLuasLingkaran)

	//Soal2
	fmt.Println("Masukkan Sudut Segitiga : ")
	fmt.Scan(&sudut1, &sudut2)
	sudutTerakhir(sudut1, sudut2)

	//Soal3
	fmt.Println("Masukkan inputan yang akan diratakan : ")
	fmt.Scan(&av1, &av2, &av3)
	computerAverage(av1, av2, av3)

	//Soal4
	fmt.Println("Masukkan Tahun yang akan dicek : ")
	fmt.Scan(&tahun)
	tahunKabisat(tahun)
}

// Nomor 1
func luasLingkaran(nilai float32) {
	hasil := (nilai * nilai) * phi

	fmt.Printf("Output Luas Lingkaran : %8.2f", hasil)
	println()
}

// Nomor2

// hitung sudut segitiga terakhir :
// contoh : suduhSegitiga(30, 60)
// output :
//
//	sudut1 : 30
//	sudut2 : 60
//	sudut3 : 90
func sudutTerakhir(nilai1, nilai2 int) {
	totalsudut := 180
	output := totalsudut - (nilai1 + nilai2)
	println("Maka nilai ketiga sudut : ", nilai1, nilai2, output)
}

//Nomor3

// Hitung nilai rata2 dari 3 inputan untuk function computeraverage
// contoh : computerAverage(15,15,25)
// output : average : 18.333

// contoh : computerAverage(12.5, 9.7, 6.98)
// output : average : 9.727

func computerAverage(average1, average2, average3 float32) {
	// n, err := fmt.Scan(average1, average2, average3)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	output := (average1 + average2 + average3) / 3

	fmt.Printf("Average : %8.3f", output)
	println()
}

//Nomor4
// buat function tahunKabisat dengan inputan integer
// contoh : tahunKabisat(2002)
// output : Tahun 2002 bukan tahun tahunKabisat

// contoh : tahunKabisat(2012)
// output : tahun 2012

func tahunKabisat(cekTahun int) {
	if cekTahun%4 == 0 && cekTahun%100 != 0 && cekTahun%400 != 0 {
		fmt.Printf("Tahun %d tahun kabisat", cekTahun)
	} else {
		fmt.Printf("Tahun %d bukan tahun kabisat", cekTahun)
	}
}
