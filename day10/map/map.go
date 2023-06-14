package main

import (
	"fmt"
)

func initMap() {
	// store nil
	emps := map[string]int{}
	fmt.Println(emps)

	//using make
	products := make(map[string]float64)

	products["TV"] = 90.00
	products["Laptop"] = 123.98

	products["Meja"] = 98.00
	products["Laptop"] = 100.98

	fmt.Println("Price TV : ", products["TV"])

	// literal value
	students := map[string]float64{
		"Widi": 98.99,
		"Budi": 90.76,
		"Dedi": 87.65,
	}

	for key, v := range students {
		fmt.Println("Key : ", key, "Value : ", v)
	}

}

func deleteMap() {
	students := map[string]float64{
		"Widi": 98.99,
		"Budi": 90.76,
		"Dedi": 87.65,
		"Yuli": 97.88,
	}

	delete(students, "Dedi")

	// check key value widi
	if value, ok := students["Yuli"]; ok {
		fmt.Println("Yuli already exist : ", value)
		delete(students, "Yuli")
		fmt.Println(value, " has been deleted")
	} else {
		fmt.Println("No store value")
	}

	fmt.Println(students)
}

func totalFruits(fruit []string, vege []string) map[string]int {
	mergeArray := []string{}
	//variadic operator
	mergeArray = append(mergeArray, fruit...)
	mergeArray = append(mergeArray, vege...)

	//declare mapFruit to store key, value
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

	// sorting slice
	//sort.Strings(mergeArray)

	//declare mapFruit to store key, value
	mapFruit := map[string]int{}

	for key := range mergeArray {
		mapFruit[mergeArray[key]]++
	}

	//declare new variable []string
	sameArray := []string{}
	diffArray := []string{}

	for key, v := range mapFruit {
		if v > 1 {
			sameArray = append(sameArray, key)
		} else {
			diffArray = append(diffArray, key)
		}
	}

	return sameArray, diffArray
}

func main() {
	/* result := totalFruits(
		[]string{"Apel", "Jambu", "Melon", "Jambu"},
		[]string{"Tomat", "Apel", "Melon", "Durian"})
	fmt.Println(result) */

	sameArray, diffArray := diffFruit(
		[]string{"Apel", "Jambu", "Melon", "Jambu"},
		[]string{"Tomat", "Apel", "Melon", "Durian"})

	fmt.Println("Same Array:", sameArray)
	fmt.Println("Different Array :", diffArray)
	//deleteMap()
	//initMap()
}
