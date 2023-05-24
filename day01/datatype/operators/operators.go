package main

func divisionModulus() {
	number := 15
	bagi := number / 3
	sisa := number % 3
	println(bagi, sisa)
}

// logikal operator
func comparison() {
	first := 100
	const second = 200          //nilai konstanta
	equal := first == second    //jika nilai sama tampung disini
	notEqual := first != second // jika nilai tdk sama tampung disini
	lessThan := first < second
	lessThanOrEqual := first <= second
	greater := first > second
	greaterThanOrEqual := first >= second

	println(first)
	println(equal)
	println(notEqual)
	println(lessThan)
	println(lessThanOrEqual)
	println(greater)
	println(greaterThanOrEqual)
}

func main() {
	divisionModulus()
	comparison()
}
