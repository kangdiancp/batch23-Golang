package main

import (
	"fmt"
	"strings"
	"unicode"
)

func wordString() {
	word := "Hello World"
	fmt.Println("Length : ", len((word)))
	fmt.Println("To Uppercase : ", strings.ToUpper(word))

	ello := word[6:10]
	foundW := strings.Index(word, "W")

	fmt.Println("Hello substr : ", ello)
	fmt.Println("Found w at position : ", foundW)
}

func basicLoop() {
	counter := 0
	for {
		fmt.Println("Counter: ", counter)
		counter++
		if counter == 5 {
			break
		}
	}
}

func confitionLoop() {
	counter := 0
	for counter <= 5 {
		fmt.Println("Counter: ", counter)
		counter++
		// if counter == 5 {
		// 	break
		// }
	}
}

func initLoop() {
	for Counter := 0; Counter < 5; Counter++ {
		if Counter == 2 {
			continue
		}
		fmt.Println("Counter: ", Counter)
	}
}

func rangeIndexLoop() {
	words := "Golang"
	for index, value := range words {
		fmt.Println("Index : ", index, value, string(value))
	}
	for _, value := range words {
		fmt.Println("Index : ", value, string(value))
	}
}

// input :BootCamp gOlang
func countUpperChar(words string) {
	countUpper := 0
	spaceUpper := 0
	for _, v := range words {
		if unicode.IsUpper(v) {
			countUpper++
		}
		if unicode.IsSpace(v) {
			spaceUpper++
		}
	}
	fmt.Println(words, "Char Uppercase :", countUpper)
	fmt.Println(words, "Char Spacecase :", spaceUpper)
}

func switchInitLoop() {
	for counter := 0; counter < 20; counter++ {
		switch counter / 2 {
		case 2, 3, 5, 7:
			fmt.Println("Prime Number", counter/2)
			break
		default:
			fmt.Println("Non Prime Number", counter/2)
			break
		}
	}
}

func switchValLoop() {
	for counter := 0; counter < 20; counter++ {
		switch val := counter; val {
		case 2, 3, 5, 7, 11, 13, 17:
			fmt.Println("Prime Number", val)
			break
		default:
			fmt.Println("Non Prime Number", val)
			break
		}
	}
}

func replaceMiddleCharWith(input string) {
	word := strings.Split(input, " ")
	output := ""
	for i := 0; i < len(word); i++ {
		words := word[i]
		if len(words) > 2 {
			middle := words[1 : len(words)-1]
			star := strings.Repeat("*", len(middle))
			words = string(words[0]) + star + string(words[len(words)-1])
		}
		output += words + " "

	}
	fmt.Println("Results : ", output)
}
func main() {
	// wordString()
	// basicLoop()
	// confitionLoop()
	// initLoop()
	rangeIndexLoop()
	countUpperChar("BootCamp gOlang")
	// 	switchInitLoop()
	// 	switchValLoop()
}
