package main

func divisionModulus() {
	number := 1234
	bagi := number / 3
	sisa := number % 10
	println("Hasil Bagi = ", bagi)
	println("Hasil sisa = ", sisa)
}

// logical comparison
func comparison() {
	first := 100
	const second = 200
	equal := first == second
	notEqual := first != second
	lessThan := first < second
	lessThanOrEqual := first <= second
	greater := first > second
	greaterThanOrEqual := first >= second

	println("Hasil Comparisson= ", first)
	println("Hasil Comparisson= ", equal)
	println("Hasil Comparisson= ", notEqual)
	println("Hasil Comparisson= ", lessThan)
	println("Hasil Comparisson= ", lessThanOrEqual)
	println("Hasil Comparisson= ", greater)
	println("Hasil Comparisson= ", greaterThanOrEqual)
}
func main() {
	divisionModulus()
	comparison()
}
