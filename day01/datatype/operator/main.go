package main

func main() {
	divisionModulus()
	comparison()
}

func divisionModulus() {
	number := 15
	bagi := number / 3
	sisa := number % 3
	println(bagi, "\n", sisa)
}

// logikal operator
func comparison() {
	first := 100
	const second = 200
	equal := first == second
	notEqual := first != second
	lessThan := first < second
	lessThanorEqual := first <= second
	greater := first > second
	greaterThanOrEqual := first >= second

	println(first)
	println(second)
	println(equal)
	println(notEqual)
	println(lessThan)
	println(lessThanorEqual)
	println(greater)
	println(greater)
	println(greaterThanOrEqual)

}
