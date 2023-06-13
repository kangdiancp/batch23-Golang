package main

import "fmt"

type Product struct {
	name, category string
	price          float32
}

type OrderProduct struct {
	product    Product
	totalOrder int
}

func initProduct() {
	//literal value
	prod := Product{
		name:     "laptop",
		category: "eletronic",
		price:    250,
	}
	prod.name = "dell laptop"
	prod.price = 355.98
	fmt.Println(prod)

	//default zero value
	var tv Product
	tv.name = "sharp"

	meja := Product{}
	fmt.Println(tv, meja)
}

func initProductPointer() {
	//built-in new function, result pointer struct
	var phone = new(Product)
	phone.name = "samsung"
	phone.category = "hp"
	phone.price = 235.98

	fmt.Println(phone)

	//using &
	var oppo = &Product{}
	oppo.name = "oppo"
	oppo.category = "Hp"
	oppo.price = 185

	fmt.Println(oppo)
}

func initOrderProduct() {
	myOrder := OrderProduct{
		product:    Product{"iphone", "HP", 355},
		totalOrder: 1,
	}
	fmt.Println(myOrder)
}

func compareStruct() {
	// var oppo = &Product{}
	// oppo.name = "oppo"
	// oppo.category = "Hp"
	// oppo.price = 185

	// var iphone = &Product{}
	// iphone.name = "iphone"
	// iphone.category = "Hp"
	// iphone.price = 185

	oppo := Product{
		name: "oppo", category: "hp", price: 185,
	}

	iphone := Product{
		name: "oppo", category: "hp", price: 185,
	}

	fmt.Println("result :", oppo == iphone)
}

//anonymous Struct
func printStruct(val struct {
	name, category string
	price          float64
}) {
	fmt.Println("name :", val.name, "category :", val.category, "price :", val.price)
}

func initSliceProduct() []Product {
	oppo := Product{
		name: "oppo", category: "hp", price: 185,
	}

	iphone := Product{
		name: "iphone", category: "hp", price: 265,
	}
	product := []Product{oppo, iphone}

	product = append(product,
		Product{name: "samsung", category: "hp", price: 256})

	return product
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
	//printStruct()
	fmt.Println(findProductByName(initSliceProduct(), "samsung"))
}
