package main

import "fmt"

type Product struct {
	name, category string
	price          float64
	supplier       Supplier
}
type Supplier struct {
	id           int
	supplierName string
}

func copyValue() {
	p1 := Product{name: "samsung", category: "hp", price: 256}

	p2 := p1
	p1.name = "oppo"

	fmt.Println("p1 : ", p1)
	fmt.Println("p2 : ", p2)
}

func copyPointer() {
	p1 := Product{name: "samsung", category: "hp", price: 256}

	p2 := &p1
	p1.name = "oppo"

	fmt.Println("p1 : ", p1)
	fmt.Println("p2 : ", *p2)
}

func calcTax(product *Product) *Product {
	if product.price > 100 {
		product.price += product.price * 0.2
	}
	return product
	//fmt.Println("Total Price + Tax : ", product)
}

//func constructor
func newProduct(name, category string, price float64, supplier *Supplier) *Product {
	//discount := 10
	return &Product{name, category, price - (price * 0.1), *supplier}
}

//method
func (product *Product) calcDiscount(discount float64) float64 {
	if product.price > 100 {
		return product.price - (product.price * discount)
	}
	return product.price
}

func main() {
	supplier := &Supplier{1, "agent laptop"}
	//inital constructor
	p1 := newProduct("dell", "laptop", 345, supplier)
	p2 := newProduct("acer", "laptop", 150, supplier)

	p1.price = p1.calcDiscount(0.5)
	p2.price = p2.calcDiscount(0.8)

	product := []Product{*p1, *p2}
	fmt.Println(product)

	// prod := Product{name: "samsung", category: "hp", price: 256}
	// fmt.Println(calcTax(&prod))

	// copyValue()
	// copyPointer()
}
