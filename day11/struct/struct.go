package main

import "fmt"

func main() {
	// initProduct()
	// initProductPointer()
	// initOrderProduct()
	// compareStruct()
	// printStruct()
	// fmt.Println(initSliceProduct())
	fmt.Println(findProductByName(initSliceProduct(), "Samsung"))
}

type Product struct {
	name, category string
	price          float64
}

type OrderProduct struct {
	product    Product //embedded fields
	totalOrder int     //reguler fields
}

func initProduct() {
	prod := Product{
		name:     "Phone",
		category: "Elektronik",
		price:    100,
	}
	prod.name = "MacBook"
	prod.price = 700
	fmt.Println(prod)

	//default zero value
	var tv Product
	tv.name = "Sharp"

	meja := Product{}
	fmt.Println(tv, meja)
}

func initProductPointer() {
	//built-in new function, result pointer struct

	var phone = new(Product)
	phone.name = "Samsung"
	phone.category = "Hp"
	phone.price = 235.98

	fmt.Println(phone)

	//using & for make a new product
	var oppo = &Product{}
	oppo.name = "Oppo"
	oppo.price = 100
	oppo.category = "Phone"

	fmt.Println(oppo)
}

func initOrderProduct() {
	myOrder := OrderProduct{
		product:    Product{"Iphone", "Phone", 100},
		totalOrder: 1,
	}
	fmt.Println(myOrder)
}

func compareStruct() {
	// var oppo = &Product{}
	// oppo.name = "Oppo"
	// oppo.price = 100
	// oppo.category = "Phone"

	// var iphone = &Product{}
	// iphone.name = "Iphone"
	// iphone.price = 600
	// iphone.category = "Phone"

	//Diatas tidak bisa melakukan compare

	oppo := Product{
		name: "Oppo", category: "Phone", price: 100,
	}
	iphone := Product{
		name: "Oppo", category: "Phone", price: 100,
	}

	fmt.Println("result : ", oppo == iphone)
}

//annonymous struct
func printStruct(val struct {
	name, category string
	price          float64
}) {
	fmt.Println("Name : ", val.name, "Category : ", val.category, "Price : ", val.price)
}

func initSliceProduct() []Product {
	oppo := Product{
		name: "Oppo", category: "Phone", price: 100,
	}
	iphone := Product{
		name: "Oppo", category: "Phone", price: 100,
	}

	products := []Product{oppo, iphone}

	products = append(products,
		Product{name: "Samsung", category: "HP", price: 255})

	return products
}

func findProductByName(products []Product, search string) Product {
	for _, val := range products {
		if val.name == search {
			return val
		}
	}
	return Product{}
}
