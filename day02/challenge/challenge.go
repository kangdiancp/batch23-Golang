package main

import (
	"fmt"
	"strings"
	"unicode"
)

func countVocalConsonant(word string) {
	countVocal := 0
	countConsonant := 0

	for _, value := range word {
		switch value {
		case 'a', 'A', 'i', 'I', 'u', 'U', 'e', 'E', 'o', 'O':
			countVocal++
		case ' ':
			break
		default:
			countConsonant++
		}
	}
	fmt.Println("Vocal : ", countVocal)
	fmt.Println("Consonant : ", countConsonant)
}

func countWordWithSpace(word string) {
	countSpace := 0

	for _, v := range word {
		if unicode.IsSpace(v) {
			countSpace++
		}
	}
	fmt.Println("Total Word : ", countSpace+1)

	/*countSpace := 0
	var countWord int
	countWords := len(strings.Fields(word))
	fmt.Println(countWords)

	for _, v := range word {
		if unicode.IsSpace(v) {
			countSpace++
			countWord = countWords - countSpace
			countSpace += countWord
		}
	}
	fmt.Println("Jumlah Kata : ", countSpace)
	*/

}

func countWordWithUpper(word string) {
	countUpper := 0

	for _, v := range word {
		if unicode.IsUpper(v) {
			countUpper++
		}
	}
	fmt.Println("Total Word : ", countUpper)
}

func findWord(word string, search string) {
	pisah := strings.Split(word, " ")
	jumlah := 0

	for _, nilai := range pisah {
		if strings.ToLower(nilai) == strings.ToLower(search) {
			jumlah++
		}
	}

	fmt.Printf("Word %s found %d times", search, jumlah)
	println()
}

func replaceMiddleCharWithStar(input string) {
	words := strings.Split(input, " ")
	hasil := ""

	for i := 0; i < len(words); i++ {
		word := words[i]

		if len(word) > 2 {
			middle := word[:len(word)-1]
			star := strings.Repeat("*", len(middle))

			word = string(word[0]) + star + string(word[len(word)-1])
		}

		hasil += word + " "
	}
	fmt.Println(hasil)

}

func main() {
	countVocalConsonant("Rockbye Baby")
	countWordWithSpace("How Many Word In This Sentence")
	countWordWithUpper("HowManyWordInThisSentence")
	findWord("bootcamp Go di codeid, go run Go", "go")
	replaceMiddleCharWithStar("Aku suka join codeid bootcamp")
}
