package main

import "fmt"

//global variable
var salary = 1_000_000 //_untuk memisahkan angka

func varDeclare() {
	//invit variable default zero value
	var a int       //default value :0 variable bisa di isi ulang dan dipanggil lagi
	var word string //default value :"" (const tdk bisa di isi ulang)

	// variable type
	var x int = 10

	// multiple varibale with the same typ
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
	fmt.Println("a:", a)
	fmt.Println("x:", x)
	fmt.Println("y:", y)
	fmt.Println("z:", z)
	fmt.Println("str:", str)
	fmt.Println(k, q, xy, zy, msg, f, g, word)
}

func shortVariable() {
	//declare variable using :=operator
	var a = 10 // untype variable
	x := 10    //ini namanya statement

	z, y := 12, "Golang"

	//re-assign value for variable x
	x = 100
	println(a, x, z, y)
}

func main() {
	fmt.Println("hello kusum")
	varDeclare()
	shortVariable()
}
