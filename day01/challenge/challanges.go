package main

import (
	"fmt"
	"strconv"
)

func luasLingkaran(jari int) {
	const phi = 3.14159
	var jarinya = float64(jari)
	var luas float64 = jarinya * jarinya * phi
	format := strconv.FormatFloat(luas, 'f', 3, 64)
	fmt.Println("luas Lingkaran: ", format)

}

func sudutSegitiga(sudut1 int, sudut2 int) {

	var sudut3 int

	sudut3 = 180 - (sudut1 + sudut2)
	fmt.Printf("sudut1: %d \n", sudut1)
	fmt.Printf("sudut2: %d \n", sudut2)
	fmt.Printf("sudut3: %d \n", sudut3)

}

func computeAverege(n1 float64, n2 float64, n3 float64) {
	var average float64
	average = (n1 + n2 + n3) / 3
	fmt.Printf("Average: %8.3f", average)
}

func tahunKabisat(th int) {
	if th%4 == 0 && (th%400 == 0 || th%100 != 0) {
		fmt.Printf("Tahun %d tahun Kabisat", th)
	} else {
		fmt.Printf("Tahun %d bukan tahun Kabisat", th)

	}

}

func main() {
	//luasLingkaran(10)
	//sudutSegitiga(30, 60)
	//computeAverege(15, 15, 25)
	//computeAverege(12.5, 9.7, 6.98)
	tahunKabisat(2200)
}
