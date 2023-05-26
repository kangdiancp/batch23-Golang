package main

import (
	"fmt"
	"strings"
	"unicode"
)

func findWords(word string, search string) {
	//case
	ubahWords := strings.ToLower(word)
	//temukan go
	foundgo := strings.Count(ubahWords, search)
	fmt.Printf("word go found %d times", foundgo)
	println()
}

func countWordWithSpace(word string) {
	space := 1
	for _, value := range word {
		if unicode.IsSpace(value) {
			space++
		}
	}
	fmt.Println(word, "total word : ", space)
}

func countVocalConsonant(word string) {
	countVocal := 0
	countConsonant := 0
	for _, value := range word {
		switch value {
		case 'a', 'i', 'u', 'e', 'o':
			countVocal++
		default:
			countConsonant++
			if unicode.IsSpace(value) {
				countConsonant--
			}
		}
	}
	fmt.Printf("Vocal : %d, Consonant : %d", countVocal, countConsonant)
	println()
}

func countWord(words string) {
	countWord := 0
	for _, v := range words {
		if unicode.IsUpper(v) {
			countWord++
		}
	}
	fmt.Println(words, "total word : ", countWord)
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
	findWords("bootcamp Go di codeid, go run Go", "go")
	replaceMiddleCharWithStar("Aku Suka Join CodeId BootCamp")
	countWordWithSpace("How Many Word In This Sentence")
	countWord("HowManyWordInThisSentence")
	countVocalConsonant("Rockbye Baby")
}
