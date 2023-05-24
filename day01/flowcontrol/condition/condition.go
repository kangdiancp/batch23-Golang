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
	number = bisa dibagi 3 & 2 -> fizzbuzz
	number = bisa dibagi 3 -> fizz
	number = bisa dibagi 2 -> buzz
	number = tidak bisa dibagi 3 & 2 -> unknown


*/

func fizzBuzz(number int) {
	if number%3 == 0 && number%2 == 0 {
		println("fizzbuzz")
	} else if number%3 == 0 {
		println("fizz")
	} else if number%2 == 0 {
		println("buzz")
	} else {
		println("unknown")
	}
}

func randomNumber() {
	random := rand.Intn(10)
	println(random)

	//range output : 1-10
	//min := 10
	//min := 20
	randomRange := rand.Intn(20-10) + 10
	println("randomMin :", randomRange)
}

func initialIf() {
	//random := rand.Intn(10)
	if random := rand.Intn(10); random < 9 {
		println(random, "low")
	} else if random >= 9 {
		println(random, "high")
	}

	/* if random < 5 {
		println("low")
	} else if random >= 5 {
		println("high")
	} */
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
	fmt.Println("input 3 angka :")
	var angka1, angka2, angka3 int
	n, err := fmt.Scan(&angka1, &angka2, &angka3)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("jawaban :", angka1, angka2, angka3, n)
}

func main() {
	oddOrEven(11)
	fizzBuzz(6)
	randomNumber()
	initialIf()
	guessNUmber()
	multipleArgs()
}
