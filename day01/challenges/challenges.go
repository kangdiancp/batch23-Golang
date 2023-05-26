package main

import (
	"fmt"
	"math/rand"
	"strconv"
	"strings"
	"unicode"
)

const phi = 3.14159

func main() {
	//Soal1
	// var inputLuasLingkaran float32
	// //Soal2
	// var sudut1, sudut2 int
	// //Soal3
	// var av1, av2, av3 float32
	// //Soal4
	// var tahun int

	// //Soal1
	// fmt.Println("Masukkan Nilai lingkaran : ")
	// fmt.Scan(&inputLuasLingkaran)
	// luasLingkaran(inputLuasLingkaran)

	// //Soal2
	// fmt.Println("Masukkan Sudut Segitiga : ")
	// fmt.Scan(&sudut1, &sudut2)
	// sudutTerakhir(sudut1, sudut2)

	// //Soal3
	// fmt.Println("Masukkan inputan yang akan diratakan : ")
	// fmt.Scan(&av1, &av2, &av3)
	// computerAverage(av1, av2, av3)

	// //Soal4
	// fmt.Println("Masukkan Tahun yang akan dicek : ")
	// fmt.Scan(&tahun)
	// tahunKabisat(tahun)

	//Soal 5
	// displayTime()

	//Soal 6
	// secondsToDays(1000000)

	//Soal 7
	// reverseNumber(1234)
	//Soal 8
	// var game int
	// fmt.Println("Masukkan Pilihan anda : (1)Gunting (2)Batu (3)Kertas")
	// fmt.Scan(&game)
	// gameGunting(game)

	//Soal 9
	// var angka1, angka2, angka3 int
	// fmt.Println("Masukkan 3 angka yg akan disorting : ")
	// fmt.Scan(&angka1, &angka2, &angka3)
	// sortTigaAngka(angka1, angka2, angka3)

	//Soal 10
	// sumNumber(1234)

	//Soal11
	// findWord("bootcamp Go di codeid, go run Go and go", "go")
	//Soal12
	replaceMiddleCharWithStar("aku suka join codeid bootcamp")

	//Soal 13
	// countVocalConsonant("Rockbye Baby")

	//Soal 14
	// countWordWithUpper("HowManyWordInThisgueSentence")

	//Soal 15
	// countWordWithSpace(" How Many Word In This Sentence yudha ")
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

//Nomor 5
// Buat function displyaTime dalam menit dan detik, dengan inputan dalam detik
// contoh : displayMinuteSec(60)
// output : Result : 1 minutes 0 seconds

// contoh : displayMinuteSec(12345)
// output : Result : 205 Minutes 45 Seconds

func displayTime() {
	displayMinuteSec := 12345
	var minute int
	var second int

	//Cara 1
	println("Cara 1")
	if displayMinuteSec%60 == 0 {
		minute = displayMinuteSec / 60
		fmt.Printf("Result : %d minutes 0 seconds", minute)
	} else if displayMinuteSec%60 != 0 {
		minute = displayMinuteSec / 60
		second = displayMinuteSec % 60
		fmt.Printf("Result : %d minutes %d seconds ", minute, second)
		println()
	}

	//Cara 2
	println("Cara 2")
	minute = displayMinuteSec / 60
	second = displayMinuteSec % 60
	fmt.Printf("Result : %d minutes %d seconds", minute, second)
	println()
}

func secondsToDays(number int) {
	hari := number / 86400
	modhari := number % 86400
	jam := modhari / 3600
	modjam := modhari % 3600
	menit := modjam / 60
	detik := modjam % 60

	fmt.Printf("Result : %d hari %d jam %d minutes %d seconds", hari, jam, menit, detik)
}

func reverseNumber(number int) {
	var angka1, angka2, angka3, angka4, sisa int

	if number >= 1000 && number < 10000 {
		angka1 = number / 1000
		sisa = number % 1000
		angka2 = sisa / 100
		sisa = sisa % 100
		angka3 = sisa / 10
		angka4 = sisa % 10

		reversed := strconv.Itoa(angka4) + strconv.Itoa(angka3) + strconv.Itoa(angka2) + strconv.Itoa(angka1)
		fmt.Printf("Reverse : %d --> %s", number, reversed)

	} else if number < 1000 || number >= 10000 {
		println("Input Number harus dalam range 1000 - 10000")
	}

}

//nomor 8

func gameGunting(inputPil int) {
	var hasilComp string
	pil1 := "Gunting"
	pil2 := "Batu"
	pil3 := "Kertas"

	compPil := rand.Intn(3)

	if compPil == 1 {
		hasilComp = "Gunting"
		if compPil == 1 && inputPil == 1 {
			fmt.Printf("Anda Seri")
		} else if compPil == 1 && inputPil == 2 {
			fmt.Printf("Anda Menang, computer[%s] vs kamu [%v]", hasilComp, pil2)
		} else if compPil == 1 && inputPil == 3 {
			fmt.Printf("Anda Kalah, computer[%s] vs kamu [%v]", hasilComp, pil3)
		}
	} else if compPil == 2 {
		hasilComp = "Batu"
		if compPil == 2 && inputPil == 2 {
			fmt.Printf("Anda Seri")
		} else if compPil == 2 && inputPil == 1 {
			fmt.Printf("Anda Menang, computer[%s] vs kamu [%v]", hasilComp, pil1)
		} else if compPil == 2 && inputPil == 3 {
			fmt.Printf("Anda Kalah computer[%s] vs kamu [%v]", hasilComp, pil3)
		}
	} else {
		hasilComp = "Kertas"
		if compPil == 3 && inputPil == 3 {
			fmt.Printf("Anda Seri")
		} else if compPil == 3 && inputPil == 1 {
			fmt.Printf("Anda Menang, computer[%s] vs kamu [%v]", hasilComp, pil1)
		} else if compPil == 3 && inputPil == 2 {
			fmt.Printf("Anda Kalah computer[%s] vs kamu [%v]", hasilComp, pil3)
		}
	}
}

// nomor 9
func sortTigaAngka(angka1, angka2, angka3 int) {
	if angka1 < angka2 && angka1 < angka3 && angka2 < angka3 {
		fmt.Printf("%d , %d, %d", angka1, angka2, angka3)
	} else if angka1 > angka2 && angka1 < angka3 && angka2 < angka3 {
		fmt.Printf("%d, %d, %d", angka2, angka1, angka3)
	} else if angka1 > angka2 && angka1 > angka3 && angka2 < angka3 {
		fmt.Printf("%d, %d, %d", angka2, angka3, angka1)
	} else if angka1 < angka2 && angka1 > angka3 && angka2 > angka3 {
		fmt.Printf("%d, %d, %d", angka3, angka1, angka2)
	} else if angka1 > angka2 && angka1 > angka3 && angka2 > angka3 {
		fmt.Printf("%d, %d, %d", angka3, angka2, angka1)
	} else if angka1 < angka2 && angka1 < angka3 && angka2 > angka3 {
		fmt.Printf("%d, %d, %d", angka1, angka3, angka2)
	}
}

// Nomor 10
func sumNumber(nilai int) {
	var angka1, angka2, angka3, angka4, sisa int
	if nilai < 1000 || nilai >= 10000 {
		fmt.Printf("\n Input Number Harus Dalam Range 1000 - 10000")
	} else {
		angka1 = nilai / 1000
		sisa = nilai % 1000

		angka2 = sisa / 100
		sisa = sisa % 100

		angka3 = sisa / 10
		angka4 = sisa % 10

		sum := angka1 + angka2 + angka3 + angka4
		fmt.Printf("\n Sum Number : %d -> %d", nilai, sum)
	}
}

//Nomor11

func findWord(words, search string) {
	// countWords := 0
	countGo := 0
	//case 1
	ubahWord := strings.ToLower(words)
	//untuk menghitung jumlah go
	foundGo := strings.Count(ubahWord, search)
	fmt.Printf("Word go found %d times", foundGo)
	println()

	//case2
	sliceWordtospace := strings.Split(ubahWord, " ")
	countWords := len(sliceWordtospace)
	for i := 0; i < countWords; i++ {
		if sliceWordtospace[i] == search {
			countGo++
		}
	}
	fmt.Printf("Word go found %d times - case 2", countGo)
	println()
}

// Nomor 12
func replaceMiddleCharWithStar(word string) {
	sliceWordfromSpace := strings.Split(word, " ")
	countWords := len(sliceWordfromSpace)
	output := ""

	for i := 0; i < countWords; i++ {
		var hasil string
		word := sliceWordfromSpace[i]
		if len(word) > 2 {
			sliceMiddle := word[1 : len(word)-1]
			// println(sliceMiddle)
			replaceStarinMiddle := strings.Repeat("*", len(sliceMiddle))
			hasil = string(word[0]) + replaceStarinMiddle + string(word[len(word)-1])
			// fmt.Println("Result : ", hasil)

		}
		output += hasil + " "
	}
	fmt.Println("Result : ", output)
}

// Nomor 13
func countVocalConsonant(word string) {
	countVocal := 0
	countConsonant := 0
	for _, value := range word {
		switch value {
		case 'a', 'i', 'u', 'e', 'o':
			countVocal++
		default:
			countConsonant++
			if unicode.IsSpace(value) {
				countConsonant--
			}
		}
	}
	fmt.Printf("Vocal : %d, Consonant : %d", countVocal, countConsonant)
	println()
}

// Nomor 14
func countWordWithUpper(word string) {
	countUpper := 0

	for _, value := range word {
		if unicode.IsUpper(value) {
			countUpper++
		}
	}
	fmt.Printf("Total Word : %d", countUpper)
	println()
}

// Nomor 15
func countWordWithSpace(word string) {
	countSpace := 0
	//var countWordWithoutSpace int
	countWords := len(strings.Fields(word))
	println(countWords)

	for _, value := range word {
		if unicode.IsSpace(value) {
			countSpace++
			countWordWithoutSpace := countWords - countSpace
			countSpace += countWordWithoutSpace
		}
	}
	fmt.Printf("Total Word : %d", countSpace)
}
