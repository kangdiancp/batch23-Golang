package main

import "fmt"

func initArrayString() {
	//initial array tanpa inisialisasi
	var emps [3]string
	emps[0] = "budi"
	emps[1] = "asep"
	emps[2] = "charlie"

	fmt.Println(emps)
	fmt.Println(emps[1])
	displayArray(emps)

	//short variable
	bts := [3]string{"jungkok", "suga", "jimin"}
	bts[1] = "lisa"
	fmt.Println(bts)

	//[...] operaor
	student := [...]string{"Asep", "budi", "charlie", "dedy"}
	fmt.Println(student)
}

func displayArray(arr [3]string) {
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
	initArrayString()
	compareArray()
}
