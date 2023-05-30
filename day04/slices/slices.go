package main

import (
	"fmt"
	"reflect"
	"sort"
)

func initSlices() {
	//init slice with literal value
	num := []int{1, 2, 3, 4, 5}
	fmt.Println(num)

	//using method to create a slice array
	tech := make([]string, 3, 6)
	tech[0] = "c#"
	tech[1] = "java"
	tech[2] = "go"
	//tech[3] = "python"

	//append to add new element
	tech = append(tech, "pythoon", "sql", "oracle", "excel", "java8")
	fmt.Println(tech)
}

func appendOneSlice() {
	tech := make([]string, 3, 6)
	tech[0] = "c#"
	tech[1] = "java"
	tech[2] = "go"

	framework := []string{"springboot", "flutter"}
	tech = append(tech, framework...)
	fmt.Println(tech)
}

func createSliceFromArray() {
	tech := [6]string{"java", "golang", "c#", "oracle", "postgre", "mongodb"}
	prog := tech[0:3]
	database := tech[3:]

	fmt.Println(prog, database)
}

func copySlice() {
	tech := [6]string{"java", "golang", "c#", "oracle", "postgre", "mongodb"}
	pl := tech[3:]          //fetch element oracle, postgre, mongodb
	db := make([]string, 2) //create slice variable db
	// use method copy, args1:destination,args:source
	copy(db, pl)
	db[1] = "aws"
	fmt.Println(tech, db, pl)
}

func replaceElementSlice() {
	tech := []string{"java", "golang", "c#", "oracle"}
	replaceEl := []string{"vb", "pascal"}
	//replace slice with slice
	copy(tech[0:2], replaceEl)
	fmt.Println(tech)
}

func sortSlice() {
	tech := []string{"java", "golang", "c#", "oracle"}
	sort.Strings(tech)

	fmt.Println(tech)

	for _, v := range tech[0:3] {
		fmt.Println(v)
	}
}

func compareSlices() {
	emp1 := []string{"asep", "upin", "ipin"}
	emp2 := []string{"asep", "upil", "ipin"}

	compare := reflect.DeepEqual(emp1, emp2)
	fmt.Println("Result compare slide : ", compare)
}

func main() {
	// initSlices()
	// appendOneSlice()
	// createSliceFromArray()
	// copySlice()
	// replaceElementSlice()
	// sortSlice()
	compareSlices()
}
