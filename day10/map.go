package main

import (
	"fmt"
	"sort"
)

func initMap() {
	// store nill
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
		"widi" : 98.99,
		"Budi" : 90.76,
		"Dedi" : 87.65,
	}

	for key, v := range students {
		fmt.Println("Key : ", key, "Value : ", v)
	}
}

func deleteMap(){
	students := map[string]float64{
		"widi" : 98.99,
		"Budi" : 90.76,
		"Dedi" : 87.65,
		"Yuli" : 97.88,
	}

	delete(students, "Dedi")
	fmt.Println(students)

	// check key value widi
	if value, ok := students["Yuli"]; ok{
		fmt.Println("Yuli Already Exist : ", value)
		delete(students, "Yuli")
		fmt.Println(ok, "has been Deleted")
	}else {
		fmt.Println("No store Value")
	}

	fmt.Println(students)
}

func totalFruits(fruit []string, vege []string) map[string]int{
	mergearray := []string{}
	// variadic operator
	mergearray = append(mergearray, fruit...)
	mergearray = append(mergearray, vege...)

	mapfruit := map[string]int{}

	// declare mapfruit to store key, value
	for key := range mergearray {
		mapfruit[mergearray[key]]++
	}
	return mapfruit
}

func diffFruit(fruit []string, vege []string) ([]string, []string){
	mergearray := []string{}
	// variadic operator
	mergearray = append(mergearray, fruit...)
	mergearray = append(mergearray, vege...)

	//sorting slice map
	sort.Strings(mergearray)

	mapfruit := map[string]int{}

	// declare mapfruit to store key, value
	for key := range mergearray {
		mapfruit[mergearray[key]]++
	}
	fmt.Println(mapfruit)

	// declare new variable array []string
	sameArray := []string{}
	diffArray := []string{}
	
	for key, v := range mapfruit {
		if v > 1{
			sameArray = append(sameArray, key)
		}else {
			diffArray = append(diffArray, key)
		}
	}
	sort.Strings(sameArray)
	sort.Strings(diffArray)
	return sameArray, diffArray
}

func main() {
	// initMap()
	// deleteMap()
	// result := totalFruits([]string{"Apel", "Jambu", "Melon", "Jambu"}, 
	// 			[]string{"Tomat", "Apel", "Melon", "Durian"})
	// fmt.Println(result)

	sameArray, diffArray := diffFruit([]string{"Apel", "Jambu", "Melon", "Jambu"}, 
			  						  []string{"Tomat", "Apel", "Melon", "Durian"})
	fmt.Println("Same Array : ",sameArray)
	fmt.Println("Different Array : ", diffArray)
}