package main

import (
	"fmt"
	"sort"
)

func initMap() {
	// Store Nil
	emps := map[string]int{}
	fmt.Println(emps)

	fmt.Println()

	// Using make
	prod := make(map[string]float64, 0)

	prod["TV"] = 96.69
	prod["Laptop"] = 123.45

	prod["Meja"] = 98.89
	prod["Laptop"] = 100.001

	fmt.Println("Price TV : ", prod["TV"])
	fmt.Println("Price Laptop : ", prod["Laptop"])

	fmt.Println()

	// Literal Value
	stud := map[string]float64{
		"Fatah": 100,
		"Adjik": 99.99,
		"Ralif": 88.88,
	}

	for i, v := range stud {
		fmt.Println("Key : ", i, ", Value : ", v)
	}

}

func deleteMap() {
	stud := map[string]float64{
		"Fatah": 100,
		"Adjik": 99.99,
		"Ralif": 88.88,
		"Yuni":  89.98,
	}

	// Method Del (Map)
	delete(stud, "Ralif")

	// Check Key Value
	if i, ok := stud["Yuni"]; ok {
		fmt.Println("Yuni already exist : ", i)
		delete(stud, "Yuni")
		fmt.Println("Yuni has been deleted!")
	} else {
		fmt.Println("No Store Value!")
	}

	fmt.Println(stud)
}

func totalFruites(fruits []string, vegans []string) map[string]int {
	mergeArray := []string{}

	// Variadic Operator
	mergeArray = append(mergeArray, fruits...)
	mergeArray = append(mergeArray, vegans...)

	// Declare mapFruit to store key, value
	mapFruits := map[string]int{}

	for i := range mergeArray {
		mapFruits[mergeArray[i]]++
	}

	return mapFruits
}

func diffFruites(fruits []string, vegans []string) ([]string, []string) {
	mergeArray := []string{}

	// Variadic Operator
	mergeArray = append(mergeArray, fruits...)
	mergeArray = append(mergeArray, vegans...)

	// Sorting Slice
	sort.Strings(mergeArray)

	// Declare mapFruit to store key, value
	mapFruits := map[string]int{}

	for i := range mergeArray {
		mapFruits[mergeArray[i]]++
	}

	// Declare new var []string
	sameArray := []string{}
	diffArray := []string{}

	for i, v := range mapFruits {
		if v >= 2 {
			sameArray = append(sameArray, i)
		} else {
			diffArray = append(diffArray, i)
		}
	}

	return sameArray, diffArray
}

func main() {
	initMap()

	fmt.Println()

	deleteMap()

	fmt.Println()

	result := totalFruites(
		[]string{"Mangga", "Apel", "Pisang", "Melon"},
		[]string{"Tomat", "Timun", "Apel", "Mangga"},
	)
	fmt.Println(result)

	fmt.Println()

	sameArray, diffArray := diffFruites(
		[]string{"Mangga", "Apel", "Pisang", "Melon"},
		[]string{"Tomat", "Timun", "Apel", "Mangga"},
	)
	fmt.Println("Same Array : ", sameArray)
	fmt.Println("Different Array : ", diffArray)
}
