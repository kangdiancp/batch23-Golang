package main

import (
	"fmt"
	"sort"
)

func initMap() {
	//store nil
	emps := map[string]int{}
	fmt.Println(emps)

	//using make
	products := make(map[string]float64)
	products["TV"] = 90.00
	products["Laptop"] = 123.98
	products["Meja"] = 98.00
	products["Laptop"] = 100.98

	fmt.Println("price TV :", products["Laptop"])

	//literal value
	students := map[string]float64{
		"widi": 98.99,
		"bayu": 90.76,
		"ari":  87.65,
	}

	for key, v := range students {
		fmt.Println("key :", key, "value :", v)
	}

}

func deleteMap() {
	students := map[string]float64{
		"widi": 98.99,
		"bayu": 90.76,
		"ari":  87.65,
		"yudi": 97.88,
	}
	delete(students, "bayu")

	//check key value widi
	if value, ok := students["yudi"]; ok {
		fmt.Println("widi already exist:", value)
		delete(students, "yuni")
		fmt.Println(value, "has been deleted")
	} else {
		fmt.Println("no store value")
	}

	fmt.Println(students)
}

func totalFruits(fruit []string, vege []string) ([]string, []string) {
	mergeArray := []string{}
	//variadic operator
	mergeArray = append(mergeArray, fruit...)
	mergeArray = append(mergeArray, vege...)

	//declare mapFruit to store key,value
	mapFruit := map[string]int{}
	for key := range mergeArray {
		mapFruit[mergeArray[key]]++
	}

	//declare new variable []string
	sameArray := []string{}
	diffArray := []string{}

	return sameArray, diffArray

}

func diffFruits(fruit []string, vege []string) ([]string, []string) {
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
	sameArray, diffArray := diffFruits([]string{"apel", "jambu", "markisa", "jambu"},
		[]string{"tomat", "jagung", "apel", "wortel"})
	fmt.Println("same array:", sameArray)
	fmt.Println("different array :", diffArray)

	// fmt.Println(totalFruits([]string{"apel", "jambu", "markisa", "jambu"},
	// 	[]string{"tomat", "jagung", "apel", "wortel"}))
	//initMap()
	//deleteMap()

}
