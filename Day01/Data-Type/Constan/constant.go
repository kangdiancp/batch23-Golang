package main

//const global
const phi = 3.78

func declareConstant() {
	// Constant Type
	const price float64 = 20.98

	//Constant Untype
	const tax = 11.99

	//Multiple Constant
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

	println(seq1, seq2, seq3)
}

func main() {
	declareConstant()
	declareIotas()
}
