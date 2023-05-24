package main

import (
	"fmt"
	"math"
	"strconv"
)

func intToFloat() {
	x := 315                   //tipe data intger nih bos
	var d float64 = float64(x) //float di isi parameter yang bisa di isi ke dalam nilai d
	println(x)
	println(d)
	fmt.Printf("Int To Float %f", d)       //memasukan data to ke float f dan d
	fmt.Printf("\nInt To Float: %8.2f", d) //.2f memberi 2 angka dibelakang koma
	fmt.Printf("\nInteger base 10: %d", x) // base 10 itu untuk integer baca dokumentasi go ya
	fmt.Printf("\nOctal: %b", d)           // oktal

	//format alternatif menggunakan format strconv
	format := strconv.FormatFloat(d, 'f', 5, 64)
	println("\nformat :", format)

	//fungsi SprintF: untuk merubah angka ke string tapi gayanya tetep sprintf,
	sprintFormat := fmt.Sprintf("SprintF : %8.2f", d)
	println(sprintFormat)

}

func floatToInt() {
	f := 509.98
	i := int64(f)
	formatInt := strconv.FormatInt(i, 10)
	fmt.Println("\nFloat To Int:", i)
	fmt.Println("Float To Int:", formatInt)
}

// fungsi integer ke string
func intToString() {
	price := 45
	totalPrice := 198.78
	//concat
	output := "Total Price :" + strconv.Itoa(price) //iota mengubah integer ke string
	//ditampung ke temp := int(totalPrice)
	//cara 1 pake tampungan
	//total := "Total Price :" + strconv.Itoa(temp)
	//cara 2 langsung diubah dari temp ke integer
	total := "Total Price :" + strconv.Itoa(int(totalPrice)) //mengubah total float ke int, maka itoa bisa dipake
	fmt.Println(output)
	fmt.Println(total)
}

func totalPrice() {
	price := 98.9
	qty := 5
	totalPrice := price * float64(qty)
	fmt.Println("Total Price :", totalPrice)
	//untuk pembulatan bisa menggunakan fungsi yang dibawah
	fmt.Println("Total Price :", math.Round(totalPrice))
}

// cara 1
func parseToInt() {
	val1 := "98"
	//fungsi dibawah harus menggunakan 2 fungsi, 64 itu untuk ukuran bitnya
	//fungsi err digunakan jika terjadi error (if err not error, err return nil)
	int1, err := strconv.ParseInt(val1, 10, 64)
	if err == nil { //nil=tdk eror, jika eror sama dg nil, maka seterusnya
		fmt.Println("Parsed value :", int1)
	} else {
		fmt.Println("Error/can not parse :", val1, err)
	}
}

// cara ke 2
func parseToFloat() {
	val1 := "98.98"
	int1, err := strconv.ParseFloat(val1, 64)
	if err == nil {
		fmt.Println("Parsed value :", int1)
	} else {
		fmt.Println("Error/can not parse :", val1, err)
	}
}

func parseStringToBool() {
	val1 := "true"
	val2 := "false"
	val3 := "not true"
	//fungsi dibawah harus menggunakan 2 fungsi, 64 itu untuk ukuran bitnya
	//fungsi err digunakan jika terjadi error (if err not error, err return nil)
	bool1, err1 := strconv.ParseBool(val1)
	bool2, err2 := strconv.ParseBool(val2)
	bool3, err3 := strconv.ParseBool(val3)

	fmt.Println("Bool1", bool1, err1)
	fmt.Println("Bool2", bool2, err2)
	fmt.Println("Bool3", bool3, err3)

	if err1 == nil {
		fmt.Println("Bool1 :", bool1)
	} else {
		fmt.Println("Error/can not parse bool1 :", bool1, err1)
	}
	if err2 == nil {
		fmt.Println("Bool2 :", bool2)
	} else {
		fmt.Println("Error/can not parse bool2 :", bool2, err2)
	}
	if err3 == nil {
		fmt.Println("Bool3 :", bool3)
	} else {
		fmt.Println("Error/can not parse bool3 :", bool3, err3)
	}

}

// jangan lupa kalo mau ngerun, harus ganti direktori dulu
// cd .., terus ke cd conversion, terus tinggal go run conversion aja langsung jadi
func main() {
	intToFloat()
	floatToInt()
	intToString()
	totalPrice()
	parseToInt()
	parseToFloat()
	parseStringToBool()
}

//conversi : dari integer ke float, float ke integer
//parsing : dari string ke integer & float, bollean ke string itu passing
