package main

import (
	"fmt"
	"log"
	"math/rand"
)

func oddOrEven(number int) {
	if number%2 == 0 {
		println("Even number")
	} else {
		println("Odd number")
	}
}

// fizzbuzz
/*
	number = bisa dibagi 3 & 5 -> FizzBuzz
	number = bisa dibagi 3 -> Fizz
	number = bisa dibagi 5 -> Buzz
	number = ga bisa dibagi 3 & 5 -> unknown
*/
func fizzBuzz(number int) {
	if number%3 == 0 && number%5 == 0 {
		println("FizzBuzz")
	} else if number%3 == 0 {
		println("Fizz")
	} else if number%5 == 0 {
		println("Buzz")
	} else {
		println("Unknown")
	}
}

func randomNumber() {
	//range output  : 0 - 9
	random := rand.Intn(10)
	println(random)

	//range output : 10 - 20
	min := 10
	max := 20
	randomRange := rand.Intn(max-min) + min
	println(randomRange)
}

func initialIf() {
	// random := rand.Intn(10)

	if random := rand.Intn(10); random < 5 {
		println(random, "low")
	} else if random >= 5 {
		println(random, "high")
	}

	// if random < 5 {
	// 	println("low")
	// } else if random >= 5 {
	// 	println("high")
	// }
}

func guessNumber() {
	angka1 := rand.Intn(10)
	angka2 := rand.Intn(6)
	angka3 := angka1 - angka2

	fmt.Printf("Berapa hasil dari %d - %d = ? ", angka1, angka2)

	// println("Berapa hasil dari 7 -3 =?")

	//declare variable jawab untuk nampung inputan dari console
	var jawab int
	_, err := fmt.Scan(&jawab)

	if err != nil {
		log.Fatal(err)
	}
	// println("Jawaban : ", jawab)

	if jawab == angka3 {
		println("Jawaban Anda Benar")
	} else {
		println("Jawaban Anda Salah")
	}

}

func multipleArgs() {
	fmt.Println("Input 3 angka : ")
	var angka1, angka2, angka3 int
	n, err := fmt.Scan(&angka1, &angka2, &angka3)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("jawaban : ", angka1, angka2, angka3, n)
}

func main() {
	oddOrEven(8)
	fizzBuzz(7)
	randomNumber()
	initialIf()
	guessNumber()
	multipleArgs()
}
