package main

import (
	"fmt"
	"sort"
)

func initMap() {
	emp := map[string]int{}
	fmt.Println(emp)

	prod := make(map[string]float64)
	prod["TV"] = 90.00
	prod["Laptop"] = 123.98

	fmt.Println("Price TV : ", prod["TV"])

	//literal value
	stud := map[string]float64{
		"widi": 98.99,
		"budi": 90.76,
		"dedi": 87.65,
	}
	for key, v := range stud {
		fmt.Println("Key : ", key, "value: ", v)
	}
}

func deleteMap() {
	stud := map[string]float64{
		"widi": 98.99,
		"budi": 90.76,
		"dedi": 87.65,
		"yuli": 97.88,
	}

	delete(stud, "dedi")

	//check key value widi
	if value, ok := stud["yuli"]; ok {
		fmt.Println("yuli already exist : ", value)
		delete(stud, "yuli")
		fmt.Println(value, "has been deleted")
	} else {
		fmt.Println("No store value")
	}

	fmt.Println(stud)
}

func totalFruit(fruit []string, vege []string) map[string]int {
	mergeArray := []string{}
	mergeArray = append(mergeArray, fruit...)
	mergeArray = append(mergeArray, vege...)

	mapFruit := map[string]int{}
	for key := range mergeArray {
		mapFruit[mergeArray[key]]++
	}

	return mapFruit
}

func initDiffruit(fruit []string, vege []string) ([]string, []string) {
	mergeArray := []string{}
	mergeArray = append(mergeArray, fruit...)
	mergeArray = append(mergeArray, vege...)

	sort.Strings(mergeArray)

	mapFruit := map[string]int{}

	for key := range mergeArray {
		mapFruit[mergeArray[key]]++
	}

	//declare new variable [ string]
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
	//initMap()
	//deleteMap()
	// result := totalFruit([]string{"apel", "jambu", "melon", "jambu"}, []string{"tomat", "apel", "melon", "durian"})
	// fmt.Println(result)

	sameArray, diffArray := initDiffruit([]string{"apel", "jambu", "melon", "jambu"},
		[]string{"tomat", "apel", "melon", "durian"})
	fmt.Println("Same Array= ", sameArray)
	fmt.Println("Different Array= ", diffArray)

}
