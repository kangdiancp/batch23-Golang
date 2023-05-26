package main

import (
	"fmt"
	"strings"
	"unicode"
)

// No 1
func countVocalConsonant(word string) {
	world := strings.ToLower(word)
	vocal := 0
	consonant := 0

	for _, character := range word {
		if character == 'a' || character == 'i' || character == 'u' || character == 'e' || character == 'o' {
			vocal++
		} else if character == ' ' {
			continue
		} else {
			consonant++
		}
	}
	fmt.Println("vocal", vocal)
	fmt.Println("consonant", consonant)
	fmt.Println("world", world)
}

// No 2
func WordWithSpace(word string) {
	hitung := 1
	for _, character := range word {
		if character == ' ' {
			hitung++
		}
	}
	fmt.Println("Total Word", hitung)
}

// NO 3
func countWordwithUpper(word string) {
	jumlah_kata := 0

	for _, v := range word {
		if unicode.IsUpper(v) {
			jumlah_kata++
		}
	}
	fmt.Println(word, "Total Word :", jumlah_kata)
}

// NO 4
func findword(word string, search string) {
	kata := strings.Split(word, " ")
	hasil_value := 0
	for _, hasil := range kata {
		if strings.ToLower(hasil) == strings.ToLower(search) {
			hasil_value++
		}
	}
	fmt.Printf("Word Go Found %d Times", hasil_value)

}

func replaceMiddleCharWithStar(input string) {
	words := strings.Split(input, " ")
	output := ""

	for i := 0; i < len(words); i++ {
		word := words[i]

		if len(word) > 2 {
			middle := word[:len(word)-1]
			star := strings.Repeat("*", len(middle))

			word = string(word[0]) + star + string(word[len(word)-1])
		}

		output += word + " "
	}
	fmt.Println("Result: ", output)

}

func main() {
	//contVocalConsonant("Rockbye Baby")
	//WordWithSpace("How Many Word In This Sentence")
	//countWordwithUpper("HowManyWordInThisSentence")
	//findword("bootcamp Go di codeid, go run Go", "go")
	//replaceMiddleCharWithStar("aku suka join codeid bootcamp")

}
