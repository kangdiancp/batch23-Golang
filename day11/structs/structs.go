package main

import "fmt"

type Product struct {
	name, category string
	price          float64
}

type OrderProduct struct {
	product    Product // Embedded Value
	totalOrder int     // Reguler Field
}

func initProduct() {
	// Literal Value
	prod := Product{
		name:     "Laptop",
		category: "Electronic",
		price:    250,
	}

	prod.name = "ASUS"
	prod.price = 355.98
	fmt.Println(prod)

	// Default Zero Value
	var TV Product
	TV.name = "TCL"

	meja := Product{}
	fmt.Println(TV, meja)
}

func initProductPointer() {
	// Build-in new function,result pointer struct
	var phone = new(Product)
	phone.name = "Sumsang"
	phone.category = "SmartPhone"
	phone.price = 555.55

	fmt.Println(phone)

	// Using &
	var oppo = &Product{}
	oppo.name = "Oppo"
	oppo.category = "SmartCam"
	oppo.price = 55.555

	fmt.Println(oppo)
}

func initOrderProduct() {
	myOrder := OrderProduct{
		product:    Product{"Ipon", "SmartIG", 12345},
		totalOrder: 5,
	}

	fmt.Println(myOrder)
}

func compareStruct() {
	var oppo = &Product{}
	oppo.name = "Oppo"
	oppo.category = "SmartCam"
	oppo.price = 55.555

	var vivo = &Product{}
	vivo.name = "Vivo"
	vivo.category = "SmartCam"
	vivo.price = 56.789

	fmt.Println("Result : ", oppo == vivo)
	printStruct(*oppo)
}

// Anonymous Struct
func printStruct(val struct {
	name, category string
	price          float64
}) {
	fmt.Println("Name : ", val.name, ", Category : ", val.category, ", Price : ", val.price)
}

func initSliceProduct() []Product {
	oppo := Product{
		name: "Oppo", category: "SP", price: 111,
	}

	vivo := Product{
		name: "Vivo", category: "SP", price: 101,
	}

	products := []Product{oppo, vivo}

	products = append(products, Product{name: "Sumsang", category: "SP", price: 543})

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
	initProduct()

	fmt.Println()

	initProductPointer()

	fmt.Println()

	initOrderProduct()

	fmt.Println()

	compareStruct()

	fmt.Println()

	fmt.Println(initSliceProduct())

	fmt.Println()

	fmt.Println(findProductByName(initSliceProduct(), "Sumsang"))
}
