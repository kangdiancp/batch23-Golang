package main

import (
	"fmt"
	"reflect"
	"sort"
)

// init array slice
func initSlice() {
	// init slie with literal value
	nums := []int{1, 2, 3, 4, 5}
	fmt.Println(nums)

	// using method to create a slice array
	tech := make([]string, 3, 6)
	tech[0] = "c#"
	tech[1] = "java"
	tech[2] = "go"
	//tech[3] = "python"

	//append to add new element
	tech = append(tech, "python", "sql", "oracle", "excel", "java8")

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
	techs := []string{"java", "golang", "c#", "oracle", "postgre", "mongodb"}

	techs = append(techs, "flute", "android")

	prog := techs[0:3]
	database := techs[3:]
	database[1] = "sql server"

	fmt.Println(prog)
	fmt.Println(database)

}

func copySlice() {
	techs := []string{"java", "golang", "c#", "oracle", "postgre", "mongodb"}

	//1.fetch element oracle, postgre, mongodb
	pl := techs[3:]

	//2.create slice variable db
	db := make([]string, 2)

	//3.use method copy, args1: destination,arg2 : source
	copy(db, pl)
	db[1] = "aws"

	fmt.Println(db)
}

func replaceElementSlice() {
	techs := []string{"java", "golang", "c#", "oracle"}
	replaceEl := []string{"vb", "pascal"}

	copy(techs[0:2], replaceEl)
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
	emp1 := []string{"asep", "upin", "ipin"}
	emp2 := []string{"asep", "upin", "ipin"}

	compare := reflect.DeepEqual(emp1, emp2)

	fmt.Println("Result compare slice : ", compare)
}

func main() {
	sortSlice()
	//replaceElementSlice()
	//copySlice()
	//createSliceFromArray()
	//appendOneSlice()
	//initSlice()
}
