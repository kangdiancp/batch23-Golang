package main

import (
	"fmt"
	"math"
	"strconv"
)

func intToFloat() {
	x := 315
	var d float64 = float64(x)
	println(d, x)
	fmt.Printf("Int TO Float%f", d)
	fmt.Printf("\nInt TO Float: %8.2f", d)

	fmt.Printf("\nInteger base 10 : %d", x)
	fmt.Printf("\nOctal : %b b", d)

	//using strconv

	format := strconv.FormatFloat(d, 'f', 5, 64)
	println("format : ", format)

	// func Sprinf
	SprintFormat := fmt.Sprintf("Sprintf : %8.2f", d)
	println(SprintFormat)

}

func floaToInt() {
	f := 509.98
	i := int64(f)
	formatInt := strconv.FormatInt(i, 10)
	fmt.Println("\n Float TO Int : ", i, formatInt)

}

func intToString() {
	price := 45
	totalprice := 198.78

	//concat

	output := "Total Unit Price" + strconv.Itoa(price)
	//ditampung
	//temp : int(totalprice)
	temp := int(totalprice)
	total := "Total Price:" + strconv.Itoa(temp)
	fmt.Println(output)
	fmt.Println(total)

}

func totalprice() {
	price := 98.9
	qty := 5
	totalprice := price * float64(qty)
	fmt.Println("Total Price : ", totalprice)
	fmt.Println("Total Price : ", math.Round(totalprice))
}

func parseTOInt() {
	vall := "98"
	int1, err := strconv.ParseInt(vall, 10, 64)
	if err == nil {
		fmt.Println("parsed value", int1)

	} else {
		fmt.Println("can not parse", vall, err)
	}
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
	intToFloat()
	floaToInt()
	intToString()
	totalprice()
	parseTOInt()
	parseStringToBool()
}
