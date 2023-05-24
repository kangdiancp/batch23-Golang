package main

func divAndMod() {
	number := 15
	bagi := number / 3
	sisa := number % 3
	println(bagi, sisa)
}

// logical operator
func comparison() {
	first := 100
	const second = 200
	equal := first == second
	notEqual := first != second
	lessThan := first < second
	lessThanOrEqual := first <= second
	gretarThan := first > second
	gretarThanOrEqual := first >= second
	println(equal, notEqual, lessThan, lessThanOrEqual, gretarThan, gretarThanOrEqual)
}

func main() {
	divAndMod()
	comparison()
}
