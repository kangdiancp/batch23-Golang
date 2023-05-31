package main

import "fmt"

func arraString() {
	var ems [3]string
	ems[0] = "Budi"
	ems[1] = "Asep"
	ems[2] = "Carli"

	fmt.Println(ems)
	fmt.Println(ems[0])

	//short variabel
	displayArray(ems)

	bts := [3]string{"jongkook", "suga", "jumin"}
	bts[1] = "lisa"
	fmt.Println(bts)
	/// [....] operator
	student := [...]string{"asep", "bool", "hyuga", "dedy"}
	fmt.Println(student)
}

func compareArray() {
	emp1 := [3]string{"asep", "bool", "hyuga"}
	emp2 := [3]string{"asep", "bool", "hyuga"}

	compare := emp1 == emp2

	fmt.Println("Reuslt Cpmpare", compare)
}
func displayArray(arr [3]string) {
	for i := 0; i < len(arr); i++ {
		fmt.Println(arr[i], " ")
	}
}

func main() {
	//arraString()
	compareArray()
}
