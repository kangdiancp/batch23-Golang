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
	number = bisa dibagi 3 & 5 -> fizzBuzz
	number = bisa dibagi 3 -> fizz
	number = bisa dibagi 5 -> buzz
	number = ga bisa dibagi 3 &5 -> unknown
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
	// range output : 0-9
	random := rand.Intn(10)
	println(random)

	//range output : 1-10
	min := 10
	max := 20
	randomRange := rand.Intn(max-min) + min
	println("randomMin:", randomRange)
}

func initialIf() {
	//random := rand.Intn(10)
	if random := rand.Intn(10); random < 5 {
		println("low")
	} else if random >= 5 {
		println("high")
	}

	/* if random < 5 {
		println("low")
	} else if random >= 5 {
		println("high")
	} */
}

func guessNumber() {
	angka1 := rand.Intn(10)
	angka2 := rand.Intn(6)
	angka3 := angka1 - angka2
	str := "tebak"

	fmt.Printf("Berapa hasil %d - %d = ? jawab:%s", angka1, angka2, str)
	//println("Berapa hasil 7-3 =?")

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
	case n == 6:
		fmt.Println("Saturday")
	default:
		fmt.Println("Not today")
	}
}

func main() {
	/* 	oddOrEven(8)
	   	fizzBuzz(15)
	   	randomNumber()
	   	initialIf()
	   	guessNumber() */
	//multipleArgs()
	tebakHari(3)
}
