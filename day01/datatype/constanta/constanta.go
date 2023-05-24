package main

import (
	"fmt"
	"strconv"
)

// global constant variable
const phi = 3.14159

func declareConstant() {
	//constant type
	const price float64 = 20.98

	//constant untype
	const tax = 11.9

	//multiple constant
	const (
		discount = 0.1
		phi      = 3.14159
	)
	println(price, tax, discount, phi)

}

func declareIotas() {
	const (
		seq1 = iota
		seq2
		seq3
	)
	var circleArea = phi * 5
	println(seq1, seq2, seq3, circleArea)
}

func parseStringToBool() {
	val1 := "true"
	val2 := "false"
	val3 := "not false"

	bool1, err1 := strconv.ParseBool(val1)
	bool2, err2 := strconv.ParseBool(val2)
	bool3, err3 := strconv.ParseBool(val3)

	fmt.Println("Bool1", bool1, err1)
	fmt.Println("Bool2", bool2, err2)
	fmt.Println("Bool3", bool3, err3)

	if err1 == nil {
		fmt.Println("Bool1", bool1)

	} else {
		fmt.Println("error parse bool1", bool1, err1)

	}

}

func main() {
	declareConstant()
	declareIotas()

}
