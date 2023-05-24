package main

import (
	"fmt"
	"log"
	"math/rand"
)

func oddOrVen(number int) {

	if number%2 == 00 {
		println("even Number")
	} else {

		println("odd number")

	}
}

// fizzbuzz
/*
number= bisa dibagi 3 & 5 -> fizzbuzz
number= bisa dibagi 3 -> fizz
number= bisa dibagi 5 => buzz
number = gabisa dbagi 3 & 5 -> unknown
*/

func fizzBuz(number int) {

	if number%3 == 0 && number%5 == 0 {
		println("fizz")

	} else if number%5 == 0 {
		println("buzz")
	} else {
		println("unknown")
	}
}

func randomNumber() {
	// range output : 0-9
	random := rand.Intn(10)
	println(random)

	// range output : 1-10
	min := 10
	max := 20
	randomRange := rand.Intn(max-min) + min
	println("randomMin:", randomRange)
}

func initialif() {
	// random := rand.Intn(10)
	if random := rand.Intn(10); random < 5 {

		println("low")
	} else if random >= 5 {
		println("high")

	}
}

func guessNumber() {
	angka1 := rand.Intn(10)
	angka2 := rand.Intn(6)
	angka3 := angka1 - angka2
	fmt.Printf("berapa hasil %d-%d = ?", angka1, angka2)
	//println("berapa hasil 7-3=?")

	var jawab int
	_, err := fmt.Scan(&jawab)
	if err != nil {
		log.Fatal(err)
	}

	if jawab == angka3 {
		println("jawaban anda bener")
	} else {
		println("jawaban anda salah")

	}
	//println("jawaban :", jawab)
}

func multipleArgs() {
	fmt.Println("input 3 angka :")
	var angka1, angka2, angka3 int
	n, err := fmt.Scan(&angka1, &angka2, &angka3)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("jawaban : ", angka1, angka2, angka3, n)
}

func tebakHari(day int) {
	switch n := day; {
	case n == 0:
		fmt.Println("sunday")
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
	case n == 6:
		fmt.Println("Saturday")

	default:
		fmt.Println("not today ")
	}
}

func main() {
	oddOrVen(8)
	fizzBuz(30)
	randomNumber()
	initialif()
	initialif()
	guessNumber()
	multipleArgs()
	tebakHari(5)
}
