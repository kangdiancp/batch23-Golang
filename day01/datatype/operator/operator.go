package main

func devisionModulus() {
	number := 15
	bagi := number / 3
	sisa := number % 2
	println(bagi, sisa)
}

// logical operator
func comparision() {
	first := 100
	const second = 200
	equal := first == second
	notequal := first != second
	lessThan := first < second
	lessthanorequal := first <= second
	greaterthan := first > second
	greaterthanorequal := first >= second

	println(first)
	println(second)
	println(equal)
	println(notequal)
	println(lessThan)
	println(lessthanorequal)
	println(greaterthan)
	println(greaterthanorequal)
}

func main() {
	devisionModulus()
	comparision()

}
