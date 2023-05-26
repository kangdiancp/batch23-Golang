package main

import (
	"fmt"
	"strings"
	"unicode"
)

func wordString() {
	word := "Hello World!"

	fmt.Println("length : ", len(word))
	fmt.Println("To Uppercase : ", strings.ToUpper(word))

	ello := word[1:5]
	world := word[6:11]

	foundW := strings.Index(word, "W")

	fmt.Println("Hello Substr : ", ello)
	fmt.Println("Found W at postion : ", foundW)
	fmt.Println("World : ", world)
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
	counter := 0
	for counter <= 5 {
		fmt.Println("Counter", counter)
		counter++
		/*if counter== 5 {
			break
		}*/
	}
}

func initLoop() {
	for counter := 0; counter < 5; counter++ {
		if counter == 2 {
			//break
			continue
		}
		fmt.Println("Counter :", counter)
	}
}

func rangeIndexLoop() {
	words := "Golang"
	for index, value := range words {
		fmt.Println("Index : ", index, value, string(value))

	}

	// _ underscore : unused variable (ketika ga perlu indexnya)
	for _, value := range words {
		fmt.Println("Index : ", value, string(value))

	}
}

// input : BootCamp gOlang
func countUpperChar(words string) {
	countUpper := 0
	for _, V := range words {
		if unicode.IsUpper(V) {
			countUpper++
		}
	}
	fmt.Println(words, " Total CharUppercase :", countUpper)
}

// display prime number
func swithInitLoop() {
	for counter := 0; counter < 20; counter++ {
		switch counter / 2 {
		case 2, 3, 5, 7:
			fmt.Println("Prime number", counter/2)
			break
		default:
			fmt.Println("Non Prime number", counter/2)
			break
		}
	}
}

func swithValLoop() {
	for counter := 0; counter < 20; counter++ {
		switch val := counter / 2; val {
		case 2, 3, 5, 7:
			fmt.Println("Prime number", val)
			break
		default:
			fmt.Println("Non Prime number", val)
			break
		}
	}
}

func main() {
	wordString()
	conditionLoop()
	initLoop()
	rangeIndexLoop()
	countUpperChar("BootCamp gOlang")
	swithInitLoop()
}
