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
	p1 := Product{name: "Esia Hidayah", category: "Bata", price: 5}

	p2 := p1
	p1.name = "Cherry"

	fmt.Println("p1 : ", p1)
	fmt.Println("p2 : ", p2)
}

func copyPointer() {
	p1 := Product{name: "BB", category: "BBM", price: 55}

	p2 := &p1
	p1.name = "Nexian"

	fmt.Println("p1 : ", p1)
	fmt.Println("p2 : ", *p2)
}

func calcTax(product *Product) *Product {
	if product.price > 100 {
		product.price += product.price * 0.2
	}

	// fmt.Println("Total Price + Tax : ", product)
	return product
}

// Func Constructor
func newProduct(name, category string, price float64, supplier *Supplier) *Product {
	// discount := 10
	totalPrice := (price * 0.1)
	return &Product{name, category, price - totalPrice, supplier}
}

// Method
func (product *Product) calcDiscount(discount float64) float64 {
	if product.price > 100 {
		return product.price - (product.price * discount)
	}

	return product.price
}

func main() {
	copyValue()

	fmt.Println()

	copyPointer()

	fmt.Println()

	prod := Product{name: "BB", category: "BBM", price: 12345}
	fmt.Println(calcTax(&prod))

	fmt.Println()

	// Initial Constructor
	supplier := &Supplier{1, "Agent Laptop"}

	p1 := newProduct("SUSA", "Laptop", 333, supplier)
	p2 := newProduct("ACER", "Laptop", 222, supplier)

	products := []Product{*p1, *p2}
	fmt.Println(products)

	fmt.Println()

	fmt.Println(p1.calcDiscount(0.5))
	fmt.Println(p2.calcDiscount(0.8))
}
