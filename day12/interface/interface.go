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

type ProductList []Product

func initProduct() []Product {
	return []Product{
		{"Samsung", "HP", 145, &Supplier{}},
		{"Oppo", "HP", 55.65, &Supplier{}},
		{"Dell", "Laptop", 245.98, &Supplier{}},
	}
}

// interface header
type Info interface {
	printInfo()
	totalPrice(tax float64)
}

// implementation interface
func (p *Product) printInfo() {
	fmt.Println("Product Info : ", p)
}

func (p *ProductList) totalPrice(tax float64) map[string]float64 {
	totals := make(map[string]float64)
	for _, v := range *p {
		totals[v.category] = totals[v.category] + v.price
	}
	return totals

}

func (s *Supplier) printInfo() {
	fmt.Println("Supplier Info : ", s)
}

func main() {
	products := ProductList(initProduct())

	products[1].printInfo()

	products.totalPrice(0.5)

	fmt.Println(products.totalPrice())

	s := Supplier{id: 1, supplierName: "abc"}
	s.printInfo()

}
