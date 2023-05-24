package main

import "fmt"

func main() {

	fmt.Println(". ")
	varDeclare()
	shortVariable()
}
func varDeclare() {

	// init variable default zero value
	var a int       // default value = 0
	var word string // default value ""

	// variable type
	var x int = 10

	// multiple variable with the same type
	var y, z, k1 int = 15, 20, 40

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

	fmt.Println("a = ", a)
	fmt.Println("word = sehat ", word)
	fmt.Println(x)
	fmt.Println(y)
	fmt.Println(z)
	fmt.Println(k1)
	fmt.Println(str)
	fmt.Println(k)
	fmt.Println(q, xy, zy, msg, f, g)

}

func shortVariable() {
	var aa = 10 // untype variable
	xa := 20

	ni, sa := 20, "golang"

	fmt.Println(aa)
	fmt.Println(xa)
	fmt.Println(ni, sa)
}
