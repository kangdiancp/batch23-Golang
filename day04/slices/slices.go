package main

import (
	"fmt"
	"reflect"
	"sort"
)

// init array slice
func initSlice() {
	//init slice with literal value
	nums := []int{1, 2, 3, 4, 5}
	fmt.Println(nums)

	// using method to create a slice array
	tech := make([]string, 3, 7)
	tech[0] = "c#"
	tech[1] = "java"
	tech[2] = "go"
	//tech[3] = "pyton"

	//append to add new element
	tech = append(tech, "pyton", "sql", "oracle", "excel", "java8")
}

func appendOneSlice() {
	tech := make([]string, 3, 7)
	tech[0] = "c#"
	tech[1] = "java"
	tech[2] = "go"

	framework := []string{"springboot", "flutter"}

	tech = append(tech, framework...)
}

func createSliceFromArray() {
	techs := []string{"java", "golang", "c#", "oracle", "postgre", "mongodb"}

	techs = append(techs, "android", "flutter")

	prog := techs[0:3]
	database := techs[3:]
	database[1] = "sql server"

	fmt.Println(prog)
	fmt.Println(database)
}

func copySlice() {
	techs := []string{"java", "golang", "c#", "oracle", "postgre", "mongodb"}

	//1. fetch element oracle, postgre, mongodb
	pl := techs[3:]

	//2. create slice variable db
	db := make([]string, 2)

	//3. use method copy, args1: destination, args2: source
	copy(db, pl)
	db[1] = "aws"

	fmt.Println()
}

func replaceElementSlice() {
	techs := []string{"java", "golang", "c#", "oracle"}
	replaceE1 := []string{"vb", "pascal"}

	copy(techs[0:2], replaceE1)
	fmt.Println(techs)
}

func sortSlice() {
	techs := []string{"java", "golang", "c#", "oracle"}
	sort.Strings(techs)

	fmt.Println(techs)

	for _, v := range techs[0:3] {
		fmt.Println(v)
	}
}

func compareSlice() {
	emp1 := []string{"asep", "abdul", "jamet"}
	emp2 := []string{"asep", "abdul", "jamet"}

	compare := reflect.DeepEqual(emp1, emp2)

	fmt.Println("result compare slice : ", compare)
}

func main() {
	//initSlice()
	//appendOneSlice()
	//createSliceFromArray()
	//copySlice()
	//replaceElementSlice()
	//sortSlice()
	compareSlice()
}
