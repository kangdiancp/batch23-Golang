package main

//global constant variable
const phi = 3.14159

func declareConstant() {
	//constant type
	const price float64 = 20.98

	//constant untype
	const tax = 11.9

	//multiple constant
	const (
		discount = .1
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
}