package main

import (
	"fmt"
	"strings"
	"unicode"
)

func countVocalConsonant(word string) {
	vokal := 0
	konson := 0

	for _, char := range word {
		switch char {
		case 'a', 'i', 'u', 'e', 'o', 'A', 'I', 'U', 'E', 'O':
			vokal++
		default:
			if unicode.IsSpace(char) {
				konson--
			}
			konson++
		}
	}
	fmt.Println("Vokal: ", vokal)
	fmt.Println("Konsonan: ", konson)
}

func countWordWithSpace(word string) {

	words := 0

	sebelumKata := ' '

	for _, char := range word {
		if char != ' ' && sebelumKata == ' ' {
			words++
		}
		sebelumKata = char
	}
	fmt.Println("jumlah kata: ", words)
}

func countWordWithSpace2(word string) {

	countWords := len(strings.Fields(word))

	fmt.Println("Jumlah Kata: ", countWords)

}

func countWordWithUpper(word string) {
	words := 0
	for _, char := range word {
		if unicode.IsUpper(char) {
			words++
		}
	}
	fmt.Print("Total Word: ", words)
}

func findWord(input string, search string) {
	count := 0
	searchlow := strings.ToLower(search)
	wordlow := strings.ToLower(input)
	WordSplit := strings.Split(wordlow, " ")
	//fmt.Println(len(WordSplit))
	perChar := len(WordSplit)
	for i := 0; i < perChar; i++ {
		if WordSplit[i] == searchlow {
			count++
		}
	}
	fmt.Printf("Word go Found %d times", count)

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
	//countVocalConsonant("Rockbye Baby")
	//countWordWithSpace("How Many Word In This Sentence ")
	//countWordWithSpace2(" How Many Word In This Sentence ")
	countWordWithUpper("HowManyWordInThisSentence")
	//findWord("botcamp Go di codeid, go run Go Go Go golang", "go")
	replaceMiddleCharWithStar("Aku suka join codeid bootcamp")

}
