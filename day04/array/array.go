package main

import "fmt"

func initArrayString() {
	var emps [3]string
	emps[0] = "budi"
	emps[1] = "asep"
	emps[2] = "charlie"

	fmt.Println(emps)
	fmt.Println(emps[1])
	displayArray(emps)

	//short Variabel
	bts := [3]string {"jungkok", "suga", "jimin"}
	bts[1] = "lisa"

	fmt.Println(bts)

	//[...] Operator
	students := [...]string {"asep", "budi", "charlie", "dedi"}
	fmt.Println(students)
	//displayArray(students)
}

func displayArray(arr [3]string){
	for i := 0; i < len(arr); i++ {
		fmt.Print(arr[i], " ")
	}
}

func compareArray(){
	emp1 := [3]string{"asep", "upin", "ipin"}
	emp2 := [3]string{"asep", "upin", "ipin"}

	compare := emp1 == emp2

	fmt.Println("Result Compare", compare)
}

func main() {
	//initArrayString()
	compareArray()

}