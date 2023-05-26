package main

import (
	"fmt"
	"strings"
	"unicode"
)

func soal1CountVocalConsonant(word string) {
	word = strings.ToLower(word)
	vocal := 0
	consonant := 0

	for _, jumlah := range word {
		if jumlah == 'a' || jumlah == 'i' || jumlah == 'u' || jumlah == 'e' || jumlah == 'o'{
			vocal++
		}else if jumlah == ' '{
			continue
		}else{
			consonant++
		}
	}
	fmt.Println("Vocal : ", vocal)
	fmt.Println("Consonant : ", consonant)
}

func soal2CountWordWithSpace(word string){
	count := 0
	kata := strings.Fields(word)
	
	for index, kalimat := range kata{
		runes := []rune(kalimat)
		count++
		kata[index] = string(runes)
	}
	fmt.Println("Total Word : ", count)
}


func soal3Countword(word string){
	jumlah_kata := 0
	for _, jumlah := range word {
		if unicode.IsUpper(jumlah){
			jumlah_kata++
		}
	}
	fmt.Println("Total Word : ", jumlah_kata)
}

func soal4FindWord(word string, search string){
	kata := strings.Split(word, " ")
	jumlah_value := 0
	for  _, jumlah := range kata {
		if strings.ToLower(jumlah) == strings.ToLower(search){
			jumlah_value++
		}
	}
	fmt.Printf("Word Go Found %d Times", jumlah_value)
}

func soal5ReplaceMiddleChaeWithStar(word string){
	kata := strings.Fields(word)
	
	for index, kalimat := range kata{
		runes := []rune(kalimat)
		for char := 1; char < len(runes)-1; char++ {
			runes[char] = '*'
		}
		kata[index] = string(runes)
	}
	hasil_kalimat := strings.Join(kata, " ")
	fmt.Println("Result : ", hasil_kalimat)
}

func main(){
	//soal1CountVocalConsonant("Rockbye Baby")
	//soal2CountWordWithSpace("How many Word In This Sentences")
	//soal3Countword("HowManyWordInThisSentences")
	//soal4FindWord("bootcamp Go di codeid, go run Go", "go")
	//soal5ReplaceMiddleChaeWithStar("aku suka join codeid bootcamp")
}
