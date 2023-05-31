package main

import "fmt"

func initArrayString() {
	// initial array tanpa inisial
	var emps [3]string
	emps[0] = "budi"
	emps[1] = "asep"
	emps[2] = "charlie"

	fmt.Println(emps)

	// Short Initial
	bts := [3]string{"Jugkok", "Suga", "Jimmy"}
	bts[1] = "Lisa"

	fmt.Println(bts)

	// [...] Operator
	student := [...]string{"Asep", "Budi", "Charlie"}

	fmt.Println(student)
	displayArray(student)
}

func displayArray(arr [3]string) {
	for i := 0; i < len(arr); i++ {
		fmt.Println(arr[i], " ")
	}
}

func compareArray() {
	emp1 := [3]string{"asep", "upin", "ipin"}
	emp2 := [3]string{"asep", "upin", "Apin"}

	compare := emp1 == emp2

	fmt.Println("Result Compare", compare)

}

func main() {
	initArrayString()
	compareArray()
}
