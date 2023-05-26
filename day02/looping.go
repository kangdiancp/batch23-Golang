package main

import (
	"fmt"
	"strings"
	"unicode"
)

func wordString() {
	word := "Hello World!"

	fmt.Println("Length : ", len(word))
	fmt.Println("To Uppercase : ", strings.ToUpper(word))

	ello := word[1:]

	foundW := strings.Index(word, "W")

	World := word[6:] // Atau word[6:12] ==> [6:] = sampai index terakhir

	fmt.Println("Hello Substr : ", ello)
	fmt.Println("Found W at position : ", foundW)
	fmt.Println("World Substr : ", World)
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

func conditionalLoop() {
	counter := 0
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
			// break
			continue
		}
		fmt.Println("Counter : ", counter)
	}
}

func rangeIndexLoop() {
	words := "Golang!"

	for index, value := range words {
		fmt.Println("Index, value, string(value) : ", index, value, string(value))
	}
	// Tanpa index, menggunakan _, underscore : unused variable
	for _, value := range words {
		fmt.Println("value, string(value) : ", value, string(value))
	}
}

// Input : BootCamp GoLang!
func countUpperChar(words string) {
	countUpper := 0
	for _, v := range words {
		if unicode.IsUpper(v) {
			countUpper++
		}
		// if unicode.IsSpace(v) {
		// 	countUpper++
		// }
	}
	fmt.Println(words, "Total Char Uppercase : ", countUpper)
	// fmt.Println(words, "Total Char Space : ", countUpper)
}

// Display prime number
func switchValLoop() {
	for i := 0; i < 20; i++ {
		switch val := i / 2; val {
		case 2, 3, 5, 7:
			fmt.Println("Prime Number", val)
			break
		default:
			fmt.Println("Non Prime Number", val)
			break
		}
	}
}

func main() {
	wordString()
	basicLoop()
	conditionalLoop()
	initLoop()
	rangeIndexLoop()
	countUpperChar("BootCamp GoLang!")
	switchValLoop()
}
