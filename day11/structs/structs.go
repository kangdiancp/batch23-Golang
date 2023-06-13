package main

import "fmt"

//init struct
type Product struct {
	name, category string
	price          float64
}

type OrderProduct struct {
	product    Product //embedid fields
	totalOrder int     // regulat fields
}

func iniProduct() {
	//literal value
	prod := Product{
		name:     "laptop",
		category: "electronic",
		price:    250,
	}
	prod.name = "Dell"
	prod.price = 200
	fmt.Println(prod)

	//default zero value
	var tv Product
	tv.name = "Politron"

	meja := Product{}
	fmt.Println(tv, meja)
}

func initProdukPointer() {
	//built-in new function, result pointer struct
	var phone = new(Product)
	phone.name = "Samsung"
	phone.category = "HP"
	phone.price = 235.98
	fmt.Println(phone)

	// using &
	var oppo = &Product{}
	oppo.name = "Oppo 4S"
	oppo.category = "HP"
	oppo.price = 2342.23
	fmt.Println(oppo)
}

func initOrderProduct() {
	myOrder := OrderProduct{
		product:    Product{"Iphone", "HP", 355},
		totalOrder: 1,
	}
	fmt.Println(myOrder)
}

func compareStruct() {
	oppo := Product{
		name: "oppo", category: "hp", price: 185,
	}
	Iphone := Product{
		name: "oppo", category: "hp", price: 185,
	}

	fmt.Println("result: ", oppo == Iphone)
	printStruct(Iphone)
}

//anonymous struct
func printStruct(val struct {
	name, category string
	price          float64
}) {
	fmt.Println("Name: ", val.name, "Category: ", val.category, "Price: ", val.price)
}

func initSliceProduct() []Product {
	oppo := Product{
		name: "oppo", category: "hp", price: 185,
	}
	Iphone := Product{
		name: "Iphone", category: "hp", price: 200,
	}
	products := []Product{oppo, Iphone}

	//pakai append
	products = append(products,
		Product{name: "Samsung", category: "hp", price: 256})

	return products

}

func findProductbyName(products []Product, search string) Product {
	for _, v := range products {
		if v.name == search {
			return v
		}
	}
	return Product{}

}

func main() {
	fmt.Println(findProductbyName(initSliceProduct(), "iphone"))

	//compareStruct()
	//iniProduct()
	//initProdukPointer()
}
