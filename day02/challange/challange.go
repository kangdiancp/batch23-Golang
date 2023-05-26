package main

import (
	"fmt"
	"strings"
	"unicode"
)

// Nomor 1 : Menghitung berapa jumlah kata yang dicari
func findWord(word string, search string) {
	count := strings.Count(strings.ToUpper(word), strings.ToUpper(search))
	fmt.Printf("Kata '%s' ditemukan %d kali\n", search, count)
}

// Nomor 2 : Mengganti karakter tengah dalam kata dengan star
func replaceMiddleCharWithStar(input string) {
	words := strings.Split(input, " ")
	result := ""

	for _, word := range words {
		if len(word) > 2 {
			middle := len(word) / 2
			word = word[:middle] + "*" + word[middle+1:]
		}

		result += word + " "
	}

	fmt.Println("Result: ", result)
}

// Nomor 3 : Menghitung berapa jumlah huruf vokal dan konsonan
func countVocalConsonant(word string) {
	jumlahVokal := 0
	jumlahKonsonan := 0

	word = strings.ToLower(word)

	for _, v := range word {
		switch v {
		case 'a', 'e', 'i', 'o', 'u':
			jumlahVokal++
		default:
			if v >= 'a' && v <= 'z' {
				jumlahKonsonan++
			}
		}
	}

	fmt.Println("Vocal:", jumlahVokal)
	fmt.Println("Consonant:", jumlahKonsonan)
}

// Nomor 4 : Menghitung berapa jumlah kata yang dipisahkan oleh huruf kapital

func countWordWithUpper(word string) {
	jumlahKata := 0

	for _, value := range word {
		if unicode.IsUpper(value) {
			jumlahKata++
		}
	}

	fmt.Println("Total Word :", jumlahKata)
}

// Nomor 5: Menghitung berapa jumlah kata yang dipisahkan oleh spasi

func countWordWithSpace(word string) {
	jumlahKata := strings.Count(word, " ") + 1
	fmt.Println("Total Word :", jumlahKata)
}

func main() {
	findWord("bootcamp Go di codeid, go run Go, go far away", "go")
	replaceMiddleCharWithStar("aku suka join codeid bootcamp")
	countVocalConsonant("Rockbye Baby")
	countWordWithUpper("HowManyWordInThisSentence")
	countWordWithSpace("How Many Word In This Sentence")
}
