package main

import (
	"fmt"
	"log"
	"math/rand"
	"strconv"
)

func soal1(number float64) {
	var luas_lingkaran float64
	phi := 3.14159
    luas_lingkaran = phi * number * number
	fmt.Printf("Luas lingkaran =  %8.3f \n", luas_lingkaran)
}

func soal2(sudut1 int, sudut2 int) {
	var sudut3 int
	sudut3 = 180 - (sudut1 + sudut2)
	println("sudut 1 = ", sudut1)
	println("sudut 2 = ", sudut2)
	fmt.Println("Jumlah sudut segitiga  : ", sudut3)
}

func soal3(N1 float64, N2 float64, N3 float64){
	var Average float64
	Average = (N1 + N2 + N3) / 3
	fmt.Printf("Average : %8.3f \n", Average)
}

func soal4(tahun int){
	if tahun % 4 == 0 && (tahun % 400 == 0 || tahun % 100 != 0) {
		fmt.Printf("Tahun %d tahun kabisat \n", tahun)
	}else{
		fmt.Printf("Tahun %d bukan tahun kabisat \n", tahun)
		println("Fuck")
	}
}

func soal5(second int){
	menit := second / 60
	detik := second % 60
	fmt.Printf("%d minutes %d seconds \n", menit, detik)
}

func soal6secondToDays(detik int){
	hari := detik / 86400
	jam :=  (detik % 86400) / 3600
	menit := ((detik % 86400) % 3600) / 60
	second := detik % 60
	fmt.Printf("%d Hari %d jam %d menit %d detik \n", hari, jam, menit, second)
}

func soal7ReverseNumber(rev int){
	if rev >= 1000 && rev < 10000{
		i := 0
		for rev > 0{
			count := rev % 10
			i = (i * 10) + count
			rev /= 10
		}

		fmt.Printf("Reverse Number : %d ", i)
	}else{
		fmt.Println(" Input Number Harus Dalam Range 1000 - 10000")
	}
}

func soal7ReverseNumber2(number int){
	var angka1, angka2, angka3, angka4, sisa int
	if number < 1000 || number >= 10000{
		fmt.Printf("Input Number Harus Dalam Range 1000 - 10000")
	}else{
		angka1 = number / 1000
		sisa = number % 1000

		angka2 = sisa / 100
		sisa = sisa % 100

		angka3 = sisa / 10
		angka4 = sisa % 10

		reverse := strconv.Itoa(angka4) + strconv.Itoa(angka3) + strconv.Itoa(angka2) + strconv.Itoa(angka1)
		fmt.Printf("\n Reverse : %d -> %s", number, reverse)
	}
}

func soal8SumNumber(number int){
	var angka1, angka2, angka3, angka4, sisa int
	if number < 1000 || number >= 10000{
		fmt.Printf("\n Input Number Harus Dalam Range 1000 - 10000")
	}else{
		angka1 = number / 1000
		sisa = number % 1000

		angka2 = sisa / 100
		sisa = sisa % 100

		angka3 = sisa / 10
		angka4 = sisa % 10

		sum := angka1 + angka2 + angka3 + angka4
		fmt.Printf("\n Sum Number : %d -> %d", number, sum)
	}
}


func soal9SortTigaAngka(){
	var angka1, angka2, angka3 int
	fmt.Printf("Masukkan 3 angka : \n")
	_,err := fmt.Scan(&angka1, &angka2, &angka3)

	if angka1 < angka2 && angka1 < angka3 && angka2 < angka3 {
		fmt.Printf("%d , %d, %d", angka1, angka2, angka3)
	} else if angka1 > angka2 && angka1 < angka3 && angka2 < angka3{
		fmt.Printf("%d, %d, %d", angka2, angka1, angka3)
	} else if angka1 > angka2 && angka1 > angka3 && angka2 < angka3{
		fmt.Printf("%d, %d, %d", angka2, angka3, angka1)
	} else if angka1 < angka2 && angka1 > angka3 && angka2 > angka3 {
		fmt.Printf("%d, %d, %d", angka3, angka1, angka2)
	} else if angka1 > angka2 && angka1 > angka3 && angka2 > angka3 {
		fmt.Printf("%d, %d, %d", angka3, angka2, angka1)
	} else if angka1 < angka2 && angka1 < angka3 && angka2 > angka3 {
		fmt.Printf("%d, %d, %d", angka1, angka3, angka2)
	} else if err != nil{
		log.Fatal(err)
	}
}

func soal10GameGunting(){
	var pilihan int
	fmt.Printf("Pilih : (1)Gunting  (2)Batu  (3)Kertas\n")
	fmt.Printf("Masukkan Pilihan Anda : ")
	p,err := fmt.Scan(&pilihan)

	pil1 := "Gunting"
	pil2 := "Batu"
	pil3 := "Kertas"
	min := 1
	max := 3
	computer := rand.Intn(max-min)+ min
	if p == 1 && computer == 1 {
		fmt.Printf("Anda Seri, computer [%s] vs Kamu [%s]\n", pil1, pil1)
	} else if p == 1 && computer == 2 {
		fmt.Printf("Anda Kalah, computer [%s] vs Kamu [%s]\n", pil1, pil2)
	} else if p == 1 && computer == 3 {
		fmt.Printf("Anda Menang, computer [%s] vs Kamu [%s]\n", pil1, pil3)
	} else if p == 2 && computer == 1 {
		fmt.Printf("Anda Menang, computer [%s] vs Kamu [%s]\n", pil2, pil1)
	} else if p == 2 && computer == 2 {
		fmt.Printf("Anda Seri, computer [%s] vs Kamu [%s]\n", pil2, pil2)
	} else if p == 2 && computer == 3 {
		fmt.Printf("Anda Kalah, computer [%s] vs Kamu [%s]\n", pil2, pil3)
	} else if p == 3 && computer == 1 {
		fmt.Printf("Anda Kalah, computer [%s] vs Kamu [%s]\n", pil3, pil1)
	} else if p == 3 && computer == 2 {
		fmt.Printf("Anda Menang, computer [%s] vs Kamu [%s]\n", pil3, pil2)
	} else if p == 3 && computer == 3 {
		fmt.Printf("Anda Seri, computer [%s] vs Kamu [%s]\n", pil3, pil3)
	}else if err != nil{
		log.Fatal(err)
	}
}

func main() {
	//soal1(10)
	//soal2(30, 60)
	//soal3(15, 15, 25)
	//soal3(12.5, 9.7, 6.98)
	//soal4(2012)
	//soal4(2002)
	//soal5(60)
	//soal5(12345)
	//soal6secondToDays(1000000)
	//soal7ReverseNumber(1234)
	//soal7ReverseNumber2(12)
	soal8SumNumber(1234)
	soal8SumNumber(4545)
	soal8SumNumber(55000)
	soal9SortTigaAngka()
	soal10GameGunting()

}