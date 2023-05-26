package main

import (
	"fmt"
	"strings"
	"unicode"
)

func wordString() {

	word := "Hello Word"

	fmt.Println("length :", len(word))
	fmt.Println("To Uppercase : ", strings.ToUpper(word))

	ello := word[1:]
	world := word[6:11]

	foundtW := strings.Index(word, "W")

	fmt.Println("Hello Subtr :", ello)
	fmt.Println("Found w at Posittion :", foundtW)
	fmt.Println("word ", world)
}

func conditionLoop() {
	counter := 0
	for counter <= 5 {
		fmt.Println("counter", counter)
		counter++

		/* if counter == 5{
			break
		}*/
	}
}

func initLoop() {
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
		fmt.Println("Index : ", index, value)

	}
	//_underscore : unused variable
	for Index, value := range words {
		fmt.Println("Index :", Index, value, string(value))
	}
}

// input : BootCamp Golang
func countUpperChar(words string) {
	countUpper := 0

	for _, v := range words {
		if unicode.IsUpper(v) {
			countUpper++
		}
	}
	fmt.Println(words, "Total Char Uppercase :", countUpper)
}

func switchInitLoop() {
	for counter := 0; counter < 20; counter++ {
		switch counter / 2 {
		case 2, 3, 5, 7:
			fmt.Println("prime number", counter/2)
			break

		default:
			fmt.Println("Non Prime Number", counter/2)
			break

		}
	}

}

func main() {
	//wordString()
	//conditionLoop()
	//initLoop()
	//rangeIndexLoop()
	//countUpperChar("BootCamp Golang")
	//switchInitLoop()

}
