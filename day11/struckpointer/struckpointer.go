package main

import "fmt"

type product struct {
	name, catagory string
	Price          float64
	Supplier       *Supplier
}

type Supplier struct {
	id           int
	SupplierName string
}

func copyValue() {
	p1 := product{
		name:     "Samsung",
		catagory: "HP",
		Price:    256,
	}

	p2 := p1
	p1.name = "Oppo"

	fmt.Println("p1 : ", p1)
	fmt.Println("p2 : ", p2)
}

func copyPointer() {

	p1 := product{name: "samsung", catagory: "HP", Price: 256}

	p2 := &p1
	p1.name = "Oppo"

	fmt.Println("p1 : ", p1)
	fmt.Println("p2 : ", *p2)
}

func calcTax(product *product) *product {
	if product.Price > 100 {
		product.Price += product.Price * 0.2
	}
	return product
	//fmt.Println("Total price + Tax:", product)
}

// Method function,betwen func and function name add receiver type
func (product *product) calcDiscount(discount float64) float64 {

	if product.Price > 100 {
		return product.Price - (product.Price * discount)
	}
	return product.Price
}

//func constructor,
func newProduct(name, catagory string, price float64, supplier *Supplier) *product {
	//discound := 10
	totalprice := (price * 0.1)
	return &product{name, catagory, price - totalprice, supplier}
}

func main() {
	//prod := product{name: "Samsung", catagory: "HP", Price: 256}
	//fmt.Println(calTax(&prod))
	//copyPointer()

	supplier := &Supplier{1, "Agent Laptop"}
	//inital constructor
	p1 := newProduct("Dell", "Laptop", 345, supplier)
	p2 := newProduct("Acer", "Laptop", 150, supplier)

	p1.Price = p1.calcDiscount(0.5)
	p2.Price = p2.calcDiscount(0.8)

	p1.Supplier.SupplierName = "Dell"

	products := []product{*p1, *p2}
	fmt.Println(products)
}
