package main

import (
	"fmt"
	"reflect"
	"sort"
)

func main() {
	// initSlice()
	// appendOneSlice()
	// cobaDulu()
	// createSliceFromArray()
	// copySlice()
	// replaceElementSlice()
	sortSlice()
	compareArray()
}

// init array slice
func initSlice() {
	//init slice with literal value
	nums := []int{1, 2, 3, 4, 5}
	nums = append(nums, 6, 7)

	fmt.Println(nums)

	//using method to create a slice array
	tech := make([]string, 3, 5)
	tech[0] = "c#"
	tech[1] = "java"
	tech[2] = "go"
	//tech[3] = "python"   Sudah melebihi kapasitas, maka masuk ke function append

	//append to add new element
	tech = append(tech, "python", "c++", "javascript", "Ruby")
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
	techs := [6]string{"java", "golang", "c#", "oracle", "postgre", "mongodb"}

	prog := techs[0:3]
	database := techs[3:6] //atau techs[3:] --> sesudah titik dua menandakan semua data diambil setelah index yg dipilih

	database[1] = "sql server"

	fmt.Println(prog)
	fmt.Println(database)
	fmt.Println(techs)

}

func copySlice() {
	techs := [6]string{"java", "golang", "c#", "oracle", "postgre", "mongodb"}

	//1. fetch element oracle, postgre, mongodb
	pl := techs[3:]

	//2. create slice variable db
	//jika mau mengubah isi array, maka jangan diganggu hasil slice pl, maka buat slice baru yaitu db dan ini tergantung case juga
	db := make([]string, 2)

	//3. use method copy, args1: destination, arg2 : source
	copy(db, pl)
	db[1] = "aws"

	fmt.Println(db)
}

func replaceElementSlice() {
	techs := []string{"java", "golang", "c#", "oracle"}
	replaceElmnt := []string{"vb", "pascal"}

	copy(techs[0:2], replaceElmnt)
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

func compareArray() {
	emp1 := [3]string{"asep", "upin", "ipin"}
	emp2 := [3]string{"asep", "upin", "ipin"}

	//untuk mengecek apakah nilai dari jumlah index dan isi dari array tersebut sama atau tidak, jikalau sama, maka akan dikembalikan nilai true, dan sebaliknya
	// compare := emp1 == emp2
	//atau
	compare := reflect.DeepEqual(emp1, emp2)

	fmt.Println("Result of Compare is", compare)
}

func cobaDulu() {
	tech := make([]string, 3)
	tech[0] = "c#"
	tech[1] = "go"
	tech[2] = "php"

	fmt.Println(tech)
}
