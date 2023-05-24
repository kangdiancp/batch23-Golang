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

//fizzbuzz
/*
	number = bisa dibagi 3 & 5 -> fizzbuzz
	number = bisa dibagi 3 -> fizz
	number = bisa dibagi 5 -> buzz
	number = ga bisa dibagi 3 & 5 -> unknown
*/
func fizzBuzz(number int) {
	if number%3 == 0 && number%5 == 0 {
		println("fizzbuzz")
	} else if number%3 == 0 {
		println("fizz")
	} else if number%5 == 0 {
		println("buzz")
	} else {
		println("unknown")
	}
}

func randomNumber() {
	//range output : 0 - 9
	random := rand.Intn(10)
	println(random)

	//range output : 1 - 10
	min := 1
	max := 10
	randomRange := rand.Intn(max-min) + min
	println("randomRange: ", randomRange)
}

func initialIf() {
	if random := rand.Intn(10); random < 5 {
		println("low")
	} else if random >= 5 {
		println("high")
	}

	/*random := rand.Intn(10)
	if random < 5 {
		println("wow")
	} else if random >= 5 {
		println("high")
	}*/
}

func guessNumber() {
	angka1 := rand.Intn(10)
	angka2 := rand.Intn(6)
	angka3 := angka1 - angka2

	fmt.Printf("Berapa hasil %d - %d = ? ", angka1, angka2)
	//println("Berapa hasil 7-3 = ?")

	//declare var jawab u/ nampung inputan dari console
	var jawab int
	_, err := fmt.Scan(&jawab)
	if err != nil {
		log.Fatal(err)
	}
	if jawab == angka3 {
		println("Jawaban anda Benar")
	} else {
		println("Jawaban anda Salah")
	}
	//println("Jawaban : ", jawab)
}

func multipleArgs() {
	fmt.Println("Input 3 angka :")
	var angka1, angka2, angka3 int
	n, err := fmt.Scan(&angka1, &angka2, &angka3)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("jawaban : ", angka1, angka2, angka3, n)
}

func main() {
	oddOrEven(8)
	fizzBuzz(15)
	randomNumber()
	initialIf()
	guessNumber()
	multipleArgs()
}
