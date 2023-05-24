package main

import "fmt"

//global variable

func main() {
	fmt.Println("Hello World")
	varDeclare()
	shortVariable()
}
func varDeclare() {
	//init variable default zero value
	var a int       //default value : 0
	var word string //default value = ""

	//variable type
	var x int = 10

	//multiple variable with the same type
	var y, z int = 15, 20

	//multiple variable untype
	var k, str = 45, "golang"

	//declare multiple variable at once
	var (
		q    int
		xy       = 20
		zy   int = 12
		msg      = "golang"
		f, g string
	)

	println("a: ", a)
	println("k: ", k)
	println("str: ", str)
	println("zy: ", zy)
	println("xy: ", xy)
	println("z: ", z)
	println(q, msg, y, f, g, word, x)
}

func shortVariable() {
	//declare variable using := operator
	var a = 10 //untype variable
	x := 10

	z, y := 12, "golang"

	//reassign value for variable x
	// x = 100

	println("ShortVariable")

	print("a : ", a)
	print("\nx : ", x)
	print("\nz : ", z)
	print("\ny : ", y)

}
