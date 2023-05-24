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
	greateThanOrEqual := first >= second

	println(first)
	println(equal)
	println(notEqual)
	println(lessThan)
	println(lessThanOrEqual)
	println(greater)
	println(greateThanOrEqual)

}

func main() {
	divisionModulus()
	comparison()

}
