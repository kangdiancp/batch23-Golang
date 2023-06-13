package main

import (
	"fmt"
	"sort"
)

func initMap() {
	// store ni
	emps := map[string]int{}
	fmt.Println(emps)

	//using make
	products := make(map[string]float64)

	products["TV"] = 90.00
	products["laptop"] = 123.98
	products["Meja"] = 98.00
	products["laptop"] = 100.00

	fmt.Println("Price TV : ", products["TV"])
	//literalvalue
	students := map[string]float64{
		"Widi": 98.99,
		"Budi": 98.76,
		"Dedi": 87.65,
	}
	for key, v := range students {
		fmt.Println("key :", key, "|value :", v)
	}

}

func deleteMap() {
	students := map[string]float64{
		"Widi": 98.99,
		"Budi": 98.76,
		"Dedi": 87.65,
	}
	delete(students, "Dedi")
	//check key value widi
	if value, ok := students["Widi"]; ok {
		fmt.Println("Widi already exist : ", value)
		delete(students, "Yuli")
		fmt.Println(value, " has been delete")
	} else {
		fmt.Println("No Store Value")
	}
	fmt.Println(students)
}

func totalfruites(fruit []string, vege []string) map[string]int {
	mergeArray := []string{}
	//variadic operator
	mergeArray = append(mergeArray, fruit...)
	mergeArray = append(mergeArray, vege...)

	//declare mapFruit to storekey,value
	mapFruit := map[string]int{}
	for key := range mergeArray {
		mapFruit[mergeArray[key]]++
	}
	return mapFruit
}

func diffFruit(fruit []string, vege []string) ([]string, []string) {
	mergeArray := []string{}
	//variadic operator
	mergeArray = append(mergeArray, fruit...)
	mergeArray = append(mergeArray, vege...)

	//sorting slice
	sort.Strings(mergeArray)

	//declare mapFruit to store key,value
	mapFruit := map[string]int{}
	for key := range mergeArray {
		mapFruit[mergeArray[key]]++
	}
	// declare new variable []string
	sameArray := []string{}
	diffArray := []string{}

	for key, V := range mapFruit {
		if V > 1 {
			sameArray = append(sameArray, key)
		} else {
			diffArray = append(diffArray, key)
		}

	}
	return sameArray, diffArray
}

func main() {
	//initMap()
	//deleteMap()
	/*	result :=
		totalfruites([]string{"Apel", "Jambu", "Melon", "Jambu"},
			[]string{"Tomat", "Apel", "Melon", "Durian"})
	fmt.Println(result)*/

	sameArray, difArray :=
		diffFruit([]string{"Apel", "Jambu", "Melon", "Jambu"},
			[]string{"Tomat", "Apel", "Melon", "Durian"})
	fmt.Println(sameArray)
	fmt.Println(difArray)

}
