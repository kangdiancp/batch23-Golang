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

//init untuk call tanpa variable baru
type ProductList []Product

func initProduct() []Product {
	return []Product{
		{"samsung", "HP", 145, &Supplier{}},
		{"oppo", "HP", 55.6, &Supplier{}},
		{"dell", "laptop", 245.98, &Supplier{}},
	}
}

// interface header
type info interface {
	printInfo()
	totalPrice()
}

// implementation interface
func (p *Product) printInfo() {
	fmt.Println("product info :", p)
}

func (p *ProductList) totalPrice() map[string]float64 {
	totals := make(map[string]float64)
	for _, v := range *p {
		totals[v.category] = totals[v.category] + v.price
	}
	return totals
}

func (s *Supplier) printInfo() {
	fmt.Println("supplier info :", s)
}

func main() {
	products := ProductList(initProduct())
	products[0].printInfo()

	products.totalPrice()
	fmt.Println(products.totalPrice())

	s := Supplier{id: 1, supplierName: "abc"}
	s.printInfo()

}
