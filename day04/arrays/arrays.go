package main

import "fmt"

func initArrayString() {
	var emps [3]string
	emps[0] = "budi"
	emps[1] = "asep"
	emps[2] = "charlie"

	fmt.Println(emps)
	fmt.Println(emps[1])

	// short variable
	bts := [3]string{"jungkok", "suga", "jimin"}
	bts[1] = "lisa"
	fmt.Println(bts)

	// operator
	student := [...]string{"Asep", "Budi", "Charlie", "Dedy"}
	fmt.Println(student)

	// displayArray(student)
}

func displayArray(arr []string) {
	for i := 0; i < len(arr); i++ {
		fmt.Print(arr[i], " ")
	}
}

func compareArray() {
	emp1 := [3]string{"asep", "upin", "ipin"}
	emp2 := [3]string{"asep", "upin", "ipin"}

	compare := emp1 == emp2

	fmt.Println("Result compare", compare)
}

func main() {
	// initArrayString()
	compareArray()
}
