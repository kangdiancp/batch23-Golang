package main

import (
	"fmt"
)

func initMap() {
	// Menyimpan nilai Nill
	emps := map[string]int{}
	fmt.Println(emps)

	//Using Make
	prdocuts := make(map[string]float64)

	prdocuts["TV"] = 90.000
	prdocuts["Laptop"] = 123.98
	prdocuts["Meja"] = 98.00
	prdocuts["Laptop"] = 100.98

	fmt.Println("Price TV: ", prdocuts["TV"])

	// Literal Value
	students := map[string]float64{
		"Widi": 98.99,
		"Budi": 90.76,
		"Dedi": 80.65,
	}

	for key, v := range students {
		fmt.Println("Key : ", key, "| Value : ", v)
	}
}

func deleteMap() {
	// Literal Value
	students := map[string]float64{
		"Widi": 98.99,
		"Budi": 90.76,
		"Dedi": 80.65,
		"Yuli": 97.88,
	}

	delete(students, "Dedi")

	// Check key value  widi
	if value, ok := students["David"]; ok {
		fmt.Println("Yuli already exist : ", value)
		delete(students, "Yuli")
		fmt.Println(value, "has been deleted")
	} else {
		fmt.Println("No store value")
	}

	fmt.Println(students)
}

func totalFruits(fruit []string, vege []string) map[string]int {
	mergeArr := []string{}
	//variadic operator
	mergeArr = append(mergeArr, fruit...)
	mergeArr = append(mergeArr, vege...)

	mapFruit := map[string]int{}

	for key := range mergeArr {
		mapFruit[mergeArr[key]]++
	}

	return mapFruit
}

func diffFruit(fruit []string, vege []string) ([]string, []string) {
	mergeArr := []string{}
	//variadic operator
	mergeArr = append(mergeArr, fruit...)
	mergeArr = append(mergeArr, vege...)

	// Sorting Slice
	// sort.Strings(mergeArr)

	mapFruit := map[string]int{}

	for key := range mergeArr {
		mapFruit[mergeArr[key]]++
	}

	// declare new Variable []string
	sameArr := []string{}
	diffArr := []string{}

	for key, v := range mapFruit {
		if v > 1 {
			sameArr = append(sameArr, key)
		} else {
			diffArr = append(diffArr, key)
		}
	}

	return sameArr, diffArr
}

func main() {
	// initMap()
	// deleteMap()

	// result := totalFruits([]string{"Apel", "Jambu", "Melon", "Jambu"}, []string{"Tomat", "Apel", "Melon", "Durian"})
	// fmt.Println(result)

	sameArr, diffArr := diffFruit([]string{"Apel", "Jambu", "Melon", "Jambu"}, []string{"Tomat", "Apel", "Melon", "Durian"})
	fmt.Println(sameArr)
	fmt.Println(diffArr)

}
