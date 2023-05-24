package main

import (
	"fmt"
	"log"
	"math/rand"
)

func oddOrEven(number int) {
	if number%2 == 0 {
		println("Even Number")
	} else {
		println("Odd Number")
	}
}

// Fizz Buzz
/*
	number % 3 == 0 --> Fizz
	number % 5 == 0 --> Buzz
	number % 3 == 0 && number % 5 == 0 --> FizzBuzz
	number % 3 != 0 && number % 5 != 0 --> Unknown
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
	// range output = 0 - 9
	random := rand.Intn(10)
	println(random)

	// range output = 1 - 10
	min := 1
	max := 10
	randomRange := rand.Intn(max-min) + min
	println(randomRange)
}

func initialIf() {
	// Gaya Golang
	if random := rand.Intn(10); random < 5 {
		println("low")
	} else if random >= 5 {
		println("high")
	}

	/*
		// Gaya Lama
		random := rand.Intn(10)
		if random < 5 {
			println("low")
		} else if random >= 5 {
			println("high")
		}
	*/
}

func guessNUmber() {
	angka1 := rand.Intn(20)
	angka2 := rand.Intn(30)
	angka3 := angka1 - angka2

	fmt.Printf("Berapa hasil %d - %d = ? ", angka1, angka2)
	// println("Berapa hasil 22 - 12 = ?")

	// declare var jawab u/ nampung inputan dari console
	var jawab int
	_, err := fmt.Scan(&jawab)
	if err != nil {
		log.Fatal(err)
	}
	println("Jawaban : ", jawab)

	if jawab == angka3 {
		println("Jawaban anda benar!")
	} else {
		println("Jawaban anda salah")
	}
}

func multipleArgs() {
	fmt.Println("Input 3 angka :")
	var angka1, angka2, angka3 int
	n, err := fmt.Scan(&angka1, &angka2, &angka3)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Jawaban : ", angka1, angka2, angka3, n)
}

func main() {
	oddOrEven(8)
	fizzBuzz(10)
	randomNumber()
	initialIf()
	guessNUmber()
	multipleArgs()
}
