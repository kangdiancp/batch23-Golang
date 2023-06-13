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

func copyPointer() {
	p1 := Product{name: "Samsung", category: "Hp", price: 256}

	p2 := &p1
	p1.name = "Oppo"

	fmt.Println("p1 :", p1)
	fmt.Println("p2 :", *p2)
}

func copyValue() {
	p1 := Product{name: "Samsung", category: "Hp", price: 256}

	p2 := p1
	p1.name = "Oppo"

	fmt.Println("p1 :", p1)
	fmt.Println("p2 :", p2)
}

func calcTax(product *Product) *Product {
	if product.price > 100 {
		product.price += product.price * 0.2
	}
	return product
	//fmt.Println("Total Price + Tax : ", product)
}

//func constructor,
func newProduct(name, category string, price float64, supplier *Supplier) *Product {
	//discount := 10
	totalPrice := (price * 0.1)
	return &Product{name, category, totalPrice, supplier}
}

//method
func (product *Product) calcDiscount(discount float64) float64 {
	if product.price > 100 {
		return product.price - (product.price * discount)
	}
	return product.price
}

func main() {
	// prod := Product{name: "Samsung", category: "Hp", price: 256}
	// fmt.Println(calcTax(&prod))
	//copyPointer()
	//copyValue()

	//initial constructor
	supplier := &Supplier{1, "Agen Laptop"}
	p1 := newProduct("Dell", "Laptop", 345, supplier)
	p2 := newProduct("Acer", "Laptop", 150, supplier)
	p1.calcDiscount(0.5)
	p2.calcDiscount(0.8)
	p1.supplier.supplierName = "Dell"

	products := []Product{*p1, *p2}

	fmt.Println(products)
}
