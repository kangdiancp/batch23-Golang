package main

import "fmt"

func main() {

	declareConstant()
	declareIotas()
}

// constant type
func declareConstant() {

	const price float64 = 20.98

	//constant untype
	const tax = 11.99

	// multiple constan
	const (
		discount = 0.1
		phi      = 3.14
	)
	fmt.Println(price, tax, discount, phi)

}

func declareIotas() {
	const (
		seq1 = iota
		seq2
		seq3
	)
	fmt.Println(seq1, seq2, seq3)
}
