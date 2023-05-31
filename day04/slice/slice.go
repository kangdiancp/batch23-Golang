package main

import (
	"fmt"
	"reflect"
	"sort"
)

// Init array slice
func initSlice() {
	// Init slice with literal value
	nums := []int{1, 2, 3, 4, 5}
	fmt.Println(nums)

	// Using method to create a slice array
	tech := make([]string, 3, 6)
	tech[0] = "Java"
	tech[1] = "Golang"
	tech[2] = "C#"
	// tech[3] = "Phyton" // Error

	// Append to add new element
	tech = append(tech, "Phyton", "MySQL", "Oracle", "Excel", "Flutter")
	fmt.Println(tech)
}

func appendOneSlice() {
	tech := make([]string, 3, 6)
	tech[0] = "Java"
	tech[1] = "Golang"
	tech[2] = "C#"

	frameWork := []string{"SpringBoot", "Flutter"}

	tech = append(tech, frameWork...)

	fmt.Println(tech)
}

func createSliceFromArray() {
	techs := [6]string{"Java", "Golang", "C#", "Oracle", "Postgree", "MongoDB"}

	prog := techs[0:3]
	database := techs[3:]
	database[1] = "MsSQL"

	fmt.Println(prog)
	fmt.Println(database)
	fmt.Println(techs)
}

func copySlice() {
	techs := []string{"Java", "Golang", "C#", "Oracle", "Postgree", "MongoDB"}

	// 1. Fetch element Oracle, Postgree, MongoDB
	pl := techs[3:]

	// 2. Create slice variable db
	db := make([]string, 3)

	// 3. Use method copy || args1 : destinantion, args2 : source
	copy(db, pl)
	db[1] = "AWS"

	fmt.Println(db)
	fmt.Println(techs)
}

func replaceElementSlice() {
	techs := []string{"Java", "Golang", "C#", "Oracle"}
	replaceEl := []string{"VB.NET", "Pascal"}

	copy(techs[0:2], replaceEl)
	fmt.Println(techs)
}

func sortSlice() {
	techs := []string{"Java", "Golang", "C#", "Oracle"}

	// Method
	sort.Strings(techs)

	fmt.Println(techs)

	// Logic
	for _, v := range techs[0:3] {
		fmt.Println(v)
	}
}

func compareArray() {
	emp1 := [3]string{"Asep", "Upin", "Ipin"}
	emp2 := [3]string{"Udin", "Adin", "Indra"}

	compare := emp1 == emp2
	fmt.Println("Result Compare Array : ", compare)
}

func compareSlice() {
	emp1 := [3]string{"Asep", "Upin", "Ipin"}
	emp2 := [3]string{"Udin", "Adin", "Indra"}

	compare := reflect.DeepEqual(emp1, emp2)

	fmt.Println("Result Compare Slice : ", compare)
}

func main() {
	initSlice()
	appendOneSlice()
	createSliceFromArray()
	copySlice()
	replaceElementSlice()
	sortSlice()
	compareArray()
	compareSlice()
}
