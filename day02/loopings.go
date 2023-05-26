package main

import (
	"fmt"
	"strings"
	"unicode"
)

func wordString() {
	word := "hello word"

	fmt.Println("length : ", len(word))
	fmt.Println("to uppercase :", strings.ToUpper(word))

	ello := word[1:5]
	world := word[6:10]

	foundw := strings.Index(word, "w")

	fmt.Println("hello substr :", ello)
	fmt.Println("found w at position :", foundw)
	fmt.Println("word :", world)
}

func basicLoop() {
	counter := 0
	for counter <= 5 {
		fmt.Println("counter", counter)
		counter++
		/*if counter == 5{
			break
		}*/
	}
}

func initLoop() {
	for counter := 0; counter < 5; counter++ {
		if counter == 2 {
			continue
		}
		fmt.Println("counter :", counter)
	}
}

func rangeIndexLoop() {
	words := "golang"
	for Index, value := range words {
		fmt.Println("Index :", Index, value, string(value))
	}
	// _underscore : unused variable
	for _, value := range words {
		fmt.Println("Index :", value, string(value))
	}
}

// input : BootCamp gOlang
func countUpperChar(words string) {
	countUpper := 0
	for _, v := range words {
		if unicode.IsUpper(v) {
			countUpper++
		}
	}
	fmt.Println(words, "Char Uppercase : ", countUpper)
}

// display prime number
func switchInitLoop() {
	for counter := 0; counter < 20; counter++ {
		switch val := counter / 2; val {
		case 2, 3, 5, 7:
			fmt.Println("prime number :", val)
			break
		default:
			fmt.Println(" non prime number :", val)
			break
		}
	}
}

func main() {
	/*wordString()
	basicLoop()
	initLoop()
	rangeIndexLoop()
	countUpperChar("BootCamp gOlang")*/
	switchInitLoop()
}
