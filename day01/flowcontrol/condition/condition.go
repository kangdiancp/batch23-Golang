package main

import (
	"fmt" //Package ini bisa digunakan untuk memanggil perintah scanln untuk menginput nilai dari keyboard
	"log"

	"math/rand" // package untuk memanggil library sepertin intn dll
)

func main() {
	var nilai int
	//
	oddOrEven(8)

	//Mencari nilai fizzbuzz
	fmt.Println("Masukkan Nilai FizzBuzz Anda : ")
	fmt.Scanln(&nilai)
	fizzBuzz(nilai)

	//memberikan nilai random
	randomNumber()

	initialIf()

	guessNumber()

	multiperArgs()

	tebakHari(0)

}

func oddOrEven(number int) {
	if number%2 == 0 {
		println("Even Number")
	} else {
		println("Odd Number")
	}
}

func fizzBuzz(number int) {
	if number%3 == 0 && number%5 == 0 {
		println("FizzBuzz")
	} else if number%3 == 0 {
		println("Fizz")
	} else {
		println("Buzz")
	}
}

func randomNumber() {
	//range output 0 - 9
	random := rand.Intn(10)
	println("random : ", random)

	min := 100
	max := 200
	randomRange := rand.Intn(max-min) + min
	println("randomOutput : ", randomRange)
}

func initialIf() {
	var inputRand int
	fmt.Print("Masukkan Nilai Random : ")
	fmt.Scanln(&inputRand)

	if random := rand.Intn(inputRand); random > 10 {
		println("High")
	} else if random < 10 {
		println("Low")
	}
}

func guessNumber() {
	println("Berapa Hasil 7 - 3 = ?")

	var jawab int
	_, err := fmt.Scan(&jawab)

	if err != nil {
		log.Fatal(err)
	}

	println("Jawab : ", jawab)

	if jawab == 4 {
		println("Benar")
	} else {
		print("Coba Kembali")
	}

	//Pakai Inputan serta parameter nilai nya berdasarkan nilai random
	var hasil int
	angka1 := rand.Intn(10)
	angka2 := rand.Intn(6)

	angkaPengurangan := angka1 - angka2

	fmt.Printf("Berapa hasil %d - %d = ? ", angka1, angka2)
	fmt.Scan(&hasil)

	if hasil == angkaPengurangan {
		println("Benar")
	} else {
		println("Coba Lagi")
	}

}

func multiperArgs() {
	fmt.Println("Input 3 Angka : ")
	var angka1, angka2, angka3 int

	n, err := fmt.Scan(&angka1, &angka2, &angka3)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Jawaban : ", angka1, angka2, angka3, n)
}

func tebakHari(day int) {
	switch n := day; {
	case n == 0:
		fmt.Println("Sunday")
	case n == 1:
		fmt.Println("Monday")
	case n == 2:
		fmt.Println("Tuesday")
	case n == 3:
		fmt.Println("Wednesday")
	case n == 4:
		fmt.Println("Thursday")
	case n == 5:
		fmt.Println("Friday")
	case n == 1:
		fmt.Println("Saturday")
	default:
		fmt.Println("Not Today")
	}
}
