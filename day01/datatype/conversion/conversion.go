package main

import (
	"fmt"
	"math"
	"strconv" //sebuah package
)

func main() {
	intToFloat()
	floatToInt()
	intToString()
	totalPrice()
	parserToInt()
	parserToFloat()
	parseToBool()
}

func intToFloat() {
	x := 315
	var d float64 = float64(x)

	println(d, x)

	fmt.Printf("Int to float %f", d)
	fmt.Printf("\nInt to float %8.2f", d)
	fmt.Printf("\nInt base 10 : %d", x)
	fmt.Printf("\nOctal : %b \n", d)

	//using strconv
	format := strconv.FormatFloat(d, 'f', 5, 64)
	println("format : ", format)

	//func sprintF
	sprintFormat := fmt.Sprintf("Sprintf : %8.2f", d)
	println("Sprint Format: ", sprintFormat)
}

func floatToInt() {
	f := 509.98
	i := int64(f)
	formatInt := strconv.FormatInt(i, 10)
	fmt.Println("\nFloat to int : ", i, formatInt)
}
func intToString() {
	price := 45
	totalPrice := 198.78
	//concat

	temp := int(totalPrice)

	output := "Total Unit Price " + strconv.Itoa(price)
	total := "Total Price: " + strconv.Itoa(temp)

	fmt.Println(output)
	fmt.Println(total)
}

func totalPrice() {
	price := 98.9
	qty := 5
	totalPrice := price * float64(qty)

	fmt.Println("Total Price : ", totalPrice)
	fmt.Println("Total Price : ", math.Round(totalPrice))
}

func parserToInt() {
	val1 := "98"
	//if err not error, err return nil
	int1, err := strconv.ParseInt(val1, 10, 64)
	if err == nil {
		fmt.Println("Parsed value : ", int1)
	} else {
		fmt.Println("Can not parse", val1, err)
	}
}
func parserToFloat() {
	val1 := "98.99"
	//if err not error, err return nil
	float1, err := strconv.ParseFloat(val1, 64)
	if err == nil {
		fmt.Println("Parsed value : ", float1)
	} else {
		fmt.Println("Can not parse", val1, err)
	}
}

func parseToBool() {
	val1 := "true"
	val2 := "false"
	val3 := "true"
	bool1, err1 := strconv.ParseBool(val1)
	bool2, err2 := strconv.ParseBool(val2)
	bool3, err3 := strconv.ParseBool(val3)

	if err1 == nil && err2 == nil && err3 == nil {
		fmt.Println("Bool1, Bool2, Bool3 : ", bool1, bool2, bool3)
	} else {
		fmt.Println("Can not parse ", val1, val2, val3, err1, err2, err3)
	}
}
