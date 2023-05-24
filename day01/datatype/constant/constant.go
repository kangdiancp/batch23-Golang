package main

//global constant variable
const phi = 3.14159

func declareConstant() {
	//constant type
	const price float64 = 20.98 //type data setia tdk tergantikan aseek

	//constant untype
	const tax = 11.9

	//mutiple constant
	const (
		discount = 0.1 //untuk 0 diawal koma tdk dikasih gpp selow aja
		phi      = 3.14159
	)

	println(price, tax, discount, phi)
}

func declarIotas() {
	const (
		seq1 = iota // increment integer, cukup inisial value di seq1 aja say
		seq2
		seq3
	)
	var circleArea = phi * 5 //phi yang terpanggil yang phi atas, dan bebas dipanggil dimana aja
	println(seq1, seq2, seq3)
	println(circleArea)
}

func main() {
	declareConstant()
	declarIotas()
}
