package main

import "fmt"

// Nomor 1: Menghitung luas lingkaran
func luasLingkaran(jariJari float64) {
	luas := 3.14 * jariJari * jariJari
	fmt.Println("Luas lingkaran:", luas)
}

// Nomor 2: Menghitung sudut segitiga terakhir
func sudutSegitiga(sudut1, sudut2 int) {
	sudut3 := 180 - sudut1 - sudut2
	fmt.Println("sudut1:", sudut1)
	fmt.Println("sudut2:", sudut2)
	fmt.Println("sudut3:", sudut3)
}

// Nomor 3: Menghitung nilai rata-rata
func computeAverage(nilai1, nilai2, nilai3 float64) {
	rataRata := (nilai1 + nilai2 + nilai3) / 3
	fmt.Println("Average:", rataRata)
}

// Nomor 4: Menentukan tahun kabisat
func tahunKabisat(tahun int) {
	if tahun%4 == 0 && (tahun%100 != 0 || tahun%400 == 0) {
		fmt.Println("Tahun", tahun, "adalah tahun kabisat")
	} else {
		fmt.Println("Tahun", tahun, "bukan tahun kabisat")
	}
}

func main() {
	// Memanggil fungsi-fungsi dengan nilai (value) yang spesifik
	sudutSegitiga(30, 60)
	luasLingkaran(10)
	computeAverage(15, 15, 25)
	computeAverage(12.5, 9.7, 6.98)
	tahunKabisat(2002)
	tahunKabisat(2012)
}
