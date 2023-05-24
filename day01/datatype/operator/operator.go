package main

func divisionModulus() {
	number := 15
	bagi := number / 3
	sisa := number % 3
	println(bagi, sisa)
}

func comparison() {
	first := 100
	const second = 200
	equal := first == second
	notEqual := first != second
	lessThan := first < second
	lessThanOrEqual := first <= second
	greater := first > second
	greaterThanOrEqual := first >= second

	println("First : ", first, ",", "Second : ", second)
	println(equal, notEqual, lessThan, lessThanOrEqual, greater, greaterThanOrEqual)
}

func main() {
	divisionModulus()
	comparison()
}
