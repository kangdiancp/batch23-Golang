package main

import "fmt"

func main() {
	// copyPointer()

	// prod := Product{name: "Samsung", category: "HP", price: 255}
	// fmt.Println(calcTax(&prod))

	// myOrder := OrderProduct{
	// 	product:      Product{name: "Samsung", category: "HP", price: 255},
	// 	countOfOrder: 3,
	// 	totalOrder:   0,
	// }
	// fmt.Println(calcTotalOrder(&myOrder))

	Supplier := &Supplier{1, "Agent HandPhone"}

	p1 := newProduct("Oppo", "Phone", 255, Supplier)
	p2 := newProduct("Iphone", "Phone", 500, Supplier)

	products := []Product{*p1, *p2}

	// fmt.Println(prod.calcDiscount(0.5))
	// prod.calcDiscount(0.5)

	// fmt.Println(prod)

	fmt.Println(products)

}

type Product struct {
	name, category string
	price          float64
	Supplier       Supplier
}
type OrderProduct struct {
	product      Product //embedded fields
	countOfOrder int     //reguler fields
	totalOrder   int
}

type Supplier struct {
	id           int
	supplierName string
}

func copyPointer() {
	p1 := Product{name: "Samsung", category: "HP", price: 255}

	p2 := &p1
	p1.name = "Oppo"

	fmt.Println("p1 : ", p1)
	fmt.Println("p2 : ", *p2)
}

func calcTax(product *Product) *Product {
	if product.price > 100 {
		product.price -= product.price * 0.2
	}
	return product
	// fmt.Println("Total Price + Tax : ", product)
}

func calcTotalOrder(order *OrderProduct) *OrderProduct {
	//Contoh Buatan
	if order.product.price > 100 {
		order.product.price -= order.product.price * 0.5
	}

	order.totalOrder = int(order.product.price) * order.countOfOrder

	return order
}

func newProduct(name, category string, price float64, supplier *Supplier) *Product {
	discount := 10
	return &Product{name, category, price - (price * float64(discount) / 100), *supplier}
}

//Contoh Method function
func (product *Product) calcDiscount(discount float64) float64 {
	if product.price > 100 {
		return product.price - (product.price * discount)
	}
	return product.price
}
