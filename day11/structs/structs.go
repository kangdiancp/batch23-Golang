package main

import "fmt"

type Product struct {
	name, category string
	price          float64
}

type Orderproduct struct {
	product    Product //embedded fields
	totalOrder int     // regular fields
}

func initProduct() {
	//literal value
	prod := Product{
		name:     "Laptop",
		category: "Electronic",
		price:    250,
	}

	prod.name = "Dell Laptop"
	prod.price = 355.98
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

	//using &
	var oppo = &Product{}
	oppo.name = "Oppo"
	oppo.category = "Hp"
	oppo.price = 185

	fmt.Println(oppo)
}

func initOrderProduct() {
	myOrder := Orderproduct{
		product:    Product{"Iphone", "Hp", 355},
		totalOrder: 1,
	}

	fmt.Println(myOrder)
}

func compareStruct() {
	oppo := Product{
		name: "Oppo", category: "hp", price: 185,
	}

	iphone := Product{
		name: "Oppo", category: "hp", price: 185,
	}
	fmt.Println("result :", oppo == iphone)
}

//anonymous struct
func printStruct(val struct {
	name, category string
	price          float64
}) {
	fmt.Println("Name : ", val.name, "Category : ", val.category, "Price : ", val.price)
}

func initSliceProduct() []Product {
	oppo := Product{
		name: "Oppo", category: "hp", price: 185,
	}

	iphone := Product{
		name: "Oppo", category: "hp", price: 185,
	}
	products := []Product{oppo, iphone}

	products = append(products,
		Product{name: "Samsung", category: "Hp", price: 256})

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
	//initProduct()
	//initProductPointer()
	//initOrderProduct()
	//compareStruct()
	fmt.Println(findProductByName(initSliceProduct(), "Samsung"))
}
