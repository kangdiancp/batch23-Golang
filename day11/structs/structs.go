package main

import (
	"fmt"
)

type Product struct {
	name, category string
	price          float64
}

type OrderProduct struct {
	product    Product //embedded fields
	totalOrder int     //regular fields
}

func initProduct() {
	//literal value
	prod := Product{
		name:     "Laptop",
		category: "electronic",
		price:    250,
	}

	prod.name = "dell laptop"
	prod.price = 355.98

	fmt.Println(prod)

	//default zero value
	var TV Product
	TV.name = "panasonic"
	meja := Product{}
	fmt.Println(TV, meja)
}

func initProductPointer() {
	//built-in new function, result pointer struct
	var phone = new(Product)
	phone.name = "samsung"
	phone.category = "HP"
	phone.price = 235.99

	fmt.Println(phone)

	//using &
	var oppo = &Product{}
	oppo.name = "oppo"
	oppo.category = "hp"
	oppo.price = 555.6

	fmt.Println(oppo)
}

func initOrderProduct() {
	myOrder := OrderProduct{
		product:    Product{"iphone", "hp", 355},
		totalOrder: 1,
	}
	fmt.Println(myOrder)
}

func compareStruct() {
	oppo := Product{
		name: "oppo", category: "hp", price: 555.6,
	}

	iphone := Product{
		name: "iphone", category: "hp", price: 555.6,
	}

	fmt.Println("result:", oppo == iphone)
	printStruct(oppo)

}

// anonymous struct
func printStruct(val struct {
	name, category string
	price          float64
}) {
	fmt.Println("name :", val.name, "category:", val.category, "price:", val.price)
}

func initSliceProduct() []Product {
	oppo := Product{
		name: "oppo", category: "hp", price: 555.6,
	}

	iphone := Product{
		name: "iphone", category: "hp", price: 765,
	}

	products := []Product{oppo, iphone}

	products = append(products, Product{name: "samsung", category: "hp", price: 765})

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
	fmt.Println(findProductByName(initSliceProduct(), "iphone"))
	//compareStruct()
	//initOrderProduct()
	//initProductPointer()
	//initProduct()
}
