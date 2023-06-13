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
	p1 := Product{name: "Samsung", category: "HP", price: 256}

	p2 := p1
	p1.name = "Oppo"

	fmt.Println("p1 : ", p1)
	fmt.Println("p2 : ", p2)
}

func copyPointer() {
	p1 := Product{name: "Samsung", category: "HP", price: 256}

	p2 := &p1
	p1.name = "Oppo"

	fmt.Println("p1 : ", p1)
	fmt.Println("p2 : ", *p2)
}

func calcTax(product *Product) *Product {
	if product.price > 100 {
		product.price += product.price * 0.2
	}
	return product
}

// function constructor,
func newProduct(name, category string, price float64, supplier *Supplier) *Product {
	// discount := 10
	totalPrice := (price * 0.1)
	return &Product{name, category, price - totalPrice, supplier}
}

// method
func (product *Product) calcDiscount(discount float64) float64 {
	if product.price > 100 {
		return product.price - (product.price * discount)
	}
	return product.price
}

func main() {
	// copyValue()
	// copyPointer()

	// prod := Product{name: "Samsung", category: "HP", price: 256}
	// fmt.Println("Total Price + Tax : ", calcTax(&prod))

	// initial constructor
	supplier := &Supplier{1, "Agent Laptop"}

	p1 := newProduct("Dell", "Laptop", 345, supplier)
	p2 := newProduct("Acer", "Laptop", 150, supplier)

	p1.calcDiscount(0.5)
	p1.calcDiscount(0.8)

	products := []Product{*p1, *p2}

	fmt.Println(products)
}
