package main

func divisionModulus() {
	number := 15
	bagi := number / 3
	sisa := number % 3
	println(number, bagi, sisa)
}

// logical operator
func comparison() {
	first := 100
	const second = 200
	equal := first == second
	notEqual := first != second
	lessThan := first < second
	lessThanEOrEqual := first <= second
	greater := first < second
	greaterThanOrEqual := first >= second

	println(first)
	println(second)
	println(equal)
	println(notEqual)
	println(lessThan)
	println(lessThanEOrEqual)
	println(greater)
	println(greaterThanOrEqual)
}

func main() {
	divisionModulus()
	comparison()
}
