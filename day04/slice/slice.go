package main

import (
	"fmt"
	"sort"
)

//init array slice

func initSlice() {
	// init slice with literal value
	nums := []int{1, 2, 3, 4, 5}
	fmt.Println(nums)

	// using method to create as slice array
	tech := make([]string, 3, 6)
	tech[0] = "c#"
	tech[1] = "java"
	tech[2] = "go"
	tech[3] = "python"

	//append to add new element
	tech = append(tech, "python", "sql", "oracle", "excel", "java8")
}

func appendOneSlice() {
	tech := make([]string, 3, 6)
	tech[0] = "c#"
	tech[1] = "java"
	tech[2] = "c#"
	tech[3] = "go"

	framework := []string{"springboot", "flutter"}
	tech = append(tech, framework...)
	fmt.Println(tech)
}

func createSlicefromArray() {
	techs := []string{"java", "golang", "ch#", "oracle", "postgre", "mongodb"}

	techs = append(techs, "flute", "android")

	prog := techs[0:3]
	database := techs[3:]
	database[1] = "sql server"

	fmt.Println(prog)
	fmt.Println(database)
}

func copySlice() {
	techs := []string{"java", "goang", "c#", "oracle", "postgre", "mongodb"}

	//1.fetch element oracle,postgre,mongodb
	p1 := techs[3:]

	//2.create slice variable db
	db := make([]string, 3)

	//3. use method copy, args1 : destination,arg2 : source
	copy(db, p1)

	fmt.Println(db)
}

func replaceElementSlice() {
	techs := []string{"java", "golang", "c#", "oracle"}
	replacE1 := []string{"vb", "pascal"}

	copy(techs[0:2], replacE1)
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

func main() {
	//createSlicefromArray()
	//copySlice()
	//replaceElementSlice()
	sortSlice()

}
