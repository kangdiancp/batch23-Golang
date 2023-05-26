package main

import (
	"fmt"
	"strings"
	"unicode"
)

func wordString() {
	word := "Hello World"

	fmt.Println("Lenght : ", len(word))
	fmt.Println("To Uppercase : ", strings.ToUpper(word))
	
	ello := word[1:5]		// memotong dari index 1 sampai index 5 yang akan di tampilkan dengan menggunakan slice
	world := word[6:11]
	foundw := strings.Index(word, "W")

	fmt.Println("Hello Substr : ", ello)
	fmt.Println("Hello Substr : ", world)
	fmt.Println("Found W at Position : ", foundw)
}

func basicLoop(){
	counter := 0
	for {
		fmt.Println("Counter", counter)
		counter++
		if counter == 5 {
			break
		}
	}
}

func ConditionLoop(){
	counter := 0
	for counter <= 5 {
		fmt.Println("Counter", counter)
		counter++
				// if counter == 5 {
				// 	break
				// }
	}
}

func initLoop(){
	for counter := 0; counter < 5; counter++ {
		if counter == 2 {
			// break
			continue
		}
		fmt.Println("Counter : ", counter)
	}
}

func rangeIndexLoop(){
	words := "Golang"
	for  index, value := range words {
		fmt.Println("Index : ", index, value, string(value))
	}

	// Underscore : unused variable
	for  _, value := range words {
		fmt.Println("Index : ",  value, string(value))
	}
}

// input : BootCamp gOlang
func countUpperChar(word string){
	countUpper := 0
	for _, v := range word {
		if unicode.IsUpper(v){ // bisa menggunakan islower isspace dll untuk menghitung 
			countUpper++
		}
	}
	fmt.Println(word, "Total Char Uppercase : ", countUpper)
}

// display prime number
func switchInitLooping(){
	for counter := 0; counter < 20; counter++ {
		switch counter / 2 {
		case 2, 3, 5, 7:
			fmt.Println("Prime Number", counter/2)
		default :
			fmt.Println("Non Prime Number", counter/2)
			break
		}
	}
}

func switchValtLooping(){
	for counter := 0; counter < 20; counter++ {
		switch val := counter / 2 ; val {
		case 2, 3, 5, 7:
			fmt.Println("Prime Number", val)
		default :
			fmt.Println("Non Prime Number", val)
			break
		}
	}
}

func main() {
	//wordString()
	//basicLoop()
	//ConditionLoop()
	//initLoop()
	//rangeIndexLoop()
	//countUpperChar("BootCamp gOlang")
	//switchInitLooping()
	switchValtLooping()
}