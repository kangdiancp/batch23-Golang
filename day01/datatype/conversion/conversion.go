package main

import (
	"fmt"
	"math"
	"strconv"
)

func intToFlot() {
	x := 315
	var d float64 = float64(x)
	println(d, x)
	fmt.Printf("Int To Float %f", d)
	fmt.Printf("\nInt To Float %8.2f", d)

	fmt.Printf("\nInteger base 10 %d", x)
	fmt.Printf("\nOctal : %b", d)

	//using strconv
	format := strconv.FormatFloat(d, 'f', 5, 64)
	println("\nformat : ", format)

	// func SprintF
	sprintFormat := fmt.Sprintf("SprintF : %8.2f", d)
	println(sprintFormat)
}

func floatToInt() {
	f := 509.98
	i := int64(f)
	formatInt := strconv.FormatInt(i, 10)
	fmt.Println("\nFloat To Int : ", i)
	println(formatInt)
}

func intToString() {
	price := 45
	totalPrice := 198.78
	//concat
	output := "Total Unit Price " + strconv.Itoa(price)
	//ditampung
	//temp := int(totalprice)
	total := "Total Price : " + strconv.Itoa(int(totalPrice))
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

func parseToInt() {
	val1 := "98"
	//if err not error, err return nil
	int1, err := strconv.ParseInt(val1, 10, 64)
	if err == nil {
		fmt.Println("Parsed value : ", int1)
	} else {
		fmt.Println("Can not parse", val1, err)
	}
}

func parseToFloat() {
	val1 := "98.98"
	//if err not error, err return nil
	int1, err := strconv.ParseFloat(val1, 64)
	if err == nil {
		fmt.Println("Parsed value : ", int1)
	} else {
		fmt.Println("Can not parse", val1, err)
	}
}

func parseStringToBool() {
	val1 := "true"
	val2 := "false"
	val3 := "not true"

	bool1, err1 := strconv.ParseBool(val1)
	bool2, err2 := strconv.ParseBool(val2)
	bool3, err3 := strconv.ParseBool(val3)

	// fmt.Println("Bool2", bool2, err2)
	// fmt.Println("Bool3", bool3, err3)

	if err1 == nil {
		fmt.Println("Parse Bool1", bool1)
	} else {
		fmt.Println("Error Parse Bool1", bool1, err1)
	}

	if err2 == nil {
		fmt.Println("Parse Bool2", bool2)
	} else {
		fmt.Println("Error Parse Bool2", bool2, err2)
	}

	if err3 == nil {
		fmt.Println("Parse Bool3", bool3)
	} else {
		fmt.Println("Error Parse Bool3", bool3, err3)
	}
}

func main() {
	intToFlot()
	floatToInt()
	intToString()
	totalPrice()
	parseToInt()
	parseToFloat()
	parseStringToBool()
}
