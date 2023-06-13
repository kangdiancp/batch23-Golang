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

func copyByValue() {
	p1 := Product{name: "samsung", category: "hp", price: 765}

	p2 := p1
	p1.name = "oppo"

	fmt.Println("p1:", p1)
	fmt.Println("p2:", p2)
}

func copyByPointer() {
	p1 := Product{name: "samsung", category: "hp", price: 765}

	p2 := &p1
	p1.name = "oppo"

	fmt.Println("p1:", p1)
	fmt.Println("p2:", *p2)
}

func calcTax(product *Product) *Product {
	if product.price > 100 {
		product.price += product.price * 0.2
	}
	return product
}

//method function, between func and function name add receiver type
func (product *Product) calcDiscount(discount float64) float64 {
	if product.price > 100 {
		return product.price - (product.price * discount)
	}
	return product.price
}

//func constructor
func newProduct(name, category string, price float64, supplier *Supplier) *Product {
	//discount := 10
	totalPrice := (price * 0.1)
	return &Product{name, category, price - totalPrice, supplier}
}

func main() {
	// prod := Product{name: "samsung", category: "hp", price: 765}
	// fmt.Println(calcTax(&prod))

	supplier := &Supplier{1, "agent laptop"}

	//initial constructor
	p1 := newProduct("dell", "laptop", 345, supplier)
	p2 := newProduct("acer", "laptop", 150, supplier)

	p1.price = p1.calcDiscount(0.5)
	p2.price = p2.calcDiscount(0.8)

	p1.supplier.supplierName = "dell"

	products := []Product{*p1, *p2}
	fmt.Println(products)

	// copyByPointer()
	// copyByValue()
}
