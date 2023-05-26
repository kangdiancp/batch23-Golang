package main

import (
	"fmt"
	"strings"
	"unicode"
)

func countVocalConsonant() {
	word := "Rockbye Baby"
	vocal := 0
	consonant := 0
	for _, v := range word {
		switch v {
		case 'a', 'i', 'u', 'e', 'o', 'A', 'I', 'U', 'E', 'O':
			vocal++
		default:
			if unicode.IsSpace(v) {
				consonant--
			}
			consonant++
		}
	}
	fmt.Println(word, "Char Vokal case :", vocal)
	fmt.Println(word, "Char Consolent case :", consonant)
}

func countWordWithUpper() {
	word := "HowManyWordInThisSentence"
	count := 0
	for _, v := range word {
		if unicode.IsUpper(v) {
			count++
		}
	}
	fmt.Println("Char Uppercase :", count)
}

func countWordWithSpace() {
	word := "How Many Word In This Sentence kamu"
	count := 1
	for _, v := range word {
		if unicode.IsSpace(v) {
			count++
		}
	}
	fmt.Println("Char Spacecase :", count)
}

func findWord(word, search string) {
	wordToLower := strings.ToLower(word)
	searchToLower := strings.ToLower(search)

	count := strings.Count(wordToLower, searchToLower)
	fmt.Printf("word %s found %d times", search, count)
}

func replaceMiddleCharWithStar(word string) {
	wordSplit := strings.Split(word, " ")

	for i, v := range wordSplit {
		// fmt.Println("\n Result : ", v, len(v))
		if len(v) > 2 {
			// var firstLetter, sensorLetter, endLetter string
			firstLetter := v[:1]
			sensorLetter := strings.Repeat("*", len(v)-2)
			endLetter := v[len(v)-1:]

			sensorWord := firstLetter + sensorLetter + endLetter
			wordSplit[i] = sensorWord
		}
	}
	fmt.Println("\nResult : ", strings.Join(wordSplit, " "))
}
func main() {
	// countVocalConsonant()
	// countWordWithUpper()
	//countWordWithSpace()
	replaceMiddleCharWithStar("aku suka join codeid bootcamp")
	// findWord("Bootcamp Go di codeid, go run Golang", "go")
}
