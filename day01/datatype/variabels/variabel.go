package main

import "fmt"

func main() {
	fmt.Println("hello world")
	varDeclare()
	shortVariable()

}

func varDeclare() {
	// init variabel default zero values
	var a int       // default value :0
	var word string //default value : ""
	// varible type
	var x int = 10

	// multiple variable with the same type
	var y, z int = 15, 20

	// multiple variable untype
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
	println("x :", x)
	println("y :", y)
	println("z :", z)
	println("k :", k)
	println("str :", str)
	println(q, xy, zy, msg, f, word, g)

}

func shortVariable() {
	//declare variable using := operator
	var a = 10 //untype variable
	x := 10
	z, y := 12, "GOLANG"
	//re-assign value for variable x
	//x = 100
	print(a, x, z, y)
}
