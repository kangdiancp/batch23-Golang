package main

import "fmt"

func main() {
	// initMap()
	// deleteMap()
	result := totalFruits(
		[]string{"Apel", "Mangga", "Jambu", "Melon", "Jambu"},
		[]string{"Tomat", "Apel", "Melon", "Jambu", "Mangga"})
	fmt.Println(result)

	sameArray, diffArray := diffFruits(
		[]string{"Apel", "Mangga", "Jambu", "Melon", "Jambu"},
		[]string{"Tomat", "Apel", "Melon", "Jambu", "Mangga"})
	fmt.Println(sameArray)
	fmt.Println(diffArray)
}

func initMap() {
	//store nil
	emps := map[string]int{}
	fmt.Println(emps)

	//using make
	products := make(map[string]float64)

	products["TV"] = 90.00
	products["Laptop"] = 123.98

	products["Meja"] = 89.00
	products["Laptop"] = 100

	fmt.Println("Price Tv : ", products["Laptop"])

	students := map[string]float64{
		"Widi": 98.99,
		"Budi": 90.76,
		"Dedi": 87.65,
	}

	for key, value := range students {
		fmt.Println("Key : ", key, "Value : ", value)
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

	if value, ok := students["widi"]; ok {
		fmt.Println("Widi Already Exist : ", value)
		delete(students, "Budi")
		fmt.Println(value, "Has been deleted")
	} else {
		fmt.Println("No Store Value")
	}
	fmt.Println(students)
}

func totalFruits(fruit []string, vege []string) map[string]int {
	mergeArray := []string{}

	mergeArray = append(mergeArray, fruit...)
	mergeArray = append(mergeArray, vege...)

	//declare mapFruit to store key
	mapFruit := map[string]int{}
	for key := range mergeArray {
		mapFruit[mergeArray[key]]++
	}

	return mapFruit
}
func diffFruits(fruit []string, vege []string) ([]string, []string) {
	mergeArray := []string{}

	mergeArray = append(mergeArray, fruit...)
	mergeArray = append(mergeArray, vege...)

	//declare mapFruit to store key
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
