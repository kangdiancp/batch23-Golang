package main

import (
	"fmt"
	"strings"
	"unicode"
)

func main() {

	//Soal11
	findWord("bootcamp Go di codeid, go run Go and go", "go")
	//Soal12
	replaceMiddleCharWithStar("aku suka join codeid bootcamp")

	//Soal 13
	countVocalConsonant("Rockbye Baby")

	//Soal 14
	countWordWithUpper("HowManyWordInThisgueSentence")

	//Soal 15
	countWordWithSpace(" How Many Word In This Sentence yudha ")

	cobaCoba("how many word in this sentence")
}

//Nomor11

func findWord(words, search string) {
	// countWords := 0
	countGo := 0
	//case 1
	ubahWord := strings.ToLower(words)
	//untuk menghitung jumlah go
	foundGo := strings.Count(ubahWord, search)
	fmt.Printf("Word go found %d times", foundGo)
	println()

	//case2
	sliceWordtospace := strings.Split(ubahWord, " ")
	countWords := len(sliceWordtospace)
	for i := 0; i < countWords; i++ {
		if sliceWordtospace[i] == search {
			countGo++
		}
	}
	fmt.Printf("Word go found %d times - case 2", countGo)
	println()
}

// Nomor 12
func replaceMiddleCharWithStar(word string) {
	sliceWordfromSpace := strings.Split(word, " ")
	countWords := len(sliceWordfromSpace)
	var output string

	//case 1
	for i := 0; i < countWords; i++ {
		var hasil string
		word := sliceWordfromSpace[i]
		if len(word) > 2 {
			sliceMiddle := word[1 : len(word)-1]
			// println(sliceMiddle)
			replaceStarinMiddle := strings.Repeat("*", len(sliceMiddle))
			hasil = string(word[0]) + replaceStarinMiddle + string(word[len(word)-1])
			// fmt.Println("Result : ", hasil)

		}
		output += hasil + " "
	}
	fmt.Println("Result : ", output)

	//case 2
	// for i, value := range {

	// }
}

// Nomor 13
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

// Nomor 14
func countWordWithUpper(word string) {
	countUpper := 0

	for _, value := range word {
		if unicode.IsUpper(value) {
			countUpper++
		}
	}
	fmt.Printf("Total Word : %d", countUpper)
	println()
}

// Nomor 15
func countWordWithSpace(word string) {
	countSpace := 0
	//var countWordWithoutSpace int
	countWords := len(strings.Fields(word))
	println(countWords)

	for _, value := range word {
		if unicode.IsSpace(value) {
			countSpace++
			countWordWithoutSpace := countWords - countSpace
			countSpace += countWordWithoutSpace
		}
	}
	fmt.Printf("Total Word : %d", countSpace)
}

// Cobacoba
func cobaCoba(word string) {

}
