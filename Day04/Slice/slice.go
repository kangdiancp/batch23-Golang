package main

import (
	"fmt"
	"reflect"
)

// init array slice
func initSlice() {
	// init slice with literal value
	// nums := []int{1, 2, 3, 4, 5}
	// fmt.Println(nums)

	// // Using method to create a slice array
	tech := make([]string, 3, 4)

	tech[0] = "C#"
	tech[1] = "Java"
	tech[2] = "Go"
	// tech[3] = "Python"

	//Append to add new element
	tech = append(tech, "Python", "PHP", "C++")
	fmt.Print(tech)
}

func appendSlice() {
	tech := make([]string, 3, 6)
	tech[0] = "C#"
	tech[1] = "Java"
	tech[2] = "Go"

	frameword := []string{"Springboot", "Fluter"}

	tech = append(tech, frameword...)

	fmt.Println(cap(tech))
}

// func createSliceFromArray() {
// 	techs := [6]string{"Java", "golang", "C#", "oracle", "postgre", "MongoDB"}

// 	techs = append(techs, "Flute", "Android")

// 	pro := techs[0:3]
// 	database := techs[3:]
// 	database[1] = "sql server"

// 	fmt.Println(pro)
// 	fmt.Println(database)
// }

func copySlice() {
	techs := []string{"Java", "golang", "C#", "oracle", "postgre", "MongoDB"}
	// 1. Fetch element oracle, postfre, mongodb
	pl := techs[3:]

	// 2. Create slice variable db
	db := make([]string, 3)

	// 3. Use method copy: destination, arg2 : source
	copy(db, pl)
	db[1] = "aws"

	fmt.Println(db)
	fmt.Println(techs)
}

// func replaceSlice() {
// 	techs := []string{"Java", "golang", "C#", "oracle", "postgre", "MongoDB"}
// 	replaceE1 := []string{"vb", "pascal"}

// 	copy(techs[0:2], replaceE1)
// 	fmt.Println(techs)
// }

// func sorSlice() {
// 	techs := []string{"Java"}
// 	sort.Strings(techs)
// 	fmt.Println(techs)

// 	for _, v := range techs[0:3] {
// 		fmt.Println(v)
// 	}
// }

func compareSlice() {
	emp1 := [3]string{"asep", "upin", "ipin"}
	emp2 := [3]string{"asep", "upin", "ipin"}

	compare := reflect.DeepEqual(emp1, emp2)

	fmt.Println("Result compare slice : ", compare)
}

func main() {
	// initSlice()
	// appendSlice()
	// createSliceFromArray()
	// copySlice()
	// replaceSlice()
	// sorSlice()
	compareSlice()
}
