package main

import (
	"fmt"
	"strings"
	"unicode"
)

// Soal 1
func findWord(word string, search string) {
	count := 0
	words := strings.Fields(word)

	for _, w := range words {
		if strings.ToLower(w) == strings.ToLower(search) {
			count++
		}
	}
	fmt.Printf("Jumlah kata yang ditemukan: %d\n", count)
}

// Soal 2
// func replaceMiddleCharWithStar(words string) {
// 	runes := []rune(words)
// 	for i := 1; i < len(runes)-1; i++ {
// 		if runes[i-1] == ' ' || runes[i+1] == ' ' {
// 			continue
// 		}
// 		if runes[i] != ' ' {
// 			runes[i] = '*'
// 		}

// 	}
// 	println(string(runes))
// }

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

// Soal 3
func countVocalConsonant(word string) {
	countVocal := 0
	countConsonant := 0
	for _, value := range word {
		switch value {
		case 'a', 'e', 'i', 'o', 'u', 'A', 'E', 'I', 'O', 'U':
			countVocal++
		case ' ':
		default:
			countConsonant++
		}
	}
	fmt.Println("Vocal : ", countVocal)
	fmt.Println("Consonant : ", countConsonant)
}

// Soal 4
func countWordWithUpper(words string) {
	countUpper := 0
	for _, v := range words {
		if unicode.IsUpper(v) {
			countUpper++
		}
	}
	fmt.Println(" char Uppercase : ", countUpper)
}

// Soal 5
func countWordWithSpace(word string) {
	count := 0
	words := strings.Fields(word)

	for _, v := range words {
		if v != "" {
			count++
		}
	}
	fmt.Println("Total word : ", count)
}

func main() {
	// countWordWithSpace("How Many Word In This Sentece")
	countWordWithUpper("HowManyWordInThisSpace")
	// replaceMiddleCharWithStar("aku suka join codeid bootcamp")
	// countVocalConsonant("Rockbye Baby")
	// findWord("bootcamp Go di codeid, go run Go", "go")
}
