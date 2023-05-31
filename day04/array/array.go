package main

import (
	"fmt"
	"reflect"
)

func initArrayString() {
	var emps [3]string
	emps[0] = "budi"
	emps[1] = "asep"
	emps[2] = "charlie"

	fmt.Println(emps)
	fmt.Println(emps[1])
	displayArray(emps)

	// short variable

	bts := [3]string{"jugkok", "surga", "jimin"}
	bts[1] = "lisa"
	fmt.Println(bts)

	//[...]operator
	student := [...]string{"asep", "budi", "charlie", "dedy"}
	fmt.Println(student)

}

func displayArray(arr [3]string) {
	for i := 0; i < len(arr); i++ {
		fmt.Println(arr[i], " ")
	}
}

func compareArray() {
	emp1 := [3]string{"asep", "upin", "ipin"}
	emp2 := [3]string{"asep", "upin", "ipin"}

	compare := reflect.DeepEqual(emp1, emp2)
	fmt.Println("result compare slice:", compare)
}
func main() {
	compareArray()

}
