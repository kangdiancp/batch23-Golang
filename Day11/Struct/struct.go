package main

import "fmt"

type Product struct {
	name, category string
	price          float64
}

type OrderProduct struct {
	product    Product // embedded fields
	totalOrder int     // Regular fields
}

func initProduct() {
	prod := Product{
		name:     "Laptop",
		category: "Elektronik",
		price:    250,
	}

	prod.name = "Dell Laptop"
	prod.price = 355.98

	// Default zero value
	var tv Product
	tv.name = "sharp"

	meja := Product{}
	fmt.Println(tv, meja)
}

func initProductPointer() {
	// Built-in new function, result pointer struct
	var phone = new(Product)
	phone.name = "Samsung"
	phone.category = "HP"
	phone.price = 234.98

	fmt.Println(phone)

	// Using &
	var oppo = &Product{}
	oppo.name = "Oppo"
	oppo.category = "HP"
	oppo.price = 2342.23

	fmt.Println(oppo)
}

func initOrderProduct() {
	myOrder := OrderProduct{
		product:    Product{"iPhone", "HP", 355},
		totalOrder: 1,
	}
	fmt.Println(myOrder)
}

func compareStruct() {
	var oppo = Product{name: "Oppo", category: "HP", price: 185}
	var iphone = Product{name: "Oppo", category: "HP", price: 185}

	fmt.Println("Result : ", oppo == iphone)
}

// Anonymous struct
func printStruct(val struct {
	name, category string
	price          float64
}) {
	fmt.Println("Name :", val.name, "Category : ", val.category, "Price : ", val.price)
}

func initSliceProduct() []Product {
	oppo := Product{
		name: "Oppo", category: "hp", price: 185,
	}

	iphone := Product{
		name: "Iphone", category: "hp", price: 200,
	}

	products := []Product{oppo, iphone}

	// Pakai append
	products = append(products, Product{name: "Samsung", category: "hp", price: 256})

	return products
}

func findProductByName(products []Product, search string) Product {

	for _, v := range products {
		if v.name == search {
			return v
		}
	}
	return Product{}
}

func main() {
	// initProduct()
	// initProductPointer()
	// initOrderProduct()
	// compareStruct()
	fmt.Println(findProductByName(initSliceProduct(), "Samsung"))
}
