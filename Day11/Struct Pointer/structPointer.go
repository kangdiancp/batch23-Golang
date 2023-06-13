package main

import "fmt"

type Product struct {
	name, category string
	price          float64
	supplier       *Supplier
}
type Supplier struct {
	id           int
	supplierName string
}

func copyValue() {
	p1 := Product{name: "Samsung", category: "hp", price: 256}
	p2 := p1
	p1.name = "Oppo"
	fmt.Println("p1: ", p1)
	fmt.Println("p2: ", p2)
}

func copyPointer() {
	p1 := Product{name: "Samsung", category: "hp", price: 256}
	p2 := &p1
	p1.name = "Oppo"
	fmt.Println("p1: ", p1)
	fmt.Println("p2: ", *p2)
}

func pajak(product *Product) *Product {
	if product.price > 100 {
		product.price += product.price * 0.2
	}
	return product
}

//func constructor
func newProduct(name, category string, price float64, supplier *Supplier) *Product {
	discount := 10.0
	totalPrice := price - (price * discount / 100)

	return &Product{name, category, totalPrice, supplier}
}

// method
func (product *Product) calcDisc(discount float64) float64 {
	if product.price > 100 {

		return product.price - (product.price * discount)
	}
	return product.price
}

func main() {

	// init constructor
	supplier := &Supplier{1, "Agent Laptop"}
	p1 := newProduct("Dell", "Laptop", 345, supplier)
	p2 := newProduct("Dell", "Laptop", 150, supplier)
	p1.calcDisc(0.5)
	p1.calcDisc(0.8)
	products := []Product{*p1, *p2}
	fmt.Println(products)

	// prod := Product{name: "Samsung", category: "hp", price: 256}

	// fmt.Println(pajak(&prod))

	//copyValue()
	//copyPointer()
}
structPointer