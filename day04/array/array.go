package main

import "fmt"

func initArrayString() {
	var students [3]string
	students[0] = "Fatah"
	students[1] = "Adjik"
	students[2] = "Alip"

	fmt.Println(students)
	fmt.Println(students[1])
	displayArray(students[:])

	// Short Variable
	cagur := [3]string{"Deni", "Wendi", "Nardji"}
	cagur[2] = "Rigen"
	// cagur[4] = "Indra" // Can't
	fmt.Println(cagur)

	// [...] Operator
	emps := [...]string{"Asep", "Joko", "Doni", "Dendy"}
	fmt.Println(emps)
}

func displayArray(arr []string) {
	for i := 0; i < len(arr); i++ {
		fmt.Print(arr[i], " ")
	}
}

func main() {
	initArrayString()
}
