package main

import "fmt"

func main() {
	initArrayString()

}

func initArrayString() {
	var emps [3]string
	emps[0] = "budi"
	emps[1] = "asep"
	emps[2] = "charlie"

	fmt.Println(emps[1])

	//declare using short variable
	bts := [3]string{"Wahyu", "dodi", "Ramadhan"}
	bts[1] = "lisa"

	fmt.Println(bts)

	displayArray(emps)

	//[...]operator
	student := [...]string{"asep", "damar", "diman"}
	fmt.Println(student)
}

func displayArray(arr [3]string) {
	for i := 0; i < len(arr); i++ {
		fmt.Print(arr[i], " ")
	}
}
