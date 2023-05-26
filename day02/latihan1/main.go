package main

import (
	"fmt"
	"strings"
	"unicode"
)

func main() {

	wordString()
	basicLoop()
	conditionLoop()
	initLoop()
	initLoop1()
	rangeIndexLoop()
	rangeIndexLoop1()
	countUpperChar("BootCamp gOlang")
	countUpperChar1("BootCamp gOlang")
	switchInitLoop()
	switchValLoop()
}

func wordString() {

	word := "Hello World"

	fmt.Println("panjang : ", len(word))
	fmt.Println("To upperCase : ", strings.ToUpper(word))

	ello := word[1:]
	ello1 := word[6:11]

	foundW := strings.Index(word, "W")

	fmt.Println("Hello Substr : ", ello)
	fmt.Println("found W at position : ", foundW)
	fmt.Println("Hello Substr : ", ello1)

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
		fmt.Println("counter", counter)
		counter++
		/*
		 	if counter == 5 {
		 break }
		*/

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

func initLoop1() {
	for counter := 0; counter < 5; counter++ {
		if counter == 2 {
			break
		}
		fmt.Println("Counter : ", counter)
	}
}

func rangeIndexLoop() {
	words := "Golang"

	for index, value := range words {
		fmt.Println("index : ", index, value, string(value))

	}

}

func rangeIndexLoop1() {
	words := "Golang"

	for _, value := range words {
		fmt.Println("tanpa index : ", value, string(value))

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

	fmt.Println("total Huruf Besar ", countUpper)

}

func countUpperChar1(words string) {
	countUpper := 0

	for _, v := range words {
		if unicode.IsSpace(v) {
			countUpper++
		}
	}

	fmt.Println("total spasi ", countUpper)

}

func switchInitLoop() {
	for counter := 0; counter < 20; counter++ {
		switch counter := counter / 2; counter {
		case 2, 3, 5, 7:
			fmt.Println("Prime number", counter/2)
			break
		default:
			fmt.Println("Non Prime Number", counter/2)
			break
		}
	}

}

func switchValLoop() {
	for counter := 0; counter < 20; counter++ {
		switch val := counter / 2; val {
		case 2, 3, 5, 7:
			fmt.Println("1 Prime number", val)
			break
		default:
			fmt.Println("2 Non Prime Number", val)
			break
		}
	}

}
