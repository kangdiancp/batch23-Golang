package main

import "fmt"

func shortVariable() {
	// declare varible using := operator
	var a = 10 //untype variable
	x := 10
	z, y := 12, "golang"

	// re-assign value for variable x
	x = 100
	println(a, x, z, y)
}

func main() {
	fmt.Println("hallo world")
	varDeclare()
	shortVariable()
}
func varDeclare() {
	//int variable dafault xero value
	var a int       // default value 0
	var word string // value

	// variable type
	var x int = 10

	// mutiple variable with same type
	var y, z int = 15, 20

	// mutiple variable untype
	var k, str = 45, "golang"

	// mutiple  variable at once
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
	println(q, xy, zy, msg, f, g, word)
}
