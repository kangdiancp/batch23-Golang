package main

import "fmt"

// global variable
var salary = 1_000_000_000

// Contoh Error (tdk bisa menggunakan shorthand)
// tax := 45

func varDeclare() {
	// Init Variable default zero value
	var a int       // default value = 0
	var word string // default value = ""

	// variale type
	var x int = 10

	// multiple variable with the same type
	var y, z int = 15, 20

	// multiple variable untype
	var k, str = 45, "golang"

	// declare multiple variable at once
	var (
		q    int
		xy       = 20
		zy   int = 12
		msg      = "golang"
		f, g string
	)

	println("a : ", a)
	println("x : ", x)
	println("y : ", y)
	println("z : ", z)
	println("k : ", k)
	println("str : ", str)
	println("q : ", q)
	println("xy : ", xy)
	println("zy : ", zy)
	println("msg : ", msg)
	println("f : ", f)
	println("g : ", g)
	println("word : ", word)
}

func shortVariable() {
	// declare variable using := operator
	var a = 10 // untype variable
	x := 10

	z, y := 12, "Golang"

	// re-assign value for variable x
	x = 100
	println(a, x, z, y)
}

func main() {
	fmt.Println("Hello World!")
	varDeclare()
	shortVariable()
}
