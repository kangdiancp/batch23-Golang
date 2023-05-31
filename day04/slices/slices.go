package main

import (
	"fmt"
	"reflect"
	"sort"
)

// init array slice (array dinamis)
func initSlice() {
	// slice tipe data (array)/ init slice with literal value
	number := []int{1, 2, 3, 4, 5}
	fmt.Println(number)

	//using method to create a slice array
	tech := make([]string, 3)
	tech[0] = "c++s"
	tech[1] = "typescript"
	tech[2] = "go"
	// tech[3] = "python"

	//append to add new element
	tech = append(tech, "pyhton", "sql", "oracle", "excel", "java8")
	fmt.Println(tech)
}

func appendOneSlice() {
	tech := make([]string, 3, 6)
	tech[0] = "c++s"
	tech[1] = "typescript"
	tech[2] = "go"

	framework := []string{"springboot", "flutter"}
	tech = append(tech, framework...)
	fmt.Println(tech)
}

func createSliceFromArray() {
	techs := [6]string{"java", "golang", "c#", "oracle", "postgre", "mongodb"}

	prog := techs[0:3]
	database := techs[3:]

	fmt.Println(prog)
	fmt.Println(database)
}

func copySlice() {
	techs := []string{"java", "golang", "c#", "oracle", "postgre", "mongodb"}
	//1.fetch element oracle, postgre, mongodb
	pl := techs[3:]

	//2.create slice variable db
	db := make([]string, 3)

	//3.use method copy, args1:= destination, arg2 :source
	copy(db, pl)
	db[1] = "aws"

	fmt.Println(db)
}

func replaceElementSlice() {
	techs := []string{"java", "golang", "c#", "oracle"}
	replaceEl := []string{"vb", "pascal"}

	copy(techs[0:1], replaceEl)
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

func displayArray(arr [3]string) {
	for i := 0; i < len(arr); i++ {
		fmt.Print(arr[i], " ")
	}
}

func compareArray() {
	emp1 := [3]string{"rose", "upin", "ipin"}
	emp2 := [3]string{"rose", "upin", "ipin"}

	compare := emp1 == emp2
	fmt.Println("Result compare", compare)
}

func compareSlice() {
	emp1 := []string{"rose", "upin", "ipin"}
	emp2 := []string{"rose", "upin", "ipin"}

	compare := reflect.DeepEqual(emp1, emp2)
	fmt.Println("Result compare", compare)
}

func main() {
	// initSlice()
	// appendOneSlice()
	// createSliceFromArray()
	// copySlice()
	// replaceElementSlice()
	// // sortSlice()
	// displayArray()
	compareArray()
	// compareSlice()
}
