package main

import "fmt"

func initArrayString() {
	// INITIAL ARRAY TANPA INISIALISASI
	var emps [3]string
	emps[0] = "budi"
	emps[1] = "asep"
	emps[2] = "charlie"

	fmt.Println(emps)
	fmt.Println(emps[1])
	displayArray(emps)

	//SHORT VARIABLE
	bts := [3]string{"jungkok, jimin, suga"}
	bts[1] = "lisa"
	fmt.Println(bts)

	//[...] operator
	student := [...]string{"asep", "budi", "charlie", "dandi"}
	fmt.Println(student)
}

func displayArray(arr [3]string) {
	for i := 0; i < len(arr); i++ {
		fmt.Println(arr[i], " ")
	}
}

func compareArray() {
	emp1 := [3]string{"asep", "abdul", "jamet"}
	emp2 := [3]string{"asep", "abdul", "jamet"}

	compare := emp1 == emp2

	fmt.Println("result compare", compare)
}

func main() {
	//initArrayString()
	compareArray()
}
