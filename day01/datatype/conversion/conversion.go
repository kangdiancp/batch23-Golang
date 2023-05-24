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
	parseToFloat()
	parseToBool()
}

func intToFloat() {
	x := 315
	var d float64 = float64(x)

	println(d, x)

	fmt.Printf("Int to float %f \n", d)
	fmt.Printf("Int to float %8.2f \n", d)
	fmt.Printf("Int base 10 : %d \n", x)
	fmt.Printf("Octal : %b \n", d)

	// using strconv
	format:= strconv.FormatFloat(d, 'f', 5, 64)
	println("format : ", format)

	// func sprintf
	sprintFormat := fmt.Sprintf("sprintf : %8.2f \n", d)
	println("sprint Format : ", sprintFormat)
}

func floatToInt(){
	f := 509.98
	i := int64(f)
	formatInt := strconv.FormatInt(i, 10)
	fmt.Println("Float to Int : ", i, formatInt);
}

func intToString(){
	price := 45
	totalPrice := 178.98
	//concat
	output := "Total Price : " + strconv.Itoa(price)
	//ditampung
	temp := int(totalPrice)
	total := "Total Price : " + strconv.Itoa(int(totalPrice))
	total1 := "Total Price : " + strconv.Itoa(temp)
	fmt.Println(output)
	fmt.Println(total)
	fmt.Println(total1)
}

func totalPrice(){
	price := 98.9
	qty := 6
	totalPrice := price * float64(qty)
	totalprice := price * float64(qty)
	fmt.Println("Total Price : ", totalPrice)
	fmt.Println("Total Price : ", math.Round(totalprice))
}

func parseToInt(){
	val1 := "98"
	// if err not error, err return nil
	int1, err := strconv.ParseInt(val1, 10, 64)
	if err == nil{
		fmt.Println("Parsed Value : ", int1)
	} else{
		fmt.Println("Can  not parse", val1, err)
	}
}

func parseToFloat(){
	val1 := "98.99"
	// if err not error, err return nil
	float1, err := strconv.ParseFloat(val1, 64)
	if err == nil {
		fmt.Println("Parse Value : ", float1)
	} else {
		fmt.Println("can not parse", val1, err)
	}
}

func parseToBool(){
	val1 := "true"
	val2 := "false"
	val3 := "true"
	bool1, err1 := strconv.ParseBool(val1)
	bool2, err2 := strconv.ParseBool(val2)
	bool3, err3 := strconv.ParseBool(val3)

	if err1 == nil {
		fmt.Println("Bool1 : ", bool1)
	}else{
		fmt.Println("can not parse", val1)
	}
	if err2 == nil {
		fmt.Println("Bool2 : ", bool2)
	}else{
		fmt.Println("can not parse",val2)
	}
	if err3 == nil {
		fmt.Println("Bool3 : ", bool3)
	}else{
		fmt.Println("can not parse", val3)
	}
}