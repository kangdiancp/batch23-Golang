package main

func divisionModulus() {
	number := 15
	bagi := number / 3
	sisa := number % 2
	println(bagi, sisa)
}

// Logical Operator
func comparison() {
	first := 100
	const second = 200
	equal := first == second
	notEqual := first != second
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
