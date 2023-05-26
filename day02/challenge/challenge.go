package main

import (
	"fmt"
	"strings"
	"unicode"
)

// soal nomor 1
func findWord(word string, search string) {
	// countWord := 0
	count := strings.Count(strings.ToLower(word), search)
	strings.ToLower(search)
	fmt.Printf("\nword '%s' found %d times", search, count)
	println()
}

// func findWord(word string, search string) {
// 	count :=
// 	word := strings.ToLower(word)
// 	search := strings.ToLower(search)
// 	words := strings.Fields(search)

// 	word := "go"
// 	search := "bootcamp Go di codeid, go run Go"
// 	hasil := findWord(word, search)
// 	fmt.Println(hasil)

// 	for _, w := range words {
// 		if w == word {
// 			count++
// 		}
// 	}
// 	return
// 	fmt.Sprintf("Word '%s' found %d times.", word, count)
// }

// no 2
func replaceMiddleCharWithStar(input string) {
	word := strings.Split(input, " ")
	output := ""

	for i := 0; i < len(word); i++ {
		words := word[i]
		if len(words) > 2 {
			middle := words[1 : len(words)-1]
			star := strings.Repeat("*", len(middle))
			words = string(words[0]) + star + string(words[len(words)-1])

		}
		output += words + " "

	}
	fmt.Println("Results :", output)
}

// no 3
func countVocalConsonant(word string) {
	vocal := 0
	consonant := 0
	for _, alfabet := range word {
		switch alfabet {
		case 'a', 'e', 'i', 'o', 'u':
			vocal++
		case 'A', 'E', 'I', 'O', 'U':
			vocal++
		default:
			consonant++
			if unicode.IsSpace(alfabet) {
				consonant--
			}

		}
	}
	fmt.Println("Vocal : ", vocal)
	fmt.Println("Consonant : ", consonant)
	return
}

// no 4
func countWordWithUpper(word string) {
	countUpper := 0
	for _, v := range word {
		if unicode.IsUpper(v) {
			countUpper++
		}
	}
	fmt.Println("Total Word after Slice Upper :", countUpper)
}

// no 5
func countWordWithSpace(word string) {
	countSpace := 0
	countwordwithspace := len(word)
	for _, v := range word {
		if unicode.IsSpace(v) {
			countSpace++
			gapWordinSpace := countwordwithspace - countSpace
			countSpace += gapWordinSpace
		}
	}
	fmt.Println("Total Word without Space :", countSpace)
}

func main() {
	// findWord("bootcamp Go di Codeid, go run Go", "go")
	replaceMiddleCharWithStar("aku suka join codeid bootcamp")
	// countVocalConsonant("Rockye Baby")
	// countWordWithUpper("HowManyWordInThisSentence")
	// countWordWithSpace("How Many Word In This Sentence")

}
