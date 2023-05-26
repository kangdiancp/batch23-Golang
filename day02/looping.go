package main

import (
	"fmt"
	"strings"
	"unicode"
)

func wordString() {
	word := "Hello Word"

	fmt.Println("length :", len((word)))
	fmt.Println("To Upercase :", strings.ToUpper((word)))

	ello := word[1:]    //index dihitung dari angka 0
	world := word[6:11] //length dihitung dari angka 1

	foundW := strings.Index(word, "W")
	fmt.Println("Hello Subtr :", ello)
	fmt.Println("Found W at position :", foundW)
	fmt.Println("World :", world)
}

func basicLoop() {
	counter := 0
	for counter <= 5 {
		fmt.Println("Counter", counter)
		counter++
		// if counter == 5 {
		// 	break
		// }
	}
}

func conditionLoop() {
	counter := 0
	for counter <= 5 {
		fmt.Println("Counter", counter)
		counter++
		// if counter == 5 {
		// 	break
		// }
	}
}

func intiLoop() {
	for counter := 0; counter < 5; counter++ {
		if counter == 2 {
			continue
		}
		fmt.Println("Counter :", counter)
	}
}

func rangeIndexLoop() {
	words := "Golang"
	for index, value := range words {
		fmt.Println("Index :", index, value, string(value))
	}
	//_ undercose untuk unsued variable
	for _, value := range words {
		fmt.Println("Index :", value, string(value))
	}
}

// input : Bootcamp Golang
func countUpperCarh(words string) {
	countUpper := 0
	for _, value := range words {
		if unicode.IsUpper(value) {
			countUpper++
		}
	}
	fmt.Println(words, "Total Char Uppercase :", countUpper)
}

func swithInitLoop() {
	for counter := 0; counter < 20; counter++ {
		switch counter / 2 {
		case 2, 3, 5, 7:
			fmt.Println("Prime number :", counter/2)
			break
		default:
			fmt.Println("Not prime number :", counter/2)
			break
		}
	}
}

func swithValLoop() {
	for counter := 0; counter < 20; counter++ {
		switch val := counter / 2; val {
		case 2, 3, 5, 7:
			fmt.Println("Prime number :", val)
			break
		default:
			fmt.Println("Not prime number :", val)
			break
		}
	}
}

func main() {
	wordString()
	// basicLoop()
	// conditionLoop()
	// intiLoop()
	// rangeIndexLoop()
	// countUpperCarh("Bootcamp golang")
	// swithInitLoop()
	// swithValLoop()
}
