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
	p1 := Product{
		name:     "Samsung",
		category: "HP",
		price:    256,
	}

	p2 := p1
	p1.name = "Oppo"

	fmt.Println("p1 : ", p1)
	fmt.Println("p2 : ", p2)
}

func copyPointer() {
	p1 := Product{
		name:     "Samsung",
		category: "HP",
		price:    256,
	}

	p2 := &p1
	p1.name = "Oppo"

	fmt.Println("p1 : ", p1)
	fmt.Println("p2 : ", *p2)
}

func calculateTax(product *Product) *Product {
	if product.price > 100 {
		product.price += product.price * 0.2
	}
	return product
	// fmt.Println("Total Price + Tax : ", product)
}

//func constructor
func newProduct(name, category string, price float64, supplier *Supplier) *Product {
	totalPrice := (price * 0.1)
	return &Product{name, category, price - totalPrice, supplier}
}

//method function, between func and function name add receiver type
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
	// fmt.Println(calculateTax(&prod))

	//initial constructor
	supplier := &Supplier{1, "Agent Laptop"}
	p1 := newProduct("Dell", "Laptop", 345, supplier)
	p2 := newProduct("Acer", "Laptop", 150, supplier)

	p1.supplier.supplierName = "Dell"

	p1.price = p1.calcDiscount(0.5)
	p2.price = p2.calcDiscount(0.8)

	products := []Product{*p1, *p2}
	fmt.Println(products)

}
