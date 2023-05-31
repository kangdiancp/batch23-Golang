package main

import (
	"fmt"
	"reflect"
	"sort"
)

func intSlice() {
	tech := make([]string, 3, 6)
	tech[0] = "c"
	tech[1] = "java"
	tech[2] = "go"
	//tech[3] = "python"

	tech = append(tech, "python", "sql", "oracle", "exel", "javas")
}

//gabung 2 array

func appendArray() {
	tech := make([]string, 3, 6)
	tech[0] = "c"
	tech[1] = "java"
	tech[2] = "go"

	framework := []string{"golang", "flutter"}

	tech = append(tech, framework...)

	fmt.Println(tech)
}

func misahArray() {
	tech := []string{"java", "golang", "c#", "oracle", "postgre", "mongodb"}

	tech = append(tech, "sql")
	//tech [0,3] 0= idex mulai mengambil, 3 = sampai index mana yang mau diambil
	prog := tech[0:3]
	database := tech[3:6]

	fmt.Println(prog)
	fmt.Println(database)
}

func copyArray() {
	tech := []string{"java", "golang", "c#", "oracle", "postgre", "mongodb"}
	//potong technya
	pl := tech[3:]

	//bikin variable array dan copykan plnya
	db := make([]string, 3)
	//metode copy, args 1= destination, args 2=source
	copy(db, pl)
	//pl[1] = "sons"

	db[1] = "aws"

	fmt.Println(tech)
	//fmt.Println(pl)
	fmt.Println(db)
}

func replaceElementSlice() {
	tech := []string{"java", "golang", "c#", "oracle"}

	replaceEl := []string{"db", "pascal"}

	copy(tech[0:2], replaceEl)
	fmt.Println(replaceEl)
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

func kompareSlice() {
	emp1 := []string{"asep", "bool", "hyuga"}
	emp2 := []string{"asep", "bool", "hyuga"}

	compare := reflect.DeepEqual(emp1, emp2)

	fmt.Println("Reuslt Cpmpare", compare)
}

func main() {
	//intSlice()
	//appendArray()
	//misahArray()
	//copyArray()
	//replaceElementSlice()
	//sortSlice()
	kompareSlice()
}
