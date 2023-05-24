package main

import (
	"fmt"
	"log"
	"math/rand"
)

func oddOrEvent(number int) {
	if number%2 == 0 {
		println("Event Number")
	} else {
		println("Odd Number")
	}
}

// FizzBuzz
func FizzBuzz(number int) {
	if number%3 == 0 && number%5 == 0 {
		println("FizzBuzz")

	} else if number%3 == 0 {
		println("Fizz")
	} else if number%5 == 0 {
		println("Buzz")
	} else {
		println("Unkown")
	}

}

func randomNumber() {
	random := rand.Intn(10)
	println(random)

	min := 10
	max := 201
	randomNumber := rand.Intn(max-min) + min
	println(randomNumber)
}

func initialif() {
	//random := rand.Intn(10)
	//println(random)
	if random := rand.Intn(10); random < 5 {
		println(random)
		println("low")
	} else {
		println(random)
		println("big")
	}
	/*if random < 5 {
		println("kecil")
	} else if random >= 5 {
		println("besar")
	}*/

}

func guessNumber() {
	angka1 := rand.Intn(10)
	angka2 := rand.Intn(6)
	angka3 := angka1 - angka2
	fmt.Printf("Berapa Hasil dari %d - %d = ", angka1, angka2)

	//print("Berapa hasilnya 7-3? ")
	//untuk menampung jawaban

	var jawab int
	_, err := fmt.Scan(&jawab)
	if err != nil {
		log.Fatal(err)
	}
	if jawab == angka3 {
		println("Jawaban Benar")
	} else {
		println("Jawaban Salah")

	}

}

func multipleArgs() {
	fmt.Println("input 3 angka: ")

	var angka1, angka2, angka3 int
	n, err := fmt.Scan(&angka1, &angka2, &angka3)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Jawaban: ", angka1, angka2, angka3, n)

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
		fmt.Println("Wenesday")
	case n == 4:
		fmt.Println("Thrusday")
	case n == 5:
		fmt.Println("Friday")
	case n == 6:
		fmt.Println("Saturday")
	default:
		fmt.Println("Not Today")
	}

}

func main() {
	//oddOrEvent(7)
	//FizzBuzz(15)
	//randomNumber()
	//initialif()
	//guessNumber()
	//multipleArgs()
	tebakHari(5)
}
