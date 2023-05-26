package main

import (
	"fmt"
	"strings"
	"unicode"
)

func wordString() {
	word := "Hello World"
	fmt.Println("length : ", len((word)))
	fmt.Println("To Uppercase : ", strings.ToUpper(word))

	ello := word[1:5]
	words := word[6:11]

	fontW := strings.Index(word, "W")

	fmt.Println("Hello Subtr : ", ello)
	fmt.Println("Found w at posittion : ", fontW)
	fmt.Println("world : ", words)
}

func basicLoop() {
	counter := 0
	for {
		fmt.Println("Counter ", counter)
		counter++
		if counter == 5 {
			break
		}
	}
}

func condition() {
	count := 0
	for count <= 5 {
		fmt.Println("Counter ", count)
		count++
	}
}

func initLoop() {
	for count := 5; count < 5; count++ {
		if count == 2 {
			break
		}
		fmt.Println("Count : ", count)
	}
}

func rangeIndexLoop() {
	words := "Golang"
	// for index, value := range words {
	// 	fmt.Println("Index : ", index, value, string(value))
	// }

	// _ Underscore = unused variable
	for index, value := range words {
		fmt.Println("Index : ", index, value, string(value))
	}
}

// input = BootCamp g0lang
func countUpperChar(words string) {
	countUpper := 0
	for _, v := range words {
		if unicode.IsUpper(v) {
			countUpper++
		}
	}
	fmt.Println(words, " char Uppercase : ", countUpper)
}

// display prime number
func switchLoop() {
	for count := 0; count < 20; count++ {
		switch count / 2 {
		case 2, 3, 5, 7:
			fmt.Println("Prime number", count/2)
		default:
			fmt.Println("Non Prime Number", count/2)
		}
	}
}

func switchValLoop() {
	for count := 0; count < 20; count++ {
		switch val := count / 2; val {
		case 2, 3, 5, 7:
			fmt.Println("Prime number", val)
		default:
			fmt.Println("Non Prime Number", val)
		}
	}
}

func main() {
	// switchLoop()
	// switchValLoop()
	// rangeIndexLoop()
	// countUpperChar("BootCamp gOlang")
	wordString()
	condition()
	// basicLoop()
	// initLoop()
}
