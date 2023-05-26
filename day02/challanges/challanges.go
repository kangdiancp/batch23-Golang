package main

import (
	"fmt"
	"strings"
	"unicode"
)

// 1.
func countWordWithSpace(word string) {
	wordSplit := strings.Split(word, " ")
	count := len(wordSplit)

	fmt.Println("(1.) Total Word : ", count)
	// Total Word : 6
}

// 2.
func countWordWithUpper(word string) {
	count := 0

	for _, v := range word {
		if unicode.IsUpper(v) {
			count++
		}
	}

	fmt.Println("\n(2.) Total Word : ", count)
	// Total Word : 6
}

// 3.
func countVocalConsonant(word string) {
	count := 0

	for _, v := range word {
		switch v {
		case 'a', 'i', 'u', 'e', 'o', 'A', 'I', 'U', 'E', 'O':
			count++
		}
	}

	fmt.Println("\n(3.) Vocal : ", count)

	count1 := 0

	for _, v := range word {
		if unicode.IsSpace(v) {
			count1++
		}
	}

	lengthWord := len(word)
	fmt.Println(" ==> Consonant : ", lengthWord-count-count1)
	// Vocal : 3
	// Consonant : 8
}

// 4.
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

	fmt.Println("\n(4.) Result : ", strings.Join(wordSplit, " "))
	// Result : a*u s**a j**n c****d b******p
}

// 5.
func findWord(word string, search string) {
	wordToLower := strings.ToLower(word)
	searchToLower := strings.ToLower(search)

	count := strings.Count(wordToLower, searchToLower)

	// wordSplit := strings.Split(word, " ")

	// count := 0
	// for i, v := range v {
	// 	if wordToLower == searchToLower {
	// 		count := strings.Count(wordToLower, searchToLower)
	// 	}
	// }

	fmt.Printf("\n(5.) Output : Word %s found %d times", search, count)
	// Output : Word go found 3 times
}

func main() {
	countWordWithSpace("How Many Word In This Sentence?")
	countWordWithUpper("HowManyWordInThisSentence?")
	countVocalConsonant("Rockbye Baby")
	replaceMiddleCharWithStar("aku suka join codeid bootcamp")
	findWord("bootcamp Go di codeid, go run Go", "go")
}
