package main

import (
	"fmt"
	"math/rand"
)

func oddOrEven(number int) {
	if number%2 == 0 {
		println("Even Number")
	} else {
		println("Odd Number")
	}
}

/*
	fizzbuzz

number = bisa dibagi 3 & 5 -> fizzBuzz
number = bisa dibagi 3 -> fizz
number = bisa dibagi 5 -> buzz
number = ga bisa dibagi 3 & 5 -> unkwon
*/
func fizzBuzzz(number int) {
	if number%3 == 0 && number%5 == 0 {
		println("fizzBuzz")
	} else if number%3 == 0 {
		println("fizz")
	} else if number%5 == 0 {
		println("buzz")
	} else {
		println("unkwon")
	}
}

// Random Number
func randomNumber() {
	// range output : 0 - 9
	random := rand.Intn(10)
	println(random)
}

func initialIf() {
	// random := rand.Intn(10)
	if random := rand.Intn(10); random < 5 {
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

// func guessNumber() {
// 	angka1 := rand.Intn(10)
// 	angka2 := rand.Intn(6)
// 	angka3 := angka1 - angka2

// 	fmt.Printf("Berpa hasil %d - %d = ?", angka1, angka2)
// 	// println("Berpa hasil 7-3 = ?")

// 	// Declare var jawab u/ nampung inputan dari console
// 	var jawab int
// 	_, err := fmt.Scan(&jawab)
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	if jawab == angka3 {
// 		println("Jawaban anda Benar")
// 	} else {
// 		println("Jawaban anda salah")
// 	}

// 	// println("jawab : ", jawab)
// }

// func multipleArgs() {
// 	fmt.Println("Input 3 angka : ")

// 	var angka1, angka2, angka3 int
// 	n, err := fmt.Scan(&angka1, &angka2, &angka3)
// if err != nil {
// 		log.Fatal(err)
// 	}

// 	fmt.Println("Jawaban : ", angka1, angka2, angka3, n)
// }

func tebakHari(day int) {
	switch n := day; {
	case n == 0:
		fmt.Println("sunday")
	case n == 1:
		fmt.Println("monday")
	case n == 2:
		fmt.Println("Thuesday")
	case n == 3:
		fmt.Println("wednesday")
	case n == 4:
		fmt.Println("thursday")
	case n == 5:
		fmt.Println("friday")
	case n == 6:
		fmt.Println("saturday")
	default:
		fmt.Println("Not today")
	}
}

func main() {
	// oddOrEven(8)
	// fizzBuzzz(15)
	// randomNumber()
	// initialIf()
	// guessNumber()
	// multipleArgs()
	tebakHari(2)
}
