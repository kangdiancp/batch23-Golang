package main

import (
	"fmt"
	"log"
	"math/rand"
)

func main() {
	oddOrEven(8)
	fizzBuzz(15)
	randomNumber()
	initialIf()
	guessNumber()
	multipleArgs()
}

func oddOrEven(number int) {
	if number%2 == 0 {
		println("Even number")
	} else {
		println("odd number")
	}
}

//FizzBuzz
/*
	number = bisa dibagi 3 & 5  > fizzBuzz
	number = bisa dibagi 3 > fizz
	number = bisa dibagi 5 > 5
	number = ga bisa dibagi 3 & 5 > unknown

*/

func fizzBuzz(number int) {

	if number%3 == 0 && number%5 == 0 {
		println("fizzBuzz") //jika kondisi memenuhi maka output yg keluar ini
	} else if number%3 == 0 {
		println("fizz") //jika angka 3 tdk sama dg 0 maka output ini
	} else if number%5 == 0 {
		println("Buzz") // jika angka 5 tdk sama dg 0 maka output ini
	} else {
		println("unknown") //jika kondisi tdk terpenuhi maka output ini
	}
}

func randomNumber() {
	//ragerandom
	random := rand.Intn(10)
	println(random)

	//range output : 1-10
	min := 1
	max := 20
	randomrange := rand.Intn(max-min) + min
	println("randomrange:", randomrange)

}

func initialIf() {
	//	random := rand.Intn(10)
	if random := rand.Intn(10); random < 5 {
		println("low")
	} else if random >= 5 {
		println("high")
	}

	/*
		if random < 5 {
			println("low")
		} else if random >= 5 {
			println("high")
		}
	*/
}

func guessNumber() {
	angka1 := rand.Intn(10)
	angka2 := rand.Intn(6)
	angka3 := angka1 - angka2

	fmt.Printf("berapa hasil dari %d - %d = ? ", angka1, angka2)

	var jawab int
	_, err := fmt.Scan(&jawab)
	if err != nil {
		log.Fatal(err)
	}
	if jawab == angka3 {
		println("benar")
	} else {
		println("salah")
	}

	println("jawaban : ", jawab)
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
