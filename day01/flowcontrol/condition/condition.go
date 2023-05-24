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

// fizzBuzz
/*
	number = bisa dibagi 3 & 5 -> fizzBuzz
	number = bisa dibagi 3 -> fizz
	number = bisa dibagi 5 -> buzz
	number = ga bisa dibagi 3 & 5 -> unknown
*/
func fizzBuzz(number int) {
	if number%3 == 0 && number%5 == 0 {
		println("fizzBuzz")
	} else if number%3 == 0 {
		println("fizz")
	} else if number%5 == 0 {
		println("Buzz")
	} else {
		println("Unknown")
	}
}

func randomNumber() {
	// range output : 0 - 9
	random := rand.Intn(10)
	println(random)

	//range output : 1 - 10
	min := 100
	max := 200
	randomRange := rand.Intn(max - min) + min
	println(randomRange)
}

func initialIf(){
	//random := rand.Intn(10)
	
	if random := rand.Intn(10); random < 5 {
		println("low")
	}else if random >= 5 {
		println("high")
	}
	
	/*if random < 5 {
		println("low")
	}else if random >= 5 {
		println("high")
	} */
}

func guestNumber(){
	angka1 := rand.Intn(10)
	angka2 := rand.Intn(6)
	angka3 := angka1 - angka2

	fmt.Printf("Berapa Hasil %d - %d = ?\n", angka1, angka2)
	//println("Berapa hasil dari 7 - 3 = ?")

	// declare variable jawab untuk nampung inputan dari console 
	var jawab int
	_,err := fmt.Scan(&jawab)
	if err != nil{
		log.Fatal(err)
	}
	println("Jawaban : ", jawab)

	if jawab == angka3 {
		println("Jawaban Anda Benar")
	} else {
		println("Jawaban Salah")
	}
}

func multipleArgs(){
	fmt.Println("Input 3 angka : ")
	var angka1, angka2, angka3 int
	n, err := fmt.Scan(&angka1, &angka2, &angka3)
	if err != nil{
		log.Fatal(err)
	}
	fmt.Println("Jawaban : ", angka1, angka2, angka3, n)
}

func tebakHari(day int){
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
		fmt.Println("Not Today")
	};
}

func main() {
	//oddOrEvent(8)
	//fizzBuzz(15)
	//randomNumber()
	//initialIf()
	//guestNumber()
	//multipleArgs()
	tebakHari(5)
}