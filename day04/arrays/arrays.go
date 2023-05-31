package main

import "fmt"

//array
func initArrayString() {
	var emp [3]string
	emp[0] = "kusuma"
	emp[1] = "lintang"
	emp[2] = "saputra"

	fmt.Println(emp)
	fmt.Println(emp[1])
	displayArray(emp[:])

	//short variable
	bts := [3]string{"bintoro", "yusuf", "afif"}
	bts[1] = "syafiq"

	fmt.Println(bts)

	//[...] operator
	student := [...]string{"rio", "albert", "rangga"}
	fmt.Println(student)
	// displayArray(student)
}

func displayArray(array []string) {
	for i := 0; i < len(array); i++ {
		fmt.Println(array[i], " ")
	}
}

func main() {
	initArrayString()
}
