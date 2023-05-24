package main

import (
	"fmt"
	"log"
	"math/rand"
)

func addOrEvent(number int) {
	if number%2 == 0 {
		println("Event Number")
	} else {
		println("Add Number")
	}
}

//fizzbuz
/*
	number = bisa dibagi 3 dan 5 -> fizzbuzz
	number = dibagi 3 -> fizz
	number = dibagi 5 -> buzz
	number = ga bisa dibagi 3&5 -> unknown
*/
func fizzBuzz(number int) {
	if number%3 == 0 && number%5 == 0 {
		println("fizzbuzz")
	} else if number%3 == 0 {
		println("fizz")
	} else if number%5 == 0 {
		println("buss")
	} else {
		println("unknown")
	}
}

func randomNumber() {
	random := rand.Intn(10) //range 0-9
	println(random)

	//range output : 1-10
	min := 100
	max := 200
	randomRange := rand.Intn(max-min) + min
	println("RandomMin:", randomRange)
}

func initialif() {
	var input int
	fmt.Println("Masukkan nilai random: ")
	fmt.Scanln(&input)

	// random := rand.Intn(10)
	if random := rand.Intn(input); random < 5 {
		println("low")
	} else if random >= 5 {
		println("high")
	}

	// if random < 5 {
	// 	println("low")
	// } else if random >= 5 {
	// 	println("high")
	// }
}

func guessNumber() {
	// println("Berapa Hasil 7-3 = ?")
	angka1 := rand.Intn(10)
	angka2 := rand.Intn(6)
	angka3 := angka1 - angka2

	fmt.Printf("Berapa hasil %d-%d = ", angka1, angka2)
	//declare var jawab u/ nampung inputan dari console
	var jawab int
	_, err := fmt.Scan(&jawab)
	if err != nil {
		log.Fatal(err)
	}
	if jawab == angka3 {
		println("Jawaban yang Benar")
	} else {
		println("Coba lagi, Tetap semangat")
	}
}

func multiArgs() {
	fmt.Println("Input 3 angka : ")
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
		println("Sunday")
	case n == 1:
		println("Monday")
	case n == 2:
		println("Tuesday")
	case n == 3:
		println("Wednesday")
	case n == 4:
		println("Thursday")
	case n == 5:
		println("Friday")
	case n == 0:
		println("Saturday")
	}
}

func main() {
	// addOrEvent(8)
	// fizzBuzz(1)
	// randomNumber()
	// initialif()
	guessNumber()
	// multiArgs()
	// tebakHari(5)
}
