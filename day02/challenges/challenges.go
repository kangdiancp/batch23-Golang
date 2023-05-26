package main

import (
	"fmt"
	"strings"
	"unicode"
)

func countWordWithSpace(word string) {
	// countSpace := 0
	// for _, space := range word {
	// 	if unicode.IsSpace(space) {
	// 		countSpace++
	// 	}
	// }
	// fmt.Println("Total word : ", countSpace+1)

	words := strings.Fields(word)
	fmt.Println("Total word : ", len(words))
}

func countWordWithUpper(word string) {
	countUpper := 0
	for _, upper := range word {
		if unicode.IsUpper(upper) {
			countUpper++
		}
	}
	fmt.Println("Total word : ", countUpper)
}

func countVocalConsonant(word string) {
	counter := 0
	for _, vocal := range word {
		switch vocal {
		case 'a', 'i', 'u', 'e', 'o', 'A', 'I', 'U', 'E', 'O':
			counter++
		}
	}
	fmt.Println("Vocal : ", counter)

	countSpace := 0
	for _, space := range word {
		if unicode.IsSpace(space) {
			countSpace++
		}
	}
	words := len((word))
	fmt.Println("Consonant : ", words-countSpace-counter)
}

func replaceMiddleCharWithStar(word string) {
	words := strings.Split(word, " ")

	for i, word := range words {
		if len(word) > 2 {
			replaceWord := word[:1] + strings.Repeat("*", len(word)-2) + word[len(word)-1:]
			words[i] = replaceWord
		}
	}

	fmt.Println(strings.Join(words, " "))
}

func findWord(word string, search string) {

	words := strings.Split(word, " ")
	counter := 0
	for _, word := range words {
		if strings.ToLower(word) == strings.ToLower(search) {
			counter++
		}
	}

	fmt.Printf("Word %s found %d times", search, counter)
}

func main() {
	countWordWithSpace("How Many Word In This Sentence")
	countWordWithUpper("HowManyWordInThisSentence")
	countVocalConsonant("Rockbye Baby")
	replaceMiddleCharWithStar("aku suka join codeid bootcamp")
	findWord("bootcamp Go di codeid, go run GO", "go")

}
