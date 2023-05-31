package main

import (
	"fmt"
	"reflect"
	"sort"
)

// init array slice
func initSlice(){
	// init slices with literal value
	nums := []int{1, 2, 3, 4, 5}
	fmt.Println(nums)

	// using method
	tech := make([]string, 3, 6)
	tech[0] = "C#"
	tech[1] = "Java"
	tech[2] = "Go"
	//tech[3] = "python"

	// append to add new element
	tech = append(tech, "python", "sql", "oracle", "excell", "java8")
	fmt.Println(tech)
}

func appendOneslice(){
	tech := make([]string, 3, 6)
	tech[0] = "C#"
	tech[1] = "Java"
	tech[2] = "Go"

	framework := []string {"springboot", "flutter"}

	tech = append(tech, framework...)

	fmt.Println(tech)
}

func createSliceFromArray(){
	techs := []string{"java", "golang", "c#", "oracle", "postgre", "mongodb"}

	prog := techs[0:3]
	database := techs[3:] // reference yang dapat merubah nilai  induknya
	database[1] = "sql server"

	fmt.Println(prog)
	fmt.Println(database)
	fmt.Println(techs)
}

func copySlice(){
	techs := []string{"java", "golang", "c#", "oracle", "postgre", "mongodb"}
	// fetch 3 element oracle, postgre, ,mongodb
	pl := techs[3:]

	// create slice variable db
	db := make([]string, 3)

	// use method copy, args1 : destination, args2 : source
	copy(db, pl)
	db[1] = "aws"

	fmt.Println(db)
	fmt.Println(techs)
}

func replaceElementSlice(){
	techs := []string{"java", "golang", "c#", "oracle", "postgre", "mongodb"}
	replaceE1 := []string{"vb", "pascal"}

	copy(techs[0:2], replaceE1)
	fmt.Println(techs)
}

func sortSlice(){
	techs := []string{"java", "golang", "c#", "oracle"}
	sort.Strings(techs)

	fmt.Println(techs)

	for _, v := range techs[0:] {
		fmt.Println(v)
	}
}

func compareSlice(){
	emp1 := []string{"asep", "upin", "ipin"}
	emp2 := []string{"asep", "upin", "ipin"}

	compare := reflect.DeepEqual(emp1, emp2)

	fmt.Println("Result Compare", compare)
}

func main() {
	//initSlice()
	//appendOneslice()
	//createSliceFromArray()
	//copySlice()
	//replaceElementSlice()
	//sortSlice()
	compareSlice()
}