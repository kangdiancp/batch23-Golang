package main

//global constant
const phy = 3.14159

func decladeConstant() {
	//constant type
	const price float64 = 20.98

	//untype constant
	const tax = 11.9

	// multiple constat
	const (
		discount = .1
		phy      = 3.14159
	)

	println(price, tax, discount, phy)

}

func declareIotas() {
	const (
		seq1 = iota
		seq2
		seq3
	)
	var circleArea = phy * 5
	println(circleArea)
	println(seq1, seq2, seq3)

}

func main() {
	decladeConstant()
	declareIotas()
}
