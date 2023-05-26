package main

import (
	"fmt"
	"strings"
	"unicode"
)

func wordString() {
	word := "Hello World"

	fmt.Println("length : ", len(word))
	fmt.Println("To uppercase : ", strings.ToUpper(word))

	ello := word[1:5]
	world := word[6:11]
	foundW := strings.Index(word, "W")

	fmt.Println("Hello Substr : ", ello)
	fmt.Println("Hello Substr : ", world)
	fmt.Println("Found W at position : ", foundW)
}

func basicLoop() {
	counter := 0
	for {
		fmt.Println("Counter : ", counter)
		counter++
		if counter == 5 {
			break
		}
	}
}

func conditionLoop() {
	counter := 0
	for counter <= 5 {
		fmt.Println("Counter : ", counter)
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
			// break
		}
		fmt.Println("Counter : ", counter)
	}
}

func rangeIndexLoop() {
	words := "Golang"
	for index, value := range words {
		fmt.Println("Index : ", index, value, string(value))
	}

	// _ Underscrore : unused variable
	for _, value := range words {
		fmt.Println("Index : ", value, string(value))
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
	fmt.Println(words, " Total Char Uppercase : ", countUpper)
}

// Display prime number
func swithInitLoop() {
	for counter := 0; counter < 20; counter++ {
		switch counter / 2 {
		case 2, 3, 5, 7:
			fmt.Println("Prime number ", counter/2)
			break
		default:
			fmt.Println("Non Prime number ", counter/2)
			break
		}
	}
}

func swithValLoop() {
	for counter := 0; counter < 20; counter++ {
		switch val := counter / 2; val {
		case 2, 3, 5, 7:
			fmt.Println("Prime number ", val)
			break
		default:
			fmt.Println("Non Prime number ", val)
			break
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
	// swithInitLoop()
	// swithValLoop()

}
