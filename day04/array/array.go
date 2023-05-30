package main

import "fmt"

func initArrayString() {
	var emp [3]string
	emp[0] = "budi"
	emp[1] = "asep"
	emp[2] = "charlie"

	fmt.Println(emp)
	fmt.Println(emp[1])
	displayArray(emp)

	//short variabel
	bts := [3]string{"jungkok", "suga", "jimin"}
	bts[1] = "lisa"
	fmt.Println(bts)

	//[...]operator
	student := [...]string{"asep", "budi", "charlie"}
	fmt.Println(student)
}

func displayArray(arr [3]string) {
	for i := 0; i < len(arr); i++ {
		fmt.Print(arr[i], " ")
	}
}

func compareArray() {
	emp1 := [3]string{"asep", "upin", "ipin"}
	emp2 := [3]string{"asep", "upil", "ipin"}

	compare := emp1 == emp2
	fmt.Println(compare)
}

func main() {
	initArrayString()
	compareArray()
}
