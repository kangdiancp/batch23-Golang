package main

// global constant variable
const phi = 3.14159

func declareConstant() {
	//constant type
	const price float64 = 20.98

	//constant untype
	const tax = 11.99

	//mutiple constant
	const (
		discount = 0.1
		phi      = 3.14159
	)

	println(price, tax, discount, phi)
}

func declareIotas() {
	const (
		seq1 = iota
		seq2
		seq3
	)
	var circleArea = phi * 5
	println(seq1, seq2, seq3, circleArea)
}

func main() {
	declareConstant()
	declareIotas()
}
