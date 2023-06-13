package main

import (
	"fmt"
	"strings"
	"unicode"
)

func basicloop() {
	counter := 0
	for counter <= 5 {
		fmt.Println("counter", counter)
		counter++
	}
}

func wordString() {
	word := "Hello Word"

	fmt.Println("length :", len(word))
	fmt.Println("to uppercase :", strings.ToUpper(word))

	ello := word[1:5]
	world := word[6:10]

	foundW := strings.Index(word, "W")

	fmt.Println("hello substr :", ello)
	fmt.Println("found w at position :", foundW)
	fmt.Println("word :", world)

}

func initloop() {
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
	for _, value := range words {
		fmt.Println("Index :", value, string(value))
	}

}

func counterUpperChar(words string) {
	counterUpper := 0
	for _, v := range words {
		if unicode.IsUpper(v) {
			counterUpper++
		}
	}
	fmt.Println(words, "Char Uppercase :", counterUpper)

}

func switchInitLoop() {
	for counter := 0; counter < 20; counter++ {
		switch val := counter / 2; val {
		case 2, 3, 5, 7:
			fmt.Println(" prime number :", val)
			break
		default:
			fmt.Println(" non prime number :", val)
			break
		}
	}
}

func main() {
	//wordString()
	//basicloop()
	initloop()
	//rangeIndexLoop()
	//counterUpperChar("Bootcamp Golang")
	//switchInitLoop()

}
