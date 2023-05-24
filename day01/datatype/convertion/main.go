package main

import (
	"fmt"
	"math"
	"strconv"
)

func main() {

	intToFloat()
	floatToInt()
	intToString()
	totalPrice()
	parseToInt()
	parseTofloat()
	parseStringToBool()
}

func intToFloat() {
	x := 315
	var d float64 = float64(x)
	println(d, x)
	fmt.Printf("int To Float %f", d)

	fmt.Printf("\n int To Float %8.2f", d)

	//format alternatif
	Format := strconv.FormatFloat(d, 'f', 5, 64)
	println("\n format: ", Format)

	//func SprintF
	sprintFormat := fmt.Sprintf("sprintF : %8.2f", d)
	println("sprint Format :", sprintFormat)

}

func floatToInt() {
	f := 509.98
	i := int64(f)
	formatInt := strconv.FormatInt(i, 10)

	fmt.Println("float to int : ", i, formatInt)

}

func intToString() {

	price := 45
	totalPrice := 198.78
	//concat
	output := "total Unit Price : " + strconv.Itoa(price)

	temp := int(totalPrice)
	total := "total price :" + strconv.Itoa(temp)
	fmt.Println(output)
	fmt.Println(total)

}

func totalPrice() {
	price := 98.9
	qty := 5
	totalPrice := price * float64(qty)
	fmt.Println("total price : ", totalPrice)
	fmt.Println("total Price: ", math.Round(totalPrice))
}

func parseToInt() {
	val1 := "98"
	//
	int1, err := strconv.ParseInt(val1, 10, 64)
	if err == nil {
		fmt.Println("parsed value : ", int1)
	} else {
		fmt.Println("can not parse", val1, err)

	}
}
func parseTofloat() {
	val1 := "98.98"
	//
	int1, err := strconv.ParseFloat(val1, 64)
	if err == nil {
		fmt.Println("parsed value : ", int1)
	} else {
		fmt.Println("can not parse", val1, err)

	}
}

func parseStringToBool() {
	val1 := "true"
	val2 := "false"
	val3 := "not true"

	bool1, err1 := strconv.ParseBool(val1)
	bool2, err2 := strconv.ParseBool(val2)
	bool3, err3 := strconv.ParseBool(val3)

	fmt.Println("Bool1", bool1, err1)
	fmt.Println("Bool2", bool2, err2)
	fmt.Println("Bool3", bool3, err3)

	if err1 == nil {
		fmt.Println("bool1", bool1)
	} else {
		fmt.Println("error parse bool 1", bool1, err1)
	}

	if err2 == nil {
		fmt.Println("bool1", bool2)
	} else {
		fmt.Println("error parse bool 1", bool2, err2)
	}

	if err3 == nil {
		fmt.Println("bool1", bool3)
	} else {
		fmt.Println("error parse bool 1", bool3, err3)
	}
}
