package main

import "fmt"

type Products struct {
	name, category string
	price          float64
}

func initProduct() {
	//literal value
	prod := Products{
		name:     "Laptop",
		category: "Electronic",
		price:    250,
	}

	prod.name = "Dell Laptop"
	prod.price = 355.98
	fmt.Println(prod)

	//default xero value
	var tv Products
	tv.name = "Sharp"

	meja := Products{}
	fmt.Println(tv, meja)
}

func initProductPointer(){
	// built in new function
	var phone = new(Products)
	phone.name = "Samsung"
	phone.category = "HP"
	phone.price = 235.98

	fmt.Println(phone)

	//using &
	var oppo = &Products{}
	oppo.name = "Oppo"
	oppo.category = "HP"
	oppo.price = 185

	fmt.Println(oppo)
}

type OrderProduct struct{
	product Products // embedded fields
	totalOrder int   // regular field
}

func initOrderProduct(){
	myOrder := OrderProduct{
		product: Products{"Iphone", "HP", 355},
		totalOrder: 1,
	}

	fmt.Println(myOrder)
}

func compareStruct(){
	// var oppo = &products{}
	// oppo.name = "iphone"
	// oppo.category = "HP"
	// oppo.price = 185

	// var iphone = &products{}
	// iphone.name = "iphone"
	// iphone.category = "HP"
	// iphone.price = 185

	oppo := Products{
		name: "iphone", category: "HP",price: 185,
	}
	
	iphone := Products{
		name: "iphone", category: "HP",price: 185,
	}

	fmt.Println("Result : ", oppo == iphone)
}

//anonymous struct
func printStruct(val struct{
	name, category string
	price float64
}){
	fmt.Println("Name : ", val.name, "Category : ", val.category, "Price : ", val.price)
}

func initSliceProduct()[]Products{
	oppo := Products{
		name: "oppo", category: "HP",price: 185,
	}
	
	iphone := Products{
		name: "iphone", category: "HP",price: 185,
	}

	Product:= []Products{oppo, iphone}

	Product = append(Product, Products{name: "Samsung", category: "HP", price: 256})

	return Product
}

func findProductByName(products []Products, search string) Products{
	
	for _, v := range products {
		if v.name == search{
			return v
		}
	}
	return Products{}
}

func main() {
	// initProduct()
	// initProductPointer()
	// initOrderProduct()
	// compareStruct()
	fmt.Println(findProductByName(initSliceProduct(), "Samsung"))
}