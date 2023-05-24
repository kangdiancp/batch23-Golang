package main

import "fmt"

// global variable

func main() {
	fmt.Println("Hello World")
	varDeclare()
	shortVariable()
}

func varDeclare(){
	//init variable default zero value
	var a int			// default value = 0
	var word string		// default value = ""

	// Variable type
	var x int = 10

	// multiple variable with same type
	var y, z int = 15, 20

	// multiple variable untype
	var k, str = 45, "golang"

	// declare multiple variables at onces
	var (
		q int
		xy = 20
		zy int = 12
		msg = "golang"
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

func shortVariable(){
	// Declare variable using := operator
	var a = 10	// untype variable
	x := 10		// short variable

	z, y := 12, "golang"


	//reassign value for variable x
	// x = 100

	print("a : ", a)
	print("x : ", x)
	print("z : ", z)
	print("y : ", y)
}