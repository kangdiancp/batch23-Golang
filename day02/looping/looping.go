package main

import (
	"fmt"
	"strings"
	"unicode"
)

func wordString() {
	word := "Hello World"
	fmt.Println("Lenght: ", len(word))
	fmt.Println("To Upper: ", strings.ToUpper((word)))

	//memotong char
	ello := word[6:11]

	roundw := strings.Index(word, "W")
	fmt.Println("Hello Substr: ", ello)
	fmt.Println("Found W at Posisi: ", roundw)

}

func basicLooping() {
	counter := 0
	for {
		fmt.Println("Counter", counter)

		if counter == 5 {
			break
		}
		counter++
	}
}

func conditionLooping() {
	counter := 0
	for counter <= 5 {
		fmt.Println("Counter", counter)
		counter++

	}
}

func initloop() {
	for Counter := 0; Counter < 5; Counter++ {
		if Counter == 2 {
			continue
		}
		fmt.Println("Counter ", Counter)
	}
}

func rangeIndexLoop() {
	words := "golang"

	for index, value := range words {
		fmt.Println("index: ", index, value, string(value))
	}
	// underscore unused variable
	for _, value := range words {
		fmt.Println("index: ", value, string(value))
	}

}

func countUpperChar(words string) {
	countUpper := 0

	for _, v := range words {
		if unicode.IsUpper(v) {
			countUpper++
		}
	}
	fmt.Println(words, "char Uppercas Totalnya adalah: ", countUpper)
}

func countSpaceChar(words string) {
	countUpper := 0

	for _, v := range words {
		if unicode.IsLower(v) {
			countUpper++
		}
	}
	fmt.Println(words, "char Spacenya adalah: ", countUpper)
}

// tampilkan bilangan prima
func switchinitloop() {
	for counter := 0; counter < 20; counter++ {
		switch counter / 2 {
		case 2, 3, 5, 7:
			fmt.Println("Prime NUmber", counter/2)
			break
		default:
			fmt.Println("Non Prime NUmber", counter/2)
			break

		}
	}
}

func switchVallloop() {
	for counter := 0; counter < 20; counter++ {
		switch val := counter; val {
		case 2, 3, 5, 7, 11, 13, 17, 19:
			fmt.Println("Prime Number", val)
			break
		default:
			fmt.Println("Non Prime Number", val)
			break

		}
	}
}

func main() {
	//wordString()
	//basicLooping()
	//conditionLooping()
	//initloop()
	//rangeIndexLoop()
	//countUpperChar("BootCamP GolaNG")
	//countSpaceChar("BootCamP GolaNG Go")
	switchVallloop()
}
