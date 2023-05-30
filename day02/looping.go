package main

import (
	"fmt"
	"strings"
	"unicode"
)

func wordString() {
	word := "Hello World"
	fmt.Println("length : ", len(word))
	fmt.Println("To Uppercase : ", strings.ToUpper(word))

	ello := word[1:5]
	//nilai 1 untuk letak huruf yg dicari, kalau 5 length atau panjang yg akan diambil dihitung dari index pertama
	//012345678910
	//HELLO WORLD
	//word[1:]cari di kata word dengan dimulai dari index 1 sampai seterus nya -> Hasil "ello world"

	foundW := strings.Index(word, "W")

	fmt.Println("Hellow W at Position : ", foundW)
	fmt.Println("Hello, substr : ", ello)
}

func basicLoop() {
	counter := 0
	for {
		fmt.Println("Counter", counter)
		counter++
		if counter == 5 {
			break
		}
	}
}

func conditionLoop() {
	counter := 5
	for counter <= 5 {
		fmt.Println("Counter", counter)
		counter++
		// if counter == 5 {
		// 	break
		// }
	}
}

func initLoop() {
	for counter := 0; counter < 5; counter++ {
		if counter == 2 {
			continue
		}
		fmt.Println("Counter : ", counter)
	}
}

func rangeIndexLoop() {
	words := "Golang"
	for index, value := range words {
		fmt.Println("Index : ", index, value, string(value))
	}
	//_ underscore : unused variable
	for _, value := range words {
		fmt.Println("Index : ", value, string(value))
	}
}

func countUpperChar(words string) {
	countUpper := 0

	for _, v := range words {
		if unicode.IsUpper(v) { //isupper, islower, isspace
			countUpper++
		}
	}
	fmt.Println(words, ", Total Char Uppercase : ", countUpper)
}

func switchInitLoop() {
	for counter := 0; counter < 20; counter++ {
		switch val := counter; val {
		case 2, 3, 5, 7, 11, 13:
			fmt.Println("Prime number ", val)
		default:
			fmt.Println("Non Prime Number ", val)
		}
	}
}

func main() {
	// wordString()
	// basicLoop()
	// conditionLoop()
	// initLoop()
	// rangeIndexLoop()
	// countUpperChar("BootCamp Golang")
	// switchInitLoop()
}
